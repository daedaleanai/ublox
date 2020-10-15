package nmea

import "time"

type Header string

func (h Header) TalkerID() TalkerID {
	if len(h) > 2 && h[0] == 'G' {
		return TalkerID(h[2])
	}
	return 0
}

// Datum reference
type GxDTM struct {
	Header
	Datum    string
	SubDatum string
	Lat_deg  float64
	North    Wind
	Lon_deg  float64
	East     Wind
	Alt_m    float64
	RefDatum string
}

// Poll a standard message (Talker ID Gx)
// GBQ, GLQ, GNQ, GPQ
type GxGxQ struct {
	Header
	MsgID string
}

// GNSS satellite fault detection
type GxGBS struct {
	Header
	TimeOfDay time.Duration // `nmea:"hhmmss.sssss"`
	ErrLat_m  float64
	ErrLon_m  float64
	ErrAlt_m  float64
	SVID      int
	Prob      float64
	Bias      float64
	StdDev    float64
	SystemID  int // V>4.10
	SignalID  int // V>4.10
}

// Global positioning system fix data
type GxGGA struct {
	Header
	TimeOfDay   time.Duration
	Lat_deg     float64 `nmea:"ddmm.mmmmm"`
	North       Wind
	Lon_deg     float64 `nmea:"dddmm.mmmmm"`
	East        Wind
	Quality     PosMode `nmea:"quality"`
	NumSV       int
	HDOP        float64
	Alt_m       float64
	AltUnit     string
	Sep_m       float64
	SepUnit     string
	DiffAge_s   float64
	DiffStation int
}

func (msg *GxGGA) LatLon() (lat_deg, lon_deg float64) {
	return msg.North.Sign(msg.Lat_deg), msg.East.Sign(msg.Lon_deg)
}

// Latitude and longitude, with time of position fix and status
type GxGGL struct {
	Header
	Lat_deg   float64 `nmea:"ddmm.mmmmm"`
	North     Wind
	Lon_deg   float64 `nmea:"dddmm.mmmmm"`
	East      Wind
	TimeOfDay time.Duration
	Status    Status
	PosMode   PosMode
}

func (msg *GxGGL) LatLon() (lat_deg, lon_deg float64) {
	return msg.North.Sign(msg.Lat_deg), msg.East.Sign(msg.Lon_deg)
}

// GNSS fix data
type GxGNS struct {
	Header
	TimeOfDay   time.Duration
	Lat_deg     float64 `nmea:"ddmm.mmmmm"`
	North       Wind
	Lon_deg     float64 `nmea:"dddmm.mmmmm"`
	East        Wind
	PosMode     PosMode
	NumSV       int
	HDOP        float64
	Alt_m       float64
	Sep_m       float64
	DiffAge_s   float64
	DiffStation int
	NavStatus   Status // V>4.10
}

func (msg *GxGNS) LatLon() (lat_deg, lon_deg float64) {
	return msg.North.Sign(msg.Lat_deg), msg.East.Sign(msg.Lon_deg)
}

// GNSS range residuals
type GxGRS struct {
	Header
	TimeOfDay time.Duration
	Mode      string      // 1 = computed after fix
	Residual  [12]float64 // matches last GSA sequence
	SystemID  int         // V>4.10
	SignalID  int         // V>4.10
}

// GNSS DOP and active satellites
type GxGSA struct {
	Header
	OpMode   OpMode
	NavMode  NavMode
	SVID     [12]int
	PDOP     float64
	HDOP     float64
	VDOP     float64
	SystemID int // V>4.10
}

// GNSS pseudorange error statistics
type GxGST struct {
	Header
	TimeOfDay  time.Duration
	RangeRMS_m float64
	StdMajor_m float64
	StdMinor_m float64
	Orient_deg float64
	StdLat_m   float64
	StdLon_m   float64
	StdAlt_m   float64
}

// GNSS satellites in view
type GxGSV struct {
	Header
	MsgCnt int
	MsgNum int
	NumSV  int
	SVInfo []struct {
		SVID    int
		Elv_deg float64
		Az_deg  float64
		CNO_db  float64
	}
	SignalID int // V>4.10  note: after the variable part!
}

// Recommended minimum data
type GxRMC struct {
	Header
	TimeOfDay time.Duration
	Valid     Status
	Lat_deg   float64 `nmea:"ddmm.mmmmm"`
	North     Wind
	Lon_deg   float64 `nmea:"dddmm.mmmmm"`
	East      Wind
	Speed_kts float64
	COG_deg   float64
	Date      time.Time // `nmea:"ddmmyy"`
	MV_deg    float64
	MVEast    Wind
	PosMode   PosMode
	NavStatus Status
}

func (msg *GxRMC) LatLon() (lat_deg, lon_deg float64) {
	return msg.North.Sign(msg.Lat_deg), msg.East.Sign(msg.Lon_deg)
}

// Text transmission
type GxTXT struct {
	Header
	MsgCnt  int
	MsgNum  int
	TxtType TxtType
	Text    string
}

// Dual ground/water distance
type GxVLW struct {
	Header
	TWD_nmi float64
	TWDUnit string
	WD_nm   float64
	WDUnit  string
	TGD_nmi float64
	TGDUnit string
	GD_nmi  float64
	GDUnit  string
}

// Course over ground and ground speed
type GxVTG struct {
	Header
	COGT_deg float64
	COGTUnit string
	COGM_deg float64
	COGMUnit string
	SOG_kn   float64
	SOGNUnit string
	SOG_kmh  float64
	SOGKUnit string
	PosMode  PosMode
}

// Time and date
type GxZDA struct {
	Header
	Time  time.Duration
	Day   int
	Month int
	Year  int
	LTZH  int
	LTZN  int
}
