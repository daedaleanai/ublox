package ubx

//go:generate stringer -output=strings_navpvt.go  -trimprefix NavPvt -type=NavPvtFixType

// values of the NavPvt FixType byte
// TODO splice into the generated code
type NavPvtFixType byte

const (
	NavPvtNoFix NavPvtFixType = iota
	NavPvtDeadReckoning
	NavPvtFix2D
	NavPvtFix3D
	NavPvtGNSS
	NavPvtTimeOnly
)
