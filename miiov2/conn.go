package miio

import "net"

type Conn struct {
	c *net.UDPConn
}

func Connect(addr *net.UDPAddr) (*Conn, error) {
	c, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}
	return &Conn{c}, nil
}

func (c *Conn) Write(p packet) error {
	_, err := c.c.Write(p)
	return err
}

func (c *Conn) Read() (packet, error) {
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
