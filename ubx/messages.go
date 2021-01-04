package ubx

type Message interface {
	ClassID() uint16
}

type RawMessage struct {
	CID  uint16
	Data []byte
}

func (msg *RawMessage) ClassID() uint16 { return msg.CID }

// 32.8.1 UBX-ACK-ACK (0x05 0x01)

// 32.8.1.1 Message acknowledged
// Description          Message acknowledged
// Type                 Output
// Output upon processing of an input message. A UBX-ACK-ACK is sent as soon as possible but at least within one second.
type AckAck struct {
	ClsID byte // Class ID of the Acknowledged Message
	MsgID byte // Message ID of the Acknowledged Message
}

func (AckAck) ClassID() uint16 { return 0x0105 }

// 32.8.2 UBX-ACK-NAK (0x05 0x00)

// 32.8.2.1 Message not acknowledged
// Description          Message not acknowledged
// Type                 Output
// Output upon processing of an input message. A UBX-ACK-NAK is sent as soon as possible but at least within one second.
type AckNak struct {
	ClsID byte // Class ID of the Acknowledged Message
	MsgID byte // Message ID of the Acknowledged Message
}

func (AckNak) ClassID() uint16 { return 0x0005 }

// 32.10.3 UBX-CFG-CFG (0x06 0x09)

// 32.10.3.1 Clear, save and load configurations
// Description          Clear, save and load configurations
// Type                 Command
// See Receiver configuration for a detailed description on how receiver
// configuration should be used. The three masks are made up of individual bits,
// each bit indicating the sub-section of all configurations on which the
// corresponding action shall be carried out. The reserved bits in the masks must
// be set to '0'. For detailed information refer to the Organization of the
// configuration sections. Note that commands can be combined. The sequence of
// execution is clear, save, load.
type CfgCfg1 struct {
	ClearMask CfgCfgClearMask // Mask with configuration sub-sections to clear (i.e. load default configurations to permanent configurations in non-volatile memory) (see graphic below)
	SaveMask  CfgCfgClearMask // Mask with configuration sub-sections to save (i.e. save current configurations to non-volatile memory), see ID description of clearMask
	LoadMask  CfgCfgClearMask // Mask with configuration sub-sections to load (i.e. load permanent configurations from non-volatile memory to currentconfigurations), see ID description of clearMask
}

func (CfgCfg1) ClassID() uint16 { return 0x0906 }

type CfgCfg2 struct {
	ClearMask  CfgCfgClearMask  // Mask with configuration sub-sections to clear (i.e. load default configurations to permanent configurations in non-volatile memory) (see graphic below)
	SaveMask   CfgCfgClearMask  // Mask with configuration sub-sections to save (i.e. save current configurations to non-volatile memory), see ID description of clearMask
	LoadMask   CfgCfgClearMask  // Mask with configuration sub-sections to load (i.e. load permanent configurations from non-volatile memory to currentconfigurations), see ID description of clearMask
	DeviceMask CfgCfgDeviceMask // Mask which selects the memory devices for this command. (see graphic below)
}

func (CfgCfg2) ClassID() uint16 { return 0x0906 }

//go:generate stringer -output=strings_cfgcfg.go  -trimprefix CfgCfg -type=CfgCfgClearMask,CfgCfgDeviceMask

//  UBX-CFG-CFG Bitfield clearMask
type CfgCfgClearMask uint32

const (
	CfgCfgIoPort   CfgCfgClearMask = 0x1    // Communications port settings. Modifying this sub-section results in an IO system reset. Because of this undefined data may be output for a short period of time after receiving the message.
	CfgCfgMsgConf  CfgCfgClearMask = 0x2    // Message configuration
	CfgCfgInfMsg   CfgCfgClearMask = 0x4    // INF message configuration
	CfgCfgNavConf  CfgCfgClearMask = 0x8    // Navigation configuration
	CfgCfgRxmConf  CfgCfgClearMask = 0x10   // Receiver Manager configuration
	CfgCfgSenConf  CfgCfgClearMask = 0x100  // Sensor interface configuration (not supported in protocol versions less than 19)
	CfgCfgRinvConf CfgCfgClearMask = 0x200  // Remote inventory configuration
	CfgCfgAntConf  CfgCfgClearMask = 0x400  // Antenna configuration
	CfgCfgLogConf  CfgCfgClearMask = 0x800  // Logging configuration
	CfgCfgFtsConf  CfgCfgClearMask = 0x1000 // FTS configuration. Only applicable to the FTS product variant.
)

