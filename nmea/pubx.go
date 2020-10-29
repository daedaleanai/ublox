package nmea

import (
	"strconv"
	"time"
)

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

var parseNavStat = map[string]NavStat{"NF": NF, "DR": DR, "G2": G2, "G3": G3, "D2": D2, "D3": D3, "RK": RK, "TT": TT}

// Set protocols and baud rate
type PUBXConfig struct {
	Header
	PUBXType    PUBXType
	PortID      int
	InProto     uint32
	OutProto    uint32
	BaudRate    int
	AutoBauding int
}

// Lat/Long position data
type PUBXPosition struct {
	Header
	PUBXType  PUBXType
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
	PUBXType PUBXType
	MsgID    string
	Rddc     int
	Rus1     int
	Rus2     int
	Rusb     int
	Rspi     int
	Reserved int
}

type PUBXSVInfo struct { // message parser magically uses preceding int as length
	SVID    int
	Status  SatStat
	Az_deg  float64
	Elv_deg float64
	CNO_db  float64
	Lock_s  float64
}

// Satellite status
type PUBXSVStatus struct {
	Header
	PUBXType PUBXType
	NumSV    int
	SVInfo   []PUBXSVInfo
}

// var lenght messages must implement custom decode()
func (msg *PUBXSVStatus) decode(fields []string) (err error) {
	if len(fields) > 0 {
		msg.Header = Header(fields[0])
	}
	if len(fields) > 1 {
		v := 0
		v, err = strconv.Atoi(fields[1])
		msg.PUBXType = PUBXType(v)
	}
	if err == nil && len(fields) > 2 {
		msg.NumSV, err = strconv.Atoi(fields[2])
	}
	if err != nil {
		return err
	}
	msg.SVInfo = make([]PUBXSVInfo, msg.NumSV)
	for i := 0; i < msg.NumSV; i++ {
		if len(fields) < 3+6*i {
			break
		}
		if err = decodeMsg(&msg.SVInfo[i], fields[3+6*i:8+6*i]); err != nil {
			return err
		}
	}
	return nil
}

// Time of day and clock information
type PUBXTime struct {
	Header
	PUBXType    PUBXType
	TimeOfDay   time.Duration
	Date        time.Time
	UTCTow_s    float64
	UTCWeek     int
	Leap_s      float64 `nmea:"leap"` // skip trailing D, todo: store it
	ClkBias_ns  int
	ClkDrift_e9 float64 // ns/s
	TpGran_ns   int
}
