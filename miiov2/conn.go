// Package miio implements miio protocol based on:
// - https://github.com/OpenMiHome/mihome-binary-protocol/blob/master/doc/PROTOCOL.md
// - https://sspai.com/post/68306
package miio

import (
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
)

const chanBufferSize = 10

// Conn represents a connection to a miio device.
type Conn struct {
	remoteAddr *net.UDPAddr
	token      [16]byte
	pc         *payloadCipher

	readCh  chan packet
	writeCh chan packet

	mu        sync.Mutex
	c         *net.UDPConn
	clockSkew int64
	id        uint32
}

// Connect creates a lazy connection to the miio device.
//
// As miio protocol uses UDP, dialing is a one-party operation. We won't know if
// the configuration is correct or the miio device is online until we send a
// packet.
func Connect(addr *net.UDPAddr, token string) (*Conn, error) {
	c := &Conn{
		remoteAddr: addr,

		readCh:  make(chan packet, chanBufferSize),
		writeCh: make(chan packet, chanBufferSize),
	}
	n, err := hex.Decode(c.token[:], []byte(token))
	if err != nil {
		return nil, err
	}
	if n != len(c.token) {
		return nil, fmt.Errorf("invalid token: %v", err)
	}
	c.pc = newPayloadCipher(c.token)
	return c, nil
}

func (c *Conn) writePacket(p packet) error {
	_, err := c.c.Write(p)
	return err
}

func (c *Conn) readPacket() (packet, error) {
	p := make(packet, 1<<16-1)
	n, err := c.c.Read(p)
	if err != nil {
		return nil, err
	}
	p = p[:n]
	if err := p.Verify(); err != nil {
		return nil, err
	}
	return p, nil
}

func (c *Conn) Close() error {
	return c.c.Close()
}

func (c *Conn) dial() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.c != nil {
		return nil
	}
	conn, err := net.DialUDP("udp", nil, c.remoteAddr)
	if err != nil {
		return err
	}
	c.c = conn
	if err := c.handshake(); err != nil {
		return err
	}
	return nil
}

func (c *Conn) handshake() error {
	if _, err := c.c.Write(helloPacket); err != nil {
		return err
	}
	p := make(packet, len(helloPacket))
	if _, err := io.ReadFull(c.c, p); err != nil {
		return err
	}
	// TODO: verify hello response packet
	c.id = p.Header().DeviceID()
	c.clockSkew = int64(p.Header().Timestamp()) - now().Unix()
	return nil
}

func (c *Conn) ID() string {
	return strconv.FormatUint(uint64(c.id), 10)
}

func (c *Conn) Write(payload []byte) error {
	p := make(packet, 1<<16-1)
	h := p.Header()
	h.SetMagic()
	h.SetDeviceID(c.id)
	h.SetPayloadSize(0)
	h.SetTimestamp(uint32(now().Unix() + c.clockSkew))
	p.SetPayload(c.pc.Encrypt(p.Payload(), payload))
	p.SignPayload(c.token)
	return c.writePacket(p)
}