// 	UBX-CFG-CFG  Bitfield deviceMask
type CfgCfgDeviceMask byte

const (
	CfgCfgDevBBR      CfgCfgDeviceMask = 0x1  // Battery backed RAM
	CfgCfgDevFlash    CfgCfgDeviceMask = 0x2  // Flash
	CfgCfgDevEEPROM   CfgCfgDeviceMask = 0x4  // EEPROM
	CfgCfgDevSpiFlash CfgCfgDeviceMask = 0x10 // SPI Flash
)

// 32.10.10 UBX-CFG-HNR (0x06 0x5C)

// 32.10.10.1 High navigation rate settings
// Description          High navigation rate settings
// Type                 Get/set
// The u-blox receivers support high rates of navigation update up to 30 Hz. The
// navigation solution output UBX-NAV-HNR will not be aligned to the top of a
// second.
// The update rate has a direct influence on the power consumption. The more
// fixes that are required, the more CPU power and communication resources are
type CfgHnr struct {
	HighNavRate byte    // Hz        Rate of navigation solution output
	Reserved1   [3]byte // -         Reserved
}

func (CfgHnr) ClassID() uint16 { return 0x5C06 }

// 32.10.14 UBX-CFG-MSG (0x06 0x01)

// 32.10.14.1 Poll a message configuration
// Description          Poll a message configuration
// Type                 Poll Request
type CfgMsg1 struct {
	MsgClass byte // Message class
	MsgID    byte // Message identifier
}

func (CfgMsg1) ClassID() uint16 { return 0x0106 }

// 32.10.14.2 Set message rate(s)
// Description          Set message rate(s)
// Type                 Get/set
// Get/set message rate configuration (s) to/from the receiver.
// See also section How to change between protocols.
// Send rate is relative to the event a message is registered on. For example, if the
// rate of a navigation message is set to 2, the message is sent every second
// navigation solution. For configuring NMEA messages, the section NMEA
// Messages Overview describes class and identifier numbers used.
type CfgMsg2 struct {
	MsgClass byte    // Message class
	MsgID    byte    // Message identifier
	Rate     [6]byte // Send rate on I/O port (6 ports)
}

func (CfgMsg2) ClassID() uint16 { return 0x0106 }

// 32.10.14.3 Set message rate
// Description          Set message rate
// Type                 Get/set
// Set message rate configuration for the current port.
// See also section How to change between protocols.
type CfgMsg3 struct {
	MsgClass byte // Message class
	MsgID    byte // Message identifier
	Rate     byte // Send rate on current portSend rate on I/O port (6 ports)
}

func (CfgMsg3) ClassID() uint16 { return 0x0106 }

// 32.10.23 UBX-CFG-RATE (0x06 0x08)

// 32.10.23.1 Navigation/measurement rate settings
// Description          Navigation/measurement rate settings
// Type                 Get/set
// This message allows the user to alter the rate at which navigation solutions (and
// the measurements that they depend on) are generated by the receiver. The
// calculation of the navigation solution will always be aligned to the top of a
// second zero (first second of the week) of the configured reference time system.
// (Navigation period is an integer multiple of the measurement period in protocol versions greater than 17).
// Each measurement triggers the measurements generation and, if available, raw data output.
// The navRate value defines that every nth measurement triggers a navigation epoch.
// The update rate has a direct influence on the power consumption. The more
// fixes that are required, the more CPU power and communication resources are required.
// For most applications a 1 Hz update rate would be sufficient.
// When using power save mode, measurement and navigation rate can differ from the values configured here.
// See Measurement and navigation rate with power save mode for details.
type CfgRate struct {
	MeasRate uint16         //     ms        The elapsed time between GNSS measurements, which defines the rate
	NavRate  uint16         //     cycles    The ratio between the number of measurements and the number of navigation solutions
	TimeRef  CfgRateTimeRef //     -         The time system to which measurements are aligned
}

