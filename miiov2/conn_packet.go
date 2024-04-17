package miio

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/binary"
	"fmt"
)

const (
	magic      = 0x2131
	headerSize = 0x20
)

type header [headerSize]byte

func (h *header) SetMagic() {
	binary.BigEndian.PutUint16(h[0:2], magic)
}

func (h *header) SetPayloadSize(l uint16) {
	binary.BigEndian.PutUint16(h[2:4], l+headerSize)
}

func (h *header) SetUnknonw1(v uint32) {
	binary.BigEndian.PutUint32(h[4:8], v)
}

func (h *header) SetDeviceID(id uint32) {
	binary.BigEndian.PutUint32(h[8:12], id)
}

func (h *header) SetTimestamp(ts uint32) {
	binary.BigEndian.PutUint32(h[12:16], ts)
}

func (h *header) ClearChecksum() {
	for i := 16; i < 32; i++ {
		h[i] = 0
	}
}

func (h *header) SetChecksum(c [16]byte) {
	copy(h[16:32], c[:])
}

func (h *header) String() string {
	return fmt.Sprintf("PayloadSize:%d Unknown:%d Did:%d Stamp:%d",
		h.PayloadSize(), h.Unknown1(), h.DeviceID(), h.Timestamp())
}

func (h *header) Magic() uint16 {
	return binary.BigEndian.Uint16(h[0:2])
}

func (h *header) size() uint16 {
	return binary.BigEndian.Uint16(h[2:4])
}

func (h *header) PayloadSize() uint16 {
	return h.size() - headerSize
}

func (h *header) Unknown1() uint32 {
	return binary.BigEndian.Uint32(h[4:8])
}

func (h *header) DeviceID() uint32 {
	return binary.BigEndian.Uint32(h[8:12])
}

func (h *header) Timestamp() uint32 {
	return binary.BigEndian.Uint32(h[12:16])
}

func (h *header) GetChecksum() (res [16]byte) {
	copy(res[:], h[16:32])
	return
}

type packet []byte

var (
	errShort         = fmt.Errorf("short packet")
	errMagicMismatch = fmt.Errorf("invalid magic")
	errSizeMismatch  = fmt.Errorf("invalid magic")
)

func (p packet) Verify() error {
	if len(p) < headerSize {
		return errShort
	}
	if p.Header().Magic() != magic {
		return errMagicMismatch
	}
	if p.Header().size() != uint16(len(p)) {
		return errSizeMismatch
	}
	return nil
}

func (p packet) String() string {
	return fmt.Sprintf("HEADER: %s PAYLOAD: %x", p.Header(), p.Payload())
}

func (p packet) Header() *header {
	return (*header)(p)
}

func (p packet) Payload() []byte {
	return p[headerSize:][:p.Header().PayloadSize()]
}

func (p *packet) SetPayload(payload []byte) {
	fmt.Println(len(payload))
	p.Header().SetPayloadSize(uint16(len(payload)))
	*p = (*p)[:p.Header().size()]
	copy(p.Payload(), payload)
}

func (p packet) SignPayload(token [16]byte) {
	h := p.Header()
	h.SetChecksum(token)
	h.SetChecksum(md5.Sum(p))
}

func (p packet) VerifyPayload(token [16]byte) error {
	h := p.Header()
	want := h.GetChecksum()
	h.SetChecksum(token)
	if got := md5.Sum(p); got != want {
		return fmt.Errorf("checksum mismatch: got %x, want %x", got, want)
	}
	return nil
}

var helloPacket = packet{
	0x21, 0x31, 0x00, 0x20,
	0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff,
}

type payloadCipher struct {
	b  cipher.Block
	iv []byte
}

func newPayloadCipher(token [16]byte) *payloadCipher {
	key := md5.Sum(token[:])
	h := md5.New()
	h.Write(key[:])
	h.Write(token[:])
	iv := h.Sum(nil)
	b, err := aes.NewCipher(key[:])
	if err != nil {
		panic(err)
	}
	return &payloadCipher{
		b:  b,
		iv: iv,
	}
}

func (c *payloadCipher) Encrypt(dst, src []byte) []byte {
	e := cipher.NewCBCEncrypter(c.b, c.iv)
	src = pkcs7padding(src)
	if dst == nil {
		dst = make([]byte, len(src))
	} else {
		dst = dst[:len(src)]
	}
	e.CryptBlocks(dst, src)
	return dst[:len(src)]
}

func (c *payloadCipher) Decrypt(dst, src []byte) ([]byte, error) {
	if len(src)%16 != 0 {
		return nil, fmt.Errorf("invalid length")
	}
	d := cipher.NewCBCDecrypter(c.b, c.iv)
	if dst == nil {
		dst = make([]byte, len(src))
	} else {
		dst = dst[:len(src)]
	}
	d.CryptBlocks(dst, src)
	return pkcs7unpadding(dst)
}

func pkcs7padding(b []byte) []byte {
	l := len(b)
	pad := 16 - l%16
	for i := 0; i < pad; i++ {
		b = append(b, byte(pad))
	}
	return b
}

func pkcs7unpadding(b []byte) ([]byte, error) {
	if len(b) == 0 {
		return nil, fmt.Errorf("empty")
	}
	pad := b[len(b)-1]
	if pad > 16 {
		return nil, fmt.Errorf("invalid padding")
	}
	for _, b := range b[len(b)-int(pad):] {
		if b != pad {
			return nil, fmt.Errorf("invalid padding")
		}
	}
	return b[:len(b)-int(pad)], nil
}
