package miio

import (
	"log"
	"net"
	"testing"
)

func TestConn(t *testing.T) {
	c, err := Connect(&net.UDPAddr{
		IP:   net.IPv4(192, 168, 1, 168),
		Port: 54321,
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Write(helloPacket); err != nil {
		t.Fatal(err)
	}
	p, err := c.Read()
	if err != nil {
		t.Fatal(err)
	}
	t.Error(p.Header())
}

func dispatch(ch <-chan packet) {
	for p := range ch {
		log.Printf("%x", []byte(p.Header()[:]))
	}
}
