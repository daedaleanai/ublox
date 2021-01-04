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
	case (AckAck{}).ClassID():
		msg = &AckAck{}
	case (AckNak{}).ClassID():
		msg = &AckNak{}
	case (CfgCfg1{}).ClassID():
		msg = &CfgCfg1{}
	case (CfgCfg2{}).ClassID():
		msg = &CfgCfg2{}
	case (CfgHnr{}).ClassID():
		msg = &CfgHnr{}
	case (CfgMsg1{}).ClassID():
		msg = &CfgMsg1{}
	case (CfgMsg2{}).ClassID():
		msg = &CfgMsg2{}
	case (CfgMsg3{}).ClassID():
		msg = &CfgMsg3{}
	case (CfgRate{}).ClassID():
		msg = &CfgRate{}
	case (CfgPrt1{}).ClassID():
		msg = &CfgPrt1{}
	case (NavPvt{}).ClassID():
		msg = &NavPvt{}
	default:
	}
	if msg != nil {
		err = binary.Read(buf, binary.LittleEndian, msg)
	} else {
		msg = &RawMessage{CID: header.ClassID, Data: append([]byte(nil), frame[6:len(frame)-2]...)}
	}

	return msg, err

}
