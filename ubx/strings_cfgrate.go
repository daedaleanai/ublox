// Code generated by "stringer -output=strings_cfgrate.go -trimprefix CfgRate -type=CfgRateTimeRef"; DO NOT EDIT.

package ubx

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CfgRateUTC-0]
	_ = x[CfgRateGPS-1]
	_ = x[CfgRateGLONASS-2]
	_ = x[CfgRateBeiDou-3]
	_ = x[CfgRateGalileo-4]
}

const _CfgRateTimeRef_name = "UTCGPSGLONASSBeiDouGalileo"

var _CfgRateTimeRef_index = [...]uint8{0, 3, 6, 13, 19, 26}

func (i CfgRateTimeRef) String() string {
	if i >= CfgRateTimeRef(len(_CfgRateTimeRef_index)-1) {
		return "CfgRateTimeRef(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CfgRateTimeRef_name[_CfgRateTimeRef_index[i]:_CfgRateTimeRef_index[i+1]]
}
