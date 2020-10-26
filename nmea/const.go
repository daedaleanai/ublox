package nmea

//go:generate stringer -output=strings.go -type=TalkerID,Status,PosMode,OpMode,NavMode,Wind,TxtType,PUBXType,SatStat,NavStat

const VERSION = 410 // NMEA Protocol version 4.10

type TalkerID int

const (
	GPS     TalkerID = 'P' // GPS, SBAS, QZSS
	GLONASS TalkerID = 'L' // GLONASS
	GALILEO TalkerID = 'A' // Galileo
	GBEIDOU TalkerID = 'B' // BeiDou (NMEA 4.11)
	GANY    TalkerID = 'N' // Any combination of GNSS
)

// Possible values for status: V = Data invalid, A = Data valid
// GLL, RMC status
type Status int

const (
	Valid   Status = 'A'
	Invalid Status = 'V'
)

// Possible values for posMode: N = No fix, A = Autonomous GNSS fix, D = Differential GNSS fix, R = RTK fixed  F = RTK float, E = Estimated/Dead reckoning fix
// Missing values will be mapped to NoFix
// GLL, VTG RMC, GNS posMode
type PosMode int

const (
	NoFix         PosMode = 'N'
	Autonomous    PosMode = 'A'
	Differential  PosMode = 'D'
	RTKFixed      PosMode = 'R'
	RTKFloat      PosMode = 'F'
	DeadReckoning PosMode = 'E'
)

// Possible values for quality: 0 = No fix, 1 = Autonomous GNSS fix, 2 = Differential GNSS fix, 4 = RTK fixed, 5 = RTK float, 6 = Estimated/Dead reckoning fix
// GGA quality can be mapped to PosMode
var quality2PosMode = "NAD.RFE"

type OpMode int

const (
	Manual    OpMode = 'M'
	Automatic OpMode = 'A'
)

// Possible values for navMode: 1 = No fix, 2 = 2D fix, 3 = 3D fix  // Pre V4.10
// GSA navMode
type NavMode int

const (
	FixNo NavMode = 1
	Fix2D NavMode = 2
	Fix3D NavMode = 3
)

// Wind encodes the compass direction of latitudes and longitudes
type Wind int

const (
	North Wind = 'N'
	South Wind = 'S'
	East  Wind = 'E'
	West  Wind = 'W'
)

// Sign returns -latlon for South or West, and latlon for all other values
func (w Wind) Sign(latlon float64) float64 {
	if w == South || w == West {
		return -latlon
	}
	return latlon
}

type TxtType int

const (
	Error   TxtType = 0
	Warning TxtType = 1
	Notice  TxtType = 2
	User    TxtType = 7
)
