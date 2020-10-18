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
	if frame[end+1] != hexChar[x>>4] || frame[end+2] != hexChar[x&0xf] { // also lowercase?
		return nil, fmt.Errorf("Expected %02X found %s", x, frame[end+1:end+3])
	}
	fields := strings.Split(string(frame[start+1:end]), ",")
	msg = mkMsg(fields[0])
	if msg == nil && fields[0] == "PUBX" && len(fields) > 2 {
		t, err := strconv.ParseInt(fields[1], 10, 0)
		if err == nil {
			msg = mkPUBX(PUBXType(t))
		}
	}
	if msg != nil {
		err = decodeMsg(msg, fields)
		return msg, err
	}

	return fields, nil
}

// https://play.golang.org/p/_GK6kdk1SdV
func decodeMsg(msg interface{}, fields []string) error {

	// override reflection based one for GxGSV and PUBXSVStatus which have variable length irregularities
	if dmsg, ok := msg.(interface{ Decode([]string) error }); ok {
		return dmsg.Decode(fields)
	}

	if reflect.ValueOf(msg).Kind() != reflect.Ptr {
		return fmt.Errorf("message must be a pointer to struct")
	}
	rmsg := reflect.ValueOf(msg).Elem()
	if rmsg.Kind() != reflect.Struct {
		return fmt.Errorf("message must be a pointer to struct")
	}

	for i, f := 0, 0; i < rmsg.NumField() && f < len(fields); i++ {

		v := rmsg.Field(i)

		// array of int, float, as in GxGRS and GxGSA
		if v.Kind() == reflect.Array {
			n := v.Type().Len()
			if n >= len(fields[f:]) {
				n = len(fields[f:])
			}
			switch v.Type().Elem().Kind() {
			case reflect.Int:
				for ii, s := range fields[f : f+n] {
					if s == "" {
						continue
					}
					vv, err := strconv.ParseInt(s, 10, 0)
					if err != nil {
						return err
					}
					v.Index(ii).SetInt(vv)
				}
			case reflect.Float64:
				for ii, s := range fields[f : f+n] {
					if s == "" {
						continue
					}
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
			return fmt.Errorf("Don't know how to parse %v field %d: %v", rmsg.Type(), i, rmsg.Type().Field(i))
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
	case "GLL":
		return &GxGLL{}
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

func mkPUBX(t PUBXType) interface{} {
	switch t {
	case CONFIG:
		return &PUBXConfig{}
	case POSITION:
		return &PUBXPosition{}
	case RATE:
		return &PUBXRate{}
	case SVSTATUS:
		return &PUBXSVStatus{}
	case TIME:
		return &PUBXTime{}
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
