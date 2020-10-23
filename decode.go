// Package ublox provides methods to encode and decode u-Blox 8 / M8 NMEA and UBX messages
// as documented in
//		UBX-13003221 - R20 u-blox 8 / u-blox M8 Receiver description Including protocol specification
//      https://www.u-blox.com/sites/default/files/products/documents/u-blox8-M8_ReceiverDescrProtSpec_%28UBX-13003221%29.pdf

package ublox

import (
	"bufio"
	"bytes"
	"io"

	"github.com/daedaleanai/ublox/nmea"
	"github.com/daedaleanai/ublox/ubx"
)

// A Decoder scans an io stream into UBX (0xB5-0x62 separated) or NMEA ("$xxx,,,,*FF\r\n") frames.
// If you have an unmixed stream of NMEA-only data you can use nmea.Decode() on bufio.Scanner.Bytes() directly.
type Decoder struct {
	s *bufio.Scanner
}

// NewDecoder creates a new bufio Scanner with a splitfunc that can handle both UBX and NMEA frames.
func NewDecoder(r io.Reader) *Decoder {
	d := bufio.NewScanner(r)
	d.Split(splitFunc)
	return &Decoder{s: d}
}

// Assume we're either at the start of an NMEA sentence or at the start of a UBX message
// if not, skip to the first $ or UBX SOM.
func splitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) == 0 {
		return 0, nil, nil
	}

	switch data[0] {
	case '$':
		return bufio.ScanLines(data, atEOF)

	case 0xB5:
		if len(data) < 8 {
			if atEOF {
				return len(data), nil, io.ErrUnexpectedEOF
			}
			return 0, nil, nil
		}

		sz := 8 + int(data[4]) + int(data[5])*256
		if data[1] == 0x62 {
			if sz <= len(data) {
				return sz, data[:sz], nil
			}
			if sz <= bufio.MaxScanTokenSize {
				return 0, nil, nil
			}
		}
	}

	// resync to SOM or $
	data = data[1:]
	i1 := bytes.IndexByte(data, '$')
	if i1 < 0 {
		i1 = len(data)
	}

	i2 := bytes.IndexByte(data, 0xB5)
	if i2 < 0 {
		i2 = len(data)
	}
	if i1 > i2 {
		i1 = i2
	}
	return 1 + i1, nil, nil
}

// Decode reads on NMEA or UBX frame and calls nmea.Decode or ubx.Decode accordingly to parse the message.
func (d *Decoder) Decode() (msg interface{}, err error) {
	if !d.s.Scan() {
		if err = d.s.Err(); err == nil {
			err = io.EOF
		}
		return nil, err
	}

	switch d.s.Bytes()[0] {
	case '$':
		return nmea.Decode(d.s.Bytes())
	case 0xB5:
		return ubx.Decode(d.s.Bytes())
	}
	panic("impossible frame")
}
