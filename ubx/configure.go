// +build ignore

package main

import (
	"os"

	"github.com/daedaleanai/ublox/ubx"	
)


func main (

	ubx.Encode(os.Stdout, ubx.CfgMsg1{MsgClass: NavPvt{}.Descriptor().Class(), MsgID: NavPvt{}.Descriptor().ID() })

)