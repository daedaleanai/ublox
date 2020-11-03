// +build ignore

package main

import (
	"os"

	. "github.com/daedaleanai/ublox/ubx"
)

func main() {

	// disable all NMEA on all ports
	Encode(os.Stdout, CfgMsg2{MsgClass: 0xF0, MsgID: 0x0}) // GGA
	Encode(os.Stdout, CfgMsg2{MsgClass: 0xF0, MsgID: 0x1}) // GLL
	Encode(os.Stdout, CfgMsg2{MsgClass: 0xF0, MsgID: 0x2}) // GSA
	Encode(os.Stdout, CfgMsg2{MsgClass: 0xF0, MsgID: 0x3}) // GSV
	Encode(os.Stdout, CfgMsg2{MsgClass: 0xF0, MsgID: 0x4}) // RMC
	Encode(os.Stdout, CfgMsg2{MsgClass: 0xF0, MsgID: 0x5}) // VTG
	Encode(os.Stdout, CfgMsg2{MsgClass: 0xF0, MsgID: 0x6}) // GRS
	Encode(os.Stdout, CfgMsg2{MsgClass: 0xF0, MsgID: 0x7}) // GST
	Encode(os.Stdout, CfgMsg2{MsgClass: 0xF0, MsgID: 0x8}) // ZDA
	Encode(os.Stdout, CfgMsg2{MsgClass: 0xF0, MsgID: 0x9}) // GBS
	Encode(os.Stdout, CfgMsg2{MsgClass: 0xF0, MsgID: 0xD}) // GNS
	Encode(os.Stdout, CfgMsg2{MsgClass: 0xF0, MsgID: 0xF}) // VLW

	Encode(os.Stdout, CfgRate{MeasRate: 62, NavRate: 1}) // 62ms (16Hz) measurement

	Encode(os.Stdout, CfgMsg3{MsgClass: 0x1, MsgID: 0x7, Rate: 1}) // NAV-PVT

	Encode(os.Stdout, CfgCfg1{SaveMask: CfgCfgMsgConf})

}
