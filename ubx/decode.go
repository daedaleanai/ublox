package ubx

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

var (
	errInvalidFrame  = errors.New("invalid UBX frame")
	errInvalidChkSum = errors.New("invalid UBX checksum")
)

type Message interface {
	classID() uint16
}

type RawMessage struct {
	ClassID uint16
	Data    []byte
}

func (msg *RawMessage) classID() uint16 { return msg.ClassID }

func Decode(frame []byte) (msg Message, err error) {

	buf := bytes.NewReader(frame)

	var header struct {
		Preamble uint16
		ClassID  uint16
		Length   uint16
	}

	if err := binary.Read(buf, binary.LittleEndian, &header); err != nil {
		return nil, err
	}

	if header.Preamble != 0x62B5 {
		return nil, errInvalidFrame
	}

	if buf.Len()+2 < int(header.Length) {
		return nil, io.ErrShortBuffer
	}

	var a, b byte
	for _, v := range frame[2 : header.Length+6] {
		a += byte(v)
		b += a
	}

	if frame[header.Length+6] != a || frame[header.Length+7] != b {
		return nil, errInvalidChkSum
	}

	msg = mkMsg(header.ClassID, header.Length, frame[6:len(frame)-2])

	if msg != nil {
		err = decode(bytes.NewReader(frame[6:len(frame)-2]), msg)
	} else {
		msg = &RawMessage{ClassID: header.ClassID, Data: append([]byte(nil), frame[6:len(frame)-2]...)}
	}

	return msg, err

}

func decode(r io.Reader, v interface{}) error { return nil }
