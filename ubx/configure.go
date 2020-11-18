// +build ignore

package main

import (
	"log"
	"os"

	. "github.com/daedaleanai/ublox/ubx"
)

func main() {

	// disable all NMEA on all ports
	for _, msg := range []Message{
		CfgMsg2{MsgClass: 0xF0, MsgID: 0x0},         // GGA
		CfgMsg2{MsgClass: 0xF0, MsgID: 0x1},         // GLL
		CfgMsg2{MsgClass: 0xF0, MsgID: 0x2},         // GSA
		CfgMsg2{MsgClass: 0xF0, MsgID: 0x3},         // GSV
		CfgMsg2{MsgClass: 0xF0, MsgID: 0x4},         // RMC
		CfgMsg2{MsgClass: 0xF0, MsgID: 0x5},         // VTG
		CfgMsg2{MsgClass: 0xF0, MsgID: 0x6},         // GRS
		CfgMsg2{MsgClass: 0xF0, MsgID: 0x7},         // GST
		CfgMsg2{MsgClass: 0xF0, MsgID: 0x8},         // ZDA
		CfgMsg2{MsgClass: 0xF0, MsgID: 0x9},         // GBS
		CfgMsg2{MsgClass: 0xF0, MsgID: 0xD},         // GNS
		CfgMsg2{MsgClass: 0xF0, MsgID: 0xF},         // VLW
		CfgRate{MeasRate_ms: 62, NavRate_cycles: 1}, // 62ms (16Hz) measurement
		CfgMsg1{MsgClass: 0x1, MsgID: 0x7, Rate: 1}, // NAV-PVT
	} {
		b, err := Encode(msg)
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.Write(b)
	}
}
