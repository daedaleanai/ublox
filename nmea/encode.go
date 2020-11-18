package nmea

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// Encode can serialize a slice of string fields or a struct
func Encode(msg interface{}) (buf []byte, err error) {
	var fields []string
	if f, ok := msg.([]string); ok {
		fields = f
	} else {
		rmsg := reflect.ValueOf(msg)
		if rmsg.Kind() == reflect.Ptr {
			rmsg = rmsg.Elem()
		}
		if rmsg.Kind() != reflect.Struct {
			return nil, fmt.Errorf("message must be a struct, *struct, or []string, got %T", msg)
		}
		var err error
		fields, err = encodeMsg(msg)
		if err != nil {
			return nil, err
		}
	}

	s := strings.Join(fields, ",")
	var x byte
	for _, v := range s {
		x ^= byte(v)
	}
	var b bytes.Buffer
	_, err = fmt.Fprintf(&b, "$%s*%02X\r\n", s, x)
	return b.Bytes(), err
}

// msg is known to be a struct
func encodeMsg(msg interface{}) (fields []string, err error) {
	rmsg := reflect.ValueOf(msg)
	for i := 0; i < rmsg.NumField(); i++ {
		switch fld := rmsg.Field(i).Interface().(type) {
		case Header:
			if fld == "" {
				// fill in from type
			}
		case PUBXType:

		}
	}
	return nil, nil
}
