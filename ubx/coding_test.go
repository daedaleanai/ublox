package ubx

import (
	"reflect"
	"testing"
)

func testType(t *testing.T, tc Message) {

	// hack.  CfgPrt1 is default, but the PortID subtype field zero is CfgPrt4
	if m, ok := tc.(*CfgPrt1); ok {
		m.PortID = 0xff
	}

	buf, err := Encode(tc)
	if err != nil {
		t.Errorf("Encoding %T: %v", tc, err)
		return
	}
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
