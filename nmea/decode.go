package nmea

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	errInvalidFrame  = errors.New("invalid NMEA frame")
	errInvalidChkSum = errors.New("invalid NMEA checksum")
)

const hexChar = "0123456789ABCDEF"

type RawMessage []string

func Decode(frame []byte) (msg interface{}, err error) {
	start, end := bytes.IndexByte(frame, '$'), bytes.LastIndexByte(frame, '*')
	if start == -1 || end == -1 || end+2 >= len(frame) {
		return nil, errInvalidFrame
	}
	var x byte
	for _, v := range frame[start+1 : end] {
		x ^= byte(v)
	}
	// fast uppercase
	if frame[end+1] >= 'a' {
		frame[end+1] -= 'a' - 'A'
	}
	if frame[end+2] >= 'a' {
		frame[end+2] -= 'a' - 'A'
	}
	if frame[end+1] != hexChar[x>>4] || frame[end+2] != hexChar[x&0xf] { // also lowercase?
		return nil, errInvalidChkSum
	}
	fields := strings.Split(string(frame[start+1:end]), ",")
	msg = mkMsg(fields[0])
	if msg == nil && fields[0] == "PUBX" && len(fields) > 2 {
		t, err := strconv.ParseInt(fields[1], 10, 0)
		if err == nil {
			msg = NewPUBX(PUBXType(t))
		}
	}
	if msg != nil {
		err = decodeMsg(msg, fields)
		return msg, err
	}

	return RawMessage(fields), nil
}

// https://play.golang.org/p/_GK6kdk1SdV
func decodeMsg(msg interface{}, fields []string) error {
	if reflect.ValueOf(msg).Kind() != reflect.Ptr {
		return fmt.Errorf("message must be a pointer to struct")
	}
	rmsg := reflect.ValueOf(msg).Elem()
	if rmsg.Kind() != reflect.Struct {
		return fmt.Errorf("message must be a pointer to struct")
	}

	for i, f := 0, 0; i < rmsg.NumField() && f < len(fields); i, f = i+1, f+1 {

		v := rmsg.Field(i)

		// array of int, float
		if v.Kind() == reflect.Array {
			n := v.Type().Len()
			if n >= len(fields[f:]) {
				n = len(fields[f:])
			}
			switch v.Type().Elem().Kind() {
			case reflect.Int:
				for ii, s := range fields[f : f+n] {
					vv, err := strconv.ParseInt(s, 10, 0)
					if err != nil {
						return err
					}
					v.Index(ii).SetInt(vv)
				}
			case reflect.Float64:
				for ii, s := range fields[f : f+n] {
					vv, err := strconv.ParseFloat(s, 64)
					if err != nil {
						return err
					}
					v.Index(ii).SetFloat(vv)
				}
			}
			f += n
			continue
		}

		// slice of struct, prefixed by int, GxGSV, PUBXSVStatus
		if v.Kind() == reflect.Slice && v.Type().Elem().Kind() == reflect.Struct {
			// preceding value must be an int which is the length
			if i < 2 || rmsg.Field(i-1).Kind() != reflect.Int || rmsg.Field(i-1).Int() < 0 {
				return fmt.Errorf("%s %s can't parse slice if not preceded by valid length", rmsg.Type(), rmsg.Type().Field(i))
			}
			n := int(rmsg.Field(i - 1).Int())
			nf := v.Type().Elem().NumField()
			if nf == 0 {
				return fmt.Errorf("%s %s can't parse slice of empty structs", rmsg.Type(), rmsg.Type().Field(i))
			}
			if n*nf >= len(fields[f:]) {
				n = len(fields) / nf
			}

			v.Set(reflect.MakeSlice(v.Type().Elem(), n, n))
			for ii := 0; ii < n; ii++ {
				if err := decodeMsg(v.Index(ii).Addr().Interface(), fields[f:f+nf]); err != nil {
					return err
				}
				f += nf
			}

			continue
		}

		// scalar types
		s := fields[f]
		f++

		if s == "" {
			continue
		}

		switch v.Interface().(type) {
		case Header, string:
			v.SetString(s)

		case time.Duration:
			hhmmss, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return err
			}
			v.Set(reflect.ValueOf(parseHHMMSS(hhmmss)))

		case time.Time:
			ddmmyy, err := time.Parse("020106", s)
			if err != nil {
				return err
			}
			v.Set(reflect.ValueOf(ddmmyy))

		case float64:
			vv, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return err
			}
			switch rmsg.Type().Field(i).Tag.Get("nmea") {
			case "ddmm.mmmmm", "dddmm.mmmmm":
				vv = parseDDMM(vv)
			}
			v.SetFloat(vv)

		case PosMode:
			switch rmsg.Type().Field(i).Tag.Get("nmea") {
			case "quality":
				vv, err := strconv.ParseInt(s, 10, 0)
				if err != nil || vv < 0 || int(vv) >= len(quality2PosMode) {
					return err
				}
				v.SetInt(int64(quality2PosMode[vv]))
			default:
				v.SetInt(int64(s[0])) // ignore trailing characters
			}

		case Status, Wind, OpMode:
			v.SetInt(int64(s[0])) // ignore trailing characters

		case int, TxtType, NavMode:
			vv, err := strconv.ParseInt(s, 10, 0)
			if err != nil {
				return err
			}
			v.SetInt(vv)

		default:
			return fmt.Errorf("Don't know how to parse %s field %d: %s", rmsg.Type(), i, rmsg.Type().Field(i))
		}
	}
	return nil
}

// Return a new GxXYZ NMEA message struct corresponding to the given header
func mkMsg(h string) interface{} {
	if Header(h).TalkerID() == 0 {
		return nil
	}
	switch h[2:] {
	case "DTM":
		return &GxDTM{}
	case "GBQ", "GLQ", "GNQ", "GPQ":
		return &GxGxQ{}
	case "GBS":
		return &GxGBS{}
	case "GGA":
		return &GxGGA{}
	case "GGL":
		return &GxGGL{}
	case "GNS":
		return &GxGNS{}
	case "GRS":
		return &GxGRS{}
	case "GSA":
		return &GxGSA{}
	case "GST":
		return &GxGST{}
	case "GSV":
		return &GxGSV{}
	case "RMC":
		return &GxRMC{}
	case "TXT":
		return &GxTXT{}
	case "VLW":
		return &GxVLW{}
	case "VTG":
		return &GxVTG{}
	case "ZDA":
		return &GxZDA{}
	}
	return nil
}

func parseHHMMSS(hhmmss float64) time.Duration {
	hh := int(hhmmss / 10000.)
	hhmmss -= 10000. * float64(hh)
	mm := int(hhmmss) / 100.
	hhmmss -= 100. * float64(mm)
	return time.Duration(hh)*time.Hour + time.Duration(mm)*time.Minute + time.Duration(hhmmss*1E9)*time.Nanosecond
}

func parseDDMM(ddmm float64) float64 {
	dd := float64(int(ddmm / 100.))
	ddmm -= 100. * dd
	return dd + ddmm/60.
}
