package miio

import (
	"bytes"
	"fmt"
	"net"
	"sync"
)

type ConnV2 struct {
	remote *net.UDPAddr

	mu        sync.RWMutex
	controlCh chan any
	c         *net.UDPConn

	txCh chan packet
	rxCh chan packet
}


func (c *ConnV2) controller() {
	for {
		select
	}
}

func (c *ConnV2) sender() {
	for p := range c.txCh {
		// grab an active conn, or create one.
		if err := c.writePacket(p); err != nil {
			c.controlCh <- writeError{err}
			break
		}
	}
}

func (c *ConnV2) receiver() {
	for {
		p, err := c.readPacket()
		if err != nil {
			// the current conn is dead.
			c.controlCh <- readError{err}
			break
		}
		if isHello := p.Header().Unknown1 == 0xffffffff; isHello {
			c.controlCh <- recivedHello{time.Now()}
			continue
		}
		select {
		case c.rxCh <- p:
		default:
			log.Printf("rxCh is full, dropping packet: %v", p)
		}
	}
}

func (c *CnnnV2) hello() error {
	if err := c.writePacket(helloPacket); err != nil {
		return err
	}
	p, err := c.readPacket()
	if err != nil {
		return err
	}
	if !bytes.Equal(p, helloPacket) {
		return fmt.Errorf("invalid hello response: %v", p)
	}
	return nil
}
