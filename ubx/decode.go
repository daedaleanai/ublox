package ubx

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"reflect"
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

func decode(r io.Reader, msg interface{}) (err error) {
	//	log.Printf("Decoding %T %#v", msg, msg)

	defer func() {
		if x, _ := recover().(error); x != nil {
			err = x
		}
	}()

	readN := func(n int) (v uint64) {
		var buf [8]byte
		if _, err := io.ReadFull(r, buf[:n]); err != nil {
			panic(err)
		}
		for i := 0; i < n; i++ {
			v |= uint64(buf[i]) << uint(8*i)
		}
		return
	}

	switch v := msg.(type) {
	case *byte:
		*v = byte(readN(1))
		return nil
	case *int8:
		*v = int8(readN(1))
		return nil
	case *uint16:
		*v = uint16(readN(2))
		return nil
	case *int16:
		*v = int16(readN(2))
		return nil
	case *uint32:
		*v = uint32(readN(4))
		return nil
	case *int32:
		*v = int32(readN(4))
		return nil
	case *uint64:
		*v = uint64(readN(8))
		return nil
	case *int64:
		*v = int64(readN(8))
		return nil
	case *float32:
		*v = math.Float32frombits(uint32(readN(4)))
		return nil
	case *float64:
		*v = math.Float64frombits(uint64(readN(8)))
		return nil

	case []byte:
		_, err := io.ReadFull(r, v) // msg must already have size
		return err
	case *string: // read rest of r
		b, err := ioutil.ReadAll(r)
		*v = string(b)
		return err
	}

	v := reflect.Indirect(reflect.ValueOf(msg))
	// fields that are type Xn are handled here as uint<8*n>
	switch v.Kind() {
	case reflect.Uint8:
		var vv uint8
		err := decode(r, &vv)
		v.SetUint(uint64(vv))
		return err
	case reflect.Uint16:
		var vv uint16
		err := decode(r, &vv)
		v.SetUint(uint64(vv))
		return err
	case reflect.Uint32:
		var vv uint32
		err := decode(r, &vv)
		v.SetUint(uint64(vv))
		return err
	case reflect.Uint64:
		var vv uint64
		err := decode(r, &vv)
		v.SetUint(uint64(vv))
		return err
	case reflect.String:
		b, err := ioutil.ReadAll(r)
		v.SetString(string(b))
		return err

	case reflect.Ptr:
		v.Set(reflect.New(v.Type().Elem()))
		return decode(r, v.Interface())

	case reflect.Array, reflect.Slice:
		l := v.Len()
		for i := 0; i < l; i++ {
			if err := decode(r, v.Index(i).Addr().Interface()); err != nil {
				return err
			}
		}
		return nil

	case reflect.Struct:
		t := v.Type()
		l := v.NumField()
		for i := 0; i < l; i++ {
			if err := decode(r, v.Field(i).Addr().Interface()); err != nil {
				return err
			}
			// if the field is a NumXXX for the XXX []... bit, set it to the length here
			if s := t.Field(i).Tag.Get("len"); s != "" {
				for ii := l - 1; ii > i; ii-- {
					if t.Field(ii).Name == s && t.Field(ii).Type.Kind() == reflect.Slice {
						sz := int(v.Field(i).Uint())
						if sz != 0 {
							v.Field(ii).Set(reflect.MakeSlice(t.Field(ii).Type, sz, sz))
						}
						break
					}
				}
			}
		}
		return nil
	}

	panic(fmt.Errorf("Cannot decode field of type %T (%v)", msg, v.Kind()))
}