func (CfgRate) ClassID() uint16 { return 0x0806 }

//go:generate stringer -output=strings_cfgrate.go  -trimprefix CfgRate -type=CfgRateTimeRef

type CfgRateTimeRef uint16

const (
	CfgRateUTC     CfgRateTimeRef = 0
	CfgRateGPS     CfgRateTimeRef = 1
	CfgRateGLONASS CfgRateTimeRef = 2
	CfgRateBeiDou  CfgRateTimeRef = 3
	CfgRateGalileo CfgRateTimeRef = 4
)

// 32.13 UBX-INF (0x04)
// Information Messages: i.e. Printf-Style Messages, with IDs such as Error, Warning, Notice.
// Messages in the INF class are used to output strings in a printf style from the firmware or
// application code. All INF messages have an associated type to indicate the kind of message.
// 32.13.1 UBX-INF-DEBUG (0x04 0x04)
// 32.13.2 UBX-INF-ERROR (0x04 0x00)
// 32.13.3 UBX-INF-NOTICE (0x04 0x02)
// 32.13.4 UBX-INF-TEST (0x04 0x03)
// 32.13.5 UBX-INF-WARNING (0x04 0x01)

// Description          ASCII output with debug contents
// Type                 Output
// This message has a variable length payload, representing an ASCII string.
// N*1            CH        -          str                  -         ASCII Character
type InfDebug string
type InfError string
type InfNotice string
type InfTest string
type InfWarning string

func (InfDebug) ClassID() uint16   { return 0x0404 }
func (InfError) ClassID() uint16   { return 0x0004 }
func (InfNotice) ClassID() uint16  { return 0x0204 }
func (InfTest) ClassID() uint16    { return 0x0304 }
func (InfWarning) ClassID() uint16 { return 0x0104 }

func (msg InfDebug) encode() string   { return string(msg) }
func (msg InfError) encode() string   { return string(msg) }
func (msg InfNotice) encode() string  { return string(msg) }
func (msg InfTest) encode() string    { return string(msg) }
func (msg InfWarning) encode() string { return string(msg) }

// 32.17.15 UBX-NAV-PVT (0x01 0x07)

