package ubx

import (
	"bytes"
	"encoding/binary"
	"io"
)

func Encode(w io.Writer, payload Message) error {
	dsc := payload.Descriptor()
	var buf bytes.Buffer
	buf.Write([]byte{0xb5, 0x62, dsc.Class(), dsc.ID(), 0, 0})
	var err error
	if s, ok := payload.(interface{ String() string }); ok {
		_, err = buf.WriteString(s.String())
	} else {
		// TODO this will break on variable length messages
		err = binary.Write(&buf, binary.LittleEndian, payload)
	}
	if err != nil {
		return err
	}
	sz := buf.Len() - 6
	buf.Bytes()[4] = uint8(sz)
	buf.Bytes()[5] = uint8(sz >> 8)

	var a, b uint8
	for _, v := range buf.Bytes()[2:] {
		a += v
		b += a
	}
	buf.Write([]byte{a, b})
	_, err = w.Write(buf.Bytes())
	return err
}
