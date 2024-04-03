package ubx

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"reflect"
	"strconv"
)

// Encode can serialize a message into a buffer
func Encode(payload Message) (buf []byte, err error) {
	var b bytes.Buffer
	b.Write([]byte{0xb5, 0x62, byte(payload.classID()), byte(payload.classID() >> 8), 0, 0})

	if err := encode(&b, payload); err != nil {
		return nil, err
	}

	sz := b.Len() - 6
	b.Bytes()[4] = byte(sz)
	b.Bytes()[5] = byte(sz >> 8)

	var x, y byte
	for _, v := range b.Bytes()[2:] {
		x += v
		y += x
	}
	b.Write([]byte{x, y})
	return b.Bytes(), nil
}

func encode(w io.Writer, msg interface{}) error {
	//	log.Printf("Encoding %T %#v", msg, msg)
	switch v := msg.(type) {
	case byte:
		_, err := w.Write([]byte{v})
		return err
	case int8:
		_, err := w.Write([]byte{byte(v)})
		return err
	case uint16:
		_, err := w.Write([]byte{byte(v), byte(v >> 8)})
		return err
	case int16:
		_, err := w.Write([]byte{byte(v), byte(v >> 8)})
		return err
	case uint32:
		_, err := w.Write([]byte{byte(v), byte(v >> 8), byte(v >> 16), byte(v >> 24)})
		return err
	case int32:
		_, err := w.Write([]byte{byte(v), byte(v >> 8), byte(v >> 16), byte(v >> 24)})
		return err
	case uint64:
		_, err := w.Write([]byte{byte(v), byte(v >> 8), byte(v >> 16), byte(v >> 24), byte(v >> 32), byte(v >> 40), byte(v >> 48), byte(v >> 56)})
		return err
	case int64:
		_, err := w.Write([]byte{byte(v), byte(v >> 8), byte(v >> 16), byte(v >> 24), byte(v >> 32), byte(v >> 40), byte(v >> 48), byte(v >> 56)})
		return err
	case float32:
		vv := math.Float32bits(v)
		_, err := w.Write([]byte{byte(vv), byte(vv >> 8), byte(vv >> 16), byte(vv >> 24)})
		return err
	case float64:
		vv := math.Float64bits(v)
		_, err := w.Write([]byte{byte(vv), byte(vv >> 8), byte(vv >> 16), byte(vv >> 24), byte(vv >> 32), byte(vv >> 40), byte(vv >> 48), byte(vv >> 56)})
		return err

	case []byte:
		_, err := w.Write(v)
		return err
	case string:
		_, err := w.Write([]byte(v))
		return err

	}

	v := reflect.Indirect(reflect.ValueOf(msg))
	// fields that are type Xn are handled here as uint<8*n>
	switch v.Kind() {
	case reflect.Uint8:
		return encode(w, uint8(v.Uint()))
	case reflect.Uint16:
		return encode(w, uint16(v.Uint()))
	case reflect.Uint32:
		return encode(w, uint32(v.Uint()))
	case reflect.Uint64:
		return encode(w, v.Uint())
	case reflect.String:
		return encode(w, v.String())

	case reflect.Array, reflect.Slice:
		l := v.Len()
		for i := 0; i < l; i++ {
			if err := encode(w, v.Index(i).Interface()); err != nil {
				return err
			}
		}
		return nil

	case reflect.Struct:
		t := v.Type()
		l := v.NumField()
		for i := 0; i < l; i++ {
			// if the field is a NumXXX for the XXX []... bit, set it to the lenght here
			if s := t.Field(i).Tag.Get("len"); s != "" {
				for ii := l - 1; ii > i; ii-- {
					if t.Field(ii).Name == s {
						v.Field(i).SetUint(uint64(v.Field(ii).Len()))
						break
					}
				}
			}
			if s := t.Field(i).Tag.Get("stf"); s != "" && s != "default" {
				n, err := strconv.ParseUint(s, 0, 8)
				if err != nil {
					log.Fatalf("invalid stf tag %v %v", t.Field(i).Tag, err)
				}
				v.Field(i).SetUint(n)
			}
			if err := encode(w, v.Field(i).Interface()); err != nil {
				return err
			}
		}
		return nil
	}

	return fmt.Errorf("Cannot encode field of type %T (%v) %v", msg, v.Kind(), msg)
}
