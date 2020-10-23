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

	switch header.ClassID {
	case (AckAck{}).Descriptor().ClassID():
		msg = &AckAck{}
	case (AckNak{}).Descriptor().ClassID():
		msg = &AckNak{}
	case (CfgCfg1{}).Descriptor().ClassID():
		msg = &CfgCfg1{}
	case (CfgCfg2{}).Descriptor().ClassID():
		msg = &CfgCfg2{}
	case (CfgHnr{}).Descriptor().ClassID():
		msg = &CfgHnr{}
	case (CfgMsg1{}).Descriptor().ClassID():
		msg = &CfgMsg1{}
	case (CfgMsg2{}).Descriptor().ClassID():
		msg = &CfgMsg2{}
	case (CfgMsg3{}).Descriptor().ClassID():
		msg = &CfgMsg3{}
	case (CfgRate{}).Descriptor().ClassID():
		msg = &CfgRate{}
	case (NavPvt{}).Descriptor().ClassID():
		msg = &NavPvt{}
	}

	err = binary.Read(buf, binary.LittleEndian, &msg)

	return msg, err

}
