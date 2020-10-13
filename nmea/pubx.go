package nmea

import "time"

// PUBX further switch content on
type PUBXHeader struct {
	NMEAHeader
	PUBXType int
}

// Set protocols and baud rate
type PUBXConfig struct {
	PUBXHeader
	PortID      int
	InProto     uint32
	OutProto    uint32
	BaudRate    int
	AutoBauding int
}

// Lat/Long position data
type PUBXPosition struct {
	PUBXHeader
	TimeOfDay time.Duration
	Lat_min   float64 `nmea:"ddmm.mmmmm"`
	North     Wind
	Lon_min   float64 `nmea:"dddmm.mmmmm"`
	East      Wind
	AltRef_m  float64
	NavStat   string // NF, DR, G2, G3, D2, D3, RK, TT TODO make ENUM
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
	PUBXHeader
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
	PUBXHeader
	NumSV  int
	SVInfo []struct {
		SVID    int
		Az_deg  float64
		Elv_deg float64
		CNO_db  float64
		Lock_s  float64
	}
}

// Time of day and clock information
type PUBXTime struct {
	PUBXHeader
	TimeOfDay   time.Duration
	Date        time.Time
	UTCTow_s    float64
	UTCWeek     int
	Leap_s      float64 `nmea:"leap"` // parse trailing D
	ClkBias_ns  int
	ClkDrift_e9 float64 // ns/s
	TpGran_ns   int
}
