package ubx

import (
	"fmt"
	"io"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func testType(t *testing.T, tc Message) {

	if err := randFill(tc); err != nil {
		t.Fatal(err)
	}

	// hack.  CfgPrt1 is default, but the PortID subtype field zero is CfgPrt4
	if m, ok := tc.(*CfgPrt1); ok {
		m.PortID = 0xff
	}

	buf, err := Encode(tc)
	if err != nil {
		t.Errorf("Encoding %T: %v", tc, err)
		return
	}

	//	t.Logf("encoded %x", buf)

	t1, err := Decode(buf)
	if err != nil {
		t.Errorf("Decoding %T: %v", tc, err)
		return
	}
	if !reflect.DeepEqual(tc, t1) {
		t.Logf("Encoded %T: %#v", tc, tc)
		t.Logf("Decoded %T: %#v", t1, t1)
		t.Errorf("Endoding/Decoding failure for %T", tc)
	}
}

var r = rand.New(rand.NewSource(time.Now().Unix()))

func randFill(msg interface{}) error {
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
		*v = r.Float32()
		return nil
	case *float64:
		*v = r.Float64()
		return nil

	case []byte:
		_, err := io.ReadFull(r, v) // msg must already have size
		return err
	case *string: // read rest of r
		*v = "test2test"
		return nil
	}

	v := reflect.Indirect(reflect.ValueOf(msg))
	// fields that are type Xn are handled here as uint<8*n>
	switch v.Kind() {
	case reflect.Uint8:
		var vv uint8
		err := randFill(&vv)
		v.SetUint(uint64(vv))
		return err
	case reflect.Uint16:
		var vv uint16
		err := randFill(&vv)
		v.SetUint(uint64(vv))
		return err
	case reflect.Uint32:
		var vv uint32
		err := randFill(&vv)
		v.SetUint(uint64(vv))
		return err
	case reflect.Uint64:
		var vv uint64
		err := randFill(&vv)
		v.SetUint(uint64(vv))
		return err
	case reflect.String:
		v.SetString("test1test")
		return nil

	case reflect.Ptr:
		v.Set(reflect.New(v.Type().Elem()))
		return randFill(v.Interface())

	case reflect.Array, reflect.Slice:
		l := v.Len()
		for i := 0; i < l; i++ {
			if err := randFill(v.Index(i).Addr().Interface()); err != nil {
				return err
			}
		}
		return nil

	case reflect.Struct:
		t := v.Type()
		l := v.NumField()
		for i := 0; i < l; i++ {
			if err := randFill(v.Field(i).Addr().Interface()); err != nil {
				return err
			}
			// if the field is a NumXXX for the XXX []... bit, set it to the length here
			if s := t.Field(i).Tag.Get("len"); s != "" {
				for ii := l - 1; ii > i; ii-- {
					if t.Field(ii).Name == s && t.Field(ii).Type.Kind() == reflect.Slice {
						sz := rand.Intn(5)
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

	return fmt.Errorf("Cannot randomize field of type %T (%v)", msg, v.Kind())
}