// 32.17.15.1 Navigation position velocity time solution
// Description          Navigation position velocity time solution
// Type                 Periodic/Polled
// This message combines position, velocity and time solution, including accuracy figures.
// Note that during a leap second there may be more or less than 60 seconds in a minute.
// See the description of leap seconds for details.
type NavPvt struct {
	ITOW_ms       uint32        //   -         GPS time of week of the navigation epoch. See the description of iTOW for details.
	Year_y        uint16        //   -         Year (UTC)
	Month_month   byte          //   -         Month, range 1..12 (UTC)
	Day_d         byte          //   -         Day of month, range 1..31 (UTC)
	Hour_h        byte          //   -         Hour of day, range 0..23 (UTC)
	Min_min       byte          //   -         Minute of hour, range 0..59 (UTC)
	Sec_s         byte          //   -         Seconds of minute, range 0..60 (UTC)
	Valid         NavPVTValid   //   -         Validity flags (see graphic below)
	TAcc_ns       uint32        //   -         Time accuracy estimate (UTC)
	Nano_ns       int32         //   -         Fraction of second, range -1e9 .. 1e9 (UTC)
	FixType       NavPVTFixType //   -         GNSSfix Type
	Flags         NavPVTFlags   //   -         Fix status flags (see graphic below)
	Flags2        NavPVTFlags2  //   -         Additional flags (see graphic below)
	NumSV         byte          //   -         Number of satellites used in Nav Solution
	Lon_dege7     int32         //   1e-7      Longitude
	Lat_dege7     int32         //   1e-7      Latitude
	Height_mm     int32         //   -         Height above ellipsoid
	HMSL_mm       int32         //   -         Height above mean sea level
	HAcc_mm       uint32        //   -         Horizontal accuracy estimate
	VAcc_mm       uint32        //   -         Vertical accuracy estimate
	VelN_mm_s     int32         //   -         NED north velocity
	VelE_mm_s     int32         //   -         NED east velocity
	VelD_mm_s     int32         //   -         NED down velocity
	GSpeed_mm_s   int32         //   -         Ground Speed (2-D)
	HeadMot_dege5 int32         //   1e-5      Heading of motion (2-D)
	SAcc_mm_s     uint32        //   -         Speed accuracy estimate
	HeadAcc_dege5 uint32        //   1e-5      Heading accuracy estimate (both motion and vehicle)
	PDOPe2        uint16        //   0.01      Position DOP
	Flags3        NavPVTFlags3  //   -         Additional flags (see graphic below)
	Reserved1     [5]byte       //   -         Reserved
	HeadVeh_dege5 int32         //   1e-5      Heading of vehicle (2-D), this is only valid when headVehValid is set, otherwise the                                                                    output is set to the heading of motion
	MagDec_dege2  int16         //   1e-2      Magnetic declination. Only supported in ADR 4.10 and later.
	MagAcc_deg2e  uint16        //   1e-2      Magnetic declination accuracy. Only supported in ADR 4.10 and later.
}

func (NavPvt) ClassID() uint16 { return 0x0701 }

//go:generate stringer -output=strings_navpvt.go  -trimprefix NavPVT -type=NavPVTFixType,NavPVTValid,NavPVTFlags,NavPVTFlags2,NavPVTFlags3

type NavPVTFixType byte

const (
	NavPVTNoFix NavPVTFixType = iota
	NavPVTDeadReckoning
	NavPVTFix2D
	NavPVTFix3D
	NavPVTGNSS
	NavPVTTimeOnly
)

type NavPVTValid byte

const (
	NavPVTValidDate     NavPVTValid = (1 << iota) // valid UTC Date (see Time Validity section for details)
	NavPVTValidTime                               // valid UTC time of day (see Time Validity section for details)
	NavPVTFullyResolved                           // UTC time of day has been fully resolved (no seconds uncertainty). Cannot be used to check if time is completely solved.
	NavPVTValidMag                                // valid magnetic declination
)

type NavPVTFlags byte

const (
	NavPVTGnssFixOK     NavPVTFlags = 1 << 0 // valid fix (i.e within DOP & accuracy masks)
	NavPVTDiffSoln      NavPVTFlags = 1 << 1 // differential corrections were applied
	NavPVTHeadVehValid  NavPVTFlags = 1 << 5 // heading of vehicle is valid, only set if the receiver is in sensor fusion mode
	NavPVTCarrSolnFloat NavPVTFlags = 1 << 6 // carrier phase range solution with floating ambiguities
	NavPVTCarrSolnFixed NavPVTFlags = 1 << 7 // carrier phase range solution with fixed ambiguities
)

type NavPVTFlags2 byte

const (
	NavPVTConfirmedAvai NavPVTFlags2 = 1 << 5 // information about UTC Date and Time of Day validity confirmation is available (see Time Validity section for details)
	NavPVTConfirmedDate NavPVTFlags2 = 1 << 6 // UTC Date validity could be confirmed (see Time Validity section for details)
	NavPVTConfirmedTime NavPVTFlags2 = 1 << 7 // UTC Time of Day could be confirmed (see Time Validity section for details)
)

type NavPVTFlags3 byte

const (
	NavPVTInvalidLlh NavPVTFlags3 = (1 << iota) // 1 = Invalid lon, lat, height and hMSL
)
