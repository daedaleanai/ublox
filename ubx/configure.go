// +build ignore

package main

import (
	"os"

	. "github.com/daedaleanai/ublox/ubx"
)

func main() {

	Encode(os.Stdout, CfgRate{MeasRate: 50, NavRate: 1}) // 50ms (20Hz) measurement, output at 2Hz

	Encode(os.Stdout, CfgMsg3{MsgClass: 0xF0, MsgID: 0x0, Rate: 0}) // GGA
	Encode(os.Stdout, CfgMsg3{MsgClass: 0xF0, MsgID: 0x1, Rate: 0}) // GLL
	Encode(os.Stdout, CfgMsg3{MsgClass: 0xF0, MsgID: 0x2, Rate: 0}) // GSA
	Encode(os.Stdout, CfgMsg3{MsgClass: 0xF0, MsgID: 0x3, Rate: 0}) // GSV
	Encode(os.Stdout, CfgMsg3{MsgClass: 0xF0, MsgID: 0x4, Rate: 0}) // RMC
	Encode(os.Stdout, CfgMsg3{MsgClass: 0xF0, MsgID: 0x5, Rate: 0}) // VTG
	Encode(os.Stdout, CfgMsg3{MsgClass: 0xF0, MsgID: 0x6, Rate: 0}) // GRS
	Encode(os.Stdout, CfgMsg3{MsgClass: 0xF0, MsgID: 0x7, Rate: 0}) // GST
	Encode(os.Stdout, CfgMsg3{MsgClass: 0xF0, MsgID: 0x8, Rate: 0}) // ZDA
	Encode(os.Stdout, CfgMsg3{MsgClass: 0xF0, MsgID: 0x9, Rate: 0}) // GBS

	Encode(os.Stdout, CfgMsg3{MsgClass: 0x1, MsgID: 0x7, Rate: 1}) // GBS

}
