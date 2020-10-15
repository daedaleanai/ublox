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
		return nil, errInvalidChksum
	}
	fields := strings.Split(string(frame[start+1:end]), ",")
	msg := mkMsg(fields[0])
	if msg == nil && fields[0] == "PUBX" && len(fields) > 2 {
		t, err := strconv.Atoi(fields[1])
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

	t := rmsg.Type()
	l := rmsg.NumField()
	if l >= len(fields) {
		l = len(fields)
	}
	for i := 0; i < l; i++ {
		if fields[i] == "" {
			continue
		}

		v := rmsg.Field(i)
		// array of int, float
		if v.Kind() == reflect.Array && v.Type().Elem().Kind() == reflect.Int


		// slice of struct, prefixed by int

		// scalar types
		switch v.Interface().(type) {
		case Header:
			v.SetString(fields[i])
		case time.Duration:
			hhmmss, err := strconv.ParseFloat(fields[i], 64)
			if err != nil {
				return fmt.Errorf("can't parse %q as hhmmss.ssss: %v", fields[i], err)
			}
			v.SetInt(parseHHMMSS(hhmmss))
		case time.Time:
			t, err := time.Parse("020106", fields[i])
			if err != nil {
				return fmt.Errorf("can't parse %q as hhmmss.ssss: %v", fields[i], err)
			}
			v.Set(reflect.ValueOf(t))
		case Status:
			v.SetInt(fields[i][0]) // ignore trailing characters
		case Wind:
			v.SetInt(fields[i][0]) // ignore trailing characters
		case OpMode:
			v.SetInt(fields[i][0]) // ignore trailing characters
		case string:
			v.SetString(fields[i])
		case int:
			vv, err := strconv.Atoi(fields[i])
			if err != nil {
				return fmt.Errorf("can't parse %q as type %s: %v", fields[i], v.Type(), err)
			}
			v.SetInt(vv)
		case float64:
			vv, err := strconv.ParseFloat(fields[i], 64)
			if err != nil {
				return fmt.Errorf("can't parse %q as type %s: %v", fields[i], v.Type(), err)
			}
			switch rmsg.Type().Field(i).Tag.Get("nmea") {
			case "ddmm.mmmmm", "dddmm.mmmmm":
				vv = parseDDMM(vv)
			}
			v.SetInt(vv)
		case PosMode:
			switch rmsg.Type().Field(i).Tag.Get("nmea") {
			case "quality":
				vv, err := strconv.Atoi(fields[i])
				if err != nil || vv < 0 || vv >= len(quality2PosMode) {
					return fmt.Errorf("can't parse %q as type %s: %v", fields[i], v.Type(), err)
				}
				v.SetInt(int(quality2PosMode[vv]))
			default:
				v.SetInt(fields[i][0]) // ignore trailing characters
			}
		case TxtType:
			vv, err := strconv.Atoi(fields[i])
			if err != nil {
				return fmt.Errorf("can't parse %q as type %s: %v", fields[i], v.Type(), err)
			}
			v.SetInt(vv)
		case NavMode:
			vv, err := strconv.Atoi(fields[i])
			if err != nil {
				return fmt.Errorf("can't parse %q as type %s: %v", fields[i], v.Type(), err)
			}
			v.SetInt(vv)
		}
	}
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
