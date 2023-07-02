package miio

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"sup/miio/devices"
	"sync"
	"time"
)

// now stab out for testing.
var now = time.Now

type Device struct {
	c     *Conn
	id    uint32
	token [16]byte
	pc    *payloadCipher

	clockSkew int64

	mu           sync.Mutex
	ongoingCalls map[int]chan<- *callResponse
}

func ConenctDevice(addr *net.UDPAddr, token string) (*Device, error) {
	c, err := Connect(addr)
	if err != nil {
		return nil, err
	}
	tk := [16]byte{}
	if n, err := hex.Decode(tk[:], []byte(token)); err != nil || n != len(tk) {
		return nil, fmt.Errorf("invalid token: %v", err)
	}
	d := &Device{
		c:            c,
		token:        tk,
		pc:           newPayloadCipher(tk),
		ongoingCalls: make(map[int]chan<- *callResponse),
	}
	// if err := d.hello(); err != nil {
	// 	c.Close()
	// 	return nil, err
	// }
	go d.listenCallResponse()
	return d, nil
}

func (d *Device) hello() error {
	if err := d.c.Write(helloPacket); err != nil {
		return err
	}
	p, err := d.c.Read()
	if err != nil {
		return err
	}
	d.id = p.Header().GetDeviceID()
	d.clockSkew = int64(p.Header().GetTimestamp()) - now().Unix()
	return nil
}

func (d *Device) ID() string {
	return strconv.FormatUint(uint64(d.id), 10)
}

func (d *Device) newPacket() packet {
	p := make(packet, 1<<16-1)
	h := p.Header()
	h.SetMagic()
	h.SetDeviceID(d.id)
	h.SetPayloadSize(0)
	h.SetTimestamp(uint32(now().Unix() + d.clockSkew))
	return p
}

func (d *Device) Call(ctx context.Context, method string, params, result interface{}) error {
	p := d.newPacket()
	buf := bytes.NewBuffer(p.Payload())
	req := &callRequest{
		ID:     rand.Intn(31),
		Method: method,
		Params: params,
	}
	if err := json.NewEncoder(buf).Encode(&req); err != nil {
		return err
	}
	log.Printf("Call %s", buf.Bytes())
	p.SetPayload(d.pc.Encrypt(p.Payload(), buf.Bytes()))
	p.SignPayload(d.token)

	ch := d.registerCall(req.ID)
	if err := d.c.Write(p); err != nil {
		return err
	}

	var resp *callResponse
	select {
	case resp = <-ch:
	case <-ctx.Done():
		d.unregisterCall(req.ID)
		return ctx.Err()
	}
	if resp.Error != nil {
		return resp.Error
	}
	return json.Unmarshal(resp.Result, &result)
}

func (d *Device) listenCallResponse() error {
	for {
		p, err := d.c.Read()
		if err != nil {
			return err
		}
		fmt.Println("listenCallResponse", p)
		if err := d.dispatchCallResponse(p); err != nil {
			fmt.Println("dispatchCallResponse", err)
		}
	}
}

type callRequest struct {
	ID     int         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

type callResponse struct {
	ID     int             `json:"id"`
	Result json.RawMessage `json:"result"`
	Error  *CallError      `json:"error,omitempty"`
}

func (d *Device) dispatchCallResponse(p packet) error {
	if err := p.VerifyPayload(d.token); err != nil {
		return err
	}
	b, err := d.pc.Decrypt(p.Payload(), p.Payload())
	if err != nil {
		return err
	}
	var resp callResponse
	if err := json.Unmarshal(b, &resp); err != nil {
		return err
	}
	log.Printf("CallResp %s", b)
	ch, ok := d.getCallHandler(resp.ID)
	if !ok {
		return fmt.Errorf("call %d: unregistered", resp.ID)
	}
	select {
	case ch <- &resp:
	default:
		return fmt.Errorf("call %d: dup", resp.ID)
	}
	return nil
}

func (d *Device) getCallHandler(id int) (ch chan<- *callResponse, ok bool) {
	d.mu.Lock()
	ch, ok = d.ongoingCalls[id]
	delete(d.ongoingCalls, id)
	d.mu.Unlock()
	return
}

func (d *Device) registerCall(id int) <-chan *callResponse {
	ch := make(chan *callResponse, 1)
	d.mu.Lock()
	d.ongoingCalls[id] = ch
	d.mu.Unlock()
	return ch
}

func (d *Device) unregisterCall(id int) {
	d.mu.Lock()
	delete(d.ongoingCalls, id)
	d.mu.Unlock()
}

type CallError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *CallError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

type DeviceInfo struct {
	FirmwareVersion string `json:"fw_ver"`
	HardwareVersion string `json:"hw_ver"`
	Model           string `json:"model"`
}

func (d *Device) Info(ctx context.Context) (*DeviceInfo, error) {
	var res DeviceInfo
	if err := d.Call(ctx, "miIO.info", nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

type PropertyID struct {
	ServiceID  int `json:"siid"`
	PropertyID int `json:"piid"`
}

type ActionID struct {
	ServiceID int `json:"siid"`
	ActionID  int `json:"aiid"`
}

type Properties []*Property

type Property struct {
	DeviceID string `json:"did"`
	PropertyID
	Value interface{} `json:"value"`
	Code  int         `json:"code"` // output only
}

func (d *Device) GetProperties(ctx context.Context, props Properties) (Properties, error) {
	var res Properties
	if err := d.Call(ctx, "get_properties", props, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (d *Device) SetProperties(ctx context.Context, props Properties) (Properties, error) {
	var res Properties
	if err := d.Call(ctx, "set_properties", props, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (d *Device) GetPropertiesV2(ctx context.Context, props any) error {
	req, err := devices.MarshalProperties(props)
	if err != nil {
		return err
	}
	var res json.RawMessage
	if err := d.Call(ctx, "get_properties", json.RawMessage(req), &res); err != nil {
		return err
	}
	return devices.UnmarshalProperties(res, props)
}

func (d *Device) SetPropertiesV2(ctx context.Context, props any) error {
	req, err := devices.MarshalProperties(props)
	if err != nil {
		return err
	}
	var res json.RawMessage
	if err := d.Call(ctx, "set_properties", json.RawMessage(req), &res); err != nil {
		return err
	}
	return devices.UnmarshalProperties(res, props)
}
