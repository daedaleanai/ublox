package nmea

import "time"

// Proprietary u-Blox NMEA sentences have NMEA header 'PUBX' and then switch on the second field
type PUBXType int

const (
	CONFIG   PUBXType = 41
	POSITION PUBXType = 0
	RATE     PUBXType = 40
	SVSTATUS PUBXType = 3
	TIME     PUBXType = 4
)

type SatStat int

const (
	NotUsed SatStat = '-'
	Used    SatStat = 'U'
	Known   SatStat = 'e'
)

type NavStat int

const (
	NF NavStat = iota + 1
	DR
	G2
	G3
	D2
	D3
	RK
	TT
)

// Set protocols and baud rate
type PUBXConfig struct {
	Header
	PUBXType
	PortID      int
	InProto     uint32
	OutProto    uint32
	BaudRate    int
	AutoBauding int
}

// Lat/Long position data
type PUBXPosition struct {
	Header
	PUBXType
	TimeOfDay time.Duration
	Lat_min   float64 `nmea:"ddmm.mmmmm"`
	North     Wind
	Lon_min   float64 `nmea:"dddmm.mmmmm"`
	East      Wind
	AltRef_m  float64
	NavStat   NavStat
	HAcc_m    float64
	VAcc_m    float64
	SOG_km_h  float64
	COG_deg   float64
	VVel_m_s  float64
	DiffAge_s float64
	HDOP      float64
	VDOP      float64
	TDOP      float64
	NumSV     int
	Reserved  int
	DR        int
}

// Set NMEA message output rate
type PUBXRate struct {
	Header
	PUBXType
	MsgID    string
	Rddc     int
	Rus1     int
	Rus2     int
	Rusb     int
	Rspi     int
	Reserved int
}

// Satellite status
type PUBXSVStatus struct {
	Header
	PUBXType
	NumSV  int
	SVInfo []struct { // message parser magically uses preceding int as length
		SVID    int
		Az_deg  float64
		Elv_deg float64
		CNO_db  float64
		Lock_s  float64
	}
}

// Time of day and clock information
type PUBXTime struct {
	Header
	PUBXType
	TimeOfDay   time.Duration
	Date        time.Time
	UTCTow_s    float64
	UTCWeek     int
	Leap_s      float64 `nmea:"leap"` // parse trailing D
	ClkBias_ns  int
	ClkDrift_e9 float64 // ns/s
	TpGran_ns   int
}
