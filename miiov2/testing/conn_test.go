package miio

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"testing"
	"time"
)

var (
	purifierIP    = net.IPv4(192, 168, 1, 241)
	purifierToken = "1269e2626b7dd566050959cdb8cab474"
	acIP          = net.IPv4(192, 168, 1, 168)
	acToken       = "9e66782fd03c1f933f40193d181134c0"
	ac2IP         = net.IPv4(192, 168, 1, 111)
	ac2Token      = "9e66782fd03c1f933f40193d181134c0"
	tvIP          = net.IPv4(192, 168, 1, 203)
	tvToken       = "6531383872346b79584e364434536370"
	fanIP         = net.IPv4(192, 168, 1, 148)
	fanToken      = "6bcf039baf2374bd44d8a455d95b7738"
)

func TestC(t *testing.T) {
	// c, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: purifierIP, Port: 54321})
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// defer c.Close()
	// c.SetReadDeadline(time.Now().Add(2 * time.Second))
	// go func() {
	// 	buf := make([]byte, 1024)
	// 	for {
	// 		n, err := c.Read(buf)
	// 		if err != nil {
	// 			t.Error(err)
	// 			continue
	// 		}
	// 		log.Printf("read %d bytes: %x", n, buf[:n])
	// 	}
	// }()
	// time.Sleep(3 * time.Second)

	c, err := dialChannel(&net.UDPAddr{
		IP:   purifierIP,
		Port: 54321,
	}, purifierToken, 65492010)
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		if err := c.listen(); err != nil {
			t.Fatal(err)
		}
	}()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := c.hello(ctx); err != nil {
		t.Fatal(err)
	}
	if err := c.Write([]byte(`{"id":879855,"method":"miIO.info","params":null}`)); err != nil {
		t.Fatal(err)
	}
	b, err := c.Read()
	if err != nil {
		t.Fatal(err)
	}
	t.Errorf("%s", b)
}

// Channel is a UDP connection to a device. It handles the connection
// encryption, and state management.
type channel struct {
	c         *net.UDPConn
	state     int
	pc        *payloadCipher
	token     [16]byte
	lastRx    time.Time
	clockSkew int64
	deviceID  uint32

	helloCh chan packet
	respCh  chan packet
}

func dialChannel(addr *net.UDPAddr, token string, deviceID uint32) (*channel, error) {
	c, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}
	var tk [16]byte
	if n, err := hex.Decode(tk[:], []byte(token)); err != nil || n != len(tk) {
		return nil, fmt.Errorf("invalid token: %v", err)
	}
	return &channel{
		c:        c,
		pc:       newPayloadCipher(tk),
		token:    tk,
		deviceID: deviceID,

		helloCh: make(chan packet),
		respCh:  make(chan packet, 1),
	}, nil
}

func (c *channel) hello(ctx context.Context) error {
	if err := c.writePacket(helloPacket); err != nil {
		return err
	}
	select {
	case p := <-c.helloCh:
		c.clockSkew = time.Now().Unix() - int64(p.Header().Timestamp())
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (c *channel) writePacket(p packet) error {
	// c.lastWrite = time.Now()
	// c.idleTimer.extend(5 * time.Minute)
	_, err := c.c.Write(p)
	return err
}

func (c *channel) readPacket() (packet, error) {
	p := make(packet, 1<<16-1)
	n, err := c.c.Read(p)
	if err != nil {
		return nil, err
	}
	p = p[:n]
	if err := p.Verify(); err != nil {
		return nil, err
	}
	if p.Header().PayloadSize() != 0 {
		if err := p.VerifyPayload(c.token); err != nil {
			return nil, err
		}
	}
	c.lastRx = time.Now()
	return p, nil
}

func (c *channel) listen() error {
	for {
		p, err := c.readPacket()
		if err != nil {
			log.Printf("%T", err)
			return err
		}
		fmt.Printf("%s\n", p.Header())
		ch := c.respCh
		if p.Header().PayloadSize() == 0 {
			ch = c.helloCh
		}
		select {
		case ch <- p:
		default:
			log.Printf("dropping packet: %v", p)
		}
	}
}

func (c *channel) Read() ([]byte, error) {
	p := <-c.respCh
	return p.Payload(), nil
}

func (c *channel) Write(b []byte) error {
	p := make(packet, 1<<16-1)
	h := p.Header()
	h.SetMagic()
	h.SetDeviceID(c.deviceID)
	h.SetTimestamp(uint32(time.Now().Unix() - c.clockSkew))
	p.SetPayload(c.pc.Encrypt(p[headerSize:], b))
	return c.writePacket(p)
}
