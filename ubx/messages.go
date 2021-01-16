// Generated Code -- DO NOT EDIT.
//go:generate go run msggen.go messages.tmpl messages.xml messages.go

package ubx

import (
	"fmt"
	"strings"
)

// Message ubx-ack-ack

// AckAck (Output) Message acknowledged
// Class/Id 0x05 0x01 (2 bytes)
// Output upon processing of an input message. A UBX-ACK-ACK is sent as soon as possible but at least within one second.
type AckAck struct {
	ClsID byte // Class ID of the Acknowledged Message
	MsgID byte // Message ID of the Acknowledged Message
}

func (AckAck) classID() uint16 { return 0x0105 }

// Message ubx-ack-nak

// AckNak (Output) Message not acknowledged
// Class/Id 0x05 0x00 (2 bytes)
// Output upon processing of an input message. A UBX-ACK-NAK is sent as soon as possible but at least within one second.
type AckNak struct {
	ClsID byte // Class ID of the Not-Acknowledged Message
	MsgID byte // Message ID of the Not-Acknowledged Message
}

func (AckNak) classID() uint16 { return 0x0005 }

// Message ubx-aid-alm (3 versions)

// AidAlm (Poll Request) Poll GPS aiding almanac data
// Class/Id 0x0b 0x30 (0 bytes)
// All UBX-AID messages are deprecated; use UBX-MGA messages instead Poll GPS aiding data (Almanac) for all 32 SVs by sending this message to the receiver without any payload. The receiver will return 32 messages of type AID- ALM as defined below.
type AidAlm struct {
}

func (AidAlm) classID() uint16 { return 0x300b }

// AidAlm1 (Poll Request) Poll GPS aiding almanac data for a SV
// Class/Id 0x0b 0x30 (1 bytes)
// All UBX-AID messages are deprecated; use UBX-MGA messages instead Poll GPS aiding data (Almanac) for an SV by sending this message to the receiver. The receiver will return one message of type AID-ALM as defined below.
type AidAlm1 struct {
	Svid byte // SV ID for which the receiver shall return its Almanac Data (Valid Range: 1 .. 32 or 51, 56, 63).
}

func (AidAlm1) classID() uint16 { return 0x300b }

// AidAlm2 (Input/Output) GPS aiding almanac input/output message
// Class/Id 0x0b 0x30 (8 or 40 bytes)
// All UBX-AID messages are deprecated; use UBX-MGA messages instead  If the WEEK Value is 0, DWRD0 to DWRD7 are not sent as the Almanac is not  available for the given SV. This may happen even if NAV-SVINFO and RXM-  SVSI are indicating almanac availability as the internal data may not represent  the content of an original broadcast almanac (or only parts thereof).  DWORD0 to DWORD7 contain the 8 words following the Hand-Over Word (  HOW ) from the GPS navigation message, either pages 1 to 24 of sub-frame 5  or pages 2 to 10 of subframe 4. See IS-GPS-200 for a full description of the  contents of the Almanac pages.  In DWORD0 to DWORD7, the parity bits have been removed, and the 24 bits of  data are located in Bits 0 to 23. Bits 24 to 31 shall be ignored.  Example: Parameter e (Eccentricity) from Almanac Subframe 4/5, Word 3, Bits  69-84 within the subframe can be found in DWRD0, Bits 15-0 whereas Bit 0 is  the LSB.
type AidAlm2 struct {
	Svid uint32 // SV ID for which this Almanac Data is (Valid Range: 1 .. 32 or 51, 56, 63).
	Week uint32 // Issue Date of Almanac (GPS week number)
	// Optional
	Dwrd [8]uint32 // Almanac Words
}

func (AidAlm2) classID() uint16 { return 0x300b }

// Message ubx-aid-aop (3 versions)

// AidAop (Poll Request) Poll AssistNow Autonomous data, all satellites
// Class/Id 0x0b 0x33 (0 bytes)
// All UBX-AID messages are deprecated; use UBX-MGA messages instead Poll AssistNow Autonomous aiding data for all GPS satellites by sending this empty message. The receiver will return an AID-AOP message (see definition below) for each GPS satellite for which data is available.
type AidAop struct {
}

func (AidAop) classID() uint16 { return 0x330b }

// AidAop1 (Poll Request) Poll AssistNow Autonomous data, one GPS satellite
// Class/Id 0x0b 0x33 (1 bytes)
// All UBX-AID messages are deprecated; use UBX-MGA messages instead Poll the AssistNow Autonomous data for the specified GPS satellite. The receiver will return a AID-AOP message (see definition below) if data is available for the requested satellite.
type AidAop1 struct {
	Svid byte // GPS SV ID for which the data is requested (valid range: 1..32).
}

func (AidAop1) classID() uint16 { return 0x330b }

// AidAop2 (Input/Output) AssistNow Autonomous data
// Class/Id 0x0b 0x33 (68 bytes)
// All UBX-AID messages are deprecated; use UBX-MGA messages instead If enabled, this message is output at irregular intervals. It is output whenever AssistNow Autonomous has produced new data for a satellite. Depending on the availability of the optional data the receiver will output either version of the message. If this message is polled using one of the two poll requests described above the receiver will send this message if AssistNow Autonomous data is available or the corresponding poll request message if no AssistNow Autonomous data is available for each satellite (i.e. svid 1..32). At the user's choice the optional data may be chopped from the payload of a previously polled message when sending the message back to the receiver. Sending a valid AID- AOP message to the receiver will automatically enable the AssistNow Autonomous feature on the receiver. See the section AssistNow Autonomous in the receiver description for details on this feature.
type AidAop2 struct {
	GnssId    byte     // GNSS identifier (see Satellite Numbering)
	SvId      byte     // Satellite identifier (see Satellite Numbering)
	Reserved1 [2]byte  // Reserved
	Data      [64]byte // assistance data
}

func (AidAop2) classID() uint16 { return 0x330b }

// Message ubx-aid-eph (3 versions)

// AidEph (Poll Request) Poll GPS aiding ephemeris data
// Class/Id 0x0b 0x31 (0 bytes)
// All UBX-AID messages are deprecated; use UBX-MGA messages instead Poll GPS Aiding Data (Ephemeris) for all 32 SVs by sending this message to the receiver without any payload. The receiver will return 32 messages of type AID- EPH as defined below.
type AidEph struct {
}

func (AidEph) classID() uint16 { return 0x310b }

// AidEph1 (Poll Request) Poll GPS aiding ephemeris data for a SV
// Class/Id 0x0b 0x31 (1 bytes)
// All UBX-AID messages are deprecated; use UBX-MGA messages instead Poll GPS Constellation Data (Ephemeris) for an SV by sending this message to the receiver. The receiver will return one message of type AID-EPH as defined below.
type AidEph1 struct {
	Svid byte // SV ID for which the receiver shall return its Ephemeris Data (Valid Range: 1 .. 32).
}

func (AidEph1) classID() uint16 { return 0x310b }

// AidEph2 (Input/Output) GPS aiding ephemeris input/output message
// Class/Id 0x0b 0x31 (8 or 104 bytes)
// All UBX-AID messages are deprecated; use UBX-MGA messages instead  SF1D0 to SF3D7 is only sent if ephemeris is available for this SV. If not, the  payload may be reduced to 8 Bytes, or all bytes are set to zero, indicating that  this SV Number does not have valid ephemeris for the moment. This may  happen even if NAV-SVINFO and RXM-SVSI are indicating ephemeris  availability as the internal data may not represent the content of an original  broadcast ephemeris (or only parts thereof).  SF1D0 to SF3D7 contain the 24 words following the Hand-Over Word ( HOW )  from the GPS navigation message, subframes 1 to 3. The Truncated TOW  Count is not valid and cannot be used. See IS-GPS-200 for a full description of  the contents of the Subframes.  In SF1D0 to SF3D7, the parity bits have been removed, and the 24 bits of data  are located in Bits 0 to 23. Bits 24 to 31 shall be ignored.  When polled, the data contained in this message does not represent the full  original ephemeris broadcast. Some fields that are irrelevant to u-blox  receivers may be missing. The week number in Subframe 1 has already been  modified to match the Time Of Ephemeris (TOE).
type AidEph2 struct {
	Svid uint32 // SV ID for which this ephemeris data is (Valid Range: 1 .. 32).
	How  uint32 // Hand-Over Word of first Subframe. This is required if data is sent to the receiver. 0 indicates that no Ephemeris Data is following.
	// Optional
	Sf1d [8]uint32 // Subframe 1 Words 3..10 (SF1D0..SF1D7)
	Sf2d [8]uint32 // Subframe 2 Words 3..10 (SF2D0..SF2D7)
	Sf3d [8]uint32 // Subframe 3 Words 3..10 (SF3D0..SF3D7)
}

func (AidEph2) classID() uint16 { return 0x310b }

// Message ubx-aid-hui (2 versions)

// AidHui (Poll Request) Poll GPS health, UTC, ionosphere parameters
// Class/Id 0x0b 0x02 (0 bytes)
// All UBX-AID messages are deprecated; use UBX-MGA messages instead -
type AidHui struct {
}

func (AidHui) classID() uint16 { return 0x020b }

// AidHui1 (Input/Output) GPS health, UTC and ionosphere parameters
// Class/Id 0x0b 0x02 (72 bytes)
// All UBX-AID messages are deprecated; use UBX-MGA messages instead This message contains a health bit mask, UTC time and Klobuchar parameters. For more information on these parameters, see the ICD-GPS-200 documentation.
type AidHui1 struct {
	Health               uint32       // Bitmask, every bit represenst a GPS SV (1- 32). If the bit is set the SV is healthy.
	UtcA0                float64      // UTC - parameter A0
	UtcA1                float64      // UTC - parameter A1
	UtcTOW               int32        // UTC - reference time of week
	UtcWNT               int16        // UTC - reference week number
	UtcLS                int16        // UTC - time difference due to leap seconds before event
	UtcWNF               int16        // UTC - week number when next leap second event occurs
	UtcDN                int16        // UTC - day of week when next leap second event occurs
	UtcLSF               int16        // UTC - time difference due to leap seconds after event
	UtcSpare             int16        // UTC - Spare to ensure structure is a multiple of 4 bytes
	KlobA0_s             float32      // [s] Klobuchar - alpha 0
	KlobA1_s_semicircle  float32      // [s/semicircle] Klobuchar - alpha 1
	KlobA2_s_semicircle2 float32      // [s/semicircle^2] Klobuchar - alpha 2
	KlobA3               float32      // [s/semicircle^3] Klobuchar - alpha 3
	KlobB0_s             float32      // [s] Klobuchar - beta 0
	KlobB1_s_semicircle  float32      // [s/semicircle] Klobuchar - beta 1
	KlobB2_s_semicircle2 float32      // [s/semicircle^2] Klobuchar - beta 2
	KlobB3               float32      // [s/semicircle^3] Klobuchar - beta 3
	Flags                AidHui1Flags // flags
}

func (AidHui1) classID() uint16 { return 0x020b }

type AidHui1Flags uint32

const (
	AidHui1HealthValid AidHui1Flags = 0x1 // Healthmask field in this message is valid
	AidHui1UtcValid    AidHui1Flags = 0x2 // UTC parameter fields in this message are valid
	AidHui1KlobValid   AidHui1Flags = 0x4 // Klobuchar parameter fields in this message are valid
)

func (v AidHui1Flags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "HealthValid")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "UtcValid")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "KlobValid")
	}
	v &^= 0x4
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-aid-ini (2 versions)

// AidIni (Poll Request) Poll GPS initial aiding data
// Class/Id 0x0b 0x01 (0 bytes)
// All UBX-AID messages are deprecated; use UBX-MGA messages instead -
type AidIni struct {
}

func (AidIni) classID() uint16 { return 0x010b }

// AidIni1 (Input/Output) Aiding position, time, frequency, clock drift
// Class/Id 0x0b 0x01 (48 bytes)
// All UBX-AID messages are deprecated; use UBX-MGA messages instead This message contains position, time and clock drift information. The position can be input in either the ECEF X/Y/Z coordinate system or as lat/lon/height. The time can either be input as inexact value via the standard communication interface, suffering from latency depending on the baud rate, or using hardware time synchronization where an accurate time pulse is input on the external interrupts. It is also possible to supply hardware frequency aiding by connecting a continuous signal to an external interrupt.
type AidIni1 struct {
	EcefXOrLat       int32        // [cm_or_deg*1e-7] WGS84 ECEF X coordinate or latitude, depending on flags below
	EcefYOrLon       int32        // [cm_or_deg*1e-7] WGS84 ECEF Y coordinate or longitude, depending on flags below
	EcefZOrAlt_cm    int32        // [cm] WGS84 ECEF Z coordinate or altitude, depending on flags below
	PosAcc_cm        uint32       // [cm] Position accuracy (stddev)
	TmCfg            AidIni1TmCfg // Time mark configuration
	WnoOrDate        uint16       // [week_or_yearMonth] Actual week number or yearSince2000/Month (YYMM), depending on flags below
	TowOrTime        uint32       // [ms_or_dayHourMinuteSec] Actual time of week or DayOfMonth/Hour/Minute/Second (DDHHMMSS), depending on flags below
	TowNs_ns         int32        // [ns] Fractional part of time of week
	TAccMs_ms        uint32       // [ms] Milliseconds part of time accuracy
	TAccNs_ns        uint32       // [ns] Nanoseconds part of time accuracy
	ClkDOrFreq       int32        // [ns/s_or_Hz*1e-2] Clock drift or frequency, depending on flags below
	ClkDAccOrFreqAcc uint32       // [ns/s_or_ppb] Accuracy of clock drift or frequency, depending on flags below
	Flags            AidIni1Flags // Bitmask with the following flags
}

func (AidIni1) classID() uint16 { return 0x010b }

type AidIni1TmCfg uint16

const (
	AidIni1FEdge AidIni1TmCfg = 0x2  // use falling edge (default rising)
	AidIni1Tm1   AidIni1TmCfg = 0x10 // time mark on extint 1 (default extint 0)
	AidIni1F1    AidIni1TmCfg = 0x40 // frequency on extint 1 (default extint 0)
)

type AidIni1Flags uint32

const (
	AidIni1Pos    AidIni1Flags = 0x1   // Position is valid
	AidIni1Time   AidIni1Flags = 0x2   // Time is valid
	AidIni1ClockD AidIni1Flags = 0x4   // Clock drift data contains valid clock drift, must not be set together with clockF
	AidIni1Tp     AidIni1Flags = 0x8   // Use time pulse
	AidIni1ClockF AidIni1Flags = 0x10  // Clock drift data contains valid frequency, must not be set together with clockD
	AidIni1Lla    AidIni1Flags = 0x20  // Position is given in lat/long/alt (default is ECEF)
	AidIni1AltInv AidIni1Flags = 0x40  // Altitude is not valid, if lla was set
	AidIni1PrevTm AidIni1Flags = 0x80  // Use time mark received before AID-INI message (default uses mark received after message)
	AidIni1Utc    AidIni1Flags = 0x400 // Time is given as UTC date/time (default is GPS wno/tow)
)

func (v AidIni1TmCfg) String() string {
	var b []string
	if v&0x2 != 0 {
		b = append(b, "FEdge")
	}
	v &^= 0x2
	if v&0x10 != 0 {
		b = append(b, "Tm1")
	}
	v &^= 0x10
	if v&0x40 != 0 {
		b = append(b, "F1")
	}
	v &^= 0x40
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v AidIni1Flags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Pos")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "Time")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "ClockD")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "Tp")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "ClockF")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "Lla")
	}
	v &^= 0x20
	if v&0x40 != 0 {
		b = append(b, "AltInv")
	}
	v &^= 0x40
	if v&0x80 != 0 {
		b = append(b, "PrevTm")
	}
	v &^= 0x80
	if v&0x400 != 0 {
		b = append(b, "Utc")
	}
	v &^= 0x400
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-ant

// CfgAnt (Get/set) Antenna control settings
// Class/Id 0x06 0x13 (4 bytes)
// This message allows the user to configure the antenna supervisor. The antenna supervisor can be used to detect the status of an active antenna and control it. It can be used to turn off the supply to the antenna in the event of a short cirquit (for example) or to manage power consumption in power save mode. Refer to antenna supervisor configuration in the Integration manual for more information regarding the behavior of the antenna supervisor. Refer to UBX-MON-HW for a description of the fields in the message used to obtain the status of the antenna. Note that not all pins can be used for antenna supervisor operation, it is recommended that you use the default pins, consult the Integration manual if you need to use other pins.
type CfgAnt struct {
	Flags CfgAntFlags // Antenna flag mask
	Pins  CfgAntPins  // Antenna pin configuration
}

func (CfgAnt) classID() uint16 { return 0x1306 }

type CfgAntFlags uint16

const (
	CfgAntSvcs      CfgAntFlags = 0x1  // Enable antenna supply voltage control signal
	CfgAntScd       CfgAntFlags = 0x2  // Enable short circuit detection
	CfgAntOcd       CfgAntFlags = 0x4  // Enable open circuit detection
	CfgAntPdwnOnSCD CfgAntFlags = 0x8  // Power down antenna supply if short circuit is detected. (only in combination with bit 1)
	CfgAntRecovery  CfgAntFlags = 0x10 // Enable automatic recovery from short state
)

type CfgAntPins uint16

const (
	CfgAntPinSwitch CfgAntPins = 0x1f   // PIO-pin used for switching antenna supply
	CfgAntPinSCD    CfgAntPins = 0x3e0  // PIO-pin used for detecting a short in the antenna supply
	CfgAntPinOCD    CfgAntPins = 0x7c00 // PIO-pin used for detecting open/not connected antenna
	CfgAntReconfig  CfgAntPins = 0x8000 // if set to one, and this command is sent to the receiver, the receiver will reconfigure the pins as specified.
)

func (v CfgAntFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Svcs")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "Scd")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "Ocd")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "PdwnOnSCD")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "Recovery")
	}
	v &^= 0x10
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgAntPins) String() string {
	var b []string
	if v&0x1f != 0 {
		b = append(b, fmt.Sprintf("PinSwitch<%b>", (v&0x1f)>>0))
	}
	v &^= 0x1f
	if v&0x3e0 != 0 {
		b = append(b, fmt.Sprintf("PinSCD<%b>", (v&0x3e0)>>5))
	}
	v &^= 0x3e0
	if v&0x7c00 != 0 {
		b = append(b, fmt.Sprintf("PinOCD<%b>", (v&0x7c00)>>10))
	}
	v &^= 0x7c00
	if v&0x8000 != 0 {
		b = append(b, "Reconfig")
	}
	v &^= 0x8000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-batch

// CfgBatch (Get/set) Get/set data batching configuration
// Class/Id 0x06 0x93 (8 bytes)
// Gets or sets the configuration for data batching. See Data Batching for more information.
type CfgBatch struct {
	Version   byte          // Message version (0x00 for this version)
	Flags     CfgBatchFlags // Flags
	BufSize   uint16        // Size of buffer in number of epochs to store
	NotifThrs uint16        // Buffer fill level that triggers PIO notification, in number of epochs stored
	PioId     byte          // PIO ID to use for buffer level notification
	Reserved1 byte          // Reserved
}

func (CfgBatch) classID() uint16 { return 0x9306 }

type CfgBatchFlags byte

const (
	CfgBatchEnable       CfgBatchFlags = 0x1  // Enable data batching
	CfgBatchExtraPvt     CfgBatchFlags = 0x4  // Store extra PVT information The fields iTOW, tAcc, numSV, hMSL, vAcc, velN, velE, velD, sAcc, headAcc and pDOP in UBX-LOG- BATCH are only valid if this flag is set.
	CfgBatchExtraOdo     CfgBatchFlags = 0x8  // Store odometer data The fields distance, totalDistance and distanceStd in UBX-LOG-BATCH are only valid if this flag is set. Note: the odometer feature itself must also be enabled.
	CfgBatchPioEnable    CfgBatchFlags = 0x20 // Enable PIO notification
	CfgBatchPioActiveLow CfgBatchFlags = 0x40 // PIO is active low
)

func (v CfgBatchFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Enable")
	}
	v &^= 0x1
	if v&0x4 != 0 {
		b = append(b, "ExtraPvt")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "ExtraOdo")
	}
	v &^= 0x8
	if v&0x20 != 0 {
		b = append(b, "PioEnable")
	}
	v &^= 0x20
	if v&0x40 != 0 {
		b = append(b, "PioActiveLow")
	}
	v &^= 0x40
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-cfg

// CfgCfg (Command) Clear, save and load configurations
// Class/Id 0x06 0x09 (12 or 13 bytes)
// See Receiver configuration for a detailed description on how receiver configuration should be used. The three masks are made up of individual bits, each bit indicating the sub-section of all configurations on which the corresponding action shall be carried out. The reserved bits in the masks must be set to '0'. For detailed information refer to the Organization of the configuration sections. Note that commands can be combined. The sequence of execution is clear, save, load.
type CfgCfg struct {
	ClearMask CfgCfgClearMask // Mask with configuration sub-sections to clear (i.e. load default configurations to permanent configurations in non-volatile memory)
	SaveMask  CfgCfgClearMask // Mask with configuration sub-sections to save (i.e. save current configurations to non-volatile memory), see ID description of clearMask
	LoadMask  CfgCfgClearMask // Mask with configuration sub-sections to load (i.e. load permanent configurations from non-volatile memory to current configurations), see ID description of clearMask
	// Optional
	DeviceMask CfgCfgDeviceMask // Mask which selects the memory devices for this command.
}

func (CfgCfg) classID() uint16 { return 0x0906 }

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

type CfgCfgDeviceMask byte

const (
	CfgCfgDevBBR      CfgCfgDeviceMask = 0x1  // Battery backed RAM
	CfgCfgDevFlash    CfgCfgDeviceMask = 0x2  // Flash
	CfgCfgDevEEPROM   CfgCfgDeviceMask = 0x4  // EEPROM
	CfgCfgDevSpiFlash CfgCfgDeviceMask = 0x10 // SPI Flash
)

func (v CfgCfgClearMask) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "IoPort")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "MsgConf")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "InfMsg")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "NavConf")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "RxmConf")
	}
	v &^= 0x10
	if v&0x100 != 0 {
		b = append(b, "SenConf")
	}
	v &^= 0x100
	if v&0x200 != 0 {
		b = append(b, "RinvConf")
	}
	v &^= 0x200
	if v&0x400 != 0 {
		b = append(b, "AntConf")
	}
	v &^= 0x400
	if v&0x800 != 0 {
		b = append(b, "LogConf")
	}
	v &^= 0x800
	if v&0x1000 != 0 {
		b = append(b, "FtsConf")
	}
	v &^= 0x1000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgCfgDeviceMask) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "DevBBR")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "DevFlash")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "DevEEPROM")
	}
	v &^= 0x4
	if v&0x10 != 0 {
		b = append(b, "DevSpiFlash")
	}
	v &^= 0x10
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-dat (2 versions)

// CfgDat (Set) Set user-defined datum
// Class/Id 0x06 0x06 (44 bytes)
// For more information see the description of Geodetic Systems and Frames.
type CfgDat struct {
	MajA_m    float64 // [m] Semi-major axis ( accepted range = 6,300, 000.0 to 6,500,000.0 meters ).
	Flat      float64 // 1.0 / flattening ( accepted range is 0.0 to 500.0 ).
	DX_m      float32 // [m] X axis shift at the origin ( accepted range is +/- 5000.0 meters ).
	DY_m      float32 // [m] Y axis shift at the origin ( accepted range is +/- 5000.0 meters ).
	DZ_m      float32 // [m] Z axis shift at the origin ( accepted range is +/- 5000.0 meters ).
	RotX_s    float32 // [s] Rotation about the X axis ( accepted range is +/- 20.0 milli-arc seconds ).
	RotY_s    float32 // [s] Rotation about the Y axis ( accepted range is +/- 20.0 milli-arc seconds ).
	RotZ_s    float32 // [s] Rotation about the Z axis ( accepted range is +/- 20.0 milli-arc seconds ).
	Scale_ppm float32 // [ppm] Scale change ( accepted range is 0.0 to 50.0 parts per million ).
}

func (CfgDat) classID() uint16 { return 0x0606 }

// CfgDat1 (Get) Get currently defined datum
// Class/Id 0x06 0x06 (52 bytes)
// Returns the parameters of the currently defined datum. If no user-defined datum has been set, this will default to WGS84.
type CfgDat1 struct {
	DatumNum  uint16  // Datum number: 0 = WGS84, 0xFFFF = user-defined
	DatumName [6]byte // ASCII string: WGS84 or USER
	MajA_m    float64 // [m] Semi-major axis ( accepted range = 6,300, 000.0 to 6,500,000.0 meters ).
	Flat      float64 // 1.0 / flattening ( accepted range is 0.0 to 500.0 ).
	DX_m      float32 // [m] X axis shift at the origin ( accepted range is +/- 5000.0 meters ).
	DY_m      float32 // [m] Y axis shift at the origin ( accepted range is +/- 5000.0 meters ).
	DZ_m      float32 // [m] Z axis shift at the origin ( accepted range is +/- 5000.0 meters ).
	RotX_s    float32 // [s] Rotation about the X axis ( accepted range is +/- 20.0 milli-arc seconds ).
	RotY_s    float32 // [s] Rotation about the Y axis ( accepted range is +/- 20.0 milli-arc seconds ).
	RotZ_s    float32 // [s] Rotation about the Z axis ( accepted range is +/- 20.0 milli-arc seconds ).
	Scale_ppm float32 // [ppm] Scale change ( accepted range is 0.0 to 50.0 parts per million ).
}

func (CfgDat1) classID() uint16 { return 0x0606 }

// Message ubx-cfg-dgnss

// CfgDgnss (Get/set) DGNSS configuration
// Class/Id 0x06 0x70 (4 bytes)
// This message allows the user to configure the DGNSS configuration of the receiver.
type CfgDgnss struct {
	DgnssMode byte    // Specifies differential mode: 2: RTK float: No attempts are made to fix ambiguities. 3: RTK fixed: Ambiguities are fixed whenever possible.
	Reserved1 [3]byte // Reserved
}

func (CfgDgnss) classID() uint16 { return 0x7006 }

// Message ubx-cfg-dosc

// CfgDosc (Get/set) Disciplined oscillator configuration
// Class/Id 0x06 0x61 (4 + N*32 bytes)
// This message allows the characteristics of the internal or external oscillator to be described to the receiver. The gainVco and gainUncertainty parameters are normally set using the calibration process initiated using UBX-TIM-VCOCAL. The behavior of the system can be badly affected by setting the wrong values, so customers are advised to only change these parameters with care.
type CfgDosc struct {
	Version   byte              // Message version (0x00 for this version)
	NumOsc    byte              `len:"Osc"` // Number of oscillators to configure (affects length of this message)
	Reserved1 [2]byte           // Reserved
	Osc       []*CfgDoscOscType // len: NumOsc
}

func (CfgDosc) classID() uint16 { return 0x6106 }

type CfgDoscOscType struct {
	OscId              byte         // Id of oscillator. 0 - internal oscillator 1 - external oscillator
	Reserved2          byte         // Reserved
	Flags              CfgDoscFlags // flags
	Freq_hzl2          uint32       // [2^-2 Hz] Nominal frequency of source
	PhaseOffset_ps     int32        // [ps] Intended phase offset of the oscillator relative to the leading edge of the time pulse
	WithTemp_ppbl8     uint32       // [2^-8 ppb] Oscillator stability limit over operating temperature range (must be > 0)
	WithAge_ppb_yearl8 uint32       // [2^-8 ppb/year] Oscillator stability with age (must be > 0)
	TimeToTemp_s       uint16       // [s] The minimum time that it could take for a temperature variation to move the oscillator frequency by 'withTemp' (must be > 0)
	Reserved3          [2]byte      // Reserved
	GainVco            int32        // [2^-16 ppb/raw LSB] Oscillator control gain/slope; change of frequency per unit change in raw control change
	GainUncertainty    byte         // Relative uncertainty (1 standard deviation) of oscillator control gain/slope
	Reserved4          [3]byte      // Reserved
}

type CfgDoscFlags uint16

const (
	CfgDoscIsCalibrated CfgDoscFlags = 0x1  // 1 if the oscillator gain is calibrated, 0 if not
	CfgDoscControlIf    CfgDoscFlags = 0x1e // Communication interface for oscillator control: 0: Custom DAC attached to receiver's I2C 1: Microchip MCP4726 (12 bit DAC) attached to receiver's I2C 2: TI DAC8571 (16 bit DAC) attached to receiver's I2C 13: 12 bit DAC attached to host 14: 14 bit DAC attached to host 15: 16 bit DAC attached to host Note that for DACs attached to the host, the host must monitor UBX-TIM-DOSC messages and pass the supplied raw values on to the DAC.
)

func (v CfgDoscFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "IsCalibrated")
	}
	v &^= 0x1
	if v&0x1e != 0 {
		b = append(b, fmt.Sprintf("ControlIf<%b>", (v&0x1e)>>1))
	}
	v &^= 0x1e
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-esfa

// CfgEsfa (Get/set) Get/set the Accelerometer (A) sensor configuration
// Class/Id 0x06 0x4c (20 bytes)
// Get/set the configuration for the accelerometer sensor required for External Sensor Fusion (ESF) based navigation. More details can be found in the Accelerometer Configuration section.
type CfgEsfa struct {
	Version             byte    // Message version (0x00 for this version)
	Reserved1           [9]byte // Reserved
	AccelRmsThdl_m_s2l6 byte    // [2^-6 m/s^2] Accelerometer RMS threshold below which automatically estimated accelerometer noise-level (accuracy) is updated.
	Frequency_hz        byte    // [Hz] Nominal accelerometer sensor data sampling frequency.
	Latency_ms          uint16  // [ms] Accelerometer sensor data latency due to e.g. CAN bus.
	Accuracy_m_s2e4     uint16  // [1e-4 m/s^2] Accelerometer sensor data accuracy.
	Reserved2           [4]byte // Reserved
}

func (CfgEsfa) classID() uint16 { return 0x4c06 }

// Message ubx-cfg-esfalg

// CfgEsfalg (Get/set) Get/set IMU-mount misalignment configuration
// Class/Id 0x06 0x56 (12 bytes)
// Get/set the IMU-mount misalignment configuration (rotation from installation- frame to the IMU-frame). A detailed description on how to compose this configuration is given in the ADR Installation section for ADR products. A detailed description on how to compose this configuration is given in the UDR Installation section for UDR products.
type CfgEsfalg struct {
	Bitfield    CfgEsfalgBitfield // Bitfield
	Yaw_dege2   uint32            // [1e-2 deg] User-defined IMU-mount yaw angle [0, 360]
	Pitch_dege2 int16             // [1e-2 deg] User-defined IMU-mount pitch angle [-90, 90]
	Roll_dege2  int16             // [1e-2 deg] User-defined IMU-mount roll angle [-180, 180]
}

func (CfgEsfalg) classID() uint16 { return 0x5606 }

type CfgEsfalgBitfield uint32

const (
	CfgEsfalgVersion      CfgEsfalgBitfield = 0xff  // Message version (0x00 for this version)
	CfgEsfalgDoAutoMntAlg CfgEsfalgBitfield = 0x100 // Only supported on certain products. Enable/disable automatic IMU-mount alignment (0: Disabled, 1: Enabled). This flag can only be used with modules containing an internal IMU.
)

func (v CfgEsfalgBitfield) String() string {
	var b []string
	if v&0xff != 0 {
		b = append(b, fmt.Sprintf("Version<%b>", (v&0xff)>>0))
	}
	v &^= 0xff
	if v&0x100 != 0 {
		b = append(b, "DoAutoMntAlg")
	}
	v &^= 0x100
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-esfg

// CfgEsfg (Get/set) Get/set the Gyroscope (G) sensor configuration
// Class/Id 0x06 0x4d (20 bytes)
// Get/set the configuration for the gyroscope sensor required for External Sensor Fusion (ESF) based navigation. More details can be found in the Gyroscope Configuration section.
type CfgEsfg struct {
	Version             byte    // Message version (0x00 for this version)
	Reserved1           [7]byte // Reserved
	TcTableSaveRate_s   uint16  // [s] Temperature-dependent gyroscope bias table saving update rate.
	GyroRmsThdl_deg_sl8 byte    // [2^-8 deg/s] Gyroscope sensor RMS threshold below which automatically estimated gyroscope noise-level (accuracy) is updated.
	Frequency_hz        byte    // [Hz] Nominal gyroscope sensor data sampling frequency.
	Latency_ms          uint16  // [ms] Gyroscope sensor data latency due to e.g. CAN bus.
	Accuracy_deg_se3    uint16  // [1e-3 deg/s] Gyroscope sensor data accuracy.
	Reserved2           [4]byte // Reserved
}

func (CfgEsfg) classID() uint16 { return 0x4d06 }

// Message ubx-cfg-esfwt

// CfgEsfwt (Get/set) Get/set wheel-tick configuration
// Class/Id 0x06 0x82 (32 bytes)
// Get/set the wheel-tick configuration for GWT or GAWT solution. Further information on the configuration parameters is given in the Automotive Dead Reckoning (ADR) chapter. This field can only be used with modules supporting analog wheel-tick signals and containing an internal IMU.
type CfgEsfwt struct {
	Version            byte           // Message version (0x00 for this version)
	Flags1             CfgEsfwtFlags1 // Flags
	Flags2             CfgEsfwtFlags2 // Flags
	Reserved1          [1]byte        // Reserved
	WtFactor           uint32         // Wheel-tick scale factor to obtain distance [m] from wheel-ticks (0 = not set)
	WtQuantError       uint32         // [1e-6 m (orm/s)] Wheel-tick quantization. If useWtSpeed is set then this is interpreted as the speed measurement error RMS.
	WtCountMax         uint32         // Wheel-tick counter maximum value (rollover - 1). If null, relative wheel-tick counts are assumed (and therefore no rollover). If not null, absolute wheel-tick counts are assumed and the value corresponds to the highest tick count value before rollover happens. If useWtSpeed is set then this value is ignored. If value is set to 1, absolute wheel-tick counts are assumed and the value will be automatic calculated if possible. It is only possible for automatic calibration to calculate wtCntMax if it can be represented as a number of set bits (i.e. 2^N). If it cannot be represented in this way it must be set to the correct absolute tick value manually.
	WtLatency_ms       uint16         // [ms] Wheel-tick data latency due to e.g. CAN bus
	WtFrequency_hz     byte           // [Hz] Nominal wheel-tick data frequency (0 = not set)
	Flags3             CfgEsfwtFlags3 // Flags
	SpeedDeadBand_cm_s uint16         // [cm/s] Speed sensor dead band (0 = not set)
	Reserved2          [10]byte       // Reserved
}

func (CfgEsfwt) classID() uint16 { return 0x8206 }

type CfgEsfwtFlags1 byte

const (
	CfgEsfwtCombineTicks CfgEsfwtFlags1 = 0x1  // Use combined rear wheel-ticks instead of the single tick
	CfgEsfwtUseWtSpeed   CfgEsfwtFlags1 = 0x10 // Use speed measurements (data type 11 in ESF-MEAS) instead of single ticks (data type 10)
	CfgEsfwtDirPinPol    CfgEsfwtFlags1 = 0x20 // Only supported on certain products. Direction pin polarity 0: High signal level means forward direction, 1: High signal level means backward direction.
	CfgEsfwtUseWtPin     CfgEsfwtFlags1 = 0x40 // Use wheel-tick pin for speed measurement.
)

type CfgEsfwtFlags2 byte

const (
	CfgEsfwtAutoWtCountMaxOff CfgEsfwtFlags2 = 0x1 // Disable automatic estimation of maximum absolute wheel-tick counter value (0: enabled, 1: disabled). See wtCountMax field description for more details. (Not supported in protocol versions less than 19)
	CfgEsfwtAutoDirPinPolOff  CfgEsfwtFlags2 = 0x2 // Only supported on certain products. Disable automatic wheel-tick direction pin polarity detection (0: enabled, 1: disabled). See dirPinPol field description for more details. (Not supported in protocol versions less than 19)
	CfgEsfwtAutoSoftwareWtOff CfgEsfwtFlags2 = 0x4 // Only supported on certain products. Disable automatic use of wheel-tick or speed data received over the software interface if available (0: enabled, 1: disabled). In this case, data coming from the hardware interface (wheel-tick pins) will automatically be ignored if wheel-tick/speed data are available from the software interface. See useWtPin field description for more details. (Not supported in protocol versions less than 19)
	CfgEsfwtAutoUseWtSpeedOff CfgEsfwtFlags2 = 0x8 // Disable automatic receiver reconfiguration for processing speed data instead of wheel-tick data if no wheel-tick data are available but speed data were detected (0: enabled, 1: disabled). See useWtSpeed field description for more details. (Not supported in protocol versions less than 19)
)

type CfgEsfwtFlags3 byte

const (
	CfgEsfwtCntBothEdges CfgEsfwtFlags3 = 0x10 // Only supported on certain products. Count both rising and falling edges on wheel-tick signal (only relevant if wheel-tick is measured by the u-blox receiver). Only turn on this feature if the wheel-tick signal has 50 % duty cycle. Turning on this feature with fixed-width pulses can lead to severe degradation of performance. Use wheel-tick pin for speed measurement. This field can only be used with modules supporting analog wheel-tick signals.
)

func (v CfgEsfwtFlags1) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "CombineTicks")
	}
	v &^= 0x1
	if v&0x10 != 0 {
		b = append(b, "UseWtSpeed")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "DirPinPol")
	}
	v &^= 0x20
	if v&0x40 != 0 {
		b = append(b, "UseWtPin")
	}
	v &^= 0x40
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgEsfwtFlags2) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "AutoWtCountMaxOff")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "AutoDirPinPolOff")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "AutoSoftwareWtOff")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "AutoUseWtSpeedOff")
	}
	v &^= 0x8
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgEsfwtFlags3) String() string {
	var b []string
	if v&0x10 != 0 {
		b = append(b, "CntBothEdges")
	}
	v &^= 0x10
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-esrc

// CfgEsrc (Get/set) External synchronization source configuration
// Class/Id 0x06 0x60 (4 + N*36 bytes)
// External time or frequency source configuration. The stability of time and frequency sources is described using different fields, see sourceType field documentation.
type CfgEsrc struct {
	Version    byte                  // Message version (0x00 for this version)
	NumSources byte                  `len:"Sources"` // Number of sources (affects length of this message)
	Reserved1  [2]byte               // Reserved
	Sources    []*CfgEsrcSourcesType // len: NumSources
}

func (CfgEsrc) classID() uint16 { return 0x6006 }

type CfgEsrcSourcesType struct {
	ExtInt               byte         // EXTINT index of this source (0 for EXTINT0 and 1 for EXTINT1)
	SourceType           byte         // Source type: 0:  none 1: frequency source; use withTemp, withAge, timeToTemp and maxDevLifeTime to describe the stability of the source 2:  time source; use offset, offsetUncertainty and jitter fields to describe the stability of the source 3:  feedback from external oscillator; stability data is taken from the external oscillator's configuration
	Flags                CfgEsrcFlags // Flags
	Freq_hzl2            uint32       // [2^-2 Hz] Nominal frequency of source
	Reserved2            [4]byte      // Reserved
	WithTemp_ppbl8       uint32       // [2^-8 ppb] Oscillator stability limit over operating temperature range (must be > 0) Only used if sourceType is 1.
	WithAge_ppb_yearl8   uint32       // [2^-8 ppb/year] Oscillator stability with age (must be > 0) Only used if sourceType is 1.
	TimeToTemp_s         uint16       // [s] The minimum time that it could take for a temperature variation to move the oscillator frequency by 'withTemp' (must be > 0) Only used if sourceType is 1.
	MaxDevLifeTime_ppb   uint16       // [ppb] Maximum frequency deviation during lifetime (must be > 0) Only used if sourceType is 1.
	Offset_ns            int32        // [ns] Phase offset of signal Only used if sourceType is 2.
	OffsetUncertainty_ns uint32       // [ns] Uncertainty of phase offset (one standard deviation) Only used if sourceType is 2.
	Jitter_ns_s          uint32       // [ns/s] Phase jitter (must be > 0) Only used if sourceType is 2.
}

type CfgEsrcFlags uint16

const (
	CfgEsrcPolarity CfgEsrcFlags = 0x1 // Polarity of signal: 0: leading edge is rising edge 1: leading edge is falling edge
	CfgEsrcGnssUtc  CfgEsrcFlags = 0x2 // Time base of timing signal: 0: GNSS - as specified in CFG-TP5 (or GPS if CFG-TP5 indicates UTC) 1: UTC Only used if sourceType is 2.
)

func (v CfgEsrcFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Polarity")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "GnssUtc")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-geofence

// CfgGeofence (Get/set) Geofencing configuration
// Class/Id 0x06 0x69 (8 + N*12 bytes)
// Gets or sets the geofencing configuration. See the Geofencing description for feature details. If the receiver is sent a valid new configuration, it will respond with a UBX-ACK- ACK message and immediately change to the new configuration. Otherwise the receiver will reject the request, by issuing a UBX-ACK-NAK and continuing operation with the previous configuration. Note that the acknowledge message does not indicate whether the PIO configuration has been successfully applied (pin assigned), it only indicates the successful configuration of the feature. The configured PIO must be previously unoccupied for successful assignment.
type CfgGeofence struct {
	Version     byte                     // Message version (0x00 for this version)
	NumFences   byte                     `len:"Fences"` // Number of geofences contained in this message. Note that the receiver can only store a limited number of geofences (currently 4).
	ConfLvl     byte                     // Required confidence level for state evaluation. This value times the position's standard deviation (sigma) defines the confidence band. 0 = no confidence required 1 = 68% 2 = 95% 3 = 99.7% 4 = 99.99%
	Reserved1   [1]byte                  // Reserved
	PioEnabled  byte                     // 1 = Enable PIO combined fence state output, 0 = disable
	PinPolarity byte                     // PIO pin polarity. 0 = Low means inside, 1 = Low means outside. Unknown state is always high.
	Pin         byte                     // PIO pin number
	Reserved2   [1]byte                  // Reserved
	Fences      []*CfgGeofenceFencesType // len: NumFences
}

func (CfgGeofence) classID() uint16 { return 0x6906 }

type CfgGeofenceFencesType struct {
	Lat_dege7  int32  // [1e-7 deg] Latitude of the geofence circle center
	Lon_dege7  int32  // [1e-7 deg] Longitude of the geofence circle center
	Radius_me2 uint32 // [1e-2 m] Radius of the geofence circle
}

// Message ubx-cfg-gnss

// CfgGnss (Get/set) GNSS system configuration
// Class/Id 0x06 0x3e (4 + N*8 bytes)
// Gets or sets the GNSS system channel sharing configuration. If the receiver is sent a valid new configuration, it will respond with a UBX-ACK- ACK message and immediately change to the new configuration. Otherwise the receiver will reject the request, by issuing a UBX-ACK-NAK and continuing operation with the previous configuration. Configuration requirements:  It is necessary for at least one major GNSS to be enabled, after applying the  new configuration to the current one.  It is also required that at least 4 tracking channels are available to each  enabled major GNSS, i.e. maxTrkCh must have a minimum value of 4 for each  enabled major GNSS.  The number of tracking channels in use must not exceed the number of  tracking channels available in hardware, and the sum of all reserved tracking  channels needs to be less than or equal to the number of tracking channels in  use. Notes:  To avoid cross-correlation issues, it is recommended that GPS and QZSS are  always both enabled or both disabled.  Polling this message returns the configuration of all supported GNSS, whether  enabled or not; it may also include GNSS unsupported by the particular  product, but in such cases the enable flag will always be unset.  See section GNSS Configuration for a discussion of the use of this message.  See section Satellite Numbering for a description of the GNSS IDs available.  Configuration specific to the GNSS system can be done via other messages (e.  g. UBX-CFG-SBAS).
type CfgGnss struct {
	MsgVer          byte                       // Message version (0x00 for this version)
	NumTrkChHw      byte                       // Number of tracking channels available in hardware (read only)
	NumTrkChUse     byte                       // (Read only in protocol versions greater than 23) Number of tracking channels to use. Must be > 0, <= numTrkChHw. If 0xFF, then number of tracking channels to use will be set to numTrkChHw.
	NumConfigBlocks byte                       `len:"ConfigBlocks"` // Number of configuration blocks following
	ConfigBlocks    []*CfgGnssConfigBlocksType // len: NumConfigBlocks
}

func (CfgGnss) classID() uint16 { return 0x3e06 }

type CfgGnssConfigBlocksType struct {
	GnssId    byte         // System identifier (see Satellite Numbering )
	ResTrkCh  byte         // (Read only in protocol versions greater than 23) Number of reserved (minimum) tracking channels for this system.
	MaxTrkCh  byte         // (Read only in protocol versions greater than 23) Maximum number of tracking channels used for this system. Must be > 0, >= resTrkChn, <= numTrkChUse and <= maximum number of tracking channels supported for this system.
	Reserved1 byte         // Reserved
	Flags     CfgGnssFlags // Bitfield of flags. At least one signal must be configured in every enabled system.
}

type CfgGnssFlags uint32

const (
	CfgGnssEnable     CfgGnssFlags = 0x1      // Enable this system
	CfgGnssSigCfgMask CfgGnssFlags = 0xff0000 // Signal configuration mask When gnssId is 0 (GPS) 0x01 = GPS L1C/A 0x10 = GPS L2C 0x20 = GPS L5 When gnssId is 1 (SBAS) 0x01 = SBAS L1C/A When gnssId is 2 (Galileo) 0x01 = Galileo E1 (not supported in protocol versions less than 18) 0x10 = Galileo E5a 0x20 = Galileo E5b When gnssId is 3 (BeiDou) 0x01 = BeiDou B1I 0x10 = BeiDou B2I 0x80 = BeiDou B2A When gnssId is 4 (IMES) 0x01 = IMES L1 When gnssId is 5 (QZSS) 0x01 = QZSS L1C/A 0x04 = QZSS L1S 0x10 = QZSS L2C 0x20 = QZSS L5 When gnssId is 6 (GLONASS) 0x01 = GLONASS L1 0x10 = GLONASS L2
)

func (v CfgGnssFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Enable")
	}
	v &^= 0x1
	if v&0xff0000 != 0 {
		b = append(b, fmt.Sprintf("SigCfgMask<%b>", (v&0xff0000)>>16))
	}
	v &^= 0xff0000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-hnr

// CfgHnr (Get/set) High navigation rate settings
// Class/Id 0x06 0x5c (4 bytes)
// The u-blox receivers support high rates of navigation update up to 30 Hz. The navigation solution output UBX-NAV-HNR will not be aligned to the top of a second.  The update rate has a direct influence on the power consumption. The more  fixes that are required, the more CPU power and communication resources are  required.  For most applications a 1 Hz update rate would be sufficient.
type CfgHnr struct {
	HighNavRate_hz byte    // [Hz] Rate of navigation solution output
	Reserved1      [3]byte // Reserved
}

func (CfgHnr) classID() uint16 { return 0x5c06 }

// Message ubx-cfg-inf (2 versions)

// CfgInf (Get/set) Information message configuration
// Class/Id 0x06 0x02 (0 + N*10 bytes)
// The value of infMsgMask[x] below is formed so that each bit represents one of the INF class messages (bit 0 for ERROR, bit 1 for WARNING and so on). For a complete list, see the Message class INF. Several configurations can be concatenated to one input message. In this case the payload length can be a multiple of the normal length. Output messages from the module contain only one configuration unit. Note that:  I/O ports 1 and 2 correspond to serial ports 1 and 2.  I/O port 0 is I2C (DDC).  I/O port 3 is USB.  I/O port 4 is SPI.  I/O port 5 is reserved for future use.
type CfgInf struct {
	Items []*CfgInfItemsType // len: N
}

func (CfgInf) classID() uint16 { return 0x0206 }

type CfgInfItemsType struct {
	ProtocolID byte                // Protocol identifier, identifying for which protocol the configuration is set/get. The following are valid protocol identifiers: 0: UBX protocol 1: NMEA protocol 2-255: Reserved
	Reserved1  [3]byte             // Reserved
	InfMsgMask [6]CfgInfInfMsgMask // A bit mask, saying which information messages are enabled on each I/O port
}

type CfgInfInfMsgMask byte

const (
	CfgInfERROR   CfgInfInfMsgMask = 0x1  // enable ERROR
	CfgInfWARNING CfgInfInfMsgMask = 0x2  // enable WARNING
	CfgInfNOTICE  CfgInfInfMsgMask = 0x4  // enable NOTICE
	CfgInfTEST    CfgInfInfMsgMask = 0x8  // enable TEST
	CfgInfDEBUG   CfgInfInfMsgMask = 0x10 // enable DEBUG
)

func (v CfgInfInfMsgMask) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "ERROR")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "WARNING")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "NOTICE")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "TEST")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "DEBUG")
	}
	v &^= 0x10
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// CfgInf1 (Poll Request) Poll configuration for one protocol
// Class/Id 0x06 0x02 (1 bytes)
// -
type CfgInf1 struct {
	ProtocolID byte // Protocol identifier, identifying the output protocol for this poll request. The following are valid protocol identifiers: 0: UBX protocol 1: NMEA protocol 2-255: Reserved
}

func (CfgInf1) classID() uint16 { return 0x0206 }

// Message ubx-cfg-itfm

// CfgItfm (Get/set) Jamming/interference monitor configuration
// Class/Id 0x06 0x39 (8 bytes)
// -
type CfgItfm struct {
	Config  CfgItfmConfig  // Interference config word
	Config2 CfgItfmConfig2 // Extra settings for jamming/interference monitor
}

func (CfgItfm) classID() uint16 { return 0x3906 }

type CfgItfmConfig uint32

const (
	CfgItfmBbThreshold   CfgItfmConfig = 0xf        // Broadband jamming detection threshold (unit = dB)
	CfgItfmCwThreshold   CfgItfmConfig = 0x1f0      // CW jamming detection threshold (unit = dB)
	CfgItfmAlgorithmBits CfgItfmConfig = 0x7ffffe00 // Reserved algorithm settings - should be set to 0x16B156 in hex for correct settings
	CfgItfmEnable        CfgItfmConfig = 0x80000000 // Enable interference detection
)

type CfgItfmConfig2 uint32

const (
	CfgItfmGeneralBits CfgItfmConfig2 = 0xfff  // General settings - should be set to 0x31E in hex for correct setting
	CfgItfmAntSetting  CfgItfmConfig2 = 0x3000 // Antenna setting, 0=unknown, 1=passive, 2=active
	CfgItfmEnable2     CfgItfmConfig2 = 0x4000 // Set to 1 to scan auxiliary bands (u-blox 8 / u-blox M8 only, otherwise ignored)
)

func (v CfgItfmConfig) String() string {
	var b []string
	if v&0xf != 0 {
		b = append(b, fmt.Sprintf("BbThreshold<%b>", (v&0xf)>>0))
	}
	v &^= 0xf
	if v&0x1f0 != 0 {
		b = append(b, fmt.Sprintf("CwThreshold<%b>", (v&0x1f0)>>4))
	}
	v &^= 0x1f0
	if v&0x7ffffe00 != 0 {
		b = append(b, fmt.Sprintf("AlgorithmBits<%b>", (v&0x7ffffe00)>>9))
	}
	v &^= 0x7ffffe00
	if v&0x80000000 != 0 {
		b = append(b, "Enable")
	}
	v &^= 0x80000000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgItfmConfig2) String() string {
	var b []string
	if v&0xfff != 0 {
		b = append(b, fmt.Sprintf("GeneralBits<%b>", (v&0xfff)>>0))
	}
	v &^= 0xfff
	if v&0x3000 != 0 {
		b = append(b, fmt.Sprintf("AntSetting<%b>", (v&0x3000)>>12))
	}
	v &^= 0x3000
	if v&0x4000 != 0 {
		b = append(b, "Enable2")
	}
	v &^= 0x4000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-logfilter

// CfgLogfilter (Get/set) Data logger configuration
// Class/Id 0x06 0x47 (12 bytes)
// This message can be used to configure the data logger, i.e. to enable/disable the log recording and to get/set the position entry filter settings. Position entries can be filtered based on time difference, position difference or current speed thresholds. Position and speed filtering also have a minimum time interval. A position is logged if any of the thresholds are exceeded. If a threshold is set to zero it is ignored. The maximum rate of position logging is 1 Hz. The filter settings will be configured to the provided values only if the 'applyAllFilterSettings' flag is set. This allows the recording to be enabled/disabled independently of configuring the filter settings. Configuring the data logger in the absence of a logging file is supported. By doing so, once the logging file is created, the data logger configuration will take effect immediately and logging recording and filtering will activate according to the configuration.
type CfgLogfilter struct {
	Version             byte              // Message version (0x01 for this version)
	Flags               CfgLogfilterFlags // Flags
	MinInterval_s       uint16            // [s] Minimum time interval between logged positions (0 = not set). This is only applied in combination with the speed and/or position thresholds. If both minInterval and timeThreshold are set, minInterval must be less than or equal to timeThreshold.
	TimeThreshold_s     uint16            // [s] If the time difference is greater than the threshold, then the position is logged (0 = not set).
	SpeedThreshold_m_s  uint16            // [m/s] If the current speed is greater than the threshold, then the position is logged (0 = not set). minInterval also applies.
	PositionThreshold_m uint32            // [m] If the 3D position difference is greater than the threshold, then the position is logged (0 = not set). minInterval also applies.
}

func (CfgLogfilter) classID() uint16 { return 0x4706 }

type CfgLogfilterFlags byte

const (
	CfgLogfilterRecordEnabled          CfgLogfilterFlags = 0x1 // 1 = enable recording, 0 = disable recording
	CfgLogfilterPsmOncePerWakupEnabled CfgLogfilterFlags = 0x2 // 1 = enable recording only one single position per PSM on/off mode wake-up period, 0 = disable once per wake-up
	CfgLogfilterApplyAllFilterSettings CfgLogfilterFlags = 0x4 // 1 = apply all filter settings, 0 = only apply recordEnabled
)

func (v CfgLogfilterFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "RecordEnabled")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "PsmOncePerWakupEnabled")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "ApplyAllFilterSettings")
	}
	v &^= 0x4
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-msg (3 versions)

// CfgMsg (Poll Request) Poll a message configuration
// Class/Id 0x06 0x01 (2 bytes)
// -
type CfgMsg struct {
	MsgClass byte // Message class
	MsgID    byte // Message identifier
}

func (CfgMsg) classID() uint16 { return 0x0106 }

// CfgMsg1 (Get/set) Set message rate
// Class/Id 0x06 0x01 (3 bytes)
// Set message rate configuration for the current port. See also section How to change between protocols.
type CfgMsg1 struct {
	MsgClass byte // Message class
	MsgID    byte // Message identifier
	Rate     byte // Send rate on current port
}

func (CfgMsg1) classID() uint16 { return 0x0106 }

// CfgMsg2 (Get/set) Set message rate(s)
// Class/Id 0x06 0x01 (8 bytes)
// Get/set message rate configuration (s) to/from the receiver. See also section How to change between protocols.  Send rate is relative to the event a message is registered on. For example, if  the rate of a navigation message is set to 2, the message is sent every second  navigation solution. For configuring NMEA messages, the section NMEA  Messages Overview describes class and identifier numbers used.
type CfgMsg2 struct {
	MsgClass byte    // Message class
	MsgID    byte    // Message identifier
	Rate     [6]byte // Send rate on I/O port (6 ports)
}

func (CfgMsg2) classID() uint16 { return 0x0106 }

// Message ubx-cfg-nav5

// CfgNav5 (Get/set) Navigation engine settings
// Class/Id 0x06 0x24 (36 bytes)
// See the Navigation Configuration Settings Description for a detailed description of how these settings affect receiver operation.
type CfgNav5 struct {
	Mask                  CfgNav5Mask // Parameters bitmask. Only the masked parameters will be applied.
	DynModel              byte        // Dynamic platform model: 0: portable 2: stationary 3: pedestrian 4: automotive 5: sea 6: airborne with <1g acceleration 7: airborne with <2g acceleration 8: airborne with <4g acceleration 9: wrist-worn watch (not supported in protocol versions less than 18) 10: bike (supported in protocol versions 19. 2)
	FixMode               byte        // Position fixing mode: 1: 2D only 2: 3D only 3: auto 2D/3D
	FixedAlt_me2          int32       // [1e-2 m] Fixed altitude (mean sea level) for 2D fix mode
	FixedAltVar_m2e4      uint32      // [1e-4 m^2] Fixed altitude variance for 2D mode
	MinElev_deg           int8        // [deg] Minimum elevation for a GNSS satellite to be used in NAV
	DrLimit_s             byte        // [s] Reserved
	PDop                  uint16      // Position DOP mask to use
	TDop                  uint16      // Time DOP mask to use
	PAcc_m                uint16      // [m] Position accuracy mask
	TAcc_m                uint16      // [m] Time accuracy mask
	StaticHoldThresh_cm_s byte        // [cm/s] Static hold threshold
	DgnssTimeout_s        byte        // [s] DGNSS timeout
	CnoThreshNumSVs       byte        // Number of satellites required to have C/N0 above cnoThresh for a fix to be attempted
	CnoThresh_dbhz        byte        // [dBHz] C/N0 threshold for deciding whether to attempt a fix
	Reserved1             [2]byte     // Reserved
	StaticHoldMaxDist_m   uint16      // [m] Static hold distance threshold (before quitting static hold)
	UtcStandard           byte        // UTC standard to be used: 0: Automatic; receiver selects based on GNSS configuration (see GNSS time bases) 3: UTC as operated by the U.S. Naval Observatory (USNO); derived from GPS time 5: UTC as combined from multiple European laboratories; derived from Galileo time 6: UTC as operated by the former Soviet Union (SU); derived from GLONASS time 7: UTC as operated by the National Time Service Center (NTSC), China; derived from BeiDou time (not supported in protocol versions less than 16).
	Reserved2             [5]byte     // Reserved
}

func (CfgNav5) classID() uint16 { return 0x2406 }

type CfgNav5Mask uint16

const (
	CfgNav5Dyn            CfgNav5Mask = 0x1   // Apply dynamic model settings
	CfgNav5MinEl          CfgNav5Mask = 0x2   // Apply minimum elevation settings
	CfgNav5PosFixMode     CfgNav5Mask = 0x4   // Apply fix mode settings
	CfgNav5DrLim          CfgNav5Mask = 0x8   // Reserved
	CfgNav5PosMask        CfgNav5Mask = 0x10  // Apply position mask settings
	CfgNav5TimeMask       CfgNav5Mask = 0x20  // Apply time mask settings
	CfgNav5StaticHoldMask CfgNav5Mask = 0x40  // Apply static hold settings
	CfgNav5DgpsMask       CfgNav5Mask = 0x80  // Apply DGPS settings
	CfgNav5CnoThreshold   CfgNav5Mask = 0x100 // Apply CNO threshold settings (cnoThresh, cnoThreshNumSVs)
	CfgNav5Utc            CfgNav5Mask = 0x400 // Apply UTC settings (not supported in protocol versions less than 16).
)

func (v CfgNav5Mask) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Dyn")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "MinEl")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "PosFixMode")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "DrLim")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "PosMask")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "TimeMask")
	}
	v &^= 0x20
	if v&0x40 != 0 {
		b = append(b, "StaticHoldMask")
	}
	v &^= 0x40
	if v&0x80 != 0 {
		b = append(b, "DgpsMask")
	}
	v &^= 0x80
	if v&0x100 != 0 {
		b = append(b, "CnoThreshold")
	}
	v &^= 0x100
	if v&0x400 != 0 {
		b = append(b, "Utc")
	}
	v &^= 0x400
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-navx5 (2 versions)

// CfgNavx5 (Get/set) Navigation engine expert settings
// Class/Id 0x06 0x23 (40 bytes)
// (Polling will send back a version 3 message in protocol versions 19.2).
type CfgNavx5 struct {
	Version               uint16         // Message version (0x0002 for this version)
	Mask1                 CfgNavx5Mask1  // First parameters bitmask. Only the flagged parameters will be applied, unused bits must be set to 0.
	Mask2                 CfgNavx5Mask2  // Second parameters bitmask. Only the flagged parameters will be applied, unused bits must be set to 0.
	Reserved1             [2]byte        // Reserved
	MinSVs                byte           // [#SVs] Minimum number of satellites for navigation
	MaxSVs                byte           // [#SVs] Maximum number of satellites for navigation
	MinCNO_dbhz           byte           // [dBHz] Minimum satellite signal level for navigation
	Reserved2             byte           // Reserved
	IniFix3D              byte           // 1 = initial fix must be 3D
	Reserved3             [2]byte        // Reserved
	AckAiding             byte           // 1 = issue acknowledgements for assistance message input
	WknRollover           uint16         // GPS week rollover number; GPS week numbers will be set correctly from this week up to 1024 weeks after this week. Setting this to 0 reverts to firmware default.
	SigAttenCompMode_dbhz byte           // [dBHz] Only supported on certain products Permanently attenuated signal compensation (0 = disabled, 255 = automatic, 1..63 = maximum expected C/N0 value)
	Reserved4             byte           // Reserved
	Reserved5             [2]byte        // Reserved
	Reserved6             [2]byte        // Reserved
	UsePPP                byte           // 1 = use Precise Point Positioning (only available with the PPP product variant)
	AopCfg                CfgNavx5AopCfg // AssistNow Autonomous configuration
	Reserved7             [2]byte        // Reserved
	AopOrbMaxErr_m        uint16         // [m] Maximum acceptable (modeled) AssistNow Autonomous orbit error (valid range = 5..1000, or 0 = reset to firmware default)
	Reserved8             [4]byte        // Reserved
	Reserved9             [3]byte        // Reserved
	UseAdr                byte           // Only supported on certain products Enable/disable ADR/UDR sensor fusion (if 0: sensor fusion is disabled - if 1: sensor fusion is enabled).
}

func (CfgNavx5) classID() uint16 { return 0x2306 }

type CfgNavx5Mask1 uint16

const (
	CfgNavx5MinMax       CfgNavx5Mask1 = 0x4    // 1 = apply min/max SVs settings
	CfgNavx5MinCno       CfgNavx5Mask1 = 0x8    // 1 = apply minimum C/N0 setting
	CfgNavx5Initial3dfix CfgNavx5Mask1 = 0x40   // 1 = apply initial 3D fix settings
	CfgNavx5WknRoll      CfgNavx5Mask1 = 0x200  // 1 = apply GPS weeknumber rollover settings
	CfgNavx5AckAid       CfgNavx5Mask1 = 0x400  // 1 = apply assistance acknowledgement settings
	CfgNavx5Ppp          CfgNavx5Mask1 = 0x2000 // 1 = apply usePPP flag
	CfgNavx5Aop          CfgNavx5Mask1 = 0x4000 // 1 = apply aopCfg (useAOP flag) and aopOrbMaxErr settings (AssistNow Autonomous)
)

type CfgNavx5Mask2 uint32

const (
	CfgNavx5Adr          CfgNavx5Mask2 = 0x40 // Apply ADR/UDR sensor fusion on/off setting (useAdr flag)
	CfgNavx5SigAttenComp CfgNavx5Mask2 = 0x80 // Only supported on certain products Apply signal attenuation compensation feature settings
)

type CfgNavx5AopCfg byte

const (
	CfgNavx5UseAOP CfgNavx5AopCfg = 0x1 // 1 = enable AssistNow Autonomous
)

func (v CfgNavx5Mask1) String() string {
	var b []string
	if v&0x4 != 0 {
		b = append(b, "MinMax")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "MinCno")
	}
	v &^= 0x8
	if v&0x40 != 0 {
		b = append(b, "Initial3dfix")
	}
	v &^= 0x40
	if v&0x200 != 0 {
		b = append(b, "WknRoll")
	}
	v &^= 0x200
	if v&0x400 != 0 {
		b = append(b, "AckAid")
	}
	v &^= 0x400
	if v&0x2000 != 0 {
		b = append(b, "Ppp")
	}
	v &^= 0x2000
	if v&0x4000 != 0 {
		b = append(b, "Aop")
	}
	v &^= 0x4000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgNavx5Mask2) String() string {
	var b []string
	if v&0x40 != 0 {
		b = append(b, "Adr")
	}
	v &^= 0x40
	if v&0x80 != 0 {
		b = append(b, "SigAttenComp")
	}
	v &^= 0x80
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgNavx5AopCfg) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "UseAOP")
	}
	v &^= 0x1
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// CfgNavx51 (Get/set) Navigation engine expert settings
// Class/Id 0x06 0x23 (44 bytes)
// -
type CfgNavx51 struct {
	Version               uint16          // Message version (0x0003 for this version)
	Mask1                 CfgNavx51Mask1  // First parameters bitmask. Only the flagged parameters will be applied, unused bits must be set to 0.
	Mask2                 CfgNavx51Mask2  // Second parameters bitmask. Only the flagged parameters will be applied, unused bits must be set to 0.
	Reserved1             [2]byte         // Reserved
	MinSVs                byte            // [#SVs] Minimum number of satellites for navigation
	MaxSVs                byte            // [#SVs] Maximum number of satellites for navigation
	MinCNO_dbhz           byte            // [dBHz] Minimum satellite signal level for navigation
	Reserved2             byte            // Reserved
	IniFix3D              byte            // 1 = initial fix must be 3D
	Reserved3             [2]byte         // Reserved
	AckAiding             byte            // 1 = issue acknowledgements for assistance message input
	WknRollover           uint16          // GPS week rollover number; GPS week numbers will be set correctly from this week up to 1024 weeks after this week. Setting this to 0 reverts to firmware default.
	SigAttenCompMode_dbhz byte            // [dBHz] Only supported on certain products Permanently attenuated signal compensation (0 = disabled, 255 = automatic, 1..63 = maximum expected C/N0 value)
	Reserved4             byte            // Reserved
	Reserved5             [2]byte         // Reserved
	Reserved6             [2]byte         // Reserved
	UsePPP                byte            // 1 = use Precise Point Positioning (only available with the PPP product variant)
	AopCfg                CfgNavx51AopCfg // AssistNow Autonomous configuration
	Reserved7             [2]byte         // Reserved
	AopOrbMaxErr_m        uint16          // [m] Maximum acceptable (modeled) AssistNow Autonomous orbit error (valid range = 5..1000, or 0 = reset to firmware default)
	Reserved8             [4]byte         // Reserved
	Reserved9             [3]byte         // Reserved
	UseAdr                byte            // Only supported on certain products Enable/disable ADR/UDR sensor fusion (if 0: sensor fusion is disabled - if 1: sensor fusion is enabled).
	Reserved10            [2]byte         // Reserved
	Reserved11            [2]byte         // Reserved
}

func (CfgNavx51) classID() uint16 { return 0x2306 }

type CfgNavx51Mask1 uint16

const (
	CfgNavx51MinMax       CfgNavx51Mask1 = 0x4    // 1 = apply min/max SVs settings
	CfgNavx51MinCno       CfgNavx51Mask1 = 0x8    // 1 = apply minimum C/N0 setting
	CfgNavx51Initial3dfix CfgNavx51Mask1 = 0x40   // 1 = apply initial 3D fix settings
	CfgNavx51WknRoll      CfgNavx51Mask1 = 0x200  // 1 = apply GPS weeknumber rollover settings
	CfgNavx51AckAid       CfgNavx51Mask1 = 0x400  // 1 = apply assistance acknowledgement settings
	CfgNavx51Ppp          CfgNavx51Mask1 = 0x2000 // 1 = apply usePPP flag
	CfgNavx51Aop          CfgNavx51Mask1 = 0x4000 // 1 = apply aopCfg (useAOP flag) and aopOrbMaxErr settings (AssistNow Autonomous)
)

type CfgNavx51Mask2 uint32

const (
	CfgNavx51Adr          CfgNavx51Mask2 = 0x40 // Apply ADR/UDR sensor fusion on/off setting (useAdr flag)
	CfgNavx51SigAttenComp CfgNavx51Mask2 = 0x80 // Only supported on certain products Apply signal attenuation compensation feature settings
)

type CfgNavx51AopCfg byte

const (
	CfgNavx51UseAOP CfgNavx51AopCfg = 0x1 // 1 = enable AssistNow Autonomous
)

func (v CfgNavx51Mask1) String() string {
	var b []string
	if v&0x4 != 0 {
		b = append(b, "MinMax")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "MinCno")
	}
	v &^= 0x8
	if v&0x40 != 0 {
		b = append(b, "Initial3dfix")
	}
	v &^= 0x40
	if v&0x200 != 0 {
		b = append(b, "WknRoll")
	}
	v &^= 0x200
	if v&0x400 != 0 {
		b = append(b, "AckAid")
	}
	v &^= 0x400
	if v&0x2000 != 0 {
		b = append(b, "Ppp")
	}
	v &^= 0x2000
	if v&0x4000 != 0 {
		b = append(b, "Aop")
	}
	v &^= 0x4000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgNavx51Mask2) String() string {
	var b []string
	if v&0x40 != 0 {
		b = append(b, "Adr")
	}
	v &^= 0x40
	if v&0x80 != 0 {
		b = append(b, "SigAttenComp")
	}
	v &^= 0x80
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgNavx51AopCfg) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "UseAOP")
	}
	v &^= 0x1
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-nmea (3 versions)

// CfgNmea (Get/set) NMEA protocol configuration (deprecated)
// Class/Id 0x06 0x17 (4 bytes)
// This message version is provided for backwards compatibility only. Use the last version listed below instead (its fields are backwards compatible with this version, it just has extra fields defined). Get/set the NMEA protocol configuration. See section NMEA Protocol Configuration for a detailed description of the configuration effects on NMEA output.
type CfgNmea struct {
	Filter      CfgNmeaFilter // filter flags
	NmeaVersion byte          // 0x23: NMEA version 2.3 0x21: NMEA version 2.1
	NumSV       byte          // Maximum number of SVs to report per TalkerId. 0: unlimited 8: 8 SVs 12: 12 SVs 16: 16 SVs
	Flags       CfgNmeaFlags  // flags
}

func (CfgNmea) classID() uint16 { return 0x1706 }

type CfgNmeaFilter byte

const (
	CfgNmeaPosFilt       CfgNmeaFilter = 0x1  // Enable position output for failed or invalid fixes
	CfgNmeaMskPosFilt    CfgNmeaFilter = 0x2  // Enable position output for invalid fixes
	CfgNmeaTimeFilt      CfgNmeaFilter = 0x4  // Enable time output for invalid times
	CfgNmeaDateFilt      CfgNmeaFilter = 0x8  // Enable date output for invalid dates
	CfgNmeaGpsOnlyFilter CfgNmeaFilter = 0x10 // Restrict output to GPS satellites only
	CfgNmeaTrackFilt     CfgNmeaFilter = 0x20 // Enable COG output even if COG is frozen
)

type CfgNmeaFlags byte

const (
	CfgNmeaCompat   CfgNmeaFlags = 0x1 // enable compatibility mode. This might be needed for certain applications when customer's NMEA parser expects a fixed number of digits in position coordinates.
	CfgNmeaConsider CfgNmeaFlags = 0x2 // enable considering mode.
)

func (v CfgNmeaFilter) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "PosFilt")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "MskPosFilt")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "TimeFilt")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "DateFilt")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "GpsOnlyFilter")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "TrackFilt")
	}
	v &^= 0x20
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgNmeaFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Compat")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "Consider")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// CfgNmea1 (Get/set) NMEA protocol configuration V0 (deprecated)
// Class/Id 0x06 0x17 (12 bytes)
// This message version is provided for backwards compatibility only. Use the last version listed below instead (its fields are backwards compatible with this version, it just has extra fields defined). Get/set the NMEA protocol configuration. See section NMEA Protocol Configuration for a detailed description of the configuration effects on NMEA output.
type CfgNmea1 struct {
	Filter       CfgNmea1Filter       // filter flags
	NmeaVersion  byte                 // 0x23: NMEA version 2.3 0x21: NMEA version 2.1
	NumSV        byte                 // Maximum number of SVs to report per TalkerId. 0: unlimited 8: 8 SVs 12: 12 SVs 16: 16 SVs
	Flags        CfgNmea1Flags        // flags
	GnssToFilter CfgNmea1GnssToFilter // Filters out satellites based on their GNSS. If a bitfield is enabled, the corresponding satellites will be not output.
	SvNumbering  byte                 // Configures the display of satellites that do not have an NMEA-defined value. Note: this does not apply to satellites with an unknown ID. 0: Strict - Satellites are not output 1: Extended - Use proprietary numbering (see Satellite Numbering)
	MainTalkerId byte                 // By default the main Talker ID (i.e. the Talker ID used for all messages other than GSV) is determined by the GNSS assignment of the receiver's channels (see UBX-CFG-GNSS). This field enables the main Talker ID to be overridden. 0: Main Talker ID is not overridden 1: Set main Talker ID to 'GP' 2: Set main Talker ID to 'GL' 3: Set main Talker ID to 'GN' 4: Set main Talker ID to 'GA' 5: Set main Talker ID to 'GB' 6: Set main Talker ID to 'GQ' (available in NMEA 4.11 or later)
	GsvTalkerId  byte                 // By default the Talker ID for GSV messages is GNSS-specific (as defined by NMEA). This field enables the GSV Talker ID to be overridden. 0: Use GNSS-specific Talker ID (as defined by NMEA) 1: Use the main Talker ID
	Version      byte                 // Message version (0x00 for this version)
}

func (CfgNmea1) classID() uint16 { return 0x1706 }

type CfgNmea1Filter byte

const (
	CfgNmea1PosFilt       CfgNmea1Filter = 0x1  // Enable position output for failed or invalid fixes
	CfgNmea1MskPosFilt    CfgNmea1Filter = 0x2  // Enable position output for invalid fixes
	CfgNmea1TimeFilt      CfgNmea1Filter = 0x4  // Enable time output for invalid times
	CfgNmea1DateFilt      CfgNmea1Filter = 0x8  // Enable date output for invalid dates
	CfgNmea1GpsOnlyFilter CfgNmea1Filter = 0x10 // Restrict output to GPS satellites only
	CfgNmea1TrackFilt     CfgNmea1Filter = 0x20 // Enable COG output even if COG is frozen
)

type CfgNmea1Flags byte

const (
	CfgNmea1Compat   CfgNmea1Flags = 0x1 // enable compatibility mode. This might be needed for certain applications when customer's NMEA parser expects a fixed number of digits in position coordinates.
	CfgNmea1Consider CfgNmea1Flags = 0x2 // enable considering mode.
)

type CfgNmea1GnssToFilter uint32

const (
	CfgNmea1Gps     CfgNmea1GnssToFilter = 0x1  // Disable reporting of GPS satellites
	CfgNmea1Sbas    CfgNmea1GnssToFilter = 0x2  // Disable reporting of SBAS satellites
	CfgNmea1Galileo CfgNmea1GnssToFilter = 0x4  // Disable reporting of Galileo satellites
	CfgNmea1Qzss    CfgNmea1GnssToFilter = 0x10 // Disable reporting of QZSS satellites
	CfgNmea1Glonass CfgNmea1GnssToFilter = 0x20 // Disable reporting of GLONASS satellites
	CfgNmea1Beidou  CfgNmea1GnssToFilter = 0x40 // Disable reporting of BeiDou satellites
)

func (v CfgNmea1Filter) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "PosFilt")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "MskPosFilt")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "TimeFilt")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "DateFilt")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "GpsOnlyFilter")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "TrackFilt")
	}
	v &^= 0x20
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgNmea1Flags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Compat")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "Consider")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgNmea1GnssToFilter) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Gps")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "Sbas")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "Galileo")
	}
	v &^= 0x4
	if v&0x10 != 0 {
		b = append(b, "Qzss")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "Glonass")
	}
	v &^= 0x20
	if v&0x40 != 0 {
		b = append(b, "Beidou")
	}
	v &^= 0x40
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// CfgNmea2 (Get/set) Extended NMEA protocol configuration V1
// Class/Id 0x06 0x17 (20 bytes)
// Get/set the NMEA protocol configuration. See section NMEA Protocol Configuration for a detailed description of the configuration effects on NMEA output.
type CfgNmea2 struct {
	Filter       CfgNmea2Filter       // filter flags
	NmeaVersion  byte                 // 0x4b: NMEA version 4.11 (not available in all products) 0x41: NMEA version 4.10 (not available in all products) 0x40: NMEA version 4.0 (not available in all products) 0x23: NMEA version 2.3 0x21: NMEA version 2.1
	NumSV        byte                 // Maximum number of SVs to report per TalkerId. 0: unlimited 8: 8 SVs 12: 12 SVs 16: 16 SVs
	Flags        CfgNmea2Flags        // flags
	GnssToFilter CfgNmea2GnssToFilter // Filters out satellites based on their GNSS. If a bitfield is enabled, the corresponding satellites will be not output.
	SvNumbering  byte                 // Configures the display of satellites that do not have an NMEA-defined value. Note: this does not apply to satellites with an unknown ID. 0: Strict - Satellites are not output 1: Extended - Use proprietary numbering (see Satellite Numbering)
	MainTalkerId byte                 // By default the main Talker ID (i.e. the Talker ID used for all messages other than GSV) is determined by the GNSS assignment of the receiver's channels (see UBX-CFG-GNSS). This field enables the main Talker ID to be overridden. 0: Main Talker ID is not overridden 1: Set main Talker ID to 'GP' 2: Set main Talker ID to 'GL' 3: Set main Talker ID to 'GN' 4: Set main Talker ID to 'GA' 5: Set main Talker ID to 'GB' 6: Set main Talker ID to 'GQ' (available in NMEA 4.11 or later)
	GsvTalkerId  byte                 // By default the Talker ID for GSV messages is GNSS-specific (as defined by NMEA). This field enables the GSV Talker ID to be overridden. 0: Use GNSS-specific Talker ID (as defined by NMEA) 1: Use the main Talker ID
	Version      byte                 // Message version (0x01 for this version)
	BdsTalkerId  [2]byte              // Sets the two characters that should be used for the BeiDou Talker ID. If these are set to zero, then the default BeiDou Talker ID will be used.
	Reserved1    [6]byte              // Reserved
}

func (CfgNmea2) classID() uint16 { return 0x1706 }

type CfgNmea2Filter byte

const (
	CfgNmea2PosFilt       CfgNmea2Filter = 0x1  // Enable position output for failed or invalid fixes
	CfgNmea2MskPosFilt    CfgNmea2Filter = 0x2  // Enable position output for invalid fixes
	CfgNmea2TimeFilt      CfgNmea2Filter = 0x4  // Enable time output for invalid times
	CfgNmea2DateFilt      CfgNmea2Filter = 0x8  // Enable date output for invalid dates
	CfgNmea2GpsOnlyFilter CfgNmea2Filter = 0x10 // Restrict output to GPS satellites only
	CfgNmea2TrackFilt     CfgNmea2Filter = 0x20 // Enable COG output even if COG is frozen
)

type CfgNmea2Flags byte

const (
	CfgNmea2Compat   CfgNmea2Flags = 0x1 // enable compatibility mode. This might be needed for certain applications when customer's NMEA parser expects a fixed number of digits in position coordinates.
	CfgNmea2Consider CfgNmea2Flags = 0x2 // enable considering mode.
	CfgNmea2Limit82  CfgNmea2Flags = 0x4 // enable strict limit to 82 characters maximum.
	CfgNmea2HighPrec CfgNmea2Flags = 0x8 // enable high precision mode. This flag cannot be set in conjunction with either compatibility mode or Limit82 mode (not supported in protocol versions less than 20.01).
)

type CfgNmea2GnssToFilter uint32

const (
	CfgNmea2Gps     CfgNmea2GnssToFilter = 0x1  // Disable reporting of GPS satellites
	CfgNmea2Sbas    CfgNmea2GnssToFilter = 0x2  // Disable reporting of SBAS satellites
	CfgNmea2Galileo CfgNmea2GnssToFilter = 0x4  // Disable reporting of Galileo satellites
	CfgNmea2Qzss    CfgNmea2GnssToFilter = 0x10 // Disable reporting of QZSS satellites
	CfgNmea2Glonass CfgNmea2GnssToFilter = 0x20 // Disable reporting of GLONASS satellites
	CfgNmea2Beidou  CfgNmea2GnssToFilter = 0x40 // Disable reporting of BeiDou satellites
)

func (v CfgNmea2Filter) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "PosFilt")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "MskPosFilt")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "TimeFilt")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "DateFilt")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "GpsOnlyFilter")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "TrackFilt")
	}
	v &^= 0x20
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgNmea2Flags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Compat")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "Consider")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "Limit82")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "HighPrec")
	}
	v &^= 0x8
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgNmea2GnssToFilter) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Gps")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "Sbas")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "Galileo")
	}
	v &^= 0x4
	if v&0x10 != 0 {
		b = append(b, "Qzss")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "Glonass")
	}
	v &^= 0x20
	if v&0x40 != 0 {
		b = append(b, "Beidou")
	}
	v &^= 0x40
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-odo

// CfgOdo (Get/set) Odometer, low-speed COG engine settings
// Class/Id 0x06 0x1e (20 bytes)
// This feature is not supported for the FTS product variant. -
type CfgOdo struct {
	Version           byte         // Message version (0x00 for this version)
	Reserved1         [3]byte      // Reserved
	Flags             CfgOdoFlags  // Odometer/Low-speed COG filter flags
	OdoCfg            CfgOdoOdoCfg // Odometer filter settings
	Reserved2         [6]byte      // Reserved
	CogMaxSpeed_m_se1 byte         // [1e-1 m/s] Speed below which course-over-ground (COG) is computed with the low-speed COG filter
	CogMaxPosAcc_m    byte         // [m] Maximum acceptable position accuracy for computing COG with the low-speed COG filter
	Reserved3         [2]byte      // Reserved
	VelLpGain         byte         // Velocity low-pass filter level, range 0..255
	CogLpGain         byte         // COG low-pass filter level (at speed < 8 m/s), range 0..255
	Reserved4         [2]byte      // Reserved
}

func (CfgOdo) classID() uint16 { return 0x1e06 }

type CfgOdoFlags byte

const (
	CfgOdoUseODO   CfgOdoFlags = 0x1 // Odometer-enabled flag
	CfgOdoUseCOG   CfgOdoFlags = 0x2 // Low-speed COG filter enabled flag
	CfgOdoOutLPVel CfgOdoFlags = 0x4 // Output low-pass filtered velocity flag
	CfgOdoOutLPCog CfgOdoFlags = 0x8 // Output low-pass filtered heading (COG) flag
)

type CfgOdoOdoCfg byte

const (
	CfgOdoProfile CfgOdoOdoCfg = 0x7 // Profile type (0=running, 1=cycling, 2=swimming, 3=car, 4=custom)
)

func (v CfgOdoFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "UseODO")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "UseCOG")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "OutLPVel")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "OutLPCog")
	}
	v &^= 0x8
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgOdoOdoCfg) String() string {
	var b []string
	if v&0x7 != 0 {
		b = append(b, fmt.Sprintf("Profile<%b>", (v&0x7)>>0))
	}
	v &^= 0x7
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-pm2 (2 versions)

// CfgPm2 (Get/set) Extended power management configuration
// Class/Id 0x06 0x3b (44 bytes)
// This feature is not supported for either the ADR or FTS products. -
type CfgPm2 struct {
	Version              byte        // Message version (0x01 for this version)
	Reserved1            byte        // Reserved
	MaxStartupStateDur_s byte        // [s] Maximum time to spend in Acquisition state. If 0: bound disabled (see maxStartupStateDur) (not supported in protocol versions less than 17).
	Reserved2            byte        // Reserved
	Flags                CfgPm2Flags // PSM configuration flags
	UpdatePeriod_ms      uint32      // [ms] Position update period. If set to 0, the receiver will never retry a fix and it will wait for external events
	SearchPeriod_ms      uint32      // [ms] Acquisition retry period if previously failed. If set to 0, the receiver will never retry a startup
	GridOffset_ms        uint32      // [ms] Grid offset relative to GPS start of week
	OnTime_s             uint16      // [s] Time to stay in Tracking state
	MinAcqTime_s         uint16      // [s] minimal search time
	Reserved3            [20]byte    // Reserved
}

func (CfgPm2) classID() uint16 { return 0x3b06 }

type CfgPm2Flags uint32

const (
	CfgPm2ExtintSel     CfgPm2Flags = 0x10    // EXTINT pin select 0 EXTINT0 1 EXTINT1
	CfgPm2ExtintWake    CfgPm2Flags = 0x20    // EXTINT pin control 0 disabled 1 enabled, keep receiver awake as long as selected EXTINT pin is 'high'
	CfgPm2ExtintBackup  CfgPm2Flags = 0x40    // EXTINT pin control 0 disabled 1 enabled, force receiver into BACKUP mode when selected EXTINT pin is 'low'
	CfgPm2LimitPeakCurr CfgPm2Flags = 0x300   // Limit peak current 00 disabled 01 enabled, peak current is limited 10 reserved 11 reserved
	CfgPm2WaitTimeFix   CfgPm2Flags = 0x400   // Wait for Timefix (see waitTimeFix) 0 wait for normal fix OK before starting on time 1 wait for time fix OK before starting on time
	CfgPm2UpdateRTC     CfgPm2Flags = 0x800   // Update Real Time Clock (see updateRTC) 0 do not wake up to update RTC. RTC is updated during normal on-time. 1 update RTC. The receiver adds extra wake-up cycles to update the RTC.
	CfgPm2UpdateEPH     CfgPm2Flags = 0x1000  // Update Ephemeris (see updateEPH) 0 do not wake up to update Ephemeris data 1 update Ephemeris. The receiver adds extra wake-up cycles to update the Ephemeris data
	CfgPm2DoNotEnterOff CfgPm2Flags = 0x10000 // Behavior of receiver in case of no fix (see doNotEnterOff) 0 receiver enters Inactive) Awaiting next search state 1 receiver does not enter (Inactive) Awaiting next search state but keeps trying to acquire a fix instead
	CfgPm2Mode          CfgPm2Flags = 0x60000 // Mode of operation (see mode) 00 ON/OFF operation (PSMOO) 01 cyclic tracking operation (PSMCT) 10 reserved 11 reserved
)

func (v CfgPm2Flags) String() string {
	var b []string
	if v&0x10 != 0 {
		b = append(b, "ExtintSel")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "ExtintWake")
	}
	v &^= 0x20
	if v&0x40 != 0 {
		b = append(b, "ExtintBackup")
	}
	v &^= 0x40
	if v&0x300 != 0 {
		b = append(b, fmt.Sprintf("LimitPeakCurr<%b>", (v&0x300)>>8))
	}
	v &^= 0x300
	if v&0x400 != 0 {
		b = append(b, "WaitTimeFix")
	}
	v &^= 0x400
	if v&0x800 != 0 {
		b = append(b, "UpdateRTC")
	}
	v &^= 0x800
	if v&0x1000 != 0 {
		b = append(b, "UpdateEPH")
	}
	v &^= 0x1000
	if v&0x10000 != 0 {
		b = append(b, "DoNotEnterOff")
	}
	v &^= 0x10000
	if v&0x60000 != 0 {
		b = append(b, fmt.Sprintf("Mode<%b>", (v&0x60000)>>17))
	}
	v &^= 0x60000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// CfgPm21 (Get/set) Extended power management configuration
// Class/Id 0x06 0x3b (48 bytes)
// This feature is not supported for either the ADR or FTS products. -
type CfgPm21 struct {
	Version               byte         // Message version (0x02 for this version) Note: the message version number is the same as for protocol versions 18 up to 22; please select correct message version based on the protocol version supported by your firmware.
	Reserved1             byte         // Reserved
	MaxStartupStateDur_s  byte         // [s] Maximum time to spend in Acquisition state. If 0: bound disabled. (see maxStartupStateDur) (not supported in protocol versions 23 to 23.01).
	Reserved2             byte         // Reserved
	Flags                 CfgPm21Flags // PSM configuration flags
	UpdatePeriod_ms       uint32       // [ms] Position update period. If set to 0, the receiver will never retry a fix and it will wait for external events .
	SearchPeriod_ms       uint32       // [ms] Acquisition retry period if previously failed. If set to 0, the receiver will never retry a startup. (not supported in protocol versions 23 to 23.01).
	GridOffset_ms         uint32       // [ms] Grid offset relative to GPS start of week (not supported in protocol versions 23 to 23.01).
	OnTime_s              uint16       // [s] Time to stay in Tracking state (not supported in protocol versions 23 to 23.01).
	MinAcqTime_s          uint16       // [s] Minimal search time
	Reserved3             [20]byte     // Reserved
	ExtintInactivityMs_ms uint32       // [ms] inactivity time out on EXTINT pin if enabled
}

func (CfgPm21) classID() uint16 { return 0x3b06 }

type CfgPm21Flags uint32

const (
	CfgPm21OptTarget      CfgPm21Flags = 0xe     // Optimization target 000 performance (default) 001 power save 010 reserved 011 reserved 100 reserved 101 reserved 110 reserved 111 reserved
	CfgPm21ExtintSel      CfgPm21Flags = 0x10    // EXTINT pin select 0 EXTINT0 1 EXTINT1
	CfgPm21ExtintWake     CfgPm21Flags = 0x20    // EXTINT pin control 0 disabled 1 enabled, keep receiver awake as long as selected EXTINT pin is 'high'
	CfgPm21ExtintBackup   CfgPm21Flags = 0x40    // EXTINT pin control 0 disabled 1 enabled, force receiver into BACKUP mode when selected EXTINT pin is 'low'
	CfgPm21ExtintInactive CfgPm21Flags = 0x80    // EXTINT pin control 0 disabled 1 enabled, force backup in case EXTINT pin is inactive for time longer than extintIncactivityMs
	CfgPm21LimitPeakCurr  CfgPm21Flags = 0x300   // Limit peak current 00 disabled 01 enabled, peak current is limited 10 reserved 11 reserved
	CfgPm21WaitTimeFix    CfgPm21Flags = 0x400   // Wait for Timefix (see waitTimeFix) 0 wait for normal fix OK before starting on time 1 wait for time fix OK before starting on time (not supported in protocol versions 23 to 23.01).
	CfgPm21UpdateRTC      CfgPm21Flags = 0x800   // Update real time clock (see updateRTC) 0 do not wake up to update RTC. RTC is updated during normal on-time. 1 update RTC. The receiver adds extra wake-up cycles to update the RTC. (not supported in protocol versions 23 to 23.01, and 32+).
	CfgPm21UpdateEPH      CfgPm21Flags = 0x1000  // Update ephemeris (see updateEPH) 0 do not wake up to update Ephemeris data 1 update Ephemeris. The receiver adds extra wake-up cycles to update the Ephemeris data.
	CfgPm21DoNotEnterOff  CfgPm21Flags = 0x10000 // Behavior of receiver in case of no fix Behavior of receiver in case of no fix (see doNotEnterOff) 0 receiver enters (Inactive) Awaiting next search state 1 receiver does not enter (Inactive) Awaiting next search state but keeps trying to acquire a fix instead (not supported in protocol versions 23 to 23.01).
	CfgPm21Mode           CfgPm21Flags = 0x60000 // Mode of operation (see mode) 00 ON/OFF operation (PSMOO) (not supported in protocol versions 23 to 23.01) 01 cyclic tracking operation (PSMCT) 10 reserved 11 reserved
)

func (v CfgPm21Flags) String() string {
	var b []string
	if v&0xe != 0 {
		b = append(b, fmt.Sprintf("OptTarget<%b>", (v&0xe)>>1))
	}
	v &^= 0xe
	if v&0x10 != 0 {
		b = append(b, "ExtintSel")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "ExtintWake")
	}
	v &^= 0x20
	if v&0x40 != 0 {
		b = append(b, "ExtintBackup")
	}
	v &^= 0x40
	if v&0x80 != 0 {
		b = append(b, "ExtintInactive")
	}
	v &^= 0x80
	if v&0x300 != 0 {
		b = append(b, fmt.Sprintf("LimitPeakCurr<%b>", (v&0x300)>>8))
	}
	v &^= 0x300
	if v&0x400 != 0 {
		b = append(b, "WaitTimeFix")
	}
	v &^= 0x400
	if v&0x800 != 0 {
		b = append(b, "UpdateRTC")
	}
	v &^= 0x800
	if v&0x1000 != 0 {
		b = append(b, "UpdateEPH")
	}
	v &^= 0x1000
	if v&0x10000 != 0 {
		b = append(b, "DoNotEnterOff")
	}
	v &^= 0x10000
	if v&0x60000 != 0 {
		b = append(b, fmt.Sprintf("Mode<%b>", (v&0x60000)>>17))
	}
	v &^= 0x60000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-pms

// CfgPms (Get/set) Power mode setup
// Class/Id 0x06 0x86 (8 bytes)
// Using UBX-CFG-PMS to set Super-E mode to 1, 2 or 4 Hz navigation rates sets minAcqTime to 180 s instead of the default 300 s in protocol version 23.01.
type CfgPms struct {
	Version         byte    // Message version (0x00 for this version)
	PowerSetupValue byte    // Power setup value 0x00 = Full power 0x01 = Balanced 0x02 = Interval 0x03 = Aggressive with 1 Hz 0x04 = Aggressive with 2 Hz 0x05 = Aggressive with 4 Hz 0xFF = Invalid (only when polling)
	Period_s        uint16  // [s] Position update period and search period. Recommended minimum period is 10 s, although the receiver accepts any value bigger than 5 s. Only valid when powerSetupValue set to Interval, otherwise must be set to '0'.
	OnTime_s        uint16  // [s] Duration of the ON phase, must be smaller than the period. Only valid when powerSetupValue set to Interval, otherwise must be set to '0'.
	Reserved1       [2]byte // Reserved
}

func (CfgPms) classID() uint16 { return 0x8606 }

// Message ubx-cfg-prt (5 versions)

// CfgPrt (Poll Request) Polls the configuration for one I/O port
// Class/Id 0x06 0x00 (1 bytes)
// Sending this message with a port ID as payload results in having the receiver return the configuration for the specified port.
type CfgPrt struct {
	PortID byte // Port identifier number (see the other versions of CFG-PRT for valid values)
}

func (CfgPrt) classID() uint16 { return 0x0006 }

// CfgPrt1 (Get/set) Port configuration for UART ports
// Class/Id 0x06 0x00 (20 bytes)
// Several configurations can be concatenated to one input message. In this case the payload length can be a multiple of the normal length (see the other versions of CFG-PRT). Output messages from the module contain only one configuration unit. Note that this message can affect baud rate and other transmission parameters. Because there may be messages queued for transmission there may be uncertainty about which protocol applies to such messages. In addition a message currently in transmission may be corrupted by a protocol change. Host data reception parameters may have to be changed to be able to receive future messages, including the acknowledge message resulting from the CFG-PRT message.
type CfgPrt1 struct {
	PortID          byte                `stf:"default"` // Port identifier number (see Integration manual for valid UART port IDs)
	Reserved1       byte                // Reserved
	TxReady         CfgPrt1TxReady      // TX ready PIN configuration
	Mode            CfgPrt1Mode         // A bit mask describing the UART mode
	BaudRate_bits_s uint32              // [Bits/s] Baud rate in bits/second
	InProtoMask     CfgPrt1InProtoMask  // A mask describing which input protocols are active. Each bit of this mask is used for a protocol. Through that, multiple protocols can be defined on a single port.
	OutProtoMask    CfgPrt1OutProtoMask // A mask describing which output protocols are active. Each bit of this mask is used for a protocol. Through that, multiple protocols can be defined on a single port.
	Flags           CfgPrt1Flags        // Flags bit mask
	Reserved2       [2]byte             // Reserved
}

func (CfgPrt1) classID() uint16 { return 0x0006 }

type CfgPrt1TxReady uint16

const (
	CfgPrt1En    CfgPrt1TxReady = 0x1    // Enable TX ready feature for this port
	CfgPrt1Pol   CfgPrt1TxReady = 0x2    // Polarity 0 High-active 1 Low-active
	CfgPrt1Pin   CfgPrt1TxReady = 0x7c   // PIO to be used (must not be in use by another function)
	CfgPrt1Thres CfgPrt1TxReady = 0xff80 // Threshold The given threshold is multiplied by 8 bytes. The TX ready PIN goes active after >= thres*8 bytes are pending for the port and going inactive after the last pending bytes have been written to hardware (0-4 bytes before end of stream). 0x000 no threshold 0x001 8byte 0x002 16byte ... 0x1FE 4080byte 0x1FF 4088byte
)

type CfgPrt1Mode uint32

const (
	CfgPrt1CharLen   CfgPrt1Mode = 0xc0   // Character length 00 5bit (not supported) 01 6bit (not supported) 10 7bit (supported only with parity) 11 8bit
	CfgPrt1Parity    CfgPrt1Mode = 0xe00  // 000 Even parity 001 Odd parity 10X No parity X1X Reserved
	CfgPrt1NStopBits CfgPrt1Mode = 0x3000 // Number of Stop bits 00 1 Stop bit 01 1.5 Stop bit 10 2 Stop bit 11 0.5 Stop bit
)

type CfgPrt1InProtoMask uint16

const (
	CfgPrt1InUbx   CfgPrt1InProtoMask = 0x1  // UBX protocol
	CfgPrt1InNmea  CfgPrt1InProtoMask = 0x2  // NMEA protocol
	CfgPrt1InRtcm  CfgPrt1InProtoMask = 0x4  // RTCM2 protocol
	CfgPrt1InRtcm3 CfgPrt1InProtoMask = 0x20 // RTCM3 protocol (not supported in protocol versions less than 20)
)

type CfgPrt1OutProtoMask uint16

const (
	CfgPrt1OutUbx   CfgPrt1OutProtoMask = 0x1  // UBX protocol
	CfgPrt1OutNmea  CfgPrt1OutProtoMask = 0x2  // NMEA protocol
	CfgPrt1OutRtcm3 CfgPrt1OutProtoMask = 0x20 // RTCM3 protocol (not supported in protocol versions less than 20)
)

type CfgPrt1Flags uint16

const (
	CfgPrt1ExtendedTxTimeout CfgPrt1Flags = 0x2 // Extended TX timeout: if set, the port will time out if allocated TX memory >=4 kB and no activity for 1. 5 s. If not set the port will time out if no activity for 1.5 s regardless on the amount of allocated TX memory .
)

func (v CfgPrt1TxReady) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "En")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "Pol")
	}
	v &^= 0x2
	if v&0x7c != 0 {
		b = append(b, fmt.Sprintf("Pin<%b>", (v&0x7c)>>2))
	}
	v &^= 0x7c
	if v&0xff80 != 0 {
		b = append(b, fmt.Sprintf("Thres<%b>", (v&0xff80)>>7))
	}
	v &^= 0xff80
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgPrt1Mode) String() string {
	var b []string
	if v&0xc0 != 0 {
		b = append(b, fmt.Sprintf("CharLen<%b>", (v&0xc0)>>6))
	}
	v &^= 0xc0
	if v&0xe00 != 0 {
		b = append(b, fmt.Sprintf("Parity<%b>", (v&0xe00)>>9))
	}
	v &^= 0xe00
	if v&0x3000 != 0 {
		b = append(b, fmt.Sprintf("NStopBits<%b>", (v&0x3000)>>12))
	}
	v &^= 0x3000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgPrt1InProtoMask) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "InUbx")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "InNmea")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "InRtcm")
	}
	v &^= 0x4
	if v&0x20 != 0 {
		b = append(b, "InRtcm3")
	}
	v &^= 0x20
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgPrt1OutProtoMask) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "OutUbx")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "OutNmea")
	}
	v &^= 0x2
	if v&0x20 != 0 {
		b = append(b, "OutRtcm3")
	}
	v &^= 0x20
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgPrt1Flags) String() string {
	var b []string
	if v&0x2 != 0 {
		b = append(b, "ExtendedTxTimeout")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// CfgPrt2 (Get/set) Port configuration for USB port
// Class/Id 0x06 0x00 (20 bytes)
// Several configurations can be concatenated to one input message. In this case the payload length can be a multiple of the normal length (see the other versions of CFG-PRT). Output messages from the module contain only one configuration unit.
type CfgPrt2 struct {
	PortID       byte                `stf:"3"` // Port identifier number (= 3 for USB port)
	Reserved1    byte                // Reserved
	TxReady      CfgPrt2TxReady      // TX ready PIN configuration
	Reserved2    [8]byte             // Reserved
	InProtoMask  CfgPrt2InProtoMask  // A mask describing which input protocols are active. Each bit of this mask is used for a protocol. Through that, multiple protocols can be defined on a single port.
	OutProtoMask CfgPrt2OutProtoMask // A mask describing which output protocols are active. Each bit of this mask is used for a protocol. Through that, multiple protocols can be defined on a single port.
	Reserved3    [2]byte             // Reserved
	Reserved4    [2]byte             // Reserved
}

func (CfgPrt2) classID() uint16 { return 0x0006 }

type CfgPrt2TxReady uint16

const (
	CfgPrt2En    CfgPrt2TxReady = 0x1    // Enable TX ready feature for this port
	CfgPrt2Pol   CfgPrt2TxReady = 0x2    // Polarity 0 High-active 1 Low-active
	CfgPrt2Pin   CfgPrt2TxReady = 0x7c   // PIO to be used (must not be in use by another function)
	CfgPrt2Thres CfgPrt2TxReady = 0xff80 // Threshold The given threshold is multiplied by 8 bytes. The TX ready PIN goes active after >= thres*8 bytes are pending for the port and going inactive after the last pending bytes have been written to hardware (0-4 bytes before end of stream). 0x000 no threshold 0x001 8byte 0x002 16byte ... 0x1FE 4080byte 0x1FF 4088byte
)

type CfgPrt2InProtoMask uint16

const (
	CfgPrt2InUbx   CfgPrt2InProtoMask = 0x1  // UBX protocol
	CfgPrt2InNmea  CfgPrt2InProtoMask = 0x2  // NMEA protocol
	CfgPrt2InRtcm  CfgPrt2InProtoMask = 0x4  // RTCM2 protocol
	CfgPrt2InRtcm3 CfgPrt2InProtoMask = 0x20 // RTCM3 protocol (not supported in protocol versions less than 20)
)

type CfgPrt2OutProtoMask uint16

const (
	CfgPrt2OutUbx   CfgPrt2OutProtoMask = 0x1  // UBX protocol
	CfgPrt2OutNmea  CfgPrt2OutProtoMask = 0x2  // NMEA protocol
	CfgPrt2OutRtcm3 CfgPrt2OutProtoMask = 0x20 // RTCM3 protocol (not supported in protocol versions less than 20)
)

func (v CfgPrt2TxReady) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "En")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "Pol")
	}
	v &^= 0x2
	if v&0x7c != 0 {
		b = append(b, fmt.Sprintf("Pin<%b>", (v&0x7c)>>2))
	}
	v &^= 0x7c
	if v&0xff80 != 0 {
		b = append(b, fmt.Sprintf("Thres<%b>", (v&0xff80)>>7))
	}
	v &^= 0xff80
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgPrt2InProtoMask) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "InUbx")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "InNmea")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "InRtcm")
	}
	v &^= 0x4
	if v&0x20 != 0 {
		b = append(b, "InRtcm3")
	}
	v &^= 0x20
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgPrt2OutProtoMask) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "OutUbx")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "OutNmea")
	}
	v &^= 0x2
	if v&0x20 != 0 {
		b = append(b, "OutRtcm3")
	}
	v &^= 0x20
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// CfgPrt3 (Get/set) Port configuration for SPI port
// Class/Id 0x06 0x00 (20 bytes)
// Several configurations can be concatenated to one input message. In this case the payload length can be a multiple of the normal length. Output messages from the module contain only one configuration unit.
type CfgPrt3 struct {
	PortID       byte                `stf:"4"` // Port identifier number (= 4 for SPI port)
	Reserved1    byte                // Reserved
	TxReady      CfgPrt3TxReady      // TX ready PIN configuration
	Mode         CfgPrt3Mode         // SPI Mode Flags
	Reserved2    [4]byte             // Reserved
	InProtoMask  CfgPrt3InProtoMask  // A mask describing which input protocols are active. Each bit of this mask is used for a protocol. Through that, multiple protocols can be defined on a single port. (The bitfield inRtcm3 is not supported in protocol versions less than 20)
	OutProtoMask CfgPrt3OutProtoMask // A mask describing which output protocols are active. Each bit of this mask is used for a protocol. Through that, multiple protocols can be defined on a single port. (The bitfield outRtcm3 is not supported in protocol versions less than 20)
	Flags        CfgPrt3Flags        // Flags bit mask
	Reserved3    [2]byte             // Reserved
}

func (CfgPrt3) classID() uint16 { return 0x0006 }

type CfgPrt3TxReady uint16

const (
	CfgPrt3En    CfgPrt3TxReady = 0x1    // Enable TX ready feature for this port
	CfgPrt3Pol   CfgPrt3TxReady = 0x2    // Polarity 0 High-active 1 Low-active
	CfgPrt3Pin   CfgPrt3TxReady = 0x7c   // PIO to be used (must not be in use by another function)
	CfgPrt3Thres CfgPrt3TxReady = 0xff80 // Threshold The given threshold is multiplied by 8 bytes. The TX ready PIN goes active after >= thres*8 bytes are pending for the port and going inactive after the last pending bytes have been written to hardware (0-4 bytes before end of stream). 0x000 no threshold 0x001 8byte 0x002 16byte ... 0x1FE 4080byte 0x1FF 4088byte
)

type CfgPrt3Mode uint32

const (
	CfgPrt3SpiMode CfgPrt3Mode = 0x6    // 00 SPI Mode 0: CPOL = 0, CPHA = 0 01 SPI Mode 1: CPOL = 0, CPHA = 1 10 SPI Mode 2: CPOL = 1, CPHA = 0 11 SPI Mode 3: CPOL = 1, CPHA = 1
	CfgPrt3FfCnt   CfgPrt3Mode = 0x3f00 // Number of bytes containing 0xFF to receive before switching off reception. Range: 0 (mechanism off) - 63
)

type CfgPrt3InProtoMask uint16

const (
	CfgPrt3InUbx   CfgPrt3InProtoMask = 0x1  // UBX protocol
	CfgPrt3InNmea  CfgPrt3InProtoMask = 0x2  // NMEA protocol
	CfgPrt3InRtcm  CfgPrt3InProtoMask = 0x4  // RTCM2 protocol
	CfgPrt3InRtcm3 CfgPrt3InProtoMask = 0x20 // RTCM3 protocol (not supported in protocol versions less than 20)
)

type CfgPrt3OutProtoMask uint16

const (
	CfgPrt3OutUbx   CfgPrt3OutProtoMask = 0x1  // UBX protocol
	CfgPrt3OutNmea  CfgPrt3OutProtoMask = 0x2  // NMEA protocol
	CfgPrt3OutRtcm3 CfgPrt3OutProtoMask = 0x20 // RTCM3 protocol (not supported in protocol versions less than 20)
)

type CfgPrt3Flags uint16

const (
	CfgPrt3ExtendedTxTimeout CfgPrt3Flags = 0x2 // Extended TX timeout: if set, the port will time out if allocated TX memory >=4 kB and no activity for 1. 5 s.
)

func (v CfgPrt3TxReady) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "En")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "Pol")
	}
	v &^= 0x2
	if v&0x7c != 0 {
		b = append(b, fmt.Sprintf("Pin<%b>", (v&0x7c)>>2))
	}
	v &^= 0x7c
	if v&0xff80 != 0 {
		b = append(b, fmt.Sprintf("Thres<%b>", (v&0xff80)>>7))
	}
	v &^= 0xff80
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgPrt3Mode) String() string {
	var b []string
	if v&0x6 != 0 {
		b = append(b, fmt.Sprintf("SpiMode<%b>", (v&0x6)>>1))
	}
	v &^= 0x6
	if v&0x3f00 != 0 {
		b = append(b, fmt.Sprintf("FfCnt<%b>", (v&0x3f00)>>8))
	}
	v &^= 0x3f00
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgPrt3InProtoMask) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "InUbx")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "InNmea")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "InRtcm")
	}
	v &^= 0x4
	if v&0x20 != 0 {
		b = append(b, "InRtcm3")
	}
	v &^= 0x20
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgPrt3OutProtoMask) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "OutUbx")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "OutNmea")
	}
	v &^= 0x2
	if v&0x20 != 0 {
		b = append(b, "OutRtcm3")
	}
	v &^= 0x20
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgPrt3Flags) String() string {
	var b []string
	if v&0x2 != 0 {
		b = append(b, "ExtendedTxTimeout")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// CfgPrt4 (Get/set) Port configuration for I2C (DDC) port
// Class/Id 0x06 0x00 (20 bytes)
// Several configurations can be concatenated to one input message. In this case the payload length can be a multiple of the normal length (see the other versions of CFG-PRT). Output messages from the module contain only one configuration unit.
type CfgPrt4 struct {
	PortID       byte                `stf:"0"` // Port identifier number (= 0 for I2C (DDC) port)
	Reserved1    byte                // Reserved
	TxReady      CfgPrt4TxReady      // TX ready PIN configuration
	Mode         CfgPrt4Mode         // I2C (DDC) Mode Flags
	Reserved2    [4]byte             // Reserved
	InProtoMask  CfgPrt4InProtoMask  // A mask describing which input protocols are active. Each bit of this mask is used for a protocol. Through that, multiple protocols can be defined on a single port. (The bitfield inRtcm3 is not supported in protocol versions less than 20)
	OutProtoMask CfgPrt4OutProtoMask // A mask describing which output protocols are active. Each bit of this mask is used for a protocol. Through that, multiple protocols can be defined on a single port. (The bitfield outRtcm3 is not supported in protocol versions less than 20)
	Flags        CfgPrt4Flags        // Flags bit mask
	Reserved3    [2]byte             // Reserved
}

func (CfgPrt4) classID() uint16 { return 0x0006 }

type CfgPrt4TxReady uint16

const (
	CfgPrt4En    CfgPrt4TxReady = 0x1    // Enable TX ready feature for this port
	CfgPrt4Pol   CfgPrt4TxReady = 0x2    // Polarity 0 High-active 1 Low-active
	CfgPrt4Pin   CfgPrt4TxReady = 0x7c   // PIO to be used (must not be in use by another function)
	CfgPrt4Thres CfgPrt4TxReady = 0xff80 // Threshold The given threshold is multiplied by 8 bytes. The TX ready PIN goes active after >= thres*8 bytes are pending for the port and going inactive after the last pending bytes have been written to hardware (0-4 bytes before end of stream). 0x000 no threshold 0x001 8byte 0x002 16byte ... 0x1FE 4080byte 0x1FF 4088byte
)

type CfgPrt4Mode uint32

const (
	CfgPrt4SlaveAddr CfgPrt4Mode = 0xfe // Slave address Range: 0x07 < slaveAddr < 0x78. Bit 0 must be 0
)

type CfgPrt4InProtoMask uint16

const (
	CfgPrt4InUbx   CfgPrt4InProtoMask = 0x1  // UBX protocol
	CfgPrt4InNmea  CfgPrt4InProtoMask = 0x2  // NMEA protocol
	CfgPrt4InRtcm  CfgPrt4InProtoMask = 0x4  // RTCM2 protocol
	CfgPrt4InRtcm3 CfgPrt4InProtoMask = 0x20 // RTCM3 protocol (not supported in protocol versions less than 20)
)

type CfgPrt4OutProtoMask uint16

const (
	CfgPrt4OutUbx   CfgPrt4OutProtoMask = 0x1  // UBX protocol
	CfgPrt4OutNmea  CfgPrt4OutProtoMask = 0x2  // NMEA protocol
	CfgPrt4OutRtcm3 CfgPrt4OutProtoMask = 0x20 // RTCM3 protocol (not supported in protocol versions less than 20)
)

type CfgPrt4Flags uint16

const (
	CfgPrt4ExtendedTxTimeout CfgPrt4Flags = 0x2 // Extended TX timeout: if set, the port will time out if allocated TX memory >=4 kB and no activity for 1. 5s.
)

func (v CfgPrt4TxReady) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "En")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "Pol")
	}
	v &^= 0x2
	if v&0x7c != 0 {
		b = append(b, fmt.Sprintf("Pin<%b>", (v&0x7c)>>2))
	}
	v &^= 0x7c
	if v&0xff80 != 0 {
		b = append(b, fmt.Sprintf("Thres<%b>", (v&0xff80)>>7))
	}
	v &^= 0xff80
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgPrt4Mode) String() string {
	var b []string
	if v&0xfe != 0 {
		b = append(b, fmt.Sprintf("SlaveAddr<%b>", (v&0xfe)>>1))
	}
	v &^= 0xfe
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgPrt4InProtoMask) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "InUbx")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "InNmea")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "InRtcm")
	}
	v &^= 0x4
	if v&0x20 != 0 {
		b = append(b, "InRtcm3")
	}
	v &^= 0x20
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgPrt4OutProtoMask) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "OutUbx")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "OutNmea")
	}
	v &^= 0x2
	if v&0x20 != 0 {
		b = append(b, "OutRtcm3")
	}
	v &^= 0x20
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgPrt4Flags) String() string {
	var b []string
	if v&0x2 != 0 {
		b = append(b, "ExtendedTxTimeout")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-pwr

// CfgPwr (Set) Put receiver in a defined power state
// Class/Id 0x06 0x57 (8 bytes)
// This message is deprecated in protocol versions greater than 17. Use UBX-CFG- RST for GNSS start/stop and UBX-RXM-PMREQ for software backup. -
type CfgPwr struct {
	Version   byte    // Message version (0x01 for this version)
	Reserved1 [3]byte // Reserved
	State     uint32  // Enter system state 0x52554E20: GNSS running 0x53544F50: GNSS stopped 0x42434B50: Software backup. USB interface will be disabled, other wakeup source is needed.
}

func (CfgPwr) classID() uint16 { return 0x5706 }

// Message ubx-cfg-rate

// CfgRate (Get/set) Navigation/measurement rate settings
// Class/Id 0x06 0x08 (6 bytes)
// This message allows the user to alter the rate at which navigation solutions (and the measurements that they depend on) are generated by the receiver. The calculation of the navigation solution will always be aligned to the top of a second zero (first second of the week) of the configured reference time system. (Navigation period is an integer multiple of the measurement period in protocol versions greater than 17).  Each measurement triggers the measurements generation and, if available,  raw data output.  The navRate value defines that every nth measurement triggers a navigation  epoch.  The update rate has a direct influence on the power consumption. The more  fixes that are required, the more CPU power and communication resources are  required.  For most applications a 1 Hz update rate would be sufficient.  When using power save mode, measurement and navigation rate can differ  from the values configured here.  See Measurement and navigation rate with power save mode for details.
type CfgRate struct {
	MeasRate_ms    uint16 // [ms] The elapsed time between GNSS measurements, which defines the rate, e. g. 100 ms => 10 Hz, 1000 ms => 1 Hz, 10000 ms => 0.1 Hz. Measurement rate should be greater than or equal to 25 ms. (Measurement rate should be greater than or equal to 50 ms in protocol versions less than 24).
	NavRate_cycles uint16 // [cycles] The ratio between the number of measurements and the number of navigation solutions, e.g. 5 means five measurements for every navigation solution. Maximum value is 127. (This parameter is ignored and the navRate is fixed to 1 in protocol versions less than 18).
	TimeRef        uint16 // The time system to which measurements are aligned: 0: UTC time 1: GPS time 2: GLONASS time (not supported in protocol versions less than 18) 3: BeiDou time (not supported in protocol versions less than 18) 4: Galileo time (not supported in protocol versions less than 18)
}

func (CfgRate) classID() uint16 { return 0x0806 }

// Message ubx-cfg-rinv

// CfgRinv (Get/set) Contents of remote inventory
// Class/Id 0x06 0x34 (1 + N*1 bytes)
// If N is greater than 30, the excess bytes are discarded.
type CfgRinv struct {
	Flags CfgRinvFlags // Flags
	Items string
}

func (CfgRinv) classID() uint16 { return 0x3406 }

type CfgRinvFlags byte

const (
	CfgRinvDump   CfgRinvFlags = 0x1 // Dump data at startup. Does not work if flag binary is set.
	CfgRinvBinary CfgRinvFlags = 0x2 // Data is binary.
)

func (v CfgRinvFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Dump")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "Binary")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-rst

// CfgRst (Command) Reset receiver / Clear backup data structures
// Class/Id 0x06 0x04 (4 bytes)
// Do not expect this message to be acknowledged by the receiver.  Newer FW version will not acknowledge this message at all.  Older FW version will acknowledge this message but the acknowledge may not  be sent completely before the receiver is reset.
type CfgRst struct {
	NavBbrMask CfgRstNavBbrMask // BBR sections to clear. The following special sets apply: 0x0000 Hot start 0x0001 Warm start 0xFFFF Cold start
	ResetMode  byte             // Reset Type 0x00 = Hardware reset (watchdog) immediately 0x01 = Controlled software reset 0x02 = Controlled software reset (GNSS only) 0x04 = Hardware reset (watchdog) after shutdown 0x08 = Controlled GNSS stop 0x09 = Controlled GNSS start
	Reserved1  byte             // Reserved
}

func (CfgRst) classID() uint16 { return 0x0406 }

type CfgRstNavBbrMask uint16

const (
	CfgRstEph    CfgRstNavBbrMask = 0x1    // Ephemeris
	CfgRstAlm    CfgRstNavBbrMask = 0x2    // Almanac
	CfgRstHealth CfgRstNavBbrMask = 0x4    // Health
	CfgRstKlob   CfgRstNavBbrMask = 0x8    // Klobuchar parameters
	CfgRstPos    CfgRstNavBbrMask = 0x10   // Position
	CfgRstClkd   CfgRstNavBbrMask = 0x20   // Clock drift
	CfgRstOsc    CfgRstNavBbrMask = 0x40   // Oscillator parameter
	CfgRstUtc    CfgRstNavBbrMask = 0x80   // UTC correction + GPS leap seconds parameters
	CfgRstRtc    CfgRstNavBbrMask = 0x100  // RTC
	CfgRstAop    CfgRstNavBbrMask = 0x8000 // Autonomous orbit parameters
)

func (v CfgRstNavBbrMask) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Eph")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "Alm")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "Health")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "Klob")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "Pos")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "Clkd")
	}
	v &^= 0x20
	if v&0x40 != 0 {
		b = append(b, "Osc")
	}
	v &^= 0x40
	if v&0x80 != 0 {
		b = append(b, "Utc")
	}
	v &^= 0x80
	if v&0x100 != 0 {
		b = append(b, "Rtc")
	}
	v &^= 0x100
	if v&0x8000 != 0 {
		b = append(b, "Aop")
	}
	v &^= 0x8000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-rxm

// CfgRxm (Get/set) RXM configuration
// Class/Id 0x06 0x11 (2 bytes)
// For a detailed description see section Power Management.
type CfgRxm struct {
	Reserved1 byte // Reserved
	LpMode    byte // Low power mode 0: Continuous mode 1: Power save mode 4: Continuous mode
}

func (CfgRxm) classID() uint16 { return 0x1106 }

// Message ubx-cfg-sbas

// CfgSbas (Get/set) SBAS configuration
// Class/Id 0x06 0x16 (8 bytes)
// This message configures the SBAS receiver subsystem (i.e. WAAS, EGNOS, MSAS). See the SBAS configuration settings description for a detailed description of how these settings affect receiver operation.
type CfgSbas struct {
	Mode      CfgSbasMode  // SBAS mode
	Usage     CfgSbasUsage // SBAS usage
	MaxSBAS   byte         // Maximum number of SBAS prioritized tracking channels (valid range: 0 - 3) to use (obsolete and superseded by UBX- CFG-GNSS in protocol versions 14+).
	Scanmode2 byte         // Continuation of scanmode bitmask below
	Scanmode1 uint32       // Which SBAS PRN numbers to search for (bitmask). If all bits are set to zero, auto-scan (i.e. all valid PRNs) are searched. Every bit corresponds to a PRN number.
}

func (CfgSbas) classID() uint16 { return 0x1606 }

type CfgSbasMode byte

const (
	CfgSbasEnabled CfgSbasMode = 0x1 // SBAS enabled (1) / disabled (0) - This field is deprecated; use UBX-CFG-GNSS to enable/disable SBAS operation
	CfgSbasTest    CfgSbasMode = 0x2 // SBAS testbed: Use data anyhow (1) / Ignore data when in test mode (SBAS msg 0)
)

type CfgSbasUsage byte

const (
	CfgSbasRange     CfgSbasUsage = 0x1 // Use SBAS GEOs as a ranging source (for navigation)
	CfgSbasDiffCorr  CfgSbasUsage = 0x2 // Use SBAS differential corrections
	CfgSbasIntegrity CfgSbasUsage = 0x4 // Use SBAS integrity information. If enabled, the receiver will only use GPS satellites for which integrity information is available.
)

func (v CfgSbasMode) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Enabled")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "Test")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgSbasUsage) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Range")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "DiffCorr")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "Integrity")
	}
	v &^= 0x4
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-senif

// CfgSenif (Get/set) I2C sensor interface configuration
// Class/Id 0x06 0x88 (6 bytes)
// -
type CfgSenif struct {
	Type    byte            // Type of interface, 0 for I2C
	Version byte            // Message version, 0 for this message
	Flags   CfgSenifFlags   // feature configuration flags
	PioConf CfgSenifPioConf // PIO configuration flags (see graphic below )
}

func (CfgSenif) classID() uint16 { return 0x8806 }

type CfgSenifFlags uint16

const (
	CfgSenifSenConn CfgSenifFlags = 0x1 // Sensor is connected to I2C interface
)

type CfgSenifPioConf uint16

const (
	CfgSenifI2cSdaPio CfgSenifPioConf = 0x1f  // PIO of the I2C SDA line Supported options:
	CfgSenifI2cSclPio CfgSenifPioConf = 0x3e0 // PIO of the I2C SCL line Supported options:
)

func (v CfgSenifFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "SenConn")
	}
	v &^= 0x1
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgSenifPioConf) String() string {
	var b []string
	if v&0x1f != 0 {
		b = append(b, fmt.Sprintf("I2cSdaPio<%b>", (v&0x1f)>>0))
	}
	v &^= 0x1f
	if v&0x3e0 != 0 {
		b = append(b, fmt.Sprintf("I2cSclPio<%b>", (v&0x3e0)>>5))
	}
	v &^= 0x3e0
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-slas

// CfgSlas (Get/set) SLAS configuration
// Class/Id 0x06 0x8d (4 bytes)
// This message configures the QZSS SLAS (Sub-meter Level Augmentation System). See the SLAS Configuration Settings Description for a detailed description of how these settings affect receiver operation. To apply SLAS corrections, QZSS operation and L1S signal tracking must be enabled see UBX-CFG-GNSS
type CfgSlas struct {
	Mode      CfgSlasMode // SLAS Mode
	Reserved1 [3]byte     // Reserved
}

func (CfgSlas) classID() uint16 { return 0x8d06 }

type CfgSlasMode byte

const (
	CfgSlasEnabled CfgSlasMode = 0x1 // Apply QZSS SLAS DGNSS corrections: Enabled (1) / Disabled (0)
	CfgSlasTest    CfgSlasMode = 0x2 // Use QZSS SLAS data when in test mode (SLAS msg 0): Use data anyhow (1) / Ignore data when in Test Mode (0)
	CfgSlasRaim    CfgSlasMode = 0x4 // Raim out measurements that are not corrected by QZSS SLAS, if at least 5 measurements are corrected: Enabled (1) / Disabled (0)
)

func (v CfgSlasMode) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Enabled")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "Test")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "Raim")
	}
	v &^= 0x4
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-smgr

// CfgSmgr (Get/set) Synchronization manager configuration
// Class/Id 0x06 0x62 (20 bytes)
// -
type CfgSmgr struct {
	Version                 byte              // Message version (0x00 for this version)
	MinGNSSFix              byte              // Minimum number of GNSS fixes before we commit to use it as a source
	MaxFreqChangeRate_ppb_s uint16            // [ppb/s] Maximum frequency change rate during disciplining. Must not exceed 30ppb/s
	MaxPhaseCorrRate_ns_s   uint16            // [ns/s] Maximum phase correction rate in coherent time pulse mode. For maximum phase correction rate in corrective time pulse mode see maxSlewRate. Note that in coherent time pulse mode phase correction is achieved by intentional frequency offset. Allowing for a high phase correction rate can result in large intentional frequency offset. Must not exceed 100ns/s
	Reserved1               [2]byte           // Reserved
	FreqTolerance_ppb       uint16            // [ppb] Limit of possible deviation from nominal before UBX-TIM-TOS indicates that frequency is out of tolerance
	TimeTolerance_ns        uint16            // [ns] Limit of possible deviation from nominal before UBX-TIM-TOS indicates that time pulse is out of tolerance
	MessageCfg              CfgSmgrMessageCfg // Sync manager message configuration
	MaxSlewRate_us_s        uint16            // [us/s] Maximum slew rate, the maximum time correction that shall be applied between locked pulses in corrective time pulse mode. To have no limit on the slew rate, set the flag disableMaxSlewRate to 1 For maximum phase correction rate in coherent time pulse mode see maxPhaseCorrRate.
	Flags                   CfgSmgrFlags      // Flags
}

func (CfgSmgr) classID() uint16 { return 0x6206 }

type CfgSmgrMessageCfg uint16

const (
	CfgSmgrMeasInternal CfgSmgrMessageCfg = 0x1 // 1 = report the estimated offset of the internal oscillator based on the oscillator model
	CfgSmgrMeasGNSS     CfgSmgrMessageCfg = 0x2 // 1 = report the internal oscillator's offset relative to GNSS
	CfgSmgrMeasEXTINT0  CfgSmgrMessageCfg = 0x4 // 1 = report the internal oscillator's offset relative to the source on EXTINT0
	CfgSmgrMeasEXTINT1  CfgSmgrMessageCfg = 0x8 // 1 = report the internal oscillator's offset relative to the source on EXTINT1
)

type CfgSmgrFlags uint32

const (
	CfgSmgrDisableInternal    CfgSmgrFlags = 0x1     // 1 = disable disciplining of the internal oscillator
	CfgSmgrDisableExternal    CfgSmgrFlags = 0x2     // 1 = disable disciplining of the external oscillator
	CfgSmgrPreferenceMode     CfgSmgrFlags = 0x4     // Reference selection preference 0 - best frequency accuracy 1 - best phase accuracy
	CfgSmgrEnableGNSS         CfgSmgrFlags = 0x8     // 1 = enable use of GNSS as synchronization source
	CfgSmgrEnableEXTINT0      CfgSmgrFlags = 0x10    // 1 = enable use of EXTINT0 as synchronization source
	CfgSmgrEnableEXTINT1      CfgSmgrFlags = 0x20    // 1 = enable use of EXTINT1 as synchronization source
	CfgSmgrEnableHostMeasInt  CfgSmgrFlags = 0x40    // 1 = enable use of host measurements on the internal oscillator as synchronization source Measurements made by the host must be sent to the receiver using a UBX-TIM-SMEAS-DATA0 message.
	CfgSmgrEnableHostMeasExt  CfgSmgrFlags = 0x80    // 1 = enable use of host measurements on the external oscillator as synchronization source Measurements made by the host must be sent to the receiver using a UBX-TIM-SMEAS-DATA0 message.
	CfgSmgrUseAnyFix          CfgSmgrFlags = 0x400   // 0 - use over-determined navigation solutions only 1 - use any fix
	CfgSmgrDisableMaxSlewRate CfgSmgrFlags = 0x800   // 0 - use the value in the field maxSlewRate for maximum time correction in corrective time pulse mode 1 - don't use the value in the field maxSlewRate
	CfgSmgrIssueFreqWarning   CfgSmgrFlags = 0x1000  // 1 = issue a warning (via UBX-TIM-TOS flag) when frequency uncertainty exceeds freqTolerance
	CfgSmgrIssueTimeWarning   CfgSmgrFlags = 0x2000  // 1 = issue a warning (via UBX-TIM-TOS flag) when time uncertainty exceeds timeTolerance
	CfgSmgrTPCoherent         CfgSmgrFlags = 0xc000  // Control time pulse coherency 0 - Coherent pulses. Time phase offsets will be corrected gradually by varying the GNSS oscillator rate within frequency tolerance limits. There will always be the correct number of GNSS oscillator cycles between time pulses. Given tight limits this may take a long time 1 - Non-coherent pulses. In this mode the receiver will correct time phase offsets as quickly as allowed by the specified maximum slew rate, in which case there may not be the expected number of GNSS oscillator cycles between time pulses. 2 - Post-initialization coherent pulses. The receiver will run in non-coherent mode as described above until the pulse timing has been corrected and PLL is active on the internal oscillator, but will then switch to coherent pulse mode.
	CfgSmgrDisableOffset      CfgSmgrFlags = 0x10000 // 1 = disable automatic storage of oscillator offset
)

func (v CfgSmgrMessageCfg) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "MeasInternal")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "MeasGNSS")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "MeasEXTINT0")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "MeasEXTINT1")
	}
	v &^= 0x8
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v CfgSmgrFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "DisableInternal")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "DisableExternal")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "PreferenceMode")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "EnableGNSS")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "EnableEXTINT0")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "EnableEXTINT1")
	}
	v &^= 0x20
	if v&0x40 != 0 {
		b = append(b, "EnableHostMeasInt")
	}
	v &^= 0x40
	if v&0x80 != 0 {
		b = append(b, "EnableHostMeasExt")
	}
	v &^= 0x80
	if v&0x400 != 0 {
		b = append(b, "UseAnyFix")
	}
	v &^= 0x400
	if v&0x800 != 0 {
		b = append(b, "DisableMaxSlewRate")
	}
	v &^= 0x800
	if v&0x1000 != 0 {
		b = append(b, "IssueFreqWarning")
	}
	v &^= 0x1000
	if v&0x2000 != 0 {
		b = append(b, "IssueTimeWarning")
	}
	v &^= 0x2000
	if v&0xc000 != 0 {
		b = append(b, fmt.Sprintf("TPCoherent<%b>", (v&0xc000)>>14))
	}
	v &^= 0xc000
	if v&0x10000 != 0 {
		b = append(b, "DisableOffset")
	}
	v &^= 0x10000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-spt

// CfgSpt (Get/set) Configure and start a sensor production test
// Class/Id 0x06 0x64 (12 bytes)
// The production test uses the built-in self-test capabilities of an attached sensor. This message is only supported if a sensor is directly connected to the u-blox receiver.
type CfgSpt struct {
	Version   byte    // Message version (0x00 for this version)
	Reserved1 byte    // Reserved
	SensorId  uint16  // ID of the sensor to be tested; see UBX- MON-SPT for defined IDs
	Reserved2 [8]byte // Reserved
}

func (CfgSpt) classID() uint16 { return 0x6406 }

// Message ubx-cfg-tmode2

// CfgTmode2 (Get/set) Time mode settings 2
// Class/Id 0x06 0x3d (28 bytes)
// This message is available only for timing receivers See the Time Mode Description for details. This message replaces the deprecated UBX-CFG-TMODE message.
type CfgTmode2 struct {
	TimeMode        byte           // Time Transfer Mode: 0 Disabled 1 Survey In 2 Fixed Mode (true position information required) 3-255 Reserved
	Reserved1       byte           // Reserved
	Flags           CfgTmode2Flags // Time mode flags
	EcefXOrLat      int32          // [cm_or_deg*1e-7] WGS84 ECEF X coordinate or latitude, depending on flags above
	EcefYOrLon      int32          // [cm_or_deg*1e-7] WGS84 ECEF Y coordinate or longitude, depending on flags above
	EcefZOrAlt_cm   int32          // [cm] WGS84 ECEF Z coordinate or altitude, depending on flags above
	FixedPosAcc_mm  uint32         // [mm] Fixed position 3D accuracy
	SvinMinDur_s    uint32         // [s] Survey-in minimum duration
	SvinAccLimit_mm uint32         // [mm] Survey-in position accuracy limit
}

func (CfgTmode2) classID() uint16 { return 0x3d06 }

type CfgTmode2Flags uint16

const (
	CfgTmode2Lla    CfgTmode2Flags = 0x1 // Position is given in LAT/LON/ALT (default is ECEF)
	CfgTmode2AltInv CfgTmode2Flags = 0x2 // Altitude is not valid, in case lla was set
)

func (v CfgTmode2Flags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Lla")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "AltInv")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-tmode3

// CfgTmode3 (Get/set) Time mode settings 3
// Class/Id 0x06 0x71 (40 bytes)
// Configures the receiver to be in Time Mode. The position referred to in this message is that of the Antenna Reference Point (ARP). See the Time Mode Description for details.
type CfgTmode3 struct {
	Version       byte           // Message version (0x00 for this version)
	Reserved1     byte           // Reserved
	Flags         CfgTmode3Flags // Receiver mode flags
	EcefXOrLat    int32          // [cm_or_deg*1e-7] WGS84 ECEF X coordinate (or latitude) of the ARP position, depending on flags above
	EcefYOrLon    int32          // [cm_or_deg*1e-7] WGS84 ECEF Y coordinate (or longitude) of the ARP position, depending on flags above
	EcefZOrAlt_cm int32          // [cm] WGS84 ECEF Z coordinate (or altitude) of the ARP position, depending on flags above
	EcefXOrLatHP  int8           // [0.1_mm_or_deg*1e-9] High-precision WGS84 ECEF X coordinate (or latitude) of the ARP position, depending on flags above. Must be in the range -99..+99. The precise WGS84 ECEF X coordinate in units of cm, or the precise WGS84 ECEF latitude in units of 1e-7 degrees, is given by ecefXOrLat + (ecefXOrLatHP * 1e-2)
	EcefYOrLonHP  int8           // [0.1_mm_or_deg*1e-9] High-precision WGS84 ECEF Y coordinate (or longitude) of the ARP position, depending on flags above. Must be in the range -99..+99. The precise WGS84 ECEF Y coordinate in units of cm, or the precise WGS84 ECEF longitude in units of 1e-7 degrees, is given by ecefYOrLon + (ecefYOrLonHP * 1e-2)
	EcefZOrAltHP  int8           // [0.1_mm] High-precision WGS84 ECEF Z coordinate (or altitude) of the ARP position, depending on flags above. Must be in the range -99..+99. The precise WGS84 ECEF Z coordinate, or altitude coordinate, in units of cm is given by ecefZOrAlt + (ecefZOrAltHP * 1e-2)
	Reserved2     byte           // Reserved
	FixedPosAcc   uint32         // [0.1_mm] Fixed position 3D accuracy
	SvinMinDur_s  uint32         // [s] Survey-in minimum duration
	SvinAccLimit  uint32         // [0.1_mm] Survey-in position accuracy limit
	Reserved3     [8]byte        // Reserved
}

func (CfgTmode3) classID() uint16 { return 0x7106 }

type CfgTmode3Flags uint16

const (
	CfgTmode3Mode CfgTmode3Flags = 0xff  // Receiver Mode: 0 Disabled 1 Survey In 2 Fixed Mode (true ARP position information required) 3-255 Reserved
	CfgTmode3Lla  CfgTmode3Flags = 0x100 // Position is given in LAT/LON/ALT (default is ECEF)
)

func (v CfgTmode3Flags) String() string {
	var b []string
	if v&0xff != 0 {
		b = append(b, fmt.Sprintf("Mode<%b>", (v&0xff)>>0))
	}
	v &^= 0xff
	if v&0x100 != 0 {
		b = append(b, "Lla")
	}
	v &^= 0x100
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-tp5 (3 versions)

// CfgTp5 (Poll Request) Poll time pulse parameters for time pulse 0
// Class/Id 0x06 0x31 (0 bytes)
// Sending this (empty / no-payload) message to the receiver results in the receiver returning a message of type UBX-CFG-TP5 with a payload as defined below for timepulse 0.
type CfgTp5 struct {
}

func (CfgTp5) classID() uint16 { return 0x3106 }

// CfgTp51 (Poll Request) Poll time pulse parameters
// Class/Id 0x06 0x31 (1 bytes)
// Sending this message to the receiver results in the receiver returning a message of type UBX-CFG-TP5 with a payload as defined below for the specified time pulse.
type CfgTp51 struct {
	TpIdx byte // Time pulse selection (0 = TIMEPULSE, 1 = TIMEPULSE2)
}

func (CfgTp51) classID() uint16 { return 0x3106 }

// CfgTp52 (Get/set) Time pulse parameters
// Class/Id 0x06 0x31 (32 bytes)
// -
type CfgTp52 struct {
	TpIdx              byte         // Time pulse selection (0 = TIMEPULSE, 1 = TIMEPULSE2)
	Version            byte         // Message version (0x01 for this version)
	Reserved1          [2]byte      // Reserved
	AntCableDelay_ns   int16        // [ns] Antenna cable delay
	RfGroupDelay_ns    int16        // [ns] RF group delay
	FreqPeriod         uint32       // [Hz_or_us] Frequency or period time, depending on setting of bit 'isFreq'
	FreqPeriodLock     uint32       // [Hz_or_us] Frequency or period time when locked to GNSS time, only used if 'lockedOtherSet' is set
	PulseLenRatio      uint32       // [us_or_2^-32] Pulse length or duty cycle, depending on 'isLength'
	PulseLenRatioLock  uint32       // [us_or_2^-32] Pulse length or duty cycle when locked to GNSS time, only used if 'lockedOtherSet' is set
	UserConfigDelay_ns int32        // [ns] User-configurable time pulse delay
	Flags              CfgTp52Flags // Configuration flags
}

func (CfgTp52) classID() uint16 { return 0x3106 }

type CfgTp52Flags uint32

const (
	CfgTp52Active         CfgTp52Flags = 0x1    // If set enable time pulse; if pin assigned to another function, other function takes precedence. Must be set for FTS variant.
	CfgTp52LockGnssFreq   CfgTp52Flags = 0x2    // If set, synchronize time pulse to GNSS as soon as GNSS time is valid. If not set, or before GNSS time is valid, use local clock. This flag is ignored by the FTS product variant; in this case the receiver always locks to the best available time/frequency reference (which is not necessarily GNSS). This flag can be unset only in Timing product variants.
	CfgTp52LockedOtherSet CfgTp52Flags = 0x4    // If set the receiver switches between the timepulse settings given by 'freqPeriodLocked' & 'pulseLenLocked' and those given by 'freqPeriod' & 'pulseLen'. The 'Locked' settings are used where the receiver has an accurate sense of time. For non-FTS products, this occurs when GNSS solution with a reliable time is available, but for FTS products the setting syncMode field governs behavior. In all cases, the receiver only uses 'freqPeriod' & 'pulseLen' when the flag is unset.
	CfgTp52IsFreq         CfgTp52Flags = 0x8    // If set 'freqPeriodLock' and 'freqPeriod' are interpreted as frequency, otherwise interpreted as period.
	CfgTp52IsLength       CfgTp52Flags = 0x10   // If set 'pulseLenRatioLock' and 'pulseLenRatio' interpreted as pulse length, otherwise interpreted as duty cycle.
	CfgTp52AlignToTow     CfgTp52Flags = 0x20   // Align pulse to top of second (period time must be integer fraction of 1s). Also set 'lockGnssFreq' to use this feature. This flag is ignored by the FTS product variant; it is assumed to be always set (as is lockGnssFreq). Set maxSlewRate and maxPhaseCorrRate fields of UBX-CFG-SMGR to 0 to disable alignment.
	CfgTp52Polarity       CfgTp52Flags = 0x40   // Pulse polarity: 0: falling edge at top of second 1: rising edge at top of second
	CfgTp52GridUtcGnss    CfgTp52Flags = 0x780  // Timegrid to use: 0: UTC 1: GPS 2: GLONASS 3: BeiDou 4: Galileo (not supported in protocol versions less than 18) This flag is only relevant if 'lockGnssFreq' and 'alignToTow' are set. Note that configured GNSS time is estimated by the receiver if locked to any GNSS system. If the receiver has a valid GNSS fix it will attempt to steer the TP to the specified time grid even if the specified time is not based on information from the constellation's satellites. To ensure timing based purely on a given GNSS, restrict the supported constellations in UBX-CFG-GNSS.
	CfgTp52SyncMode       CfgTp52Flags = 0x3800 // Sync Manager lock mode to use: 0: switch to 'freqPeriodLock' and 'pulseLenRatioLock' as soon as Sync Manager has an accurate time, never switch back to 'freqPeriod' and 'pulseLenRatio' 1: switch to 'freqPeriodLock' and 'pulseLenRatioLock' as soon as Sync Manager has an accurate time, and switch back to 'freqPeriod' and 'pulseLenRatio' as soon as time gets inaccurate This field is only relevant for the FTS product variant. This field is only relevant if the flag 'lockedOtherSet' is set.
)

func (v CfgTp52Flags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Active")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "LockGnssFreq")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "LockedOtherSet")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "IsFreq")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "IsLength")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "AlignToTow")
	}
	v &^= 0x20
	if v&0x40 != 0 {
		b = append(b, "Polarity")
	}
	v &^= 0x40
	if v&0x780 != 0 {
		b = append(b, fmt.Sprintf("GridUtcGnss<%b>", (v&0x780)>>7))
	}
	v &^= 0x780
	if v&0x3800 != 0 {
		b = append(b, fmt.Sprintf("SyncMode<%b>", (v&0x3800)>>11))
	}
	v &^= 0x3800
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-txslot

// CfgTxslot (Set) TX buffer time slots configuration
// Class/Id 0x06 0x53 (16 bytes)
// This message configures how transmit time slots are defined for the receiver interfaces. These time slots are relative to the chosen time pulse. A receiver that supports this message offers 3 time slots: nr. 0, 1 and 2. These time pulses follow each other and their associated priorities decrease in this order. The end of each can be specified in this message, the beginning is when the circularly previous slot ends (i.e. slot 0 starts when slot 2 finishes).
type CfgTxslot struct {
	Version   byte            // Message version (0x00 for this version)
	Enable    CfgTxslotEnable // Bitfield of ports for which the slots are enabled.
	RefTp     byte            // Reference timepulse source 0 - Timepulse 1 - Timepulse 2
	Reserved1 byte            // Reserved
	End1_ms   uint32          // [ms] End of timeslot in milliseconds after time pulse 1
	End2_ms   uint32          // [ms] End of timeslot in milliseconds after time pulse 2
	End3_ms   uint32          // [ms] End of timeslot in milliseconds after time pulse 3
}

func (CfgTxslot) classID() uint16 { return 0x5306 }

type CfgTxslotEnable byte

const (
	CfgTxslotDDC   CfgTxslotEnable = 0x1  // DDC/I2C
	CfgTxslotUART1 CfgTxslotEnable = 0x2  // UART 1
	CfgTxslotUART2 CfgTxslotEnable = 0x4  // UART 2
	CfgTxslotUSB   CfgTxslotEnable = 0x8  // USB
	CfgTxslotSPI   CfgTxslotEnable = 0x10 // SPI
)

func (v CfgTxslotEnable) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "DDC")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "UART1")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "UART2")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "USB")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "SPI")
	}
	v &^= 0x10
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-cfg-usb

// CfgUsb (Get/set) USB configuration
// Class/Id 0x06 0x1b (108 bytes)
// -
type CfgUsb struct {
	VendorID            uint16      // Vendor ID. This field shall only be set to registered Vendor IDs. Changing this field requires special Host drivers.
	ProductID           uint16      // Product ID. Changing this field requires special Host drivers.
	Reserved1           [2]byte     // Reserved
	Reserved2           [2]byte     // Reserved
	PowerConsumption_ma uint16      // [mA] Power consumed by the device
	Flags               CfgUsbFlags // various configuration flags
	VendorString        [32]byte    // String containing the vendor name. 32 ASCII bytes including 0-termination.
	ProductString       [32]byte    // String containing the product name. 32 ASCII bytes including 0-termination.
	SerialNumber        [32]byte    // String containing the serial number. 32 ASCII bytes including 0-termination. Changing the String fields requires special Host drivers.
}

func (CfgUsb) classID() uint16 { return 0x1b06 }

type CfgUsbFlags uint16

const (
	CfgUsbReEnum    CfgUsbFlags = 0x1 // force re-enumeration
	CfgUsbPowerMode CfgUsbFlags = 0x2 // self-powered (1), bus-powered (0)
)

func (v CfgUsbFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "ReEnum")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "PowerMode")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-esf-alg

// EsfAlg (Periodic/Polled) IMU alignment information
// Class/Id 0x10 0x14 (16 bytes)
// This message outputs the IMU alignment angles which define the rotation from the installation-frame to the IMU-frame (see the IMU-mount Misalignment section for more details). In addition, it outputs information about the automatic IMU-mount alignment (if enabled).
type EsfAlg struct {
	ITOW_ms     uint32      // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Version     byte        // Message version (0x01 for this version)
	Flags       EsfAlgFlags // Flags
	Error       EsfAlgError // Flags
	Reserved1   byte        // Reserved
	Yaw_dege2   uint32      // [1e-2 deg] IMU-mount yaw angle [0, 360]
	Pitch_dege2 int16       // [1e-2 deg] IMU-mount pitch angle [-90, 90]
	Roll_dege2  int16       // [1e-2 deg] IMU-mount roll angle [-180, 180]
}

func (EsfAlg) classID() uint16 { return 0x1410 }

type EsfAlgFlags byte

const (
	EsfAlgAutoMntAlgOn EsfAlgFlags = 0x1 // Automatic IMU-mount alignment on/off bit (0: automatic alignment is not running, 1: automatic alignment is running)
	EsfAlgStatus       EsfAlgFlags = 0xe // Status of the IMU-mount alignment (0: user-defined/fixed angles are used, 1: IMU-mount roll/pitch angles alignment is ongoing, 2: IMU-mount roll/pitch/yaw angles alignment is ongoing, 3: coarse IMU-mount alignment are used, 4: fine IMU-mount alignment are used)
)

type EsfAlgError byte

const (
	EsfAlgTiltAlgError EsfAlgError = 0x1 // IMU-mount tilt (roll and/or pitch) alignment error (0: no error, 1: error)
	EsfAlgYawAlgError  EsfAlgError = 0x2 // IMU-mount yaw alignment error (0: no error, 1: error)
	EsfAlgAngleError   EsfAlgError = 0x4 // IMU-mount misalignment Euler angle singularity error (0: no error, 1: error). If this error bit is set, the IMU-mount roll and IMU-mount yaw angles cannot uniquely be defined due to the singularity issue happening with installations mounted with a +/- 90 degrees misalignment around pitch axis. This is also known as the 'gimbal-lock' problem affecting rotations described by Euler angles.
)

func (v EsfAlgFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "AutoMntAlgOn")
	}
	v &^= 0x1
	if v&0xe != 0 {
		b = append(b, fmt.Sprintf("Status<%b>", (v&0xe)>>1))
	}
	v &^= 0xe
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v EsfAlgError) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "TiltAlgError")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "YawAlgError")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "AngleError")
	}
	v &^= 0x4
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-esf-ins

// EsfIns (Periodic/Polled) Vehicle dynamics information
// Class/Id 0x10 0x15 (36 bytes)
// This message outputs information about the vehicle dynamics. For ADR products (in protocol versions less than 19.2), the output dynamics information (angular rates and accelerations) is expressed with respect to the vehicle-frame. More information can be found in the ADR Navigation Output section. For ADR products, the output dynamics information (angular rates and accelerations) is expressed with respect to the vehicle-frame. More information can be found in the ADR Navigation Output section. For UDR products, the output dynamics information (angular rates and accelerations) are expressed with respect to the body-frame. More information can be found in the UDR Navigation Output section.
type EsfIns struct {
	Bitfield0        EsfInsBitfield0 // Bitfield
	Reserved1        [4]byte         // Reserved
	ITOW_ms          uint32          // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	XAngRate_deg_se3 int32           // [1e-3 deg/s] Compensated x-axis angular rate.
	YAngRate_deg_se3 int32           // [1e-3 deg/s] Compensated y-axis angular rate.
	ZAngRate_deg_se3 int32           // [1e-3 deg/s] Compensated z-axis angular rate.
	XAccel_m_s2e2    int32           // [1e-2 m/s^2] Compensated x-axis acceleration (gravity- free).
	YAccel_m_s2e2    int32           // [1e-2 m/s^2] Compensated y-axis acceleration (gravity- free).
	ZAccel_m_s2e2    int32           // [1e-2 m/s^2] Compensated z-axis acceleration (gravity- free).
}

func (EsfIns) classID() uint16 { return 0x1510 }

type EsfInsBitfield0 uint32

const (
	EsfInsVersion       EsfInsBitfield0 = 0xff   // Message version (0x01 for this version)
	EsfInsXAngRateValid EsfInsBitfield0 = 0x100  // Compensated x-axis angular rate data validity flag (0: not valid, 1: valid).
	EsfInsYAngRateValid EsfInsBitfield0 = 0x200  // Compensated y-axis angular rate data validity flag (0: not valid, 1: valid).
	EsfInsZAngRateValid EsfInsBitfield0 = 0x400  // Compensated z-axis angular rate data validity flag (0: not valid, 1: valid).
	EsfInsXAccelValid   EsfInsBitfield0 = 0x800  // Compensated x-axis acceleration data validity flag (0: not valid, 1: valid).
	EsfInsYAccelValid   EsfInsBitfield0 = 0x1000 // Compensated y-axis acceleration data validity flag (0: not valid, 1: valid).
	EsfInsZAccelValid   EsfInsBitfield0 = 0x2000 // Compensated z-axis acceleration data validity flag (0: not valid, 1: valid).
)

func (v EsfInsBitfield0) String() string {
	var b []string
	if v&0xff != 0 {
		b = append(b, fmt.Sprintf("Version<%b>", (v&0xff)>>0))
	}
	v &^= 0xff
	if v&0x100 != 0 {
		b = append(b, "XAngRateValid")
	}
	v &^= 0x100
	if v&0x200 != 0 {
		b = append(b, "YAngRateValid")
	}
	v &^= 0x200
	if v&0x400 != 0 {
		b = append(b, "ZAngRateValid")
	}
	v &^= 0x400
	if v&0x800 != 0 {
		b = append(b, "XAccelValid")
	}
	v &^= 0x800
	if v&0x1000 != 0 {
		b = append(b, "YAccelValid")
	}
	v &^= 0x1000
	if v&0x2000 != 0 {
		b = append(b, "ZAccelValid")
	}
	v &^= 0x2000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-esf-meas

// EsfMeas (Input/Output) External sensor fusion measurements
// Class/Id 0x10 0x02 (8 or 12 + N*4 bytes)
// Possible data types for the data field are described in the ESF Measurement Data section.
type EsfMeas struct {
	TimeTag uint32             // Time tag of measurement generated by external sensor
	Flags   EsfMeasFlags       // Flags. Set all unused bits to zero.
	Id      uint16             // Identification number of data provider
	Meas    []*EsfMeasMeasType // len: NumMeas
	// Optional
	CalibTtag_ms uint32 // [ms] Receiver local time calibrated. This field must not be supplied when calibTtagValid is set to 0.
}

func (EsfMeas) classID() uint16 { return 0x0210 }

type EsfMeasMeasType struct {
	Data EsfMeasData // data
}

type EsfMeasFlags uint16

const (
	EsfMeasTimeMarkSent   EsfMeasFlags = 0x3    // Time mark signal was supplied just prior to sending this message: 0 = none, 1 = on Ext0, 2 = on Ext1
	EsfMeasTimeMarkEdge   EsfMeasFlags = 0x4    // Trigger on rising (0) or falling (1) edge of time mark signal
	EsfMeasCalibTtagValid EsfMeasFlags = 0x8    // Calibration time tag available. Always set to zero.
	EsfMeasNumMeas        EsfMeasFlags = 0xf800 // Number of measurements contained in this message (optional, can be obtained from message size)
)

type EsfMeasData uint32

const (
	EsfMeasDataField EsfMeasData = 0xffffff   // Data
	EsfMeasDataType  EsfMeasData = 0x3f000000 // Type of data (0 = no data; 1..63 = data type)
)

func (v EsfMeasFlags) String() string {
	var b []string
	if v&0x3 != 0 {
		b = append(b, fmt.Sprintf("TimeMarkSent<%b>", (v&0x3)>>0))
	}
	v &^= 0x3
	if v&0x4 != 0 {
		b = append(b, "TimeMarkEdge")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "CalibTtagValid")
	}
	v &^= 0x8
	if v&0xf800 != 0 {
		b = append(b, fmt.Sprintf("NumMeas<%b>", (v&0xf800)>>11))
	}
	v &^= 0xf800
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v EsfMeasData) String() string {
	var b []string
	if v&0xffffff != 0 {
		b = append(b, fmt.Sprintf("DataField<%b>", (v&0xffffff)>>0))
	}
	v &^= 0xffffff
	if v&0x3f000000 != 0 {
		b = append(b, fmt.Sprintf("DataType<%b>", (v&0x3f000000)>>24))
	}
	v &^= 0x3f000000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-esf-raw

// EsfRaw (Output) Raw sensor measurements
// Class/Id 0x10 0x03 (4 + N*8 bytes)
// The message contains measurements from the active inertial sensors connected to the GNSS chip. Possible data types for the data field are accelerometer, gyroscope and temperature readings as described in the ESF Measurement Data section. Note that the rate selected in UBX-CFG-MSG is not respected. If a positive rate is selected then all raw measurements will be output. See also Raw Sensor Measurement Data.
type EsfRaw struct {
	Reserved1 [4]byte            // Reserved
	Items     []*EsfRawItemsType // len: N
}

func (EsfRaw) classID() uint16 { return 0x0310 }

type EsfRawItemsType struct {
	Data  EsfRawData // data Same as in UBX-ESF-MEAS
	STtag uint32     // sensor time tag
}

type EsfRawData uint32

const (
	EsfRawDataField EsfRawData = 0xffffff   // data
	EsfRawDataType  EsfRawData = 0xff000000 // type of data (0 = no data; 1..255 = data type)
)

func (v EsfRawData) String() string {
	var b []string
	if v&0xffffff != 0 {
		b = append(b, fmt.Sprintf("DataField<%b>", (v&0xffffff)>>0))
	}
	v &^= 0xffffff
	if v&0xff000000 != 0 {
		b = append(b, fmt.Sprintf("DataType<%b>", (v&0xff000000)>>24))
	}
	v &^= 0xff000000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-esf-status

// EsfStatus (Periodic/Polled) External sensor fusion status
// Class/Id 0x10 0x10 (16 + N*4 bytes)
// -
type EsfStatus struct {
	ITOW_ms    uint32               // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Version    byte                 // Message version (0x02 for this version)
	Reserved1  [7]byte              // Reserved
	FusionMode byte                 // Fusion mode: 0: Initialization mode: receiver is initializing some unknown values required for doing sensor fusion 1: Fusion mode: GNSS and sensor data are used for navigation solution computation 2: Suspended fusion mode: sensor fusion is temporarily disabled due to e.g. invalid sensor data or detected ferry 3: Disabled fusion mode: sensor fusion is permanently disabled until receiver reset due e.g. to sensor error More details can be found in the Fusion Modes section.
	Reserved2  [2]byte              // Reserved
	NumSens    byte                 `len:"Sens"` // Number of sensors
	Sens       []*EsfStatusSensType // len: NumSens
}

func (EsfStatus) classID() uint16 { return 0x1010 }

type EsfStatusSensType struct {
	SensStatus1 EsfStatusSensStatus1 // Sensor status, part 1
	SensStatus2 EsfStatusSensStatus2 // Sensor status, part 2
	Freq_hz     byte                 // [Hz] Observation frequency
	Faults      EsfStatusFaults      // Sensor faults
}

type EsfStatusSensStatus1 byte

const (
	EsfStatusType  EsfStatusSensStatus1 = 0x3f // Sensor data type. See section Sensor data types in the Integration manual for details.
	EsfStatusUsed  EsfStatusSensStatus1 = 0x40 // If set, sensor data is used for the current sensor fusion solution.
	EsfStatusReady EsfStatusSensStatus1 = 0x80 // If set, sensor is set up (configuration is available or not required) but not used for computing the current sensor fusion solution.
)

type EsfStatusSensStatus2 byte

const (
	EsfStatusCalibStatus EsfStatusSensStatus2 = 0x3 // 00: Sensor is not calibrated 01: Sensor is calibrating 10/11: Sensor is calibrated Good dead reckoning performance is only possible when all used sensors are calibrated. Depending on the quality of the GNSS signals and the sensor data, the sensors may take a longer time to get calibrated.
	EsfStatusTimeStatus  EsfStatusSensStatus2 = 0xc // 00: No data 01: Reception of the first byte used to tag the measurement 10: Event input used to tag the measurement 11: Time tag provided with the data
)

type EsfStatusFaults byte

const (
	EsfStatusBadMeas     EsfStatusFaults = 0x1 // Bad measurements detected
	EsfStatusBadTTag     EsfStatusFaults = 0x2 // Bad measurement time-tags detected
	EsfStatusMissingMeas EsfStatusFaults = 0x4 // Missing or time-misaligned measurements detected
	EsfStatusNoisyMeas   EsfStatusFaults = 0x8 // High measurement noise-level detected
)

func (v EsfStatusSensStatus1) String() string {
	var b []string
	if v&0x3f != 0 {
		b = append(b, fmt.Sprintf("Type<%b>", (v&0x3f)>>0))
	}
	v &^= 0x3f
	if v&0x40 != 0 {
		b = append(b, "Used")
	}
	v &^= 0x40
	if v&0x80 != 0 {
		b = append(b, "Ready")
	}
	v &^= 0x80
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v EsfStatusSensStatus2) String() string {
	var b []string
	if v&0x3 != 0 {
		b = append(b, fmt.Sprintf("CalibStatus<%b>", (v&0x3)>>0))
	}
	v &^= 0x3
	if v&0xc != 0 {
		b = append(b, fmt.Sprintf("TimeStatus<%b>", (v&0xc)>>2))
	}
	v &^= 0xc
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v EsfStatusFaults) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "BadMeas")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "BadTTag")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "MissingMeas")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "NoisyMeas")
	}
	v &^= 0x8
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-hnr-att

// HnrAtt (Periodic/Polled) Attitude solution
// Class/Id 0x28 0x01 (32 bytes)
// This message outputs the attitude solution as roll, pitch and heading angles. More details about vehicle attitude can be found in the Vehicle Attitude Output (ADR) section for ADR products. More details about vehicle attitude can be found in the Vehicle Attitude Output (UDR) section for UDR products.
type HnrAtt struct {
	ITOW_ms          uint32  // [ms] GPS time of week of the HNR epoch.
	Version          byte    // Message version (0x01 for this version)
	Reserved1        [3]byte // Reserved
	Roll_dege5       int32   // [1e-5 deg] Vehicle roll.
	Pitch_dege5      int32   // [1e-5 deg] Vehicle pitch.
	Heading_dege5    int32   // [1e-5 deg] Vehicle heading.
	AccRoll_dege5    uint32  // [1e-5 deg] Vehicle roll accuracy (if null, roll angle is not available).
	AccPitch_dege5   uint32  // [1e-5 deg] Vehicle pitch accuracy (if null, pitch angle is not available).
	AccHeading_dege5 uint32  // [1e-5 deg] Vehicle heading accuracy (if null, heading angle is not available).
}

func (HnrAtt) classID() uint16 { return 0x0128 }

// Message ubx-hnr-ins

// HnrIns (Periodic/Polled) Vehicle dynamics information
// Class/Id 0x28 0x02 (36 bytes)
// This message outputs high rate information about vehicle dynamics computed by the Inertial Navigation System (INS) during ESF-based navigation. For ADR products (in protocol versions less than 19.2), the output dynamics information (angular rates and accelerations) is expressed with respect to the vehicle-frame. More information can be found in the ADR Navigation Output section. For UDR products, the output dynamics information (angular rates and accelerations) is expressed with respect to the body-frame. More information can be found in the UDR Navigation Output section. For ADR products, the output dynamics information (angular rates and accelerations) is expressed with respect to the vehicle-frame. More information can be found in the ADR Navigation Output section.
type HnrIns struct {
	Bitfield0        HnrInsBitfield0 // Bitfield
	Reserved1        [4]byte         // Reserved
	ITOW_ms          uint32          // [ms] GPS time of week of the HNR epoch.
	XAngRate_deg_se3 int32           // [1e-3 deg/s] Compensated x-axis angular rate.
	YAngRate_deg_se3 int32           // [1e-3 deg/s] Compensated y-axis angular rate.
	ZAngRate_deg_se3 int32           // [1e-3 deg/s] Compensated z-axis angular rate.
	XAccel_m_s2e2    int32           // [1e-2 m/s^2] Compensated x-axis acceleration (with gravity).
	YAccel_m_s2e2    int32           // [1e-2 m/s^2] Compensated y-axis acceleration (with gravity).
	ZAccel_m_s2e2    int32           // [1e-2 m/s^2] Compensated z-axis acceleration (with gravity).
}

func (HnrIns) classID() uint16 { return 0x0228 }

type HnrInsBitfield0 uint32

const (
	HnrInsVersion       HnrInsBitfield0 = 0xff   // Message version (0x00 for this version)
	HnrInsXAngRateValid HnrInsBitfield0 = 0x100  // Compensated x-axis angular rate data validity flag (0: not valid, 1: valid).
	HnrInsYAngRateValid HnrInsBitfield0 = 0x200  // Compensated y-axis angular rate data validity flag (0: not valid, 1: valid).
	HnrInsZAngRateValid HnrInsBitfield0 = 0x400  // Compensated z-axis angular rate data validity flag (0: not valid, 1: valid).
	HnrInsXAccelValid   HnrInsBitfield0 = 0x800  // Compensated x-axis acceleration data validity flag (0: not valid, 1: valid).
	HnrInsYAccelValid   HnrInsBitfield0 = 0x1000 // Compensated y-axis acceleration data validity flag (0: not valid, 1: valid).
	HnrInsZAccelValid   HnrInsBitfield0 = 0x2000 // Compensated z-axis acceleration data validity flag (0: not valid, 1: valid).
)

func (v HnrInsBitfield0) String() string {
	var b []string
	if v&0xff != 0 {
		b = append(b, fmt.Sprintf("Version<%b>", (v&0xff)>>0))
	}
	v &^= 0xff
	if v&0x100 != 0 {
		b = append(b, "XAngRateValid")
	}
	v &^= 0x100
	if v&0x200 != 0 {
		b = append(b, "YAngRateValid")
	}
	v &^= 0x200
	if v&0x400 != 0 {
		b = append(b, "ZAngRateValid")
	}
	v &^= 0x400
	if v&0x800 != 0 {
		b = append(b, "XAccelValid")
	}
	v &^= 0x800
	if v&0x1000 != 0 {
		b = append(b, "YAccelValid")
	}
	v &^= 0x1000
	if v&0x2000 != 0 {
		b = append(b, "ZAccelValid")
	}
	v &^= 0x2000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-hnr-pvt

// HnrPvt (Periodic/Polled) High rate output of PVT solution
// Class/Id 0x28 0x00 (72 bytes)
// This message provides the position, velocity and time solution with high output rate. Note that during a leap second there may be more or less than 60 seconds in a minute. See the description of leap seconds for details.
type HnrPvt struct {
	ITOW_ms       uint32      // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Year_y        uint16      // [y] Year (UTC)
	Month_month   byte        // [month] Month, range 1..12 (UTC)
	Day_d         byte        // [d] Day of month, range 1..31 (UTC)
	Hour_h        byte        // [h] Hour of day, range 0..23 (UTC)
	Min_min       byte        // [min] Minute of hour, range 0..59 (UTC)
	Sec_s         byte        // [s] Seconds of minute, range 0..60 (UTC)
	Valid         HnrPvtValid // Validity Flags
	Nano_ns       int32       // [ns] Fraction of second, range -1e9 .. 1e9 (UTC)
	GpsFix        byte        // GPSfix Type, range 0..5 0x00 = No Fix 0x01 = Dead Reckoning only 0x02 = 2D-Fix 0x03 = 3D-Fix 0x04 = GPS + dead reckoning combined 0x05 = Time only fix 0x06..0xff: reserved
	Flags         HnrPvtFlags // Fix Status Flags
	Reserved1     [2]byte     // Reserved
	Lon_dege7     int32       // [1e-7 deg] Longitude
	Lat_dege7     int32       // [1e-7 deg] Latitude
	Height_mm     int32       // [mm] Height above Ellipsoid
	HMSL_mm       int32       // [mm] Height above mean sea level
	GSpeed_mm_s   int32       // [mm/s] Ground Speed (2-D)
	Speed_mm_s    int32       // [mm/s] Speed (3-D)
	HeadMot_dege5 int32       // [1e-5 deg] Heading of motion (2-D)
	HeadVeh_dege5 int32       // [1e-5 deg] Heading of vehicle (2-D)
	HAcc_mm       uint32      // [mm] Horizontal accuracy
	VAcc_mm       uint32      // [mm] Vertical accuracy
	SAcc_mm_s     uint32      // [mm/s] Speed accuracy
	HeadAcc_dege5 uint32      // [1e-5 deg] Heading accuracy
	Reserved2     [4]byte     // Reserved
}

func (HnrPvt) classID() uint16 { return 0x0028 }

type HnrPvtValid byte

const (
	HnrPvtValidDate     HnrPvtValid = 0x1 // 1 = Valid UTC Date (see Integration manual Time Validity section for details)
	HnrPvtValidTime     HnrPvtValid = 0x2 // 1 = Valid UTC Time of Day (see Integration manual Time Validity section for details)
	HnrPvtFullyResolved HnrPvtValid = 0x4 // 1 = UTC Time of Day has been fully resolved (no seconds uncertainty)
)

type HnrPvtFlags byte

const (
	HnrPvtGPSfixOK     HnrPvtFlags = 0x1  // >1 = Fix within limits (e.g. DOP & accuracy)
	HnrPvtDiffSoln     HnrPvtFlags = 0x2  // 1 = DGPS used
	HnrPvtWKNSET       HnrPvtFlags = 0x4  // 1 = Valid GPS week number
	HnrPvtTOWSET       HnrPvtFlags = 0x8  // 1 = Valid GPS time of week (iTOW & fTOW)
	HnrPvtHeadVehValid HnrPvtFlags = 0x10 // 1= Heading of vehicle is valid
)

func (v HnrPvtValid) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "ValidDate")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "ValidTime")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "FullyResolved")
	}
	v &^= 0x4
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v HnrPvtFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "GPSfixOK")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "DiffSoln")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "WKNSET")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "TOWSET")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "HeadVehValid")
	}
	v &^= 0x10
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-inf-debug

// InfDebug (Output) ASCII output with debug contents
// Class/Id 0x04 0x04 (0 + N*1 bytes)
// This message has a variable length payload, representing an ASCII string.
type InfDebug string

func (InfDebug) classID() uint16 { return 0x0404 }

// Message ubx-inf-error

// InfError (Output) ASCII output with error contents
// Class/Id 0x04 0x00 (0 + N*1 bytes)
// This message has a variable length payload, representing an ASCII string.
type InfError string

func (InfError) classID() uint16 { return 0x0004 }

// Message ubx-inf-notice

// InfNotice (Output) ASCII output with informational contents
// Class/Id 0x04 0x02 (0 + N*1 bytes)
// This message has a variable length payload, representing an ASCII string.
type InfNotice string

func (InfNotice) classID() uint16 { return 0x0204 }

// Message ubx-inf-test

// InfTest (Output) ASCII output with test contents
// Class/Id 0x04 0x03 (0 + N*1 bytes)
// This message has a variable length payload, representing an ASCII string.
type InfTest string

func (InfTest) classID() uint16 { return 0x0304 }

// Message ubx-inf-warning

// InfWarning (Output) ASCII output with warning contents
// Class/Id 0x04 0x01 (0 + N*1 bytes)
// This message has a variable length payload, representing an ASCII string.
type InfWarning string

func (InfWarning) classID() uint16 { return 0x0104 }

// Message ubx-log-batch

// LogBatch (Polled) Batched data
// Class/Id 0x21 0x11 (100 bytes)
// This message combines position, velocity and time solution, including accuracy figures. The output of this message can be requested via UBX-LOG-RETRIEVEBATCH. The content of this message is influenced by UBX-CFG-BATCH. Depending on the flags extraPvt and extraOdo some of the fields in this message may not be valid. This validity information is also indicated in this message via flags of the same name. See Data Batching for more information. Note that during a leap second there may be more or less than 60 seconds in a minute. See the description of leap seconds for details.
type LogBatch struct {
	Version         byte                 // Message version (0x00 for this version)
	ContentValid    LogBatchContentValid // Content validity flags
	MsgCnt          uint16               // Message counter; increments for each sent UBX-LOG-BATCH message.
	ITOW_ms         uint32               // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details. Only valid if extraPvt is set.
	Year_y          uint16               // [y] Year (UTC)
	Month_month     byte                 // [month] Month, range 1..12 (UTC)
	Day_d           byte                 // [d] Day of month, range 1..31 (UTC)
	Hour_h          byte                 // [h] Hour of day, range 0..23 (UTC)
	Min_min         byte                 // [min] Minute of hour, range 0..59 (UTC)
	Sec_s           byte                 // [s] Seconds of minute, range 0..60 (UTC)
	Valid           LogBatchValid        // Validity flags
	TAcc_ns         uint32               // [ns] Time accuracy estimate (UTC) Only valid if extraPvt is set.
	FracSec_ns      int32                // [ns] Fraction of second, range -1e9 .. 1e9 (UTC)
	FixType         byte                 // GNSSfix Type: 0: no fix 2: 2D-fix 3: 3D-fix
	Flags           LogBatchFlags        // Fix status flags
	Flags2          byte                 // Additional flags
	NumSV           byte                 // Number of satellites used in Nav Solution Only valid if extraPvt is set.
	Lon_dege7       int32                // [1e-7 deg] Longitude
	Lat_dege7       int32                // [1e-7 deg] Latitude
	Height_mm       int32                // [mm] Height above ellipsoid
	HMSL_mm         int32                // [mm] Height above mean sea level Only valid if extraPvt is set.
	HAcc_mm         uint32               // [mm] Horizontal accuracy estimate
	VAcc_mm         uint32               // [mm] Vertical accuracy estimate Only valid if extraPvt is set.
	VelN_mm_s       int32                // [mm/s] NED north velocity Only valid if extraPvt is set.
	VelE_mm_s       int32                // [mm/s] NED east velocity Only valid if extraPvt is set.
	VelD_mm_s       int32                // [mm/s] NED down velocity Only valid if extraPvt is set.
	GSpeed_mm_s     int32                // [mm/s] Ground Speed (2-D)
	HeadMot_dege5   int32                // [1e-5 deg] Heading of motion (2-D)
	SAcc_mm_s       uint32               // [mm/s] Speed accuracy estimate Only valid if extraPvt is set.
	HeadAcc_dege5   uint32               // [1e-5 deg] Heading accuracy estimate Only valid if extraPvt is set.
	PDOP            uint16               // Position DOP Only valid if extraPvt is set.
	Reserved1       [2]byte              // Reserved
	Distance_m      uint32               // [m] Ground distance since last reset Only valid if extraOdo is set.
	TotalDistance_m uint32               // [m] Total cumulative ground distance Only valid if extraOdo is set.
	DistanceStd_m   uint32               // [m] Ground distance accuracy (1-sigma) Only valid if extraOdo is set.
	Reserved2       [4]byte              // Reserved
}

func (LogBatch) classID() uint16 { return 0x1121 }

type LogBatchContentValid byte

const (
	LogBatchExtraPvt LogBatchContentValid = 0x1 // Extra PVT information is valid The fields iTOW, tAcc, numSV, hMSL, vAcc, velN, velE, velD, sAcc, headAcc and pDOP are only valid if this flag is set.
	LogBatchExtraOdo LogBatchContentValid = 0x2 // Odometer data is valid The fields distance, totalDistance and distanceStd are only valid if this flag is set. Note: the odometer feature itself must also be enabled.
)

type LogBatchValid byte

const (
	LogBatchValidDate LogBatchValid = 0x1 // 1 = valid UTC Date (see Time Validity section for details)
	LogBatchValidTime LogBatchValid = 0x2 // 1 = valid UTC Time of Day (see Time Validity section for details)
)

type LogBatchFlags byte

const (
	LogBatchGnssFixOK LogBatchFlags = 0x1  // 1 = valid fix (i.e within DOP & accuracy masks)
	LogBatchDiffSoln  LogBatchFlags = 0x2  // 1 = differential corrections were applied
	LogBatchPsmState  LogBatchFlags = 0x1c // Power save mode state (see Power Management) 0: PSM is not active 1: Enabled (an intermediate state before Acquisition state) 2: Acquisition 3: Tracking 4: Power optimized tracking 5: Inactive
)

func (v LogBatchContentValid) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "ExtraPvt")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "ExtraOdo")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v LogBatchValid) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "ValidDate")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "ValidTime")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v LogBatchFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "GnssFixOK")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "DiffSoln")
	}
	v &^= 0x2
	if v&0x1c != 0 {
		b = append(b, fmt.Sprintf("PsmState<%b>", (v&0x1c)>>2))
	}
	v &^= 0x1c
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-log-create

// LogCreate (Command) Create log file
// Class/Id 0x21 0x07 (8 bytes)
// This message is used to create an initial logging file and activate the logging subsystem. UBX-ACK-ACK or UBX-ACK-NAK are returned to indicate success or failure. This message does not handle activation of recording or filtering of log entries (see UBX-CFG-LOGFILTER).
type LogCreate struct {
	Version               byte            // Message version (0x00 for this version)
	LogCfg                LogCreateLogCfg // Config flags
	Reserved1             byte            // Reserved
	LogSize               byte            // Indicates the size of the log: 0 (maximum safe size): Ensures that logging will not be interrupted and enough space will be left available for all other uses of the filestore 1 (minimum size): 2 (user-defined): See 'userDefinedSize' below
	UserDefinedSize_bytes uint32          // [bytes] Sets the maximum amount of space in the filestore that can be used by the logging task. This field is only applicable if logSize is set to user-defined.
}

func (LogCreate) classID() uint16 { return 0x0721 }

type LogCreateLogCfg byte

const (
	LogCreateCircular LogCreateLogCfg = 0x1 // Log is circular (new entries overwrite old ones in a full log) if this bit set
)

func (v LogCreateLogCfg) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Circular")
	}
	v &^= 0x1
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-log-erase

// LogErase (Command) Erase logged data
// Class/Id 0x21 0x03 (0 bytes)
// This message deactivates the logging system and erases all logged data. UBX-ACK-ACK or UBX-ACK-NAK are returned to indicate success or failure.
type LogErase struct {
}

func (LogErase) classID() uint16 { return 0x0321 }

// Message ubx-log-findtime (2 versions)

// LogFindtime (Output) Response to FINDTIME request
// Class/Id 0x21 0x0e (8 bytes)
// -
type LogFindtime struct {
	Version     byte    // Message version (0x01 for this version)
	Type        byte    // Message type, 1 for response
	Reserved1   [2]byte // Reserved
	EntryNumber uint32  // Index of the first log entry with time = given time, otherwise index of the most recent entry with time < given time. If 0xFFFFFFFF, no log entry found with time <= given time. The indexing of log entries is zero-based.
}

func (LogFindtime) classID() uint16 { return 0x0e21 }

// LogFindtime1 (Input) Find index of a log entry based on a given time
// Class/Id 0x21 0x0e (12 bytes)
// This message can be used for a time-based search of a log. It can find the index of the first log entry with time equal to the given time, otherwise the index of the most recent entry with time less than the given time. This index can then be used with the UBX-LOG-RETRIEVE message to provide time-based retrieval of log entries. Searching a log is effective for a given time later than the base date (January 1st, 2004). Searching a log for a given time earlier than the base date will result in an 'entry not found' response. (Searching a log for a given time earlier than the base date will result in a UBX-ACK-NAK message in protocol versions less than 18). Searching a log for a given time greater than the last recorded entry's time will return the index of the last recorded entry. (If the logging has stopped due to lack of file space, such a search will result in a UBX-ACK-NAK message in protocol versions less than 18).
type LogFindtime1 struct {
	Version   byte    // Message version (0x00 for this version)
	Type      byte    // Message type, 0 for request
	Reserved1 [2]byte // Reserved
	Year      uint16  // Year (1-65635) of UTC time
	Month     byte    // Month (1-12) of UTC time
	Day       byte    // Day (1-31) of UTC time
	Hour      byte    // Hour (0-23) of UTC time
	Minute    byte    // Minute (0-59) of UTC time
	Second    byte    // Second (0-60) of UTC time
	Reserved2 byte    // Reserved
}

func (LogFindtime1) classID() uint16 { return 0x0e21 }

// Message ubx-log-info (2 versions)

// LogInfo (Poll Request) Poll for log information
// Class/Id 0x21 0x08 (0 bytes)
// Upon sending of this message, the receiver returns UBX-LOG-INFO as defined below.
type LogInfo struct {
}

func (LogInfo) classID() uint16 { return 0x0821 }

// LogInfo1 (Output) Log information
// Class/Id 0x21 0x08 (48 bytes)
// This message is used to report information about the logging subsystem. Note:  The reported maximum log size will be smaller than that originally specified in  LOG-CREATE due to logging and filestore implementation overheads.  Log entries are compressed in a variable length fashion, so it may be difficult  to predict log space usage with any precision.  There may be times when the receiver does not have an accurate time (e.g. if  the week number is not yet known), in which case some entries will not have a  timestamp. This may result in the oldest/newest entry time values not taking  account of these entries.
type LogInfo1 struct {
	Version                 byte           // Message version (0x01 for this version)
	Reserved1               [3]byte        // Reserved
	FilestoreCapacity_bytes uint32         // [bytes] The capacity of the filestore
	Reserved2               [8]byte        // Reserved
	CurrentMaxLogSize_bytes uint32         // [bytes] The maximum size the current log is allowed to grow to
	CurrentLogSize_bytes    uint32         // [bytes] Approximate amount of space in log currently occupied
	EntryCount              uint32         // Number of entries in the log. Note: for circular logs this value will decrease when a group of entries is deleted to make space for new ones.
	OldestYear              uint16         // Oldest entry UTC year (1-65635) or zero if there are no entries with known time
	OldestMonth             byte           // Oldest month (1-12)
	OldestDay               byte           // Oldest day (1-31)
	OldestHour              byte           // Oldest hour (0-23)
	OldestMinute            byte           // Oldest minute (0-59)
	OldestSecond            byte           // Oldest second (0-60)
	Reserved3               byte           // Reserved
	NewestYear              uint16         // Newest year (1-65635) or zero if there are no entries with known time
	NewestMonth             byte           // Newest month (1-12)
	NewestDay               byte           // Newest day (1-31)
	NewestHour              byte           // Newest hour (0-23)
	NewestMinute            byte           // Newest minute (0-59)
	NewestSecond            byte           // Newest second (0-60)
	Reserved4               byte           // Reserved
	Status                  LogInfo1Status // Log status flags
	Reserved5               [3]byte        // Reserved
}

func (LogInfo1) classID() uint16 { return 0x0821 }

type LogInfo1Status byte

const (
	LogInfo1Recording LogInfo1Status = 0x8  // Log entry recording is currently turned on
	LogInfo1Inactive  LogInfo1Status = 0x10 // Logging system not active - no log present
	LogInfo1Circular  LogInfo1Status = 0x20 // The current log is circular
)

func (v LogInfo1Status) String() string {
	var b []string
	if v&0x8 != 0 {
		b = append(b, "Recording")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "Inactive")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "Circular")
	}
	v &^= 0x20
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-log-retrieve

// LogRetrieve (Command) Request log data
// Class/Id 0x21 0x09 (12 bytes)
// This message is used to request logged data (log recording must first be disabled, see UBX-CFG-LOGFILTER). Log entries are returned in chronological order, using the messages UBX-LOG- RETRIEVEPOS and UBX-LOG-RETRIEVESTRING. If the odometer was enabled at the time a position was logged, then message UBX-LOG-RETRIEVEPOSEXTRA will also be used. The maximum number of entries that can be returned in response to a single UBX-LOG-RETRIEVE message is 256. If more entries than this are required the message will need to be sent multiple times with different startNumbers. The retrieve will be stopped if any UBX-LOG message is received. The speed of transfer can be maximized by using a high data rate and temporarily stopping the GPS processing (see UBX-CFG-RST).
type LogRetrieve struct {
	StartNumber uint32  // Index of first log entry to be transferred. If it is larger than the index of the last available log entry, then the first log entry to be transferred is the last available log entry. The indexing of log entries is zero- based.
	EntryCount  uint32  // Number of log entries to transfer in total including the first entry to be transferred. If it is larger than the log entries available starting from the first entry to be transferred, then only the available log entries are transferred followed by a UBX- ACK-NAK. The maximum is 256.
	Version     byte    // Message version (0x00 for this version)
	Reserved1   [3]byte // Reserved
}

func (LogRetrieve) classID() uint16 { return 0x0921 }

// Message ubx-log-retrievebatch

// LogRetrievebatch (Command) Request batch data
// Class/Id 0x21 0x10 (4 bytes)
// This message is used to request batched data. Batch entries are returned in chronological order, using one UBX-LOG-BATCH per navigation epoch. The speed of transfer can be maximized by using a high data rate. See Data Batching for more information.
type LogRetrievebatch struct {
	Version   byte                  // Message version (0x00 for this version)
	Flags     LogRetrievebatchFlags // Flags
	Reserved1 [2]byte               // Reserved
}

func (LogRetrievebatch) classID() uint16 { return 0x1021 }

type LogRetrievebatchFlags byte

const (
	LogRetrievebatchSendMonFirst LogRetrievebatchFlags = 0x1 // Send UBX-MON-BATCH message before sending the UBX-LOG-BATCH message(s).
)

func (v LogRetrievebatchFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "SendMonFirst")
	}
	v &^= 0x1
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-log-retrievepos

// LogRetrievepos (Output) Position fix log entry
// Class/Id 0x21 0x0b (40 bytes)
// This message is used to report a position fix log entry
type LogRetrievepos struct {
	EntryIndex    uint32 // The index of this log entry
	Lon_dege7     int32  // [1e-7 deg] Longitude
	Lat_dege7     int32  // [1e-7 deg] Latitude
	HMSL_mm       int32  // [mm] Height above mean sea level
	HAcc_mm       uint32 // [mm] Horizontal accuracy estimate
	GSpeed_mm_s   uint32 // [mm/s] Ground speed (2-D)
	Heading_dege5 uint32 // [1e-5 deg] Heading
	Version       byte   // Message version (0x00 for this version)
	FixType       byte   // Fix type: 0x01: Dead Reckoning only 0x02: 2D-Fix 0x03: 3D-Fix 0x04: GNSS + Dead Reckoning combined
	Year          uint16 // Year (1-65635) of UTC time
	Month         byte   // Month (1-12) of UTC time
	Day           byte   // Day (1-31) of UTC time
	Hour          byte   // Hour (0-23) of UTC time
	Minute        byte   // Minute (0-59) of UTC time
	Second        byte   // Second (0-60) of UTC time
	Reserved1     byte   // Reserved
	NumSV         byte   // Number of satellites used in the position fix
	Reserved2     byte   // Reserved
}

func (LogRetrievepos) classID() uint16 { return 0x0b21 }

// Message ubx-log-retrieveposextra

// LogRetrieveposextra (Output) Odometer log entry
// Class/Id 0x21 0x0f (32 bytes)
// This message is used to report an odometer log entry
type LogRetrieveposextra struct {
	EntryIndex uint32   // The index of this log entry
	Version    byte     // Message version (0x00 for this version)
	Reserved1  byte     // Reserved
	Year       uint16   // Year (1-65635) of UTC time. Will be zero if time not known
	Month      byte     // Month (1-12) of UTC time
	Day        byte     // Day (1-31) of UTC time
	Hour       byte     // Hour (0-23) of UTC time
	Minute     byte     // Minute (0-59) of UTC time
	Second     byte     // Second (0-60) of UTC time
	Reserved2  [3]byte  // Reserved
	Distance   uint32   // Odometer distance traveled since the last time the odometer was reset by a UBX- NAV-RESETODO
	Reserved3  [12]byte // Reserved
}

func (LogRetrieveposextra) classID() uint16 { return 0x0f21 }

// Message ubx-log-retrievestring

// LogRetrievestring (Output) Byte string log entry
// Class/Id 0x21 0x0d (16 + N*1 bytes)
// This message is used to report a byte string log entry
type LogRetrievestring struct {
	EntryIndex uint32 // The index of this log entry
	Version    byte   // Message version (0x00 for this version)
	Reserved1  byte   // Reserved
	Year       uint16 // Year (1-65635) of UTC time. Will be zero if time not known
	Month      byte   // Month (1-12) of UTC time
	Day        byte   // Day (1-31) of UTC time
	Hour       byte   // Hour (0-23) of UTC time
	Minute     byte   // Minute (0-59) of UTC time
	Second     byte   // Second (0-60) of UTC time
	Reserved2  byte   // Reserved
	ByteCount  uint16 `len:"Items"` // Size of string in bytes
	Items      string
}

func (LogRetrievestring) classID() uint16 { return 0x0d21 }

// Message ubx-log-string

// LogString (Command) Store arbitrary string in on-board flash
// Class/Id 0x21 0x04 (0 + N*1 bytes)
// This message can be used to store an arbitrary byte string in the on-board flash memory. The maximum length that can be stored is 256 bytes.
type LogString string

func (LogString) classID() uint16 { return 0x0421 }

// Message ubx-mga-ack

// MgaAckData0 (Output) Multiple GNSS acknowledge message
// Class/Id 0x13 0x60 (8 bytes)
// This message is sent by a u-blox receiver to acknowledge the receipt of an assistance message. Acknowledgments are enabled by setting the ackAiding parameter in the UBX- CFG-NAVX5 message. See the description of flow control for details.
type MgaAckData0 struct {
	Type            byte    // Type of acknowledgment: 0: The message was not used by the receiver (see infoCode field for an indication of why) 1: The message was accepted for use by the receiver (the infoCode field will be 0)
	Version         byte    // Message version (0x00 for this version)
	InfoCode        byte    // Provides greater information on what the receiver chose to do with the message contents: 0: The receiver accepted the data 1: The receiver does not know the time so it cannot use the data (To resolve this a UBX-MGA-INI-TIME_UTC message should be supplied first) 2: The message version is not supported by the receiver 3: The message size does not match the message version 4: The message data could not be stored to the database 5: The receiver is not ready to use the message data 6: The message type is unknown
	MsgId           byte    // UBX message ID of the acknowledged message
	MsgPayloadStart [4]byte // The first 4 bytes of the acknowledged message's payload
}

func (MgaAckData0) classID() uint16 { return 0x6013 }

// Message ubx-mga-ano

// MgaAno (Input) Multiple GNSS AssistNow Offline assistance
// Class/Id 0x13 0x20 (76 bytes)
// This message is created by the AssistNow Offline service to deliver AssistNow Offline assistance to the receiver. See the description of AssistNow Offline for details.
type MgaAno struct {
	Type      byte     // Message type (0x00 for this type)
	Version   byte     // Message version (0x00 for this version)
	SvId      byte     // Satellite identifier (see Satellite Numbering)
	GnssId    byte     // GNSS identifier (see Satellite Numbering)
	Year      byte     // years since the year 2000
	Month     byte     // month (1..12)
	Day       byte     // day (1..31)
	Reserved1 byte     // Reserved
	Data      [64]byte // assistance data
	Reserved2 [4]byte  // Reserved
}

func (MgaAno) classID() uint16 { return 0x2013 }

// Message ubx-mga-bds (5 versions)

// MgaBdsAlm (Input) BeiDou almanac assistance
// Class/Id 0x13 0x03 (40 bytes)
// This message allows the delivery of BeiDou almanac assistance to a receiver. See the description of AssistNow Online for details.
type MgaBdsAlm struct {
	Type      byte    // Message type (0x02 for this version)
	Version   byte    // Message version (0x00 for this version)
	SvId      byte    // BeiDou satellite identifier (see Satellite Numbering)
	Reserved1 byte    // Reserved
	Wna_week  byte    // [week] Almanac Week Number
	Toa_sr12  byte    // [2^12 s] Almanac reference time
	DeltaI    int16   // [2^-19 semi-circles] Almanac correction of orbit reference inclination at reference time
	SqrtA     uint32  // [2^-11 m^0.5] Almanac square root of semi-major axis
	E         uint32  // Almanac eccentricity
	Omega     int32   // [2^-23 semi-circles] Almanac argument of perigee
	M0        int32   // [2^-23 semi-circles] Almanac mean anomaly at reference time
	Omega0    int32   // [2^-23 semi-circles] Almanac longitude of ascending node of orbit plane at computed according to reference time
	OmegaDot  int32   // [2^-38 semi-circles/s] Almanac rate of right ascension
	A0_sl20   int16   // [2^-20 s] Almanac satellite clock bias
	A1_s_sl38 int16   // [2^-38 s/s] Almanac satellite clock rate
	Reserved2 [4]byte // Reserved
}

func (MgaBdsAlm) classID() uint16 { return 0x0313 }

// MgaBdsEph1 (Input) BeiDou ephemeris assistance
// Class/Id 0x13 0x03 (88 bytes)
// This message allows the delivery of BeiDou ephemeris assistance to a receiver. See the description of AssistNow Online for details.
type MgaBdsEph1 struct {
	Type       byte    // Message type (0x01 for this type)
	Version    byte    // Message version (0x00 for this version)
	SvId       byte    // BeiDou satellite identifier (see Satellite Numbering)
	Reserved1  byte    // Reserved
	SatH1      byte    // Autonomous satellite Health flag
	IODC       byte    // Issue of Data, Clock
	A2_s_s2l66 int16   // [2^-66 s/s^2] Time polynomial coefficient 2
	A1_s_sl50  int32   // [2^-50 s/s] Time polynomial coefficient 1
	A0_sl33    int32   // [2^-33 s] Time polynomial coefficient 0
	Toc_sr3    uint32  // [2^3 s] Clock data reference time
	TGD1_nse1  int16   // [1e-1 ns] Equipment Group Delay Differential
	URAI       byte    // User Range Accuracy Index
	IODE       byte    // Issue of Data, Ephemeris
	Toe_sr3    uint32  // [2^3 s] Ephemeris reference time
	SqrtA      uint32  // [2^-19 m^0.5] Square root of semi-major axis
	E          uint32  // Eccentricity
	Omega      int32   // [2^-31 semi-circles] Argument of perigee
	Deltan     int16   // [2^-43 semi-circles/s] Mean motion difference from computed value
	IDOT       int16   // [2^-43 semi-circles/s] Rate of inclination angle
	M0         int32   // [2^-31 semi-circles] Mean anomaly at reference time
	Omega0     int32   // [2^-31 semi-circles] Longitude of ascending node of orbital of plane computed according to reference time
	OmegaDot   int32   // [2^-43 semi-circles/s] Rate of right ascension
	I0         int32   // [2^-31 semi-circles] Inclination angle at reference time
	Cuc        int32   // [2^-31 semi-circles] Amplitude of cosine harmonic correction term to the argument of latitude
	Cus        int32   // [2^-31 semi-circles] Amplitude of sine harmonic correction term to the argument of latitude
	Crc_ml6    int32   // [2^-6 m] Amplitude of cosine harmonic correction term to the orbit radius
	Crs_ml6    int32   // [2^-6 m] Amplitude of sine harmonic correction term to the orbit radius
	Cic        int32   // [2^-31 semi-circles] Amplitude of cosine harmonic correction term to the angle of inclination
	Cis        int32   // [2^-31 semi-circles] Amplitude of sine harmonic correction term to the angle of inclination
	Reserved2  [4]byte // Reserved
}

func (MgaBdsEph1) classID() uint16 { return 0x0313 }

// MgaBdsHealth2 (Input) BeiDou health assistance
// Class/Id 0x13 0x03 (68 bytes)
// This message allows the delivery of BeiDou health assistance to a receiver. See the description of AssistNow Online for details.
type MgaBdsHealth2 struct {
	Type       byte       // Message type (0x04 for this type)
	Version    byte       // Message version (0x00 for this version)
	Reserved1  [2]byte    // Reserved
	HealthCode [30]uint16 // Each two-byte value represents a BeiDou SV (1-30). The 9 LSBs of each byte contain the 9 bit health code from subframe 5 pages 7,8 of the D1 message, and from subframe 5 pages 35,36 of the D1 message.
	Reserved2  [4]byte    // Reserved
}

func (MgaBdsHealth2) classID() uint16 { return 0x0313 }

// MgaBdsIono3 (Input) BeiDou ionosphere assistance
// Class/Id 0x13 0x03 (16 bytes)
// This message allows the delivery of BeiDou ionospheric assistance to a receiver. See the description of AssistNow Online for details.
type MgaBdsIono3 struct {
	Type            byte    // Message type (0x06 for this type)
	Version         byte    // Message version (0x00 for this version)
	Reserved1       [2]byte // Reserved
	Alpha0_sl30     int8    // [2^-30 s] Ionospheric parameter alpha0
	Alpha1_s_pil27  int8    // [2^-27 s/pi] Ionospheric parameter alpha1
	Alpha2_s_pi2l24 int8    // [2^-24 s/pi^2] Ionospheric parameter alpha2
	Alpha3          int8    // [2^-24 s/pi^3] Ionospheric parameter alpha3
	Beta0_sr11      int8    // [2^11 s] Ionospheric parameter beta0
	Beta1_s_pir14   int8    // [2^14 s/pi] Ionospheric parameter beta1
	Beta2_s_pi2r16  int8    // [2^16 s/pi^2] Ionospheric parameter beta2
	Beta3           int8    // [2^16 s/pi^3] Ionospheric parameter beta3
	Reserved2       [4]byte // Reserved
}

func (MgaBdsIono3) classID() uint16 { return 0x0313 }

// MgaBdsUtc4 (Input) BeiDou UTC assistance
// Class/Id 0x13 0x03 (20 bytes)
// This message allows the delivery of BeiDou UTC assistance to a receiver. See the description of AssistNow Online for details.
type MgaBdsUtc4 struct {
	Type         byte    // Message type (0x05 for this type)
	Version      byte    // Message version (0x00 for this version)
	Reserved1    [2]byte // Reserved
	A0UTC_sl30   int32   // [2^-30 s] BDT clock bias relative to UTC
	A1UTC_s_sl50 int32   // [2^-50 s/s] BDT clock rate relative to UTC
	DtLS_s       int8    // [s] Delta time due to leap seconds before the new leap second effective
	Reserved2    [1]byte // Reserved
	WnRec_week   byte    // [week] BeiDou week number of reception of this UTC parameter set (8-bit truncated)
	WnLSF_week   byte    // [week] Week number of the new leap second
	DN_day       byte    // [day] Day number of the new leap second
	DtLSF_s      int8    // [s] Delta time due to leap seconds after the new leap second effective
	Reserved3    [2]byte // Reserved
}

func (MgaBdsUtc4) classID() uint16 { return 0x0313 }

// Message ubx-mga-dbd (2 versions)

// MgaDbd (Poll Request) Poll the navigation database
// Class/Id 0x13 0x80 (0 bytes)
// Poll the whole navigation data base. The receiver will send all available data from its internal database. The receiver will indicate the finish of the transmission with a UBX-MGA-ACK. The msgPayloadStart field of the UBX-MGA-ACK message will contain a U4 representing the number of UBX-MGA-DBD-DATA* messages sent.
type MgaDbd struct {
}

func (MgaDbd) classID() uint16 { return 0x8013 }

// MgaDbd1 (Input/Output) Navigation database dump entry
// Class/Id 0x13 0x80 (12 + N*1 bytes)
// UBX-MGA-DBD messages are only intended to be sent back to the same receiver that generated them. Navigation database entry. The data fields are firmware-specific. Transmission of this type of message will be acknowledged by UBX-MGA-ACK messages, if acknowledgment has been enabled. See the description of flow control for details. The maximum payload size for firmware 2.01 onwards is 164 bytes (which makes the maximum message size 172 bytes).
type MgaDbd1 struct {
	Reserved1 [12]byte // Reserved
	Items     string
}

func (MgaDbd1) classID() uint16 { return 0x8013 }

// Message ubx-mga-flash (3 versions)

// MgaFlashAck (Output) Acknowledge last FLASH-DATA or -STOP
// Class/Id 0x13 0x21 (6 bytes)
// This message reports an ACK/NACK to the host for the last MGA-FLASH type 1 or type 2 message message received. See Flash-based AssistNow Offline for details.
type MgaFlashAck struct {
	Type      byte   `stf:"3"` // Message type (0x03 for this type)
	Version   byte   // Message version (0x00 for this version)
	Ack       byte   // Acknowledgment type. 0 - ACK: Message received and written to flash. 1 - NACK: Problem with last message, re- transmission required (this only happens while acknowledging a UBX-MGA_FLASH_ DATA message). 2 - NACK: problem with last message, give up.
	Reserved1 byte   // Reserved
	Sequence  uint16 // If acknowledging a UBX-MGA-FLASH- DATA message this is the Message sequence number being ack'ed. If acknowledging a UBX-MGA-FLASH-STOP message it will be set to 0xffff.
}

func (MgaFlashAck) classID() uint16 { return 0x2113 }

// MgaFlashData1 (Input) Transfer MGA-ANO data block to flash
// Class/Id 0x13 0x21 (6 + N*1 bytes)
// This message is used to transfer a block of MGA-ANO data from host to the receiver. Upon reception of this message, the receiver will write the payload data to its internal non-volatile memory (flash). Also, on reception of the first MGA- FLASH-DATA message, the receiver will erase the flash allocated to storing any existing MGA-ANO data. The payload can be up to 512 bytes. Payloads larger than this would exceed the receiver's internal buffering capabilities. The receiver will ACK/NACK this message using the message alternatives given below. The host shall wait for an acknowledge message before sending the next data block. See Flash-based AssistNow Offline for details.
type MgaFlashData1 struct {
	Type     byte   `stf:"1"` // Message type (0x01 for this type)
	Version  byte   // Message version (0x00 for this version)
	Sequence uint16 // Message sequence number, starting at 0 and increamenting by 1 for each MGA- FLASH-DATA message sent.
	Size     uint16 `len:"Items"` // Payload size in bytes.
	Items    string
}

func (MgaFlashData1) classID() uint16 { return 0x2113 }

// MgaFlashStop2 (Input) Finish flashing MGA-ANO data
// Class/Id 0x13 0x21 (2 bytes)
// This message is used to tell the receiver that there are no more MGA-FLASH type 1 messages coming, and that it can do any final internal operations needed to commit the data to flash as a background activity. A UBX-MGA-ACK message will be sent at the end of this process. Note that there may be a delay of several seconds before the UBX-MGA-ACK for this message is sent because of the time taken for this processing. See Flash-based AssistNow Offline for details.
type MgaFlashStop2 struct {
	Type    byte `stf:"2"` // Message type (0x02 for this type)
	Version byte // Message version (0x00 for this version)
}

func (MgaFlashStop2) classID() uint16 { return 0x2113 }

// Message ubx-mga-gal (4 versions)

// MgaGalAlm (Input) Galileo almanac assistance
// Class/Id 0x13 0x02 (32 bytes)
// This message allows the delivery of Galileo almanac assistance to a receiver. See the description of AssistNow Online for details.
type MgaGalAlm struct {
	Type        byte    // Message type (0x02 for this type)
	Version     byte    // Message version (0x00 for this version)
	SvId        byte    // Galileo Satellite identifier (see Satellite Numbering)
	Reserved1   byte    // Reserved
	Ioda        byte    // Almanac Issue of Data
	AlmWNa_week byte    // [week] Almanac reference week number
	Toa_s       uint16  // [600 s] Almanac reference time
	DeltaSqrtA  int16   // [2^-9 m^0.5] Difference with respect to the square root of the nominal semi-major axis (29 600 km)
	E           uint16  // Eccentricity
	DeltaI      int16   // [2^-14 semi-circles] Inclination at reference time relative to i0 = 56 degree
	Omega0      int16   // [2^-15 semi-circles] Longitude of ascending node of orbital plane at weekly epoch
	OmegaDot    int16   // [2^-33 semi-circles/s] Rate of change of right ascension
	Omega       int16   // [2^-15 semi-circles] Argument of perigee
	M0          int16   // [2^-15 semi-circles] Satellite mean anomaly at reference time
	Af0_sl19    int16   // [2^-19 s] Satellite clock correction bias 'truncated'
	Af1_s_sl38  int16   // [2^-38 s/s] Satellite clock correction linear 'truncated'
	HealthE1B   byte    // Satellite E1-B signal health status
	HealthE5b   byte    // Satellite E5b signal health status
	Reserved2   [4]byte // Reserved
}

func (MgaGalAlm) classID() uint16 { return 0x0213 }

// MgaGalEph1 (Input) Galileo ephemeris assistance
// Class/Id 0x13 0x02 (76 bytes)
// This message allows the delivery of Galileo ephemeris assistance to a receiver. See the description of AssistNow Online for details.
type MgaGalEph1 struct {
	Type              byte    // Message type (0x01 for this type)
	Version           byte    // Message version (0x00 for this version)
	SvId              byte    // Galileo Satellite identifier (see Satellite Numbering)
	Reserved1         byte    // Reserved
	IodNav            uint16  // Ephemeris and clock correction Issue of Data
	DeltaN            int16   // [2^-43 semi-circles/s] Mean motion difference from computed value
	M0                int32   // [2^-31 semi-circles] Mean anomaly at reference time
	E                 uint32  // Eccentricity
	SqrtA             uint32  // [2^-19 m^0.5] Square root of the semi-major axis
	Omega0            int32   // [2^-31 semi-circles] Longitude of ascending node of orbital plane at weekly epoch
	I0                int32   // [2^-31 semi-circles] Inclination angle at reference time
	Omega             int32   // [2^-31 semi-circles] Argument of perigee
	OmegaDot          int32   // [2^-43 semi-circles/s] Rate of change of right ascension
	IDot              int16   // [2^-43 semi-circles/s] Rate of change of inclination angle
	Cuc_radiansl29    int16   // [2^-29 radians] Amplitude of the cosine harmonic correction term to the argument of latitude
	Cus_radiansl29    int16   // [2^-29 radians] Amplitude of the sine harmonic correction term to the argument of latitude
	Crc_radiansl5     int16   // [2^-5 radians] Amplitude of the cosine harmonic correction term to the orbit radius
	Crs_radiansl5     int16   // [2^-5 radians] Amplitude of the sine harmonic correction term to the orbit radius
	Cic_radiansl29    int16   // [2^-29 radians] Amplitude of the cosine harmonic correction term to the angle of inclination
	Cis_radiansl29    int16   // [2^-29 radians] Amplitude of the sine harmonic correction term to the angle of inclination
	Toe_s             uint16  // [60 s] Ephemeris reference time
	Af0_sl34          int32   // [2^-34 s] SV clock bias correction coefficient
	Af1_s_sl46        int32   // [2^-46 s/s] SV clock drift correction coefficient
	Af2_s_ssquaredl59 int8    // [2^-59 s/ssquared] SV clock drift rate correction coefficient
	SisaIndexE1E5b    byte    // Signal-In-Space Accuracy index for dual frequency E1-E5b
	Toc_s             uint16  // [60 s] Clock correction data reference Time of Week
	BgdE1E5b          int16   // E1-E5b Broadcast Group Delay
	Reserved2         [2]byte // Reserved
	HealthE1B         byte    // E1-B Signal Health Status
	DataValidityE1B   byte    // E1-B Data Validity Status
	HealthE5b         byte    // E5b Signal Health Status
	DataValidityE5b   byte    // E5b Data Validity Status
	Reserved3         [4]byte // Reserved
}

func (MgaGalEph1) classID() uint16 { return 0x0213 }

// MgaGalTimeoffset2 (Input) Galileo GPS time offset assistance
// Class/Id 0x13 0x02 (12 bytes)
// This message allows the delivery of Galileo time to GPS time offset. See the description of AssistNow Online for details.
type MgaGalTimeoffset2 struct {
	Type       byte    // Message type (0x03 for this type)
	Version    byte    // Message version (0x00 for this version)
	Reserved1  [2]byte // Reserved
	A0G_sl35   int16   // [2^-35 s] Constant term of the polynomial describing the offset
	A1G_s_sl51 int16   // [2^-51 s/s] Rate of change of the offset
	T0G_s      byte    // [3600 s] Reference time for GGTO data
	Wn0G_weeks byte    // [weeks] Week Number of GGTO reference
	Reserved2  [2]byte // Reserved
}

func (MgaGalTimeoffset2) classID() uint16 { return 0x0213 }

// MgaGalUtc3 (Input) Galileo UTC assistance
// Class/Id 0x13 0x02 (20 bytes)
// This message allows the delivery of Galileo UTC assistance to a receiver. See the description of AssistNow Online for details.
type MgaGalUtc3 struct {
	Type        byte    // Message type (0x05 for this type)
	Version     byte    // Message version (0x00 for this version)
	Reserved1   [2]byte // Reserved
	A0_sl30     int32   // [2^-30 s] First parameter of UTC polynomial
	A1_s_sl50   int32   // [2^-50 s/s] Second parameter of UTC polynomial
	DtLS_s      int8    // [s] Delta time due to current leap seconds
	Tot_s       byte    // [3600 s] UTC parameters reference time of week (Galileo time)
	Wnt_weeks   byte    // [weeks] UTC parameters reference week number (the 8-bit WNt field)
	WnLSF_weeks byte    // [weeks] Week number at the end of which the future leap second becomes effective (the 8-bit WNLSF field)
	DN_days     byte    // [days] Day number at the end of which the future leap second becomes effective
	DTLSF_s     int8    // [s] Delta time due to future leap seconds
	Reserved2   [2]byte // Reserved
}

func (MgaGalUtc3) classID() uint16 { return 0x0213 }

// Message ubx-mga-glo (3 versions)

// MgaGloAlm (Input) GLONASS almanac assistance
// Class/Id 0x13 0x06 (36 bytes)
// This message allows the delivery of GLONASS almanac assistance to a receiver. See the description of AssistNow Online for details.
type MgaGloAlm struct {
	Type        byte    // Message type (0x02 for this type)
	Version     byte    // Message version (0x00 for this version)
	SvId        byte    // GLONASS Satellite identifier (see Satellite Numbering)
	Reserved1   byte    // Reserved
	N_days      uint16  // [days] Reference calender day number of almanac within the four-year period (from string 5)
	M           byte    // Type of GLONASS satellite (1 indicates GLONASS-M)
	C           byte    // Unhealthy flag at instant of almanac upload (1 indicates operability of satellite)
	Tau_sl18    int16   // [2^-18 s] Coarse time correction to GLONASS time
	Epsilon     uint16  // Eccentricity
	Lambda      int32   // [2^-20 semi-circles] Longitude of the first (within the N-day) ascending node of satellite orbit in PC-90. 02 coordinate system
	DeltaI      int32   // [2^-20 semi-circles] Correction to the mean value of inclination
	TLambda_sl5 uint32  // [2^-5 s] Time of the first ascending node passage
	DeltaT      int32   // [2^-9 s/orbital-period] Correction to the mean value of Draconian period
	DeltaDT     int8    // [2^-14 s/orbital-period^2] Rate of change of Draconian period
	H           int8    // Carrier frequency number of navigation RF signal, Range=(-7 .. 6)
	Omega       int16   // Argument of perigee
	Reserved2   [4]byte // Reserved
}

func (MgaGloAlm) classID() uint16 { return 0x0613 }

// MgaGloEph1 (Input) GLONASS ephemeris assistance
// Class/Id 0x13 0x06 (48 bytes)
// This message allows the delivery of GLONASS ephemeris assistance to a receiver. See the description of AssistNow Online for details.
type MgaGloEph1 struct {
	Type          byte    // Message type (0x01 for this type)
	Version       byte    // Message version (0x00 for this version)
	SvId          byte    // GLONASS Satellite identifier (see Satellite Numbering)
	Reserved1     byte    // Reserved
	FT            byte    // User range accuracy
	B             byte    // Health flag from string 2
	M             byte    // Type of GLONASS satellite (1 indicates GLONASS-M)
	H             int8    // Carrier frequency number of navigation RF signal, Range=(-7 .. 6), -128 for unknown
	X_kml11       int32   // [2^-11 km] X component of the SV position in PZ-90. 02 coordinate System
	Y_kml11       int32   // [2^-11 km] Y component of the SV position in PZ-90. 02 coordinate System
	Z_kml11       int32   // [2^-11 km] Z component of the SV position in PZ-90. 02 coordinate System
	Dx_km_sl20    int32   // [2^-20 km/s] X component of the SV velocity in PZ-90. 02 coordinate System
	Dy_km_sl20    int32   // [2^-20 km/s] Y component of the SV velocity in PZ-90. 02 coordinate System
	Dz_km_sl20    int32   // [2^-20 km/s] Z component of the SV velocity in PZ-90. 02 coordinate System
	Ddx_km_s2l30  int8    // [2^-30 km/s^2] X component of the SV acceleration in PZ- 90.02 coordinate System
	Ddy_km_s2l30  int8    // [2^-30 km/s^2] Y component of the SV acceleration in PZ- 90.02 coordinate System
	Ddz_km_s2l30  int8    // [2^-30 km/s^2] Z component of the SV acceleration in PZ- 90.02 coordinate System
	Tb_minutes    byte    // [15 minutes] Index of a time interval within current day according to UTC(SU)
	Gamma         int16   // Relative carrier frequency deviation
	E_days        byte    // [days] Ephemeris data age indicator
	DeltaTau_sl30 int8    // [2^-30 s] Time difference between L2 and L1 band
	Tau_sl30      int32   // [2^-30 s] SV clock bias
	Reserved2     [4]byte // Reserved
}

func (MgaGloEph1) classID() uint16 { return 0x0613 }

// MgaGloTimeoffset2 (Input) GLONASS auxiliary time offset assistance
// Class/Id 0x13 0x06 (20 bytes)
// This message allows the delivery of auxiliary GLONASS assistance (including the GLONASS time offsets to other GNSS systems) to a receiver. See the description of AssistNow Online for details.
type MgaGloTimeoffset2 struct {
	Type        byte    // Message type (0x03 for this type)
	Version     byte    // Message version (0x00 for this version)
	N_days      uint16  // [days] Reference calendar day number within the four-year period of almanac (from string 5)
	TauC_sl27   int32   // [2^-27 s] Time scale correction to UTC(SU) time
	TauGps_sl31 int32   // [2^-31 s] Correction to GPS time relative to GLONASS time
	B1_sl10     int16   // [2^-10 s] Coefficient to determine delta UT1
	B2_s_msdl16 int16   // [2^-16 s/msd] Rate of change of delta UT1
	Reserved1   [4]byte // Reserved
}

func (MgaGloTimeoffset2) classID() uint16 { return 0x0613 }

// Message ubx-mga-gps (5 versions)

// MgaGpsAlm (Input) GPS almanac assistance
// Class/Id 0x13 0x00 (36 bytes)
// This message allows the delivery of GPS almanac assistance to a receiver. See the description of AssistNow Online for details.
type MgaGpsAlm struct {
	Type        byte    // Message type (0x02 for this type)
	Version     byte    // Message version (0x00 for this version)
	SvId        byte    // GPS Satellite identifier (see Satellite Numbering)
	SvHealth    byte    // SV health information
	E           uint16  // Eccentricity
	AlmWNa_week byte    // [week] Reference week number of almanac (the 8-bit WNa field)
	Toa_sr12    byte    // [2^12 s] Reference time of almanac
	DeltaI      int16   // [2^-19 semi-circles] Delta inclination angle at reference time
	OmegaDot    int16   // [2^-38 semi-circles/s] Rate of right ascension
	SqrtA       uint32  // [2^-11 m^0.5] Square root of the semi-major axis
	Omega0      int32   // [2^-23 semi-circles] Longitude of ascending node of orbit plane
	Omega       int32   // [2^-23 semi-circles] Argument of perigee
	M0          int32   // [2^-23 semi-circles] Mean anomaly at reference time
	Af0_sl20    int16   // [2^-20 s] Time polynomial coefficient 0 (8 MSBs)
	Af1_s_sl38  int16   // [2^-38 s/s] Time polynomial coefficient 1
	Reserved1   [4]byte // Reserved
}

func (MgaGpsAlm) classID() uint16 { return 0x0013 }

// MgaGpsEph1 (Input) GPS ephemeris assistance
// Class/Id 0x13 0x00 (68 bytes)
// This message allows the delivery of GPS ephemeris assistance to a receiver. See the description of AssistNow Online for details.
type MgaGpsEph1 struct {
	Type              byte    // Message type (0x01 for this type)
	Version           byte    // Message version (0x00 for this version)
	SvId              byte    // GPS Satellite identifier (see Satellite Numbering)
	Reserved1         byte    // Reserved
	FitInterval       byte    // Fit interval flag
	UraIndex          byte    // URA index
	SvHealth          byte    // SV health
	Tgd_sl31          int8    // [2^-31 s] Group delay differential
	Iodc              uint16  // IODC
	Toc_sr4           uint16  // [2^4 s] Clock data reference time
	Reserved2         byte    // Reserved
	Af2_s_ssquaredl55 int8    // [2^-55 s/ssquared] Time polynomial coefficient 2
	Af1_s_sl43        int16   // [2^-43 s/s] Time polynomial coefficient 1
	Af0_sl31          int32   // [2^-31 s] Time polynomial coefficient 0
	Crs_ml5           int16   // [2^-5 m] Crs
	DeltaN            int16   // [2^-43 semi-circles/s] Mean motion difference from computed value
	M0                int32   // [2^-31 semi-circles] Mean anomaly at reference time
	Cuc_radiansl29    int16   // [2^-29 radians] Amplitude of cosine harmonic correction term to argument of latitude
	Cus_radiansl29    int16   // [2^-29 radians] Amplitude of sine harmonic correction term to argument of latitude
	E                 uint32  // Eccentricity
	SqrtA             uint32  // [2^-19 m^0.5] Square root of the semi-major axis
	Toe_sr4           uint16  // [2^4 s] Reference time of ephemeris
	Cic_radiansl29    int16   // [2^-29 radians] Amplitude of cos harmonic correction term to angle of inclination
	Omega0            int32   // [2^-31 semi-circles] Longitude of ascending node of orbit plane at weekly epoch
	Cis_radiansl29    int16   // [2^-29 radians] Amplitude of sine harmonic correction term to angle of inclination
	Crc_ml5           int16   // [2^-5 m] Amplitude of cosine harmonic correction term to orbit radius
	I0                int32   // [2^-31 semi-circles] Inclination angle at reference time
	Omega             int32   // [2^-31 semi-circles] Argument of perigee
	OmegaDot          int32   // [2^-43 semi-circles/s] Rate of right ascension
	Idot              int16   // [2^-43 semi-circles/s] Rate of inclination angle
	Reserved3         [2]byte // Reserved
}

func (MgaGpsEph1) classID() uint16 { return 0x0013 }

// MgaGpsHealth2 (Input) GPS health assistance
// Class/Id 0x13 0x00 (40 bytes)
// This message allows the delivery of GPS health assistance to a receiver. See the description of AssistNow Online for details.
type MgaGpsHealth2 struct {
	Type       byte     // Message type (0x04 for this type)
	Version    byte     // Message version (0x00 for this version)
	Reserved1  [2]byte  // Reserved
	HealthCode [32]byte // Each byte represents a GPS SV (1-32). The 6 LSBs of each byte contains the 6 bit health code from subframes 4/5 page 25.
	Reserved2  [4]byte  // Reserved
}

func (MgaGpsHealth2) classID() uint16 { return 0x0013 }

// MgaGpsIono3 (Input) GPS ionosphere assistance
// Class/Id 0x13 0x00 (16 bytes)
// This message allows the delivery of GPS ionospheric assistance to a receiver. See the description of AssistNow Online for details.
type MgaGpsIono3 struct {
	Type            byte    // Message type (0x06 for this type)
	Version         byte    // Message version (0x00 for this version)
	Reserved1       [2]byte // Reserved
	IonoAlpha0_sl30 int8    // [2^-30 s] Ionospheric parameter alpha0 [s]
	IonoAlpha1      int8    // [2^-27 s/semi-circle] Ionospheric parameter alpha1 [s/semi- circle]
	IonoAlpha2      int8    // [2^-24 s/(semi-circle^2)] Ionospheric parameter alpha2 [s/semi- circle^2]
	IonoAlpha3      int8    // [2^-24 s/(semi-circle^3)] Ionospheric parameter alpha3 [s/semi- circle^3]
	IonoBeta0_sr11  int8    // [2^11 s] Ionospheric parameter beta0 [s]
	IonoBeta1       int8    // [2^14 s/semi-circle] Ionospheric parameter beta1 [s/semi- circle]
	IonoBeta2       int8    // [2^16 s/(semi-circle^2)] Ionospheric parameter beta2 [s/semi- circle^2]
	IonoBeta3       int8    // [2^16 s/(semi-circle^3)] Ionospheric parameter beta3 [s/semi- circle^3]
	Reserved2       [4]byte // Reserved
}

func (MgaGpsIono3) classID() uint16 { return 0x0013 }

// MgaGpsUtc4 (Input) GPS UTC assistance
// Class/Id 0x13 0x00 (20 bytes)
// This message allows the delivery of GPS UTC assistance to a receiver. See the description of AssistNow Online for details.
type MgaGpsUtc4 struct {
	Type           byte    // Message type (0x05 for this type)
	Version        byte    // Message version (0x00 for this version)
	Reserved1      [2]byte // Reserved
	UtcA0_sl30     int32   // [2^-30 s] First parameter of UTC polynomial
	UtcA1_s_sl50   int32   // [2^-50 s/s] Second parameter of UTC polynomial
	UtcDtLS_s      int8    // [s] Delta time due to current leap seconds
	UtcTot_sr12    byte    // [2^12 s] UTC parameters reference time of week (GPS time)
	UtcWNt_weeks   byte    // [weeks] UTC parameters reference week number (the 8-bit WNt field)
	UtcWNlsf_weeks byte    // [weeks] Week number at the end of which the future leap second becomes effective (the 8-bit WNLSF field)
	UtcDn_days     byte    // [days] Day number at the end of which the future leap second becomes effective
	UtcDtLSF_s     int8    // [s] Delta time due to future leap seconds
	Reserved2      [2]byte // Reserved
}

func (MgaGpsUtc4) classID() uint16 { return 0x0013 }

// Message ubx-mga-ini (7 versions)

// MgaIniClkd (Input) Initial clock drift assistance
// Class/Id 0x13 0x40 (12 bytes)
// Supplying clock drift assistance that is inaccurate by more than the specified accuracy, may lead to substantially degraded receiver performance. This message allows the delivery of clock drift assistance to a receiver. See the description of AssistNow Online for details.
type MgaIniClkd struct {
	Type         byte    `stf:"0x20"` // Message type (0x20 for this type)
	Version      byte    // Message version (0x00 for this version)
	Reserved1    [2]byte // Reserved
	ClkD_ns_s    int32   // [ns/s] Clock drift
	ClkDAcc_ns_s uint32  // [ns/s] Clock drift accuracy
}

func (MgaIniClkd) classID() uint16 { return 0x4013 }

// MgaIniEop1 (Input) Earth orientation parameters assistance
// Class/Id 0x13 0x40 (72 bytes)
// This message allows the delivery of new earth orientation parameters (EOP) to a receiver to improve AssistNow Autonomous operation.
type MgaIniEop1 struct {
	Type             byte     `stf:"0x30"` // Message type (0x30 for this type)
	Version          byte     // Message version (0x00 for this version)
	Reserved1        [2]byte  // Reserved
	D2kRef_d         uint16   // [d] reference time (days since 1.1.2000 12.00h UTC)
	D2kMax_d         uint16   // [d] expiration time (days since 1.1.2000 12.00h UTC)
	XpP0_arcsecl30   int32    // [2^-30 arcsec] x_p t^0 polynomial term (offset)
	XpP1_arcsec_dl30 int32    // [2^-30 arcsec/d] x_p t^1 polynomial term (drift)
	YpP0_arcsecl30   int32    // [2^-30 arcsec] y_p t^0 polynomial term (offset)
	YpP1_arcsec_dl30 int32    // [2^-30 arcsec/d] y_p t^1 polynomial term (drift)
	DUT1_sl25        int32    // [2^-25 s] dUT1 t^0 polynomial term (offset)
	DdUT1_s_dl30     int32    // [2^-30 s/d] dUT1 t^1 polynomial term (drift)
	Reserved2        [40]byte // Reserved
}

func (MgaIniEop1) classID() uint16 { return 0x4013 }

// MgaIniFreq2 (Input) Initial frequency assistance
// Class/Id 0x13 0x40 (12 bytes)
// Supplying external frequency assistance that is inaccurate by more than the specified accuracy, may lead to substantially degraded receiver performance. This message allows the delivery of external frequency assistance to a receiver. See the description of AssistNow Online for details.
type MgaIniFreq2 struct {
	Type        byte             `stf:"0x21"` // Message type (0x21 for this type)
	Version     byte             // Message version (0x00 for this version)
	Reserved1   byte             // Reserved
	Flags       MgaIniFreq2Flags // Frequency reference
	Freq_hze2   int32            // [1e-2 Hz] Frequency
	FreqAcc_ppb uint32           // [ppb] Frequency accuracy
}

func (MgaIniFreq2) classID() uint16 { return 0x4013 }

type MgaIniFreq2Flags byte

const (
	MgaIniFreq2Source MgaIniFreq2Flags = 0xf  // 0: frequency available on EXTINT0 1: frequency available on EXTINT1 2-15: reserved
	MgaIniFreq2Fall   MgaIniFreq2Flags = 0x10 // use falling edge of EXTINT pulse (default rising)
)

func (v MgaIniFreq2Flags) String() string {
	var b []string
	if v&0xf != 0 {
		b = append(b, fmt.Sprintf("Source<%b>", (v&0xf)>>0))
	}
	v &^= 0xf
	if v&0x10 != 0 {
		b = append(b, "Fall")
	}
	v &^= 0x10
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// MgaIniPos_llh3 (Input) Initial position assistance
// Class/Id 0x13 0x40 (20 bytes)
// Supplying position assistance that is inaccurate by more than the specified position accuracy, may lead to substantially degraded receiver performance. This message allows the delivery of initial position assistance to a receiver in WGS84 lat/long/alt coordinates. This message is equivalent to the UBX-MGA- INI-POS_XYZ message, except for the coordinate system. See the description of AssistNow Online for details.
type MgaIniPos_llh3 struct {
	Type      byte    `stf:"1"` // Message type (0x01 for this type)
	Version   byte    // Message version (0x00 for this version)
	Reserved1 [2]byte // Reserved
	Lat_dege7 int32   // [1e-7 deg] WGS84 Latitude
	Lon_dege7 int32   // [1e-7 deg] WGS84 Longitude
	Alt_cm    int32   // [cm] WGS84 Altitude
	PosAcc_cm uint32  // [cm] Position accuracy (stddev)
}

func (MgaIniPos_llh3) classID() uint16 { return 0x4013 }

// MgaIniPos_xyz4 (Input) Initial position assistance
// Class/Id 0x13 0x40 (20 bytes)
// Supplying position assistance that is inaccurate by more than the specified position accuracy, may lead to substantially degraded receiver performance. This message allows the delivery of initial position assistance to a receiver in cartesian ECEF coordinates. This message is equivalent to the UBX-MGA-INI- POS_LLH message, except for the coordinate system. See the description of AssistNow Online for details.
type MgaIniPos_xyz4 struct {
	Type      byte    `stf:"0"` // Message type (0x00 for this type)
	Version   byte    // Message version (0x00 for this version)
	Reserved1 [2]byte // Reserved
	EcefX_cm  int32   // [cm] WGS84 ECEF X coordinate
	EcefY_cm  int32   // [cm] WGS84 ECEF Y coordinate
	EcefZ_cm  int32   // [cm] WGS84 ECEF Z coordinate
	PosAcc_cm uint32  // [cm] Position accuracy (stddev)
}

func (MgaIniPos_xyz4) classID() uint16 { return 0x4013 }

// MgaIniTime_gnss5 (Input) Initial time assistance
// Class/Id 0x13 0x40 (24 bytes)
// Supplying time assistance that is inaccurate by more than the specified time accuracy, may lead to substantially degraded receiver performance. This message allows the delivery of time assistance to a receiver in a chosen GNSS timebase. This message is equivalent to the UBX-MGA-INI-TIME_UTC message, except for the time base. See the description of AssistNow Online for details.
type MgaIniTime_gnss5 struct {
	Type      byte                `stf:"0x11"` // Message type (0x11 for this type)
	Version   byte                // Message version (0x00 for this version)
	Ref       MgaIniTime_gnss5Ref // Reference to be used to set time
	GnssId    byte                // Source of time information. Currently supported: 0: GPS time 2: Galileo time 3: BeiDou time 6: GLONASS time: week = 834 + ((N4- 1)*1461 + Nt)/7, tow = (((N4-1)*1461 + Nt) % 7) * 86400 + tod
	Reserved1 [2]byte             // Reserved
	Week      uint16              // GNSS week number
	Tow_s     uint32              // [s] GNSS time of week
	Ns_ns     uint32              // [ns] GNSS time of week, nanosecond part from 0 to 999,999,999
	TAccS_s   uint16              // [s] Seconds part of time accuracy
	Reserved2 [2]byte             // Reserved
	TAccNs_ns uint32              // [ns] Nanoseconds part of time accuracy, from 0 to 999,999,999
}

func (MgaIniTime_gnss5) classID() uint16 { return 0x4013 }

type MgaIniTime_gnss5Ref byte

const (
	MgaIniTime_gnss5Source MgaIniTime_gnss5Ref = 0xf  // 0: none, i.e. on receipt of message (will be inaccurate!) 1: relative to pulse sent to EXTINT0 2: relative to pulse sent to EXTINT1 3-15: reserved
	MgaIniTime_gnss5Fall   MgaIniTime_gnss5Ref = 0x10 // use falling edge of EXTINT pulse (default rising) - only if source is EXTINT
	MgaIniTime_gnss5Last   MgaIniTime_gnss5Ref = 0x20 // use last EXTINT pulse (default next pulse) - only if source is EXTINT
)

func (v MgaIniTime_gnss5Ref) String() string {
	var b []string
	if v&0xf != 0 {
		b = append(b, fmt.Sprintf("Source<%b>", (v&0xf)>>0))
	}
	v &^= 0xf
	if v&0x10 != 0 {
		b = append(b, "Fall")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "Last")
	}
	v &^= 0x20
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// MgaIniTime_utc6 (Input) Initial time assistance
// Class/Id 0x13 0x40 (24 bytes)
// Supplying time assistance that is inaccurate by more than the specified time accuracy, may lead to substantially degraded receiver performance. This message allows the delivery of UTC time assistance to a receiver. This message is equivalent to the UBX-MGA-INI-TIME_GNSS message, except for the time base. See the description of AssistNow Online for details.
type MgaIniTime_utc6 struct {
	Type       byte               `stf:"0x10"` // Message type (0x10 for this type)
	Version    byte               // Message version (0x00 for this version)
	Ref        MgaIniTime_utc6Ref // Reference to be used to set time
	LeapSecs_s int8               // [s] Number of leap seconds since 1980 (or 0x80 = -128 if unknown)
	Year       uint16             // Year
	Month      byte               // Month, starting at 1
	Day        byte               // Day, starting at 1
	Hour       byte               // Hour, from 0 to 23
	Minute     byte               // Minute, from 0 to 59
	Second_s   byte               // [s] Seconds, from 0 to 59
	Reserved1  byte               // Reserved
	Ns_ns      uint32             // [ns] Nanoseconds, from 0 to 999,999,999
	TAccS_s    uint16             // [s] Seconds part of time accuracy
	Reserved2  [2]byte            // Reserved
	TAccNs_ns  uint32             // [ns] Nanoseconds part of time accuracy, from 0 to 999,999,999
}

func (MgaIniTime_utc6) classID() uint16 { return 0x4013 }

type MgaIniTime_utc6Ref byte

const (
	MgaIniTime_utc6Source MgaIniTime_utc6Ref = 0xf  // 0: none, i.e. on receipt of message (will be inaccurate!) 1: relative to pulse sent to EXTINT0 2: relative to pulse sent to EXTINT1 3-15: reserved
	MgaIniTime_utc6Fall   MgaIniTime_utc6Ref = 0x10 // use falling edge of EXTINT pulse (default rising) - only if source is EXTINT
	MgaIniTime_utc6Last   MgaIniTime_utc6Ref = 0x20 // use last EXTINT pulse (default next pulse) - only if source is EXTINT
)

func (v MgaIniTime_utc6Ref) String() string {
	var b []string
	if v&0xf != 0 {
		b = append(b, fmt.Sprintf("Source<%b>", (v&0xf)>>0))
	}
	v &^= 0xf
	if v&0x10 != 0 {
		b = append(b, "Fall")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "Last")
	}
	v &^= 0x20
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-mga-qzss (3 versions)

// MgaQzssAlm (Input) QZSS almanac assistance
// Class/Id 0x13 0x05 (36 bytes)
// This message allows the delivery of QZSS almanac assistance to a receiver. See the description of AssistNow Online for details.
type MgaQzssAlm struct {
	Type        byte    // Message type (0x02 for this type)
	Version     byte    // Message version (0x00 for this version)
	SvId        byte    // QZSS Satellite identifier (see Satellite Numbering), Range 1-5
	SvHealth    byte    // Almanac SV health information
	E           uint16  // Almanac eccentricity
	AlmWNa_week byte    // [week] Reference week number of almanac (the 8-bit WNa field)
	Toa_sr12    byte    // [2^12 s] Reference time of almanac
	DeltaI      int16   // [2^-19 semi-circles] Delta inclination angle at reference time
	OmegaDot    int16   // [2^-38 semi-circles/s] Almanac rate of right ascension
	SqrtA       uint32  // [2^-11 m^0.5] Almanac square root of the semi-major axis A
	Omega0      int32   // [2^-23 semi-circles] Almanac long of asc node of orbit plane at weekly
	Omega       int32   // [2^-23 semi-circles] Almanac argument of perigee
	M0          int32   // [2^-23 semi-circles] Almanac mean anomaly at reference time
	Af0_sl20    int16   // [2^-20 s] Almanac time polynomial coefficient 0 (8 MSBs)
	Af1_s_sl38  int16   // [2^-38 s/s] Almanac time polynomial coefficient 1
	Reserved1   [4]byte // Reserved
}

func (MgaQzssAlm) classID() uint16 { return 0x0513 }

// MgaQzssEph1 (Input) QZSS ephemeris assistance
// Class/Id 0x13 0x05 (68 bytes)
// This message allows the delivery of QZSS ephemeris assistance to a receiver. See the description of AssistNow Online for details.
type MgaQzssEph1 struct {
	Type              byte    // Message type (0x01 for this type)
	Version           byte    // Message version (0x00 for this version)
	SvId              byte    // QZSS Satellite identifier (see Satellite Numbering), Range 1-5
	Reserved1         byte    // Reserved
	FitInterval       byte    // Fit interval flag
	UraIndex          byte    // URA index
	SvHealth          byte    // SV health
	Tgd_sl31          int8    // [2^-31 s] Group delay differential
	Iodc              uint16  // IODC
	Toc_sr4           uint16  // [2^4 s] Clock data reference time
	Reserved2         byte    // Reserved
	Af2_s_ssquaredl55 int8    // [2^-55 s/ssquared] Time polynomial coefficient 2
	Af1_s_sl43        int16   // [2^-43 s/s] Time polynomial coefficient 1
	Af0_sl31          int32   // [2^-31 s] Time polynomial coefficient 0
	Crs_ml5           int16   // [2^-5 m] Crs
	DeltaN            int16   // [2^-43 semi-circles/s] Mean motion difference from computed value
	M0                int32   // [2^-31 semi-circles] Mean anomaly at reference time
	Cuc_radiansl29    int16   // [2^-29 radians] Amp of cosine harmonic corr term to arg of lat
	Cus_radiansl29    int16   // [2^-29 radians] Amp of sine harmonic corr term to arg of lat
	E                 uint32  // eccentricity
	SqrtA             uint32  // [2^-19 m^0.5] Square root of the semi-major axis A
	Toe_sr4           uint16  // [2^4 s] Reference time of ephemeris
	Cic_radiansl29    int16   // [2^-29 radians] Amp of cos harmonic corr term to angle of inclination
	Omega0            int32   // [2^-31 semi-circles] Long of asc node of orbit plane at weekly epoch
	Cis_radiansl29    int16   // [2^-29 radians] Amp of sine harmonic corr term to angle of inclination
	Crc_ml5           int16   // [2^-5 m] Amp of cosine harmonic corr term to orbit radius
	I0                int32   // [2^-31 semi-circles] Inclination angle at reference time
	Omega             int32   // [2^-31 semi-circles] Argument of perigee
	OmegaDot          int32   // [2^-43 semi-circles/s] Rate of right ascension
	Idot              int16   // [2^-43 semi-circles/s] Rate of inclination angle
	Reserved3         [2]byte // Reserved
}

func (MgaQzssEph1) classID() uint16 { return 0x0513 }

// MgaQzssHealth2 (Input) QZSS health assistance
// Class/Id 0x13 0x05 (12 bytes)
// This message allows the delivery of QZSS health assistance to a receiver. See the description of AssistNow Online for details.
type MgaQzssHealth2 struct {
	Type       byte    // Message type (0x04 for this type)
	Version    byte    // Message version (0x00 for this version)
	Reserved1  [2]byte // Reserved
	HealthCode [5]byte // Each byte represents a QZSS SV (1-5). The 6 LSBs of each byte contains the 6 bit health code from subframes 4/5, data ID = 3, SV ID = 51
	Reserved2  [3]byte // Reserved
}

func (MgaQzssHealth2) classID() uint16 { return 0x0513 }

// Message ubx-mon-batch

// MonBatch (Polled) Data batching buffer status
// Class/Id 0x0a 0x32 (12 bytes)
// This message contains status information about the batching buffer. It can be polled and it can also be sent by the receiver as a response to a UBX- LOG-RETRIEVEBATCH message before the UBX-LOG-BATCH messages. See Data Batching for more information.
type MonBatch struct {
	Version       byte    // Message version (0x00 for this version)
	Reserved1     [3]byte // Reserved
	FillLevel     uint16  // Current buffer fill level, i.e. number of epochs currently stored
	DropsAll      uint16  // Number of dropped epochs since startup Note: changing the batching configuration will reset this counter.
	DropsSinceMon uint16  // Number of dropped epochs since last MON-BATCH message
	NextMsgCnt    uint16  // The next retrieved UBX-LOG-BATCH will have this msgCnt value.
}

func (MonBatch) classID() uint16 { return 0x320a }

// Message ubx-mon-gnss

// MonGnss (Polled) Information message major GNSS selection
// Class/Id 0x0a 0x28 (8 bytes)
// This message reports major GNSS selection. It does this by means of bit masks in U1 fields. Each bit in a bit mask corresponds to one major GNSS. Augmentation systems are not reported.
type MonGnss struct {
	Version      byte               // Message version (0x01for this version)
	Supported    MonGnssSupported   // A bit mask showing the major GNSS that can be supported by this receiver
	DefaultGnss  MonGnssDefaultGnss // A bit mask showing the default major GNSS selection. If the default major GNSS selection is currently configured in the efuse for this receiver, it takes precedence over the default major GNSS selection configured in the executing firmware of this receiver.
	Enabled      MonGnssEnabled     // A bit mask showing the current major GNSS selection enabled for this receiver
	Simultaneous byte               // Maximum number of concurrent major GNSS that can be supported by this receiver
	Reserved1    [3]byte            // Reserved
}

func (MonGnss) classID() uint16 { return 0x280a }

type MonGnssSupported byte

const (
	MonGnssGPSSup     MonGnssSupported = 0x1 // GPS is supported
	MonGnssGlonassSup MonGnssSupported = 0x2 // GLONASS is supported
	MonGnssBeidouSup  MonGnssSupported = 0x4 // BeiDou is supported
	MonGnssGalileoSup MonGnssSupported = 0x8 // Galileo is supported
)

type MonGnssDefaultGnss byte

const (
	MonGnssGPSDef     MonGnssDefaultGnss = 0x1 // GPS is default-enabled
	MonGnssGlonassDef MonGnssDefaultGnss = 0x2 // GLONASS is default-enabled
	MonGnssBeidouDef  MonGnssDefaultGnss = 0x4 // BeiDou is default-enabled
	MonGnssGalileoDef MonGnssDefaultGnss = 0x8 // Galileo is default-enabled
)

type MonGnssEnabled byte

const (
	MonGnssGPSEna     MonGnssEnabled = 0x1 // GPS is enabled
	MonGnssGlonassEna MonGnssEnabled = 0x2 // GLONASS is enabled
	MonGnssBeidouEna  MonGnssEnabled = 0x4 // BeiDou is enabled
	MonGnssGalileoEna MonGnssEnabled = 0x8 // Galileo is enabled
)

func (v MonGnssSupported) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "GPSSup")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "GlonassSup")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "BeidouSup")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "GalileoSup")
	}
	v &^= 0x8
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v MonGnssDefaultGnss) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "GPSDef")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "GlonassDef")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "BeidouDef")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "GalileoDef")
	}
	v &^= 0x8
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v MonGnssEnabled) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "GPSEna")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "GlonassEna")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "BeidouEna")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "GalileoEna")
	}
	v &^= 0x8
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-mon-hw

// MonHw (Periodic/polled) Hardware status
// Class/Id 0x0a 0x09 (60 bytes)
// Status of different aspect of the hardware, such as antenna, PIO/peripheral pins, noise level, automatic gain control (AGC)
type MonHw struct {
	PinSel     uint32     // Mask of pins set as peripheral/PIO
	PinBank    uint32     // Mask of pins set as bank A/B
	PinDir     uint32     // Mask of pins set as input/output
	PinVal     uint32     // Mask of pins value low/high
	NoisePerMS uint16     // Noise level as measured by the GPS core
	AgcCnt     uint16     // AGC monitor (counts SIGHI xor SIGLO, range 0 to 8191)
	AStatus    byte       // Status of the antenna supervisor state machine (0=INIT, 1=DONTKNOW, 2=OK, 3=SHORT, 4=OPEN)
	APower     byte       // Current power status of antenna (0=OFF, 1=ON, 2=DONTKNOW)
	Flags      MonHwFlags // Flags
	Reserved1  byte       // Reserved
	UsedMask   uint32     // Mask of pins that are used by the virtual pin manager
	VP         [17]byte   // Array of pin mappings for each of the 17 physical pins
	JamInd     byte       // CW jamming indicator, scaled (0 = no CW jamming, 255 = strong CW jamming)
	Reserved2  [2]byte    // Reserved
	PinIrq     uint32     // Mask of pins value using the PIO Irq
	PullH      uint32     // Mask of pins value using the PIO pull high resistor
	PullL      uint32     // Mask of pins value using the PIO pull low resistor
}

func (MonHw) classID() uint16 { return 0x090a }

type MonHwFlags byte

const (
	MonHwRtcCalib     MonHwFlags = 0x1  // RTC is calibrated
	MonHwSafeBoot     MonHwFlags = 0x2  // Safeboot mode (0 = inactive, 1 = active)
	MonHwJammingState MonHwFlags = 0xc  // Output from jamming/interference monitor (0 = unknown or feature disabled, 1 = ok - no significant jamming, 2 = warning - interference visible but fix OK, 3 = critical - interference visible and no fix)
	MonHwXtalAbsent   MonHwFlags = 0x10 // RTC xtal has been determined to be absent (not supported in protocol versions less than 18)
)

func (v MonHwFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "RtcCalib")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "SafeBoot")
	}
	v &^= 0x2
	if v&0xc != 0 {
		b = append(b, fmt.Sprintf("JammingState<%b>", (v&0xc)>>2))
	}
	v &^= 0xc
	if v&0x10 != 0 {
		b = append(b, "XtalAbsent")
	}
	v &^= 0x10
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-mon-hw2

// MonHw2 (Periodic/Polled) Extended hardware status
// Class/Id 0x0a 0x0b (28 bytes)
// Status of different aspects of the hardware such as Imbalance, Low-Level Configuration and POST Results. The first four parameters of this message represent the complex signal from the RF front end. The following rules of thumb apply:  The smaller the absolute value of the variable ofsI and ofsQ, the better.  Ideally, the magnitude of the I-part (magI) and the Q-part (magQ) of the  complex signal should be the same.
type MonHw2 struct {
	OfsI       int8    // Imbalance of I-part of complex signal, scaled (-128 = max. negative imbalance, 127 = max. positive imbalance)
	MagI       byte    // Magnitude of I-part of complex signal, scaled (0 = no signal, 255 = max. magnitude)
	OfsQ       int8    // Imbalance of Q-part of complex signal, scaled (-128 = max. negative imbalance, 127 = max. positive imbalance)
	MagQ       byte    // Magnitude of Q-part of complex signal, scaled (0 = no signal, 255 = max. magnitude)
	CfgSource  byte    // Source of low-level configuration (114 = ROM, 111 = OTP, 112 = config pins, 102 = flash image)
	Reserved1  [3]byte // Reserved
	LowLevCfg  uint32  // Low-level configuration (obsolete in protocol versions greater than 15)
	Reserved2  [8]byte // Reserved
	PostStatus uint32  // POST status word
	Reserved3  [4]byte // Reserved
}

func (MonHw2) classID() uint16 { return 0x0b0a }

// Message ubx-mon-io

// MonIo (Periodic/Polled) I/O system status
// Class/Id 0x0a 0x02 (0 + N*20 bytes)
// The size of the message is determined by the number of ports 'N' the receiver supports, i.e. on u-blox 5 the number of ports is 6.
type MonIo struct {
	Items []*MonIoItemsType // len: N
}

func (MonIo) classID() uint16 { return 0x020a }

type MonIoItemsType struct {
	RxBytes_bytes uint32  // [bytes] Number of bytes ever received
	TxBytes_bytes uint32  // [bytes] Number of bytes ever sent
	ParityErrs    uint16  // Number of 100 ms timeslots with parity errors
	FramingErrs   uint16  // Number of 100 ms timeslots with framing errors
	OverrunErrs   uint16  // Number of 100 ms timeslots with overrun errors
	BreakCond     uint16  // Number of 100 ms timeslots with break conditions
	Reserved1     [4]byte // Reserved
}

// Message ubx-mon-msgpp

// MonMsgpp (Periodic/Polled) Message parse and process status
// Class/Id 0x0a 0x06 (120 bytes)
// -
type MonMsgpp struct {
	Msg1_msgs     [8]uint16 // [msgs] Number of successfully parsed messages for each protocol on port0
	Msg2_msgs     [8]uint16 // [msgs] Number of successfully parsed messages for each protocol on port1
	Msg3_msgs     [8]uint16 // [msgs] Number of successfully parsed messages for each protocol on port2
	Msg4_msgs     [8]uint16 // [msgs] Number of successfully parsed messages for each protocol on port3
	Msg5_msgs     [8]uint16 // [msgs] Number of successfully parsed messages for each protocol on port4
	Msg6_msgs     [8]uint16 // [msgs] Number of successfully parsed messages for each protocol on port5
	Skipped_bytes [6]uint32 // [bytes] Number skipped bytes for each port
}

func (MonMsgpp) classID() uint16 { return 0x060a }

// Message ubx-mon-patch (2 versions)

// MonPatch (Poll Request) Poll request for installed patches
// Class/Id 0x0a 0x27 (0 bytes)
// -
type MonPatch struct {
}

func (MonPatch) classID() uint16 { return 0x270a }

// MonPatch1 (Polled) Installed patches
// Class/Id 0x0a 0x27 (4 + N*16 bytes)
// This message reports information about patches installed and currently enabled on the receiver. It does not report on patches installed and then disabled. An enabled patch is considered active when the receiver executes from the code space where the patch resides on. For example, a ROM patch is reported active only when the system runs from ROM.
type MonPatch1 struct {
	Version  uint16                // Message version (0x0001 for this version)
	NEntries uint16                `len:"Items"` // Total number of reported patches
	Items    []*MonPatch1ItemsType // len: NEntries
}

func (MonPatch1) classID() uint16 { return 0x270a }

type MonPatch1ItemsType struct {
	PatchInfo        MonPatch1PatchInfo // Status information about the reported patch
	ComparatorNumber uint32             // The number of the comparator
	PatchAddress     uint32             // The address that is targeted by the patch
	PatchData        uint32             // The data that is inserted at the patchAddress
}

type MonPatch1PatchInfo uint32

const (
	MonPatch1Activated MonPatch1PatchInfo = 0x1 // 1: the patch is active, 0: otherwise
	MonPatch1Location  MonPatch1PatchInfo = 0x6 // Indicates where the patch is stored. 0: eFuse, 1: ROM, 2: BBR, 3: file system
)

func (v MonPatch1PatchInfo) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Activated")
	}
	v &^= 0x1
	if v&0x6 != 0 {
		b = append(b, fmt.Sprintf("Location<%b>", (v&0x6)>>1))
	}
	v &^= 0x6
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-mon-rxbuf

// MonRxbuf (Periodic/Polled) Receiver buffer status
// Class/Id 0x0a 0x07 (24 bytes)
// -
type MonRxbuf struct {
	Pending_bytes [6]uint16 // [bytes] Number of bytes pending in receiver buffer for each target
	Usage         [6]byte   // [%] Maximum usage receiver buffer during the last sysmon period for each target
	PeakUsage     [6]byte   // [%] Maximum usage receiver buffer for each target
}

func (MonRxbuf) classID() uint16 { return 0x070a }

// Message ubx-mon-rxr

// MonRxr (Output) Receiver status information
// Class/Id 0x0a 0x21 (1 bytes)
// The receiver ready message is sent when the receiver changes from or to backup mode.
type MonRxr struct {
	Flags MonRxrFlags // Receiver status flags
}

func (MonRxr) classID() uint16 { return 0x210a }

type MonRxrFlags byte

const (
	MonRxrAwake MonRxrFlags = 0x1 // not in backup mode
)

func (v MonRxrFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Awake")
	}
	v &^= 0x1
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-mon-smgr

// MonSmgr (Periodic/Polled) Synchronization manager status
// Class/Id 0x0a 0x2e (16 bytes)
// This message reports the status of internal and external oscillators and sources as well as whether GNSS is used for disciplining.
type MonSmgr struct {
	Version   byte           // Message version (0x00 for this version)
	Reserved1 [3]byte        // Reserved
	ITOW_ms   uint32         // [ms] Time of the week
	IntOsc    MonSmgrIntOsc  // A bit mask, indicating the status of the local oscillator
	ExtOsc    MonSmgrExtOsc  // A bit mask, indicating the status of the external oscillator
	DiscSrc   byte           // Disciplining source identifier: 0: internal oscillator 1: GNSS 2: EXTINT0 3: EXTINT1 4: internal oscillator measured by the host 5: external oscillator measured by the host
	Gnss      MonSmgrGnss    // A bit mask, indicating the status of the GNSS
	ExtInt0   MonSmgrExtInt0 // A bit mask, indicating the status of the external input 0
	ExtInt1   MonSmgrExtInt1 // A bit mask, indicating the status of the external input 1
}

func (MonSmgr) classID() uint16 { return 0x2e0a }

type MonSmgrIntOsc uint16

const (
	MonSmgrIntOscState MonSmgrIntOsc = 0xf  // State of the oscillator: 0: autonomous operation 1: calibration ongoing 2: oscillator is steered by the host 3: idle state
	MonSmgrIntOscCalib MonSmgrIntOsc = 0x10 // 1 = oscillator gain is calibrated
	MonSmgrIntOscDisc  MonSmgrIntOsc = 0x20 // 1 = signal is disciplined
)

type MonSmgrExtOsc uint16

const (
	MonSmgrExtOscState MonSmgrExtOsc = 0xf  // State of the oscillator: 0: autonomous operation 1: calibration ongoing 2: oscillator is steered by the host 3: idle state
	MonSmgrExtOscCalib MonSmgrExtOsc = 0x10 // 1 = oscillator gain is calibrated
	MonSmgrExtOscDisc  MonSmgrExtOsc = 0x20 // 1 = signal is disciplined
)

type MonSmgrGnss byte

const (
	MonSmgrGnssAvail MonSmgrGnss = 0x1 // 1 = GNSS is present
)

type MonSmgrExtInt0 byte

const (
	MonSmgrExtInt0Avail    MonSmgrExtInt0 = 0x1 // 1 = signal present at this input
	MonSmgrExtInt0Type     MonSmgrExtInt0 = 0x2 // Source type: 0: frequency 1: time
	MonSmgrExtInt0FeedBack MonSmgrExtInt0 = 0x4 // This source is used as feedback of the external oscillator
)

type MonSmgrExtInt1 byte

const (
	MonSmgrExtInt1Avail    MonSmgrExtInt1 = 0x1 // 1 = signal present at this input
	MonSmgrExtInt1Type     MonSmgrExtInt1 = 0x2 // Source type: 0: frequency 1: time
	MonSmgrExtInt1FeedBack MonSmgrExtInt1 = 0x4 // This source is used as feedback of the external oscillator
)

func (v MonSmgrIntOsc) String() string {
	var b []string
	if v&0xf != 0 {
		b = append(b, fmt.Sprintf("IntOscState<%b>", (v&0xf)>>0))
	}
	v &^= 0xf
	if v&0x10 != 0 {
		b = append(b, "IntOscCalib")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "IntOscDisc")
	}
	v &^= 0x20
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v MonSmgrExtOsc) String() string {
	var b []string
	if v&0xf != 0 {
		b = append(b, fmt.Sprintf("ExtOscState<%b>", (v&0xf)>>0))
	}
	v &^= 0xf
	if v&0x10 != 0 {
		b = append(b, "ExtOscCalib")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "ExtOscDisc")
	}
	v &^= 0x20
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v MonSmgrGnss) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "GnssAvail")
	}
	v &^= 0x1
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v MonSmgrExtInt0) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "ExtInt0Avail")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "ExtInt0Type")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "ExtInt0FeedBack")
	}
	v &^= 0x4
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v MonSmgrExtInt1) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "ExtInt1Avail")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "ExtInt1Type")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "ExtInt1FeedBack")
	}
	v &^= 0x4
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-mon-spt

// MonSpt (Polled) Sensor production test
// Class/Id 0x0a 0x2f (4 + N*16 bytes)
// This message reports the state of, and measurements made during, sensor self- tests. This message can also be used to retrieve information about detected sensor(s) and driver(s) used. This message is only supported if a sensor is directly connected to the u-blox chip. This includes modules that contain IMUs. Note that this message shows the status of the last self-test since sensor startup. The self-test results are not stored in non-volatile memory.
type MonSpt struct {
	Version   byte                // Message version (0x01 for this version)
	NumSensor byte                `len:"Sensor"` // number of sensors reported in this message
	NumRes    byte                `len:"Res"`    // number of result items reported in this message
	Reserved1 byte                // Reserved
	Sensor    []*MonSptSensorType // len: NumSensor
	Res       []*MonSptResType    // len: NumRes
}

func (MonSpt) classID() uint16 { return 0x2f0a }

type MonSptSensorType struct {
	SensorId    byte         // Sensor ID The following IDs are defined, others are reserved: 1: ST LSM6DS0 6-axis IMU with temperature sensor 2: Invensense MPU6500 6-axis IMU with temperature sensor 3: Bosch BMI160 6-axis IMU with temperature sensor 7: ST LSM6DS3 6-axis IMU with temperature sensor 9: Bosch SMI130 6-axis IMU with temperature sensor 12: MPU6515, 6-axis inertial sensor from Invensense 13: ST LSM6DSL 6-axis IMU with temperature sensor 14: SMG130, 3-axis gyroscope with temperature sensor from Bosch 15: SMI230, 6-axis IMU with temperature sensor from Bosch 16: BMI260, 6-axis IMU with temperature sensor from Bosch 17: ICM330DLC, 6-axis IMU with temperature sensor from ST 18: ICM330DHCX, 6-axis IMU with 105 deg temperature sensor from ST Not all sensors are supported in any released firmware. Refer to the release notes to find out which sensor is supported by a certain firmware.
	DrvVer      MonSptDrvVer // Version information
	TestState   byte         // State of one sensor's test, it can be 0: test not yet started 1: test started but not yet finished 2: test did not finish due to error during execution 3: test finished normally, test data is available
	DrvFileName byte         // 0 if the active driver is loaded from image, last character of the file name if it is loaded from separate file.
}

type MonSptResType struct {
	SensorIdRes uint16  // Sensor ID; eligible values are the same as in sensorIdState field
	SensorType  uint16  // Sensor type and axis (if applicable) to which the result refers The following values are defined, others are reserved: 5: Gyroscope z axis 12: Gyroscope temperature 13: Gyroscope y axis 14: Gyroscope x axis 16: Accelerometer x axis 17: Accelerometer y axis 18: Accelerometer z axis 19: Barometer 22: Magnetometer x axis 23: Magnetometer y axis 24: Magnetometer z axis 25: Barometer temperature
	ResType     uint16  // The type of result stored in the value field 1: Measurement without self-test offset (raw and unscaled digital value) 2: Measurement with positive self-test offset (raw and unscaled digital value) 3: Measurement with negative self-test offset (raw and unscaled digital value) 4: Minimum off-to-positive to pass self- test, as deduced from on-chip trimming information 5: Maximum off-to-positive to pass self- test, as deduced from on-chip trimming information 6: Minimum negative-to-positive to pass self-test, as deduced from on-chip trimming information 7: Maximum negative-to-positive to pass self-test, as deduced from on-chip trimming information 8: Self-test passed; test passed if value = 1 and failed if 0. Used if the decision is read out from the sensor itself.
	Reserved2   [2]byte // Reserved
	Value       int32   // value of the specific test result
}

type MonSptDrvVer byte

const (
	MonSptDrvVerMaj MonSptDrvVer = 0xf  // Driver major version
	MonSptDrvVerMin MonSptDrvVer = 0xf0 // Driver minor version
)

func (v MonSptDrvVer) String() string {
	var b []string
	if v&0xf != 0 {
		b = append(b, fmt.Sprintf("DrvVerMaj<%b>", (v&0xf)>>0))
	}
	v &^= 0xf
	if v&0xf0 != 0 {
		b = append(b, fmt.Sprintf("DrvVerMin<%b>", (v&0xf0)>>4))
	}
	v &^= 0xf0
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-mon-txbuf

// MonTxbuf (Periodic/Polled) Transmitter buffer status
// Class/Id 0x0a 0x08 (28 bytes)
// -
type MonTxbuf struct {
	Pending_bytes [6]uint16      // [bytes] Number of bytes pending in transmitter buffer for each target
	Usage         [6]byte        // [%] Maximum usage transmitter buffer during the last sysmon period for each target
	PeakUsage     [6]byte        // [%] Maximum usage transmitter buffer for each target
	TUsage        byte           // [%] Maximum usage of transmitter buffer during the last sysmon period for all targets
	TPeakusage    byte           // [%] Maximum usage of transmitter buffer for all targets
	Errors        MonTxbufErrors // Error bitmask
	Reserved1     byte           // Reserved
}

func (MonTxbuf) classID() uint16 { return 0x080a }

type MonTxbufErrors byte

const (
	MonTxbufLimit MonTxbufErrors = 0x3f // Buffer limit of corresponding target reached
	MonTxbufMem   MonTxbufErrors = 0x40 // Memory Allocation error
	MonTxbufAlloc MonTxbufErrors = 0x80 // Allocation error (TX buffer full)
)

func (v MonTxbufErrors) String() string {
	var b []string
	if v&0x3f != 0 {
		b = append(b, fmt.Sprintf("Limit<%b>", (v&0x3f)>>0))
	}
	v &^= 0x3f
	if v&0x40 != 0 {
		b = append(b, "Mem")
	}
	v &^= 0x40
	if v&0x80 != 0 {
		b = append(b, "Alloc")
	}
	v &^= 0x80
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-mon-ver (2 versions)

// MonVer (Poll Request) Poll receiver and software version
// Class/Id 0x0a 0x04 (0 bytes)
// -
type MonVer struct {
}

func (MonVer) classID() uint16 { return 0x040a }

// MonVer1 (Polled) Receiver and software version
// Class/Id 0x0a 0x04 (40 + N*30 bytes)
// -
type MonVer1 struct {
	SwVersion [30]byte            // Nul-terminated software version string.
	HwVersion [10]byte            // Nul-terminated hardware version string
	Items     []*MonVer1ItemsType // len: N
}

func (MonVer1) classID() uint16 { return 0x040a }

type MonVer1ItemsType struct {
	Extension [30]byte // Extended software information strings. A series of nul-terminated strings. Each extension field is 30 characters long and contains varying software information. Not all extension fields may appear. Examples of reported information: the software version string of the underlying ROM (when the receiver's firmware is running from flash), the firmware version, the supported protocol version, the module identifier, the flash information structure (FIS) file information, the supported major GNSS, the supported augmentation systems. See Firmware and protocol versions for details.
}

// Message ubx-nav-aopstatus

// NavAopstatus (Periodic/Polled) AssistNow Autonomous status
// Class/Id 0x01 0x60 (16 bytes)
// This message provides information on the status of the AssistNow Autonomous subsystem on the receiver. For example, a host application can determine the optimal time to shut down the receiver by monitoring the status field for a steady 0. See the chapter AssistNow Autonomous in the receiver description for details on this feature.
type NavAopstatus struct {
	ITOW_ms   uint32             // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	AopCfg    NavAopstatusAopCfg // AssistNow Autonomous configuration
	Status    byte               // AssistNow Autonomous subsystem is idle (0) or running (not 0)
	Reserved1 [10]byte           // Reserved
}

func (NavAopstatus) classID() uint16 { return 0x6001 }

type NavAopstatusAopCfg byte

const (
	NavAopstatusUseAOP NavAopstatusAopCfg = 0x1 // AOP enabled flag
)

func (v NavAopstatusAopCfg) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "UseAOP")
	}
	v &^= 0x1
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-att

// NavAtt (Periodic/Polled) Attitude solution
// Class/Id 0x01 0x05 (32 bytes)
// This message outputs the attitude solution as roll, pitch and heading angles. More details about vehicle attitude can be found in the Vehicle Attitude Output (ADR) section for ADR products. More details about vehicle attitude can be found in the Vehicle Attitude Output (UDR) section for UDR products.
type NavAtt struct {
	ITOW_ms          uint32  // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Version          byte    // Message version (0x00 for this version)
	Reserved1        [3]byte // Reserved
	Roll_dege5       int32   // [1e-5 deg] Vehicle roll.
	Pitch_dege5      int32   // [1e-5 deg] Vehicle pitch.
	Heading_dege5    int32   // [1e-5 deg] Vehicle heading.
	AccRoll_dege5    uint32  // [1e-5 deg] Vehicle roll accuracy (if null, roll angle is not available).
	AccPitch_dege5   uint32  // [1e-5 deg] Vehicle pitch accuracy (if null, pitch angle is not available).
	AccHeading_dege5 uint32  // [1e-5 deg] Vehicle heading accuracy (if null, heading angle is not available).
}

func (NavAtt) classID() uint16 { return 0x0501 }

// Message ubx-nav-clock

// NavClock (Periodic/Polled) Clock solution
// Class/Id 0x01 0x22 (20 bytes)
// -
type NavClock struct {
	ITOW_ms   uint32 // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	ClkB_ns   int32  // [ns] Clock bias
	ClkD_ns_s int32  // [ns/s] Clock drift
	TAcc_ns   uint32 // [ns] Time accuracy estimate
	FAcc_ps_s uint32 // [ps/s] Frequency accuracy estimate
}

func (NavClock) classID() uint16 { return 0x2201 }

// Message ubx-nav-cov

// NavCov (Periodic/Polled) Covariance matrices
// Class/Id 0x01 0x36 (64 bytes)
// This message outputs the covariance matrices for the position and velocity solutions in the topocentric coordinate system defined as the local-level North (N), East (E), Down (D) frame. As the covariance matrices are symmetric, only the upper triangular part is output.
type NavCov struct {
	ITOW_ms        uint32  // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Version        byte    // Message version (0x00 for this version)
	PosCovValid    byte    // Position covariance matrix validity flag
	VelCovValid    byte    // Velocity covariance matrix validity flag
	Reserved1      [9]byte // Reserved
	PosCovNN_m2    float32 // [m^2] Position covariance matrix value p_NN
	PosCovNE_m2    float32 // [m^2] Position covariance matrix value p_NE
	PosCovND_m2    float32 // [m^2] Position covariance matrix value p_ND
	PosCovEE_m2    float32 // [m^2] Position covariance matrix value p_EE
	PosCovED_m2    float32 // [m^2] Position covariance matrix value p_ED
	PosCovDD_m2    float32 // [m^2] Position covariance matrix value p_DD
	VelCovNN_m2_s2 float32 // [m^2/s^2] Velocity covariance matrix value v_NN
	VelCovNE_m2_s2 float32 // [m^2/s^2] Velocity covariance matrix value v_NE
	VelCovND_m2_s2 float32 // [m^2/s^2] Velocity covariance matrix value v_ND
	VelCovEE_m2_s2 float32 // [m^2/s^2] Velocity covariance matrix value v_EE
	VelCovED_m2_s2 float32 // [m^2/s^2] Velocity covariance matrix value v_ED
	VelCovDD_m2_s2 float32 // [m^2/s^2] Velocity covariance matrix value v_DD
}

func (NavCov) classID() uint16 { return 0x3601 }

// Message ubx-nav-dgps

// NavDgps (Periodic/Polled) DGPS data used for NAV
// Class/Id 0x01 0x31 (16 + N*12 bytes)
// This message outputs the DGPS correction data that has been applied to the current NAV Solution. See also the notes on the RTCM protocol.
type NavDgps struct {
	ITOW_ms    uint32           // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Age_ms     int32            // [ms] Age of newest correction data
	BaseId     int16            // DGPS base station identifier
	BaseHealth int16            // DGPS base station health status
	NumCh      byte             `len:"Ch"` // Number of channels for which correction data is following
	Status     byte             // DGPS correction type status: 0x00:   none 0x01:  PR+PRR correction
	Reserved1  [2]byte          // Reserved
	Ch         []*NavDgpsChType // len: NumCh
}

func (NavDgps) classID() uint16 { return 0x3101 }

type NavDgpsChType struct {
	Svid     byte         // Satellite ID
	Flags    NavDgpsFlags // Channel number and usage
	AgeC_ms  uint16       // [ms] Age of latest correction data
	Prc_m    float32      // [m] Pseudorange correction
	Prrc_m_s float32      // [m/s] Pseudorange rate correction
}

type NavDgpsFlags byte

const (
	NavDgpsChannel  NavDgpsFlags = 0xf  // GPS channel number this SV is on. Channel numbers in the firmware greater than 15 are displayed as having channel number 15
	NavDgpsDgpsUsed NavDgpsFlags = 0x10 // 1 = DGPS used for this SV
)

func (v NavDgpsFlags) String() string {
	var b []string
	if v&0xf != 0 {
		b = append(b, fmt.Sprintf("Channel<%b>", (v&0xf)>>0))
	}
	v &^= 0xf
	if v&0x10 != 0 {
		b = append(b, "DgpsUsed")
	}
	v &^= 0x10
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-dop

// NavDop (Periodic/Polled) Dilution of precision
// Class/Id 0x01 0x04 (18 bytes)
//  DOP values are dimensionless.  All DOP values are scaled by a factor of 100. If the unit transmits a value of e.g.  156, the DOP value is 1.56.
type NavDop struct {
	ITOW_ms uint32 // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	GDOP    uint16 // Geometric DOP
	PDOP    uint16 // Position DOP
	TDOP    uint16 // Time DOP
	VDOP    uint16 // Vertical DOP
	HDOP    uint16 // Horizontal DOP
	NDOP    uint16 // Northing DOP
	EDOP    uint16 // Easting DOP
}

func (NavDop) classID() uint16 { return 0x0401 }

// Message ubx-nav-eell

// NavEell (Periodic/Polled) Position error ellipse parameters
// Class/Id 0x01 0x3d (16 bytes)
// This message outputs the error ellipse parameters for the position solutions.
type NavEell struct {
	ITOW_ms                uint32 // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Version                byte   // Message version (0x00 for this version)
	Reserved1              byte   // Reserved
	ErrEllipseOrient_dege2 uint16 // [1e-2 deg] Orientation of semi-major axis of error ellipse (degrees from true north)
	ErrEllipseMajor_mm     uint32 // [mm] Semi-major axis of error ellipse
	ErrEllipseMinor_mm     uint32 // [mm] Semi-minor axis of error ellipse
}

func (NavEell) classID() uint16 { return 0x3d01 }

// Message ubx-nav-eoe

// NavEoe (Periodic) End of epoch
// Class/Id 0x01 0x61 (4 bytes)
// This message is intended to be used as a marker to collect all navigation messages of an epoch. It is output after all enabled NAV class messages (except UBX-NAV-HNR) and after all enabled NMEA messages.
type NavEoe struct {
	ITOW_ms uint32 // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
}

func (NavEoe) classID() uint16 { return 0x6101 }

// Message ubx-nav-geofence

// NavGeofence (Periodic/Polled) Geofencing status
// Class/Id 0x01 0x39 (8 + N*2 bytes)
// This message outputs the evaluated states of all configured geofences for the current epoch's position. See the Geofencing description for feature details.
type NavGeofence struct {
	ITOW_ms   uint32                   // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Version   byte                     // Message version (0x00 for this version)
	Status    byte                     // Geofencing status 0 - Geofencing not available or not reliable 1 - Geofencing active
	NumFences byte                     `len:"Fences"` // Number of geofences
	CombState byte                     // Combined (logical OR) state of all geofences 0 - Unknown 1 - Inside 2 - Outside
	Fences    []*NavGeofenceFencesType // len: NumFences
}

func (NavGeofence) classID() uint16 { return 0x3901 }

type NavGeofenceFencesType struct {
	State byte // Geofence state 0 - Unknown 1 - Inside 2 - Outside
	Id    byte // Geofence ID (0 = not available)
}

// Message ubx-nav-hpposecef

// NavHpposecef (Periodic/Polled) High precision position solution in ECEF
// Class/Id 0x01 0x13 (28 bytes)
// See important comments concerning validity of position given in section Navigation Output Filters.
type NavHpposecef struct {
	Version      byte              // Message version (0x00 for this version)
	Reserved1    [3]byte           // Reserved
	ITOW_ms      uint32            // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	EcefX_cm     int32             // [cm] ECEF X coordinate
	EcefY_cm     int32             // [cm] ECEF Y coordinate
	EcefZ_cm     int32             // [cm] ECEF Z coordinate
	EcefXHp_mme1 int8              // [1e-1 mm] High precision component of ECEF X coordinate. Must be in the range of -99.. +99. Precise coordinate in cm = ecefX + (ecefXHp * 1e-2).
	EcefYHp_mme1 int8              // [1e-1 mm] High precision component of ECEF Y coordinate. Must be in the range of -99.. +99. Precise coordinate in cm = ecefY + (ecefYHp * 1e-2).
	EcefZHp_mme1 int8              // [1e-1 mm] High precision component of ECEF Z coordinate. Must be in the range of -99.. +99. Precise coordinate in cm = ecefZ + (ecefZHp * 1e-2).
	Flags        NavHpposecefFlags // Additional flags
	PAcc_mme1    uint32            // [1e-1 mm] Position Accuracy Estimate
}

func (NavHpposecef) classID() uint16 { return 0x1301 }

type NavHpposecefFlags byte

const (
	NavHpposecefInvalidEcef NavHpposecefFlags = 0x1 // 1 = Invalid ecefX, ecefY, ecefZ, ecefXHp, ecefYHp and ecefZHp
)

func (v NavHpposecefFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "InvalidEcef")
	}
	v &^= 0x1
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-hpposllh

// NavHpposllh (Periodic/Polled) High precision geodetic position solution
// Class/Id 0x01 0x14 (36 bytes)
// See important comments concerning validity of position given in section Navigation Output Filters. This message outputs the Geodetic position with high precision in the currently selected ellipsoid. The default is the WGS84 Ellipsoid, but can be changed with the message UBX-CFG-DAT.
type NavHpposllh struct {
	Version       byte             // Message version (0x00 for this version)
	Reserved1     [2]byte          // Reserved
	Flags         NavHpposllhFlags // Additional flags
	ITOW_ms       uint32           // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Lon_dege7     int32            // [1e-7 deg] Longitude
	Lat_dege7     int32            // [1e-7 deg] Latitude
	Height_mm     int32            // [mm] Height above ellipsoid.
	HMSL_mm       int32            // [mm] Height above mean sea level
	LonHp_dege9   int8             // [1e-9 deg] High precision component of longitude. Must be in the range -99..+99. Precise longitude in deg * 1e-7 = lon + (lonHp * 1e- 2).
	LatHp_dege9   int8             // [1e-9 deg] High precision component of latitude. Must be in the range -99..+99. Precise latitude in deg * 1e-7 = lat + (latHp * 1e-2).
	HeightHp_mme1 int8             // [1e-1 mm] High precision component of height above ellipsoid. Must be in the range -9..+9. Precise height in mm = height + (heightHp * 0.1).
	HMSLHp_mme1   int8             // [1e-1 mm] High precision component of height above mean sea level. Must be in range -9..+9. Precise height in mm = hMSL + (hMSLHp * 0.1)
	HAcc_mme1     uint32           // [1e-1 mm] Horizontal accuracy estimate
	VAcc_mme1     uint32           // [1e-1 mm] Vertical accuracy estimate
}

func (NavHpposllh) classID() uint16 { return 0x1401 }

type NavHpposllhFlags byte

const (
	NavHpposllhInvalidLlh NavHpposllhFlags = 0x1 // 1 = Invalid lon, lat, height, hMSL, lonHp, latHp, heightHp and hMSLHp
)

func (v NavHpposllhFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "InvalidLlh")
	}
	v &^= 0x1
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-nmi

// NavNmi (Periodic/Polled) Navigation message cross-check information
// Class/Id 0x01 0x28 (16 bytes)
// Information about the validity of received satellite navigation payload.
type NavNmi struct {
	ITOW_ms     uint32            // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Version     byte              // Message version (0x01 for this version)
	Reserved1   [4]byte           // Reserved
	GpsNmiFlags NavNmiGpsNmiFlags // GPS navigation message cross-check information flags.
	GpsLsFlags  NavNmiGpsLsFlags  // GPS leap second cross-check information flags.
	GalNmiFlags NavNmiGalNmiFlags // Galileo navigation message cross-check information flags.
	GalLsFlags  NavNmiGalLsFlags  // Galileo leap second cross-check information flags.
	BdsNmiFlags NavNmiBdsNmiFlags // BeiDou navigation message cross-check information flags.
	BdsLsFlags  NavNmiBdsLsFlags  // BeiDou leap second cross-check information flags.
	GloNmiFlags NavNmiGloNmiFlags // GLONASS navigation message cross- check information flags.
}

func (NavNmi) classID() uint16 { return 0x2801 }

type NavNmiGpsNmiFlags byte

const (
	NavNmiWnoCheckedGPS NavNmiGpsNmiFlags = 0x1 // 1 = week number check performed.
	NavNmiWnoInvalidGPS NavNmiGpsNmiFlags = 0x2 // 1 = week number invalid.
)

type NavNmiGpsLsFlags byte

const (
	NavNmiLsValGPS    NavNmiGpsLsFlags = 0x1  // 1 = Leap second value out of range.
	NavNmiDnRangeGPS  NavNmiGpsLsFlags = 0x2  // 1 = Day number value out of range.
	NavNmiTotRangeGPS NavNmiGpsLsFlags = 0x4  // 1 = Data reference TOW out of range.
	NavNmiLsEventGPS  NavNmiGpsLsFlags = 0x8  // 1 = Unexpected leap second event.
	NavNmiRecNowGPS   NavNmiGpsLsFlags = 0x10 // 1 = Data received this epoch.
)

type NavNmiGalNmiFlags byte

const (
	NavNmiWnoCheckedGAL NavNmiGalNmiFlags = 0x1 // 1 = week number check performed.
	NavNmiWnoInvalidGAL NavNmiGalNmiFlags = 0x2 // 1 = week number invalid.
)

type NavNmiGalLsFlags byte

const (
	NavNmiLsValGAL    NavNmiGalLsFlags = 0x1  // 1 = Leap second value out of range.
	NavNmiDnRangeGAL  NavNmiGalLsFlags = 0x2  // 1 = Day number value out of range.
	NavNmiTotRangeGAL NavNmiGalLsFlags = 0x4  // 1 = Data reference TOW out of range.
	NavNmiLsEventGAL  NavNmiGalLsFlags = 0x8  // 1 = Unexpected leap second event.
	NavNmiRecNowGAL   NavNmiGalLsFlags = 0x10 // 1 = Data received this epoch.
)

type NavNmiBdsNmiFlags byte

const (
	NavNmiWnoCheckedBDS NavNmiBdsNmiFlags = 0x1 // 1 = week number check performed.
	NavNmiWnoInvalidBDS NavNmiBdsNmiFlags = 0x2 // 1 = week number invalid.
)

type NavNmiBdsLsFlags byte

const (
	NavNmiLsValBDS    NavNmiBdsLsFlags = 0x1  // 1 = Leap second value out of range.
	NavNmiDnRangeBDS  NavNmiBdsLsFlags = 0x2  // 1 = Day number value out of range.
	NavNmiTotRangeBDS NavNmiBdsLsFlags = 0x4  // 1 = Data reference TOW out of range.
	NavNmiLsEventBDS  NavNmiBdsLsFlags = 0x8  // 1 = Unexpected leap second event.
	NavNmiRecNowBDS   NavNmiBdsLsFlags = 0x10 // 1 = Data received this epoch.
)

type NavNmiGloNmiFlags byte

const (
	NavNmiWnoCheckedGLO NavNmiGloNmiFlags = 0x1 // 1 = week number check performed.
	NavNmiWnoInvalidGLO NavNmiGloNmiFlags = 0x2 // 1 = week number invalid.
)

func (v NavNmiGpsNmiFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "WnoCheckedGPS")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "WnoInvalidGPS")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v NavNmiGpsLsFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "LsValGPS")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "DnRangeGPS")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "TotRangeGPS")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "LsEventGPS")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "RecNowGPS")
	}
	v &^= 0x10
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v NavNmiGalNmiFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "WnoCheckedGAL")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "WnoInvalidGAL")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v NavNmiGalLsFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "LsValGAL")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "DnRangeGAL")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "TotRangeGAL")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "LsEventGAL")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "RecNowGAL")
	}
	v &^= 0x10
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v NavNmiBdsNmiFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "WnoCheckedBDS")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "WnoInvalidBDS")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v NavNmiBdsLsFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "LsValBDS")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "DnRangeBDS")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "TotRangeBDS")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "LsEventBDS")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "RecNowBDS")
	}
	v &^= 0x10
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v NavNmiGloNmiFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "WnoCheckedGLO")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "WnoInvalidGLO")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-odo

// NavOdo (Periodic/Polled) Odometer solution
// Class/Id 0x01 0x09 (20 bytes)
// This message outputs the traveled distance since last reset (see UBX-NAV- RESETODO) together with an associated estimated accuracy and the total cumulated ground distance (can only be reset by a cold start of the receiver).
type NavOdo struct {
	Version         byte    // Message version (0x00 for this version)
	Reserved1       [3]byte // Reserved
	ITOW_ms         uint32  // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Distance_m      uint32  // [m] Ground distance since last reset
	TotalDistance_m uint32  // [m] Total cumulative ground distance
	DistanceStd_m   uint32  // [m] Ground distance accuracy (1-sigma)
}

func (NavOdo) classID() uint16 { return 0x0901 }

// Message ubx-nav-orb

// NavOrb (Periodic/Polled) GNSS orbit database info
// Class/Id 0x01 0x34 (8 + N*6 bytes)
// Status of the GNSS orbit database knowledge.
type NavOrb struct {
	ITOW_ms   uint32          // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Version   byte            // Message version (0x01 for this version)
	NumSv     byte            `len:"Sv"` // Number of SVs in the database
	Reserved1 [2]byte         // Reserved
	Sv        []*NavOrbSvType // len: NumSv
}

func (NavOrb) classID() uint16 { return 0x3401 }

type NavOrbSvType struct {
	GnssId   byte           // GNSS ID
	SvId     byte           // Satellite ID
	SvFlag   NavOrbSvFlag   // Information Flags
	Eph      NavOrbEph      // Ephemeris data
	Alm      NavOrbAlm      // Almanac data
	OtherOrb NavOrbOtherOrb // Other orbit data available
}

type NavOrbSvFlag byte

const (
	NavOrbHealth     NavOrbSvFlag = 0x3 // SV health: 0: unknown 1: healthy 2: not healty
	NavOrbVisibility NavOrbSvFlag = 0xc // SV health: 0: unknown 1: below horizon 2: above horizon 3: above elevation mask
)

type NavOrbEph byte

const (
	NavOrbEphUsability NavOrbEph = 0x1f // How long the receiver will be able to use the stored ephemeris data from now on: 31: The usability period is unknown 30: The usability period is more than 450 minutes 30 > n > 0: The usability period is between (n-1)*15 and n*15 minutes 0: Ephemeris can no longer be used
	NavOrbEphSource    NavOrbEph = 0xe0 // 0: not available 1: GNSS transmission 2: external aiding 3-7: other
)

type NavOrbAlm byte

const (
	NavOrbAlmUsability NavOrbAlm = 0x1f // How long the receiver will be able to use the stored almanac data from now on: 31: The usability period is unknown 30: The usability period is more than 30 days 30 > n > 0: The usability period is between n-1 and n days 0: Almanac can no longer be used
	NavOrbAlmSource    NavOrbAlm = 0xe0 // 0: not available 1: GNSS transmission 2: external aiding 3-7: other
)

type NavOrbOtherOrb byte

const (
	NavOrbAnoAopUsability NavOrbOtherOrb = 0x1f // How long the receiver will be able to use the orbit data from now on: 31: The usability period is unknown 30: The usability period is more than 30 days 30 > n > 0: The usability period is between n-1 and n days 0: Data can no longer be used
	NavOrbType            NavOrbOtherOrb = 0xe0 // Type of orbit data: 0: No orbit data available 1: AssistNow Offline data 2: AssistNow Autonomous data 3-7: Other orbit data
)

func (v NavOrbSvFlag) String() string {
	var b []string
	if v&0x3 != 0 {
		b = append(b, fmt.Sprintf("Health<%b>", (v&0x3)>>0))
	}
	v &^= 0x3
	if v&0xc != 0 {
		b = append(b, fmt.Sprintf("Visibility<%b>", (v&0xc)>>2))
	}
	v &^= 0xc
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v NavOrbEph) String() string {
	var b []string
	if v&0x1f != 0 {
		b = append(b, fmt.Sprintf("EphUsability<%b>", (v&0x1f)>>0))
	}
	v &^= 0x1f
	if v&0xe0 != 0 {
		b = append(b, fmt.Sprintf("EphSource<%b>", (v&0xe0)>>5))
	}
	v &^= 0xe0
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v NavOrbAlm) String() string {
	var b []string
	if v&0x1f != 0 {
		b = append(b, fmt.Sprintf("AlmUsability<%b>", (v&0x1f)>>0))
	}
	v &^= 0x1f
	if v&0xe0 != 0 {
		b = append(b, fmt.Sprintf("AlmSource<%b>", (v&0xe0)>>5))
	}
	v &^= 0xe0
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v NavOrbOtherOrb) String() string {
	var b []string
	if v&0x1f != 0 {
		b = append(b, fmt.Sprintf("AnoAopUsability<%b>", (v&0x1f)>>0))
	}
	v &^= 0x1f
	if v&0xe0 != 0 {
		b = append(b, fmt.Sprintf("Type<%b>", (v&0xe0)>>5))
	}
	v &^= 0xe0
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-posecef

// NavPosecef (Periodic/Polled) Position solution in ECEF
// Class/Id 0x01 0x01 (20 bytes)
// See important comments concerning validity of position given in section Navigation Output Filters.
type NavPosecef struct {
	ITOW_ms  uint32 // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	EcefX_cm int32  // [cm] ECEF X coordinate
	EcefY_cm int32  // [cm] ECEF Y coordinate
	EcefZ_cm int32  // [cm] ECEF Z coordinate
	PAcc_cm  uint32 // [cm] Position Accuracy Estimate
}

func (NavPosecef) classID() uint16 { return 0x0101 }

// Message ubx-nav-posllh

// NavPosllh (Periodic/Polled) Geodetic position solution
// Class/Id 0x01 0x02 (28 bytes)
// See important comments concerning validity of position given in section Navigation Output Filters. This message outputs the Geodetic position in the currently selected ellipsoid. The default is the WGS84 Ellipsoid, but can be changed with the message UBX- CFG-DAT.
type NavPosllh struct {
	ITOW_ms   uint32 // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Lon_dege7 int32  // [1e-7 deg] Longitude
	Lat_dege7 int32  // [1e-7 deg] Latitude
	Height_mm int32  // [mm] Height above ellipsoid
	HMSL_mm   int32  // [mm] Height above mean sea level
	HAcc_mm   uint32 // [mm] Horizontal accuracy estimate
	VAcc_mm   uint32 // [mm] Vertical accuracy estimate
}

func (NavPosllh) classID() uint16 { return 0x0201 }

// Message ubx-nav-pvt

// NavPvt (Periodic/Polled) Navigation position velocity time solution
// Class/Id 0x01 0x07 (92 bytes)
// This message combines position, velocity and time solution, including accuracy figures. Note that during a leap second there may be more or less than 60 seconds in a minute. See the description of leap seconds for details.
type NavPvt struct {
	ITOW_ms       uint32       // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Year_y        uint16       // [y] Year (UTC)
	Month_month   byte         // [month] Month, range 1..12 (UTC)
	Day_d         byte         // [d] Day of month, range 1..31 (UTC)
	Hour_h        byte         // [h] Hour of day, range 0..23 (UTC)
	Min_min       byte         // [min] Minute of hour, range 0..59 (UTC)
	Sec_s         byte         // [s] Seconds of minute, range 0..60 (UTC)
	Valid         NavPvtValid  // Validity flags
	TAcc_ns       uint32       // [ns] Time accuracy estimate (UTC)
	Nano_ns       int32        // [ns] Fraction of second, range -1e9 .. 1e9 (UTC)
	FixType       byte         // GNSSfix Type: 0: no fix 1: dead reckoning only 2: 2D-fix 3: 3D-fix 4: GNSS + dead reckoning combined 5: time only fix
	Flags         NavPvtFlags  // Fix status flags
	Flags2        NavPvtFlags2 // Additional flags
	NumSV         byte         // Number of satellites used in Nav Solution
	Lon_dege7     int32        // [1e-7 deg] Longitude
	Lat_dege7     int32        // [1e-7 deg] Latitude
	Height_mm     int32        // [mm] Height above ellipsoid
	HMSL_mm       int32        // [mm] Height above mean sea level
	HAcc_mm       uint32       // [mm] Horizontal accuracy estimate
	VAcc_mm       uint32       // [mm] Vertical accuracy estimate
	VelN_mm_s     int32        // [mm/s] NED north velocity
	VelE_mm_s     int32        // [mm/s] NED east velocity
	VelD_mm_s     int32        // [mm/s] NED down velocity
	GSpeed_mm_s   int32        // [mm/s] Ground Speed (2-D)
	HeadMot_dege5 int32        // [1e-5 deg] Heading of motion (2-D)
	SAcc_mm_s     uint32       // [mm/s] Speed accuracy estimate
	HeadAcc_dege5 uint32       // [1e-5 deg] Heading accuracy estimate (both motion and vehicle)
	PDOP          uint16       // Position DOP
	Flags3        NavPvtFlags3 // Additional flags
	Reserved1     [5]byte      // Reserved
	HeadVeh_dege5 int32        // [1e-5 deg] Heading of vehicle (2-D), this is only valid when headVehValid is set, otherwise the output is set to the heading of motion
	MagDec_dege2  int16        // [1e-2 deg] Magnetic declination. Only supported in ADR 4.10 and later.
	MagAcc_dege2  uint16       // [1e-2 deg] Magnetic declination accuracy. Only supported in ADR 4.10 and later.
}

func (NavPvt) classID() uint16 { return 0x0701 }

type NavPvtValid byte

const (
	NavPvtValidDate     NavPvtValid = 0x1 // 1 = valid UTC Date (see Time Validity section for details)
	NavPvtValidTime     NavPvtValid = 0x2 // 1 = valid UTC time of day (see Time Validity section for details)
	NavPvtFullyResolved NavPvtValid = 0x4 // 1 = UTC time of day has been fully resolved (no seconds uncertainty). Cannot be used to check if time is completely solved.
	NavPvtValidMag      NavPvtValid = 0x8 // 1 = valid magnetic declination
)

type NavPvtFlags byte

const (
	NavPvtGnssFixOK    NavPvtFlags = 0x1  // 1 = valid fix (i.e within DOP & accuracy masks)
	NavPvtDiffSoln     NavPvtFlags = 0x2  // 1 = differential corrections were applied
	NavPvtHeadVehValid NavPvtFlags = 0x20 // 1 = heading of vehicle is valid, only set if the receiver is in sensor fusion mode
	NavPvtCarrSoln     NavPvtFlags = 0xc0 // Carrier phase range solution status: 0: no carrier phase range solution 1: carrier phase range solution with floating ambiguities 2: carrier phase range solution with fixed ambiguities (not supported in protocol versions less than 20)
)

type NavPvtFlags2 byte

const (
	NavPvtConfirmedAvai NavPvtFlags2 = 0x20 // 1 = information about UTC Date and Time of Day validity confirmation is available (see Time Validity section for details) This flag is only supported in Protocol Versions 19.00, 19.10, 20.10, 20.20, 20.30, 22.00, 23.00, 23.01, 27 and 28.
	NavPvtConfirmedDate NavPvtFlags2 = 0x40 // 1 = UTC Date validity could be confirmed (see Time Validity section for details)
	NavPvtConfirmedTime NavPvtFlags2 = 0x80 // 1 = UTC Time of Day could be confirmed (see Time Validity section for details)
)

type NavPvtFlags3 byte

const (
	NavPvtInvalidLlh NavPvtFlags3 = 0x1 // 1 = Invalid lon, lat, height and hMSL
)

func (v NavPvtValid) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "ValidDate")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "ValidTime")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "FullyResolved")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "ValidMag")
	}
	v &^= 0x8
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v NavPvtFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "GnssFixOK")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "DiffSoln")
	}
	v &^= 0x2
	if v&0x20 != 0 {
		b = append(b, "HeadVehValid")
	}
	v &^= 0x20
	if v&0xc0 != 0 {
		b = append(b, fmt.Sprintf("CarrSoln<%b>", (v&0xc0)>>6))
	}
	v &^= 0xc0
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v NavPvtFlags2) String() string {
	var b []string
	if v&0x20 != 0 {
		b = append(b, "ConfirmedAvai")
	}
	v &^= 0x20
	if v&0x40 != 0 {
		b = append(b, "ConfirmedDate")
	}
	v &^= 0x40
	if v&0x80 != 0 {
		b = append(b, "ConfirmedTime")
	}
	v &^= 0x80
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v NavPvtFlags3) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "InvalidLlh")
	}
	v &^= 0x1
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-relposned

// NavRelposned (Periodic/Polled) Relative positioning information in NED frame
// Class/Id 0x01 0x3c (40 bytes)
// The NED frame is defined as the local topological system at the reference station. The relative position vector components in this message, along with their associated accuracies, are given in that local topological system. This message contains the relative position vector from the Reference Station to the Rover, including accuracy figures, in the local topological system defined at the reference station
type NavRelposned struct {
	Version        byte              // Message version (0x00 for this version)
	Reserved1      byte              // Reserved
	RefStationId   uint16            // Reference Station ID. Must be in the range 0..4095
	ITOW_ms        uint32            // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	RelPosN_cm     int32             // [cm] North component of relative position vector
	RelPosE_cm     int32             // [cm] East component of relative position vector
	RelPosD_cm     int32             // [cm] Down component of relative position vector
	RelPosHPN_mme1 int8              // [1e-1 mm] High-precision North component of relative position vector. Must be in the range -99 to +99. The full North component of the relative position vector, in units of cm, is given by relPosN + (relPosHPN * 1e-2)
	RelPosHPE_mme1 int8              // [1e-1 mm] High-precision East component of relative position vector. Must be in the range -99 to +99. The full East component of the relative position vector, in units of cm, is given by relPosE + (relPosHPE * 1e-2)
	RelPosHPD_mme1 int8              // [1e-1 mm] High-precision Down component of relative position vector. Must be in the range -99 to +99. The full Down component of the relative position vector, in units of cm, is given by relPosD + (relPosHPD * 1e-2)
	Reserved2      byte              // Reserved
	AccN_mme1      uint32            // [1e-1 mm] Accuracy of relative position North component
	AccE_mme1      uint32            // [1e-1 mm] Accuracy of relative position East component
	AccD_mme1      uint32            // [1e-1 mm] Accuracy of relative position Down component
	Flags          NavRelposnedFlags // Flags
}

func (NavRelposned) classID() uint16 { return 0x3c01 }

type NavRelposnedFlags uint32

const (
	NavRelposnedGnssFixOK   NavRelposnedFlags = 0x1  // A valid fix (i.e within DOP & accuracy masks)
	NavRelposnedDiffSoln    NavRelposnedFlags = 0x2  // 1 if differential corrections were applied
	NavRelposnedRelPosValid NavRelposnedFlags = 0x4  // 1 if relative position components and accuracies are valid
	NavRelposnedCarrSoln    NavRelposnedFlags = 0x18 // Carrier phase range solution status: 0 = no carrier phase range solution 1 = carrier phase range solution with floating ambiguities 2 = carrier phase range solution with fixed ambiguities
	NavRelposnedIsMoving    NavRelposnedFlags = 0x20 // 1 if the receiver is operating in moving baseline mode (not supported in protocol versions less than 20.3)
	NavRelposnedRefPosMiss  NavRelposnedFlags = 0x40 // 1 if extrapolated reference position was used to compute moving baseline solution this epoch (not supported in protocol versions less than 20.3)
	NavRelposnedRefObsMiss  NavRelposnedFlags = 0x80 // 1 if extrapolated reference observations were used to compute moving baseline solution this epoch (not supported in protocol versions less than 20.3)
)

func (v NavRelposnedFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "GnssFixOK")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "DiffSoln")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "RelPosValid")
	}
	v &^= 0x4
	if v&0x18 != 0 {
		b = append(b, fmt.Sprintf("CarrSoln<%b>", (v&0x18)>>3))
	}
	v &^= 0x18
	if v&0x20 != 0 {
		b = append(b, "IsMoving")
	}
	v &^= 0x20
	if v&0x40 != 0 {
		b = append(b, "RefPosMiss")
	}
	v &^= 0x40
	if v&0x80 != 0 {
		b = append(b, "RefObsMiss")
	}
	v &^= 0x80
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-resetodo

// NavResetodo (Command) Reset odometer
// Class/Id 0x01 0x10 (0 bytes)
// This message resets the traveled distance computed by the odometer (see UBX- NAV-ODO). UBX-ACK-ACK or UBX-ACK-NAK are returned to indicate success or failure.
type NavResetodo struct {
}

func (NavResetodo) classID() uint16 { return 0x1001 }

// Message ubx-nav-sat

// NavSat (Periodic/Polled) Satellite information
// Class/Id 0x01 0x35 (8 + N*12 bytes)
// This message displays information about SVs that are either known to be visible or currently tracked by the receiver. All signal related information corresponds to the subset of signals specified in Signal Identifiers.
type NavSat struct {
	ITOW_ms   uint32           // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Version   byte             // Message version (0x01 for this version)
	NumSvs    byte             `len:"Svs"` // Number of satellites
	Reserved1 [2]byte          // Reserved
	Svs       []*NavSatSvsType // len: NumSvs
}

func (NavSat) classID() uint16 { return 0x3501 }

type NavSatSvsType struct {
	GnssId    byte        // GNSS identifier (see Satellite Numbering) for assignment
	SvId      byte        // Satellite identifier (see Satellite Numbering) for assignment
	Cno_dbhz  byte        // [dBHz] Carrier to noise ratio (signal strength)
	Elev_deg  int8        // [deg] Elevation (range: +/-90), unknown if out of range
	Azim_deg  int16       // [deg] Azimuth (range 0-360), unknown if elevation is out of range
	PrRes_me1 int16       // [1e-1 m] Pseudorange residual
	Flags     NavSatFlags // Bitmask
}

type NavSatFlags uint32

const (
	NavSatQualityInd   NavSatFlags = 0x7      // Signal quality indicator: 0: no signal 1: searching signal 2: signal acquired 3: signal detected but unusable 4: code locked and time synchronized 5, 6, 7: code and carrier locked and time synchronized Note: Since IMES signals are not time synchronized, a channel tracking an IMES signal can never reach a quality indicator value of higher than 3.
	NavSatSvUsed       NavSatFlags = 0x8      // 1 = Signal in the subset specified in Signal Identifiers is currently being used for navigation
	NavSatHealth       NavSatFlags = 0x30     // Signal health flag: 0: unknown 1: healthy 2: unhealthy
	NavSatDiffCorr     NavSatFlags = 0x40     // 1 = differential correction data is available for this SV
	NavSatSmoothed     NavSatFlags = 0x80     // 1 = carrier smoothed pseudorange used
	NavSatOrbitSource  NavSatFlags = 0x700    // Orbit source: 0: no orbit information is available for this SV 1: ephemeris is used 2: almanac is used 3: AssistNow Offline orbit is used 4: AssistNow Autonomous orbit is used 5, 6, 7: other orbit information is used
	NavSatEphAvail     NavSatFlags = 0x800    // 1 = ephemeris is available for this SV
	NavSatAlmAvail     NavSatFlags = 0x1000   // 1 = almanac is available for this SV
	NavSatAnoAvail     NavSatFlags = 0x2000   // 1 = AssistNow Offline data is available for this SV
	NavSatAopAvail     NavSatFlags = 0x4000   // 1 = AssistNow Autonomous data is available for this SV
	NavSatSbasCorrUsed NavSatFlags = 0x10000  // 1 = SBAS corrections have been used for a signal in the subset specified in Signal Identifiers
	NavSatRtcmCorrUsed NavSatFlags = 0x20000  // 1 = RTCM corrections have been used for a signal in the subset specified in Signal Identifiers
	NavSatSlasCorrUsed NavSatFlags = 0x40000  // 1 = QZSS SLAS corrections have been used for a signal in the subset specified in Signal Identifiers
	NavSatPrCorrUsed   NavSatFlags = 0x100000 // 1 = Pseudorange corrections have been used for a signal in the subset specified in Signal Identifiers
	NavSatCrCorrUsed   NavSatFlags = 0x200000 // 1 = Carrier range corrections have been used for a signal in the subset specified in Signal Identifiers
	NavSatDoCorrUsed   NavSatFlags = 0x400000 // 1 = Range rate (Doppler) corrections have been used for a signal in the subset specified in Signal Identifiers
)

func (v NavSatFlags) String() string {
	var b []string
	if v&0x7 != 0 {
		b = append(b, fmt.Sprintf("QualityInd<%b>", (v&0x7)>>0))
	}
	v &^= 0x7
	if v&0x8 != 0 {
		b = append(b, "SvUsed")
	}
	v &^= 0x8
	if v&0x30 != 0 {
		b = append(b, fmt.Sprintf("Health<%b>", (v&0x30)>>4))
	}
	v &^= 0x30
	if v&0x40 != 0 {
		b = append(b, "DiffCorr")
	}
	v &^= 0x40
	if v&0x80 != 0 {
		b = append(b, "Smoothed")
	}
	v &^= 0x80
	if v&0x700 != 0 {
		b = append(b, fmt.Sprintf("OrbitSource<%b>", (v&0x700)>>8))
	}
	v &^= 0x700
	if v&0x800 != 0 {
		b = append(b, "EphAvail")
	}
	v &^= 0x800
	if v&0x1000 != 0 {
		b = append(b, "AlmAvail")
	}
	v &^= 0x1000
	if v&0x2000 != 0 {
		b = append(b, "AnoAvail")
	}
	v &^= 0x2000
	if v&0x4000 != 0 {
		b = append(b, "AopAvail")
	}
	v &^= 0x4000
	if v&0x10000 != 0 {
		b = append(b, "SbasCorrUsed")
	}
	v &^= 0x10000
	if v&0x20000 != 0 {
		b = append(b, "RtcmCorrUsed")
	}
	v &^= 0x20000
	if v&0x40000 != 0 {
		b = append(b, "SlasCorrUsed")
	}
	v &^= 0x40000
	if v&0x100000 != 0 {
		b = append(b, "PrCorrUsed")
	}
	v &^= 0x100000
	if v&0x200000 != 0 {
		b = append(b, "CrCorrUsed")
	}
	v &^= 0x200000
	if v&0x400000 != 0 {
		b = append(b, "DoCorrUsed")
	}
	v &^= 0x400000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-sbas

// NavSbas (Periodic/Polled) SBAS status data
// Class/Id 0x01 0x32 (12 + N*12 bytes)
// This message outputs the status of the SBAS sub system
type NavSbas struct {
	ITOW_ms   uint32              // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Geo       byte                // PRN Number of the GEO where correction and integrity data is used from
	Mode      byte                // SBAS Mode 0 Disabled 1 Enabled integrity 3 Enabled test mode
	Sys       int8                // SBAS System (WAAS/EGNOS/...) -1 Unknown 0 WAAS 1 EGNOS 2 MSAS 3 GAGAN 16 GPS
	Service   NavSbasService      // SBAS Services available
	Cnt       byte                `len:"Items"` // Number of SV data following
	Reserved1 [3]byte             // Reserved
	Items     []*NavSbasItemsType // len: Cnt
}

func (NavSbas) classID() uint16 { return 0x3201 }

type NavSbasItemsType struct {
	Svid      byte    // SV ID
	Flags     byte    // Flags for this SV
	Udre      byte    // Monitoring status
	SvSys     byte    // System (WAAS/EGNOS/...) same as SYS
	SvService byte    // Services available same as SERVICE
	Reserved2 byte    // Reserved
	Prc_cm    int16   // [cm] Pseudo Range correction in [cm]
	Reserved3 [2]byte // Reserved
	Ic_cm     int16   // [cm] Ionosphere correction in [cm]
}

type NavSbasService byte

const (
	NavSbasRanging     NavSbasService = 0x1  // GEO may be used as ranging source
	NavSbasCorrections NavSbasService = 0x2  // GEO is providing correction data
	NavSbasIntegrity   NavSbasService = 0x4  // GEO is providing integrity
	NavSbasTestmode    NavSbasService = 0x8  // GEO is in test mode
	NavSbasBad         NavSbasService = 0x10 // Problem with signal or broadcast data indicated
)

func (v NavSbasService) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Ranging")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "Corrections")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "Integrity")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "Testmode")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "Bad")
	}
	v &^= 0x10
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-slas

// NavSlas (Periodic/Polled) QZSS L1S SLAS status data
// Class/Id 0x01 0x42 (20 + N*8 bytes)
// This message outputs the status of the QZSS L1S SLAS sub system
type NavSlas struct {
	ITOW_ms      uint32              // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Version      byte                // Message version (0x00 for this version)
	Reserved1    [3]byte             // Reserved
	GmsLon_dege3 int32               // [1e-3 deg] Longitude of the used ground monitoring station
	GmsLat_dege3 int32               // [1e-3 deg] Latitude of the used ground monitoring station
	GmsCode      byte                // Code of the used ground monitoring station according to the QZSS SLAS Interface Specification, available from qzss.go.jp/en/
	QzssSvId     byte                // Satellite identifier of the QZS/GEO whose correction data is used (see Satellite Numbering)
	ServiceFlags NavSlasServiceFlags // Flags regarding SLAS service
	Cnt          byte                `len:"Items"` // Number of pseudorange corrections following
	Items        []*NavSlasItemsType // len: Cnt
}

func (NavSlas) classID() uint16 { return 0x4201 }

type NavSlasItemsType struct {
	GnssId    byte    // GNSS identifier (see Satellite Numbering)
	SvId      byte    // Satellite identifier (see Satellite Numbering)
	Reserved2 byte    // Reserved
	Reserved3 [3]byte // Reserved
	Prc_cm    int16   // [cm] Pseudorange correction
}

type NavSlasServiceFlags byte

const (
	NavSlasGmsAvailable    NavSlasServiceFlags = 0x1 // 1 = Ground monitoring station available
	NavSlasQzssSvAvailable NavSlasServiceFlags = 0x2 // 1 = Correction providing QZSS SV available
	NavSlasTestMode        NavSlasServiceFlags = 0x4 // 1 = Currently used QZSS SV in test mode
)

func (v NavSlasServiceFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "GmsAvailable")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "QzssSvAvailable")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "TestMode")
	}
	v &^= 0x4
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-sol

// NavSol (Periodic/Polled) Navigation solution information
// Class/Id 0x01 0x06 (52 bytes)
// This message combines position, velocity and time solution in ECEF, including accuracy figures. This message has only been retained for backwards compatibility; users are recommended to use the UBX-NAV-PVT message in preference.
type NavSol struct {
	ITOW_ms     uint32      // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	FTOW_ns     int32       // [ns] Fractional part of iTOW (range: +/- 500000). The precise GPS time of week in seconds is: (iTOW    *  1e-3)    +  (fTOW    *  1e-9)
	Week_weeks  int16       // [weeks] GPS week number of the navigation epoch
	GpsFix      byte        // GPSfix Type, range 0..5 0x00 = No Fix 0x01 = Dead Reckoning only 0x02 = 2D-Fix 0x03 = 3D-Fix 0x04 = GPS + dead reckoning combined 0x05 = Time only fix 0x06..0xff: reserved
	Flags       NavSolFlags // Fix Status Flags
	EcefX_cm    int32       // [cm] ECEF X coordinate
	EcefY_cm    int32       // [cm] ECEF Y coordinate
	EcefZ_cm    int32       // [cm] ECEF Z coordinate
	PAcc_cm     uint32      // [cm] 3D Position Accuracy Estimate
	EcefVX_cm_s int32       // [cm/s] ECEF X velocity
	EcefVY_cm_s int32       // [cm/s] ECEF Y velocity
	EcefVZ_cm_s int32       // [cm/s] ECEF Z velocity
	SAcc_cm_s   uint32      // [cm/s] Speed Accuracy Estimate
	PDOP        uint16      // Position DOP
	Reserved1   byte        // Reserved
	NumSV       byte        // Number of SVs used in Nav Solution
	Reserved2   [4]byte     // Reserved
}

func (NavSol) classID() uint16 { return 0x0601 }

type NavSolFlags byte

const (
	NavSolGPSfixOK NavSolFlags = 0x1 // 1 = Fix within limits (e.g. DOP & accuracy)
	NavSolDiffSoln NavSolFlags = 0x2 // 1 = DGPS used
	NavSolWKNSET   NavSolFlags = 0x4 // 1 = Valid GPS week number (see Time Validity section for details)
	NavSolTOWSET   NavSolFlags = 0x8 // 1 = Valid GPS time of week (iTOW & fTOW, see Time Validity section for details)
)

func (v NavSolFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "GPSfixOK")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "DiffSoln")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "WKNSET")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "TOWSET")
	}
	v &^= 0x8
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-status

// NavStatus (Periodic/Polled) Receiver navigation status
// Class/Id 0x01 0x03 (16 bytes)
// See important comments concerning validity of position given in section Navigation Output Filters.
type NavStatus struct {
	ITOW_ms uint32           // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	GpsFix  byte             // GPSfix Type, this value does not qualify a fix as valid and within the limits. See note on flag gpsFixOk below. 0x00 = no fix 0x01 = dead reckoning only 0x02 = 2D-fix 0x03 = 3D-fix 0x04 = GPS + dead reckoning combined 0x05 = Time only fix 0x06..0xff = reserved
	Flags   NavStatusFlags   // Navigation Status Flags
	FixStat NavStatusFixStat // Fix Status Information
	Flags2  NavStatusFlags2  // further information about navigation output
	Ttff_ms uint32           // [ms] Time to first fix (millisecond time tag)
	Msss_ms uint32           // [ms] Milliseconds since Startup / Reset
}

func (NavStatus) classID() uint16 { return 0x0301 }

type NavStatusFlags byte

const (
	NavStatusGpsFixOk NavStatusFlags = 0x1 // 1 = position and velocity valid and within DOP and ACC Masks.
	NavStatusDiffSoln NavStatusFlags = 0x2 // 1 = differential corrections were applied
	NavStatusWknSet   NavStatusFlags = 0x4 // 1 = Week Number valid (see Time Validity section for details)
	NavStatusTowSet   NavStatusFlags = 0x8 // 1 = Time of Week valid (see Time Validity section for details)
)

type NavStatusFixStat byte

const (
	NavStatusDiffCorr      NavStatusFixStat = 0x1  // 1 = differential corrections available
	NavStatusCarrSolnValid NavStatusFixStat = 0x2  // 1 = valid carrSoln
	NavStatusMapMatching   NavStatusFixStat = 0xc0 // map matching status: 00: none 01: valid but not used, i.e. map matching data was received, but was too old 10: valid and used, map matching data was applied 11: valid and used, map matching data was applied. In case of sensor unavailability map matching data enables dead reckoning. This requires map matched latitude/longitude or heading data.
)

type NavStatusFlags2 byte

const (
	NavStatusPsmState      NavStatusFlags2 = 0x3  // power save mode state 0: ACQUISITION [or when psm disabled] 1: TRACKING 2: POWER OPTIMIZED TRACKING 3: INACTIVE
	NavStatusSpoofDetState NavStatusFlags2 = 0x18 // Spoofing detection state (not supported in protocol versions less than 18) 0: Unknown or deactivated 1: No spoofing indicated 2: Spoofing indicated 3: Multiple spoofing indications Note that the spoofing state value only reflects the detector state for the current navigation epoch. As spoofing can be detected most easily at the transition from real signal to spoofing signal, this is also where the detector is triggered the most. I.e. a value of 1 - No spoofing indicated does not mean that the receiver is not spoofed, it simply states that the detector was not triggered in this epoch.
	NavStatusCarrSoln      NavStatusFlags2 = 0xc0 // Carrier phase range solution status: 0: no carrier phase range solution 1: carrier phase range solution with floating ambiguities 2: carrier phase range solution with fixed ambiguities
)

func (v NavStatusFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "GpsFixOk")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "DiffSoln")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "WknSet")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "TowSet")
	}
	v &^= 0x8
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v NavStatusFixStat) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "DiffCorr")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "CarrSolnValid")
	}
	v &^= 0x2
	if v&0xc0 != 0 {
		b = append(b, fmt.Sprintf("MapMatching<%b>", (v&0xc0)>>6))
	}
	v &^= 0xc0
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v NavStatusFlags2) String() string {
	var b []string
	if v&0x3 != 0 {
		b = append(b, fmt.Sprintf("PsmState<%b>", (v&0x3)>>0))
	}
	v &^= 0x3
	if v&0x18 != 0 {
		b = append(b, fmt.Sprintf("SpoofDetState<%b>", (v&0x18)>>3))
	}
	v &^= 0x18
	if v&0xc0 != 0 {
		b = append(b, fmt.Sprintf("CarrSoln<%b>", (v&0xc0)>>6))
	}
	v &^= 0xc0
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-svin

// NavSvin (Periodic/Polled) Survey-in data
// Class/Id 0x01 0x3b (40 bytes)
// This message contains information about survey-in parameters.
type NavSvin struct {
	Version   byte    // Message version (0x00 for this version)
	Reserved1 [3]byte // Reserved
	ITOW_ms   uint32  // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Dur_s     uint32  // [s] Passed survey-in observation time
	MeanX_cm  int32   // [cm] Current survey-in mean position ECEF X coordinate
	MeanY_cm  int32   // [cm] Current survey-in mean position ECEF Y coordinate
	MeanZ_cm  int32   // [cm] Current survey-in mean position ECEF Z coordinate
	MeanXHP   int8    // [0.1_mm] Current high-precision survey-in mean position ECEF X coordinate. Must be in the range -99..+99. The current survey-in mean position ECEF X coordinate, in units of cm, is given by meanX + (0.01 * meanXHP)
	MeanYHP   int8    // [0.1_mm] Current high-precision survey-in mean position ECEF Y coordinate. Must be in the range -99..+99. The current survey-in mean position ECEF Y coordinate, in units of cm, is given by meanY + (0.01 * meanYHP)
	MeanZHP   int8    // [0.1_mm] Current high-precision survey-in mean position ECEF Z coordinate. Must be in the range -99..+99. The current survey-in mean position ECEF Z coordinate, in units of cm, is given by meanZ + (0.01 * meanZHP)
	Reserved2 byte    // Reserved
	MeanAcc   uint32  // [0.1_mm] Current survey-in mean position accuracy
	Obs       uint32  // Number of position observations used during survey-in
	Valid     byte    // Survey-in position validity flag, 1 = valid, otherwise 0
	Active    byte    // Survey-in in progress flag, 1 = in-progress, otherwise 0
	Reserved3 [2]byte // Reserved
}

func (NavSvin) classID() uint16 { return 0x3b01 }

// Message ubx-nav-svinfo

// NavSvinfo (Periodic/Polled) Space vehicle information
// Class/Id 0x01 0x30 (8 + N*12 bytes)
// Information about satellites used or visible This message has only been retained for backwards compatibility; users are recommended to use the UBX-NAV-SAT message in preference.
type NavSvinfo struct {
	ITOW_ms     uint32               // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	NumCh       byte                 `len:"Ch"` // Number of channels
	GlobalFlags NavSvinfoGlobalFlags // Bitmask
	Reserved1   [2]byte              // Reserved
	Ch          []*NavSvinfoChType   // len: NumCh
}

func (NavSvinfo) classID() uint16 { return 0x3001 }

type NavSvinfoChType struct {
	Chn      byte             // Channel number, 255 for SVs not assigned to a channel
	Svid     byte             // Satellite ID, see Satellite Numbering for assignment
	Flags    NavSvinfoFlags   // Bitmask
	Quality  NavSvinfoQuality // Bitfield
	Cno_dbhz byte             // [dBHz] Carrier to Noise Ratio (Signal Strength)
	Elev_deg int8             // [deg] Elevation in integer degrees
	Azim_deg int16            // [deg] Azimuth in integer degrees
	PrRes_cm int32            // [cm] Pseudo range residual in centimeters
}

type NavSvinfoGlobalFlags byte

const (
	NavSvinfoChipGen NavSvinfoGlobalFlags = 0x7 // Chip hardware generation 0: Antaris, Antaris 4 1: u-blox 5 2: u-blox 6 3: u-blox 7 4: u-blox 8 / u-blox M8
)

type NavSvinfoFlags byte

const (
	NavSvinfoSvUsed     NavSvinfoFlags = 0x1  // SV is used for navigation
	NavSvinfoDiffCorr   NavSvinfoFlags = 0x2  // Differential correction data is available for this SV
	NavSvinfoOrbitAvail NavSvinfoFlags = 0x4  // Orbit information is available for this SV (Ephemeris or Almanac)
	NavSvinfoOrbitEph   NavSvinfoFlags = 0x8  // Orbit information is Ephemeris
	NavSvinfoUnhealthy  NavSvinfoFlags = 0x10 // SV is unhealthy / shall not be used
	NavSvinfoOrbitAlm   NavSvinfoFlags = 0x20 // Orbit information is Almanac Plus
	NavSvinfoOrbitAop   NavSvinfoFlags = 0x40 // Orbit information is AssistNow Autonomous
	NavSvinfoSmoothed   NavSvinfoFlags = 0x80 // Carrier smoothed pseudorange used
)

type NavSvinfoQuality byte

const (
	NavSvinfoQualityInd NavSvinfoQuality = 0xf // Signal Quality indicator (range 0..7). The following list shows the meaning of the different QI values: 0: no signal 1: searching signal 2: signal acquired 3: signal detected but unusable 4: code locked and time synchronized 5, 6, 7: code and carrier locked and time synchronized Note: Since IMES signals are not time synchronized, a channel tracking an IMES signal can never reach a quality indicator value of higher than 3.
)

func (v NavSvinfoGlobalFlags) String() string {
	var b []string
	if v&0x7 != 0 {
		b = append(b, fmt.Sprintf("ChipGen<%b>", (v&0x7)>>0))
	}
	v &^= 0x7
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v NavSvinfoFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "SvUsed")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "DiffCorr")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "OrbitAvail")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "OrbitEph")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "Unhealthy")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "OrbitAlm")
	}
	v &^= 0x20
	if v&0x40 != 0 {
		b = append(b, "OrbitAop")
	}
	v &^= 0x40
	if v&0x80 != 0 {
		b = append(b, "Smoothed")
	}
	v &^= 0x80
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v NavSvinfoQuality) String() string {
	var b []string
	if v&0xf != 0 {
		b = append(b, fmt.Sprintf("QualityInd<%b>", (v&0xf)>>0))
	}
	v &^= 0xf
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-timebds

// NavTimebds (Periodic/Polled) BeiDou time solution
// Class/Id 0x01 0x24 (20 bytes)
// This message reports the precise BDS time of the most recent navigation solution including validity flags and an accuracy estimate.
type NavTimebds struct {
	ITOW_ms uint32          // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	SOW_s   uint32          // [s] BDS time of week (rounded to seconds)
	FSOW_ns int32           // [ns] Fractional part of SOW (range: +/- 500000000). The precise BDS time of week in seconds is: SOW   +  fSOW    *  1e-9
	Week    int16           // BDS week number of the navigation epoch
	LeapS_s int8            // [s] BDS leap seconds (BDS-UTC)
	Valid   NavTimebdsValid // Validity Flags
	TAcc_ns uint32          // [ns] Time Accuracy Estimate
}

func (NavTimebds) classID() uint16 { return 0x2401 }

type NavTimebdsValid byte

const (
	NavTimebdsSowValid   NavTimebdsValid = 0x1 // 1 = Valid SOW and fSOW (see Time Validity section for details)
	NavTimebdsWeekValid  NavTimebdsValid = 0x2 // 1 = Valid week (see Time Validity section for details)
	NavTimebdsLeapSValid NavTimebdsValid = 0x4 // 1 = Valid leap second
)

func (v NavTimebdsValid) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "SowValid")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "WeekValid")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "LeapSValid")
	}
	v &^= 0x4
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-timegal

// NavTimegal (Periodic/Polled) Galileo time solution
// Class/Id 0x01 0x25 (20 bytes)
// This message reports the precise Galileo time of the most recent navigation solution including validity flags and an accuracy estimate.
type NavTimegal struct {
	ITOW_ms    uint32          // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	GalTow_s   uint32          // [s] Galileo time of week (rounded to seconds)
	FGalTow_ns int32           // [ns] Fractional part of the Galileo time of week (range: +/-500000000). The precise Galileo time of week in seconds is: galTow     +  fGalTow     *  1e-9
	GalWno     int16           // Galileo week number
	LeapS_s    int8            // [s] Galileo leap seconds (Galileo-UTC)
	Valid      NavTimegalValid // Validity Flags
	TAcc_ns    uint32          // [ns] Time Accuracy Estimate
}

func (NavTimegal) classID() uint16 { return 0x2501 }

type NavTimegalValid byte

const (
	NavTimegalGalTowValid NavTimegalValid = 0x1 // 1 = Valid galTow and fGalTow (see the section Time validity in the Integration manual for details)
	NavTimegalGalWnoValid NavTimegalValid = 0x2 // 1 = Valid galWno (see the section Time validity in the Integration manual for details)
	NavTimegalLeapSValid  NavTimegalValid = 0x4 // 1 = Valid leapS
)

func (v NavTimegalValid) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "GalTowValid")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "GalWnoValid")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "LeapSValid")
	}
	v &^= 0x4
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-timeglo

// NavTimeglo (Periodic/Polled) GLONASS time solution
// Class/Id 0x01 0x23 (20 bytes)
// This message reports the precise GLO time of the most recent navigation solution including validity flags and an accuracy estimate.
type NavTimeglo struct {
	ITOW_ms uint32          // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	TOD_s   uint32          // [s] GLONASS time of day (rounded to integer seconds)
	FTOD_ns int32           // [ns] Fractional part of TOD (range: +/- 500000000). The precise GLONASS time of day in seconds is: TOD   +  fTOD    *  1e-9
	Nt_days uint16          // [days] Current date (range: 1-1461), starting at 1 from the 1st Jan of the year indicated by N4 and ending at 1461 at the 31st Dec of the third year after that indicated by N4
	N4      byte            // Four-year interval number starting from 1996 (1=1996, 2=2000, 3=2004...)
	Valid   NavTimegloValid // Validity flags
	TAcc_ns uint32          // [ns] Time Accuracy Estimate
}

func (NavTimeglo) classID() uint16 { return 0x2301 }

type NavTimegloValid byte

const (
	NavTimegloTodValid  NavTimegloValid = 0x1 // 1 = Valid TOD and fTOD (see Time Validity section for details)
	NavTimegloDateValid NavTimegloValid = 0x2 // 1 = Valid N4 and Nt (see Time Validity section for details)
)

func (v NavTimegloValid) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "TodValid")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "DateValid")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-timegps

// NavTimegps (Periodic/Polled) GPS time solution
// Class/Id 0x01 0x20 (16 bytes)
// This message reports the precise GPS time of the most recent navigation solution including validity flags and an accuracy estimate.
type NavTimegps struct {
	ITOW_ms uint32          // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	FTOW_ns int32           // [ns] Fractional part of iTOW (range: +/- 500000). The precise GPS time of week in seconds is: (iTOW    *  1e-3)    +  (fTOW    *  1e-9)
	Week    int16           // GPS week number of the navigation epoch
	LeapS_s int8            // [s] GPS leap seconds (GPS-UTC)
	Valid   NavTimegpsValid // Validity Flags
	TAcc_ns uint32          // [ns] Time Accuracy Estimate
}

func (NavTimegps) classID() uint16 { return 0x2001 }

type NavTimegpsValid byte

const (
	NavTimegpsTowValid   NavTimegpsValid = 0x1 // 1 = Valid GPS time of week (iTOW & fTOW, (see Time Validity section for details)
	NavTimegpsWeekValid  NavTimegpsValid = 0x2 // 1 = Valid GPS week number (see Time Validity section for details)
	NavTimegpsLeapSValid NavTimegpsValid = 0x4 // 1 = Valid GPS leap seconds
)

func (v NavTimegpsValid) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "TowValid")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "WeekValid")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "LeapSValid")
	}
	v &^= 0x4
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-timels

// NavTimels (Periodic/Polled) Leap second event information
// Class/Id 0x01 0x26 (24 bytes)
// Information about the upcoming leap second event if one is scheduled.
type NavTimels struct {
	ITOW_ms         uint32         // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Version         byte           // Message version (0x00 for this version)
	Reserved1       [3]byte        // Reserved
	SrcOfCurrLs     byte           // Information source for the current number of leap seconds. 0: Default (hardcoded in the firmware, can be outdated) 1: Derived from time difference between GPS and GLONASS time 2: GPS 3: SBAS 4: BeiDou 5: Galileo 6: Aided data 7: Configured 255: Unknown
	CurrLs_s        int8           // [s] Current number of leap seconds since start of GPS time (Jan 6, 1980). It reflects how much GPS time is ahead of UTC time. Galileo number of leap seconds is the same as GPS. BeiDou number of leap seconds is 14 less than GPS. GLONASS follows UTC time, so no leap seconds.
	SrcOfLsChange   byte           // Information source for the future leap second event. 0: No source 2: GPS 3: SBAS 4: BeiDou 5: Galileo 6: GLONASS
	LsChange_s      int8           // [s] Future leap second change if one is scheduled. +1 = positive leap second, -1 = negative leap second, 0 = no future leap second event scheduled or no information available.
	TimeToLsEvent_s int32          // [s] Number of seconds until the next leap second event, or from the last leap second event if no future event scheduled. If > 0 event is in the future, = 0 event is now, < 0 event is in the past. Valid only if validTimeToLsEvent = 1.
	DateOfLsGpsWn   uint16         // GPS week number (WN) of the next leap second event or the last one if no future event scheduled. Valid only if validTimeToLsEvent = 1.
	DateOfLsGpsDn   uint16         // GPS day of week number (DN) for the next leap second event or the last one if no future event scheduled. Valid only if validTimeToLsEvent = 1. (GPS and Galileo DN: from 1 = Sun to 7 = Sat. BeiDou DN: from 0 = Sun to 6 = Sat.)
	Reserved2       [3]byte        // Reserved
	Valid           NavTimelsValid // Validity flags
}

func (NavTimels) classID() uint16 { return 0x2601 }

type NavTimelsValid byte

const (
	NavTimelsValidCurrLs        NavTimelsValid = 0x1 // 1 = Valid current number of leap seconds value.
	NavTimelsValidTimeToLsEvent NavTimelsValid = 0x2 // 1 = Valid time to next leap second event or from the last leap second event if no future event scheduled.
)

func (v NavTimelsValid) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "ValidCurrLs")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "ValidTimeToLsEvent")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-timeutc

// NavTimeutc (Periodic/Polled) UTC time solution
// Class/Id 0x01 0x21 (20 bytes)
// Note that during a leap second there may be more or less than 60 seconds in a minute. See the description of leap seconds for details.
type NavTimeutc struct {
	ITOW_ms     uint32          // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	TAcc_ns     uint32          // [ns] Time accuracy estimate (UTC)
	Nano_ns     int32           // [ns] Fraction of second, range -1e9 .. 1e9 (UTC)
	Year_y      uint16          // [y] Year, range 1999..2099 (UTC)
	Month_month byte            // [month] Month, range 1..12 (UTC)
	Day_d       byte            // [d] Day of month, range 1..31 (UTC)
	Hour_h      byte            // [h] Hour of day, range 0..23 (UTC)
	Min_min     byte            // [min] Minute of hour, range 0..59 (UTC)
	Sec_s       byte            // [s] Seconds of minute, range 0..60 (UTC)
	Valid       NavTimeutcValid // Validity Flags
}

func (NavTimeutc) classID() uint16 { return 0x2101 }

type NavTimeutcValid byte

const (
	NavTimeutcValidTOW    NavTimeutcValid = 0x1  // 1 = Valid Time of Week (see Time Validity section for details)
	NavTimeutcValidWKN    NavTimeutcValid = 0x2  // 1 = Valid Week Number (see Time Validity section for details)
	NavTimeutcValidUTC    NavTimeutcValid = 0x4  // 1 = Valid UTC Time
	NavTimeutcUtcStandard NavTimeutcValid = 0xf0 // UTC standard identifier. 0: Information not available 1: Communications Research Labratory (CRL), Tokyo, Japan 2: National Institute of Standards and Technology (NIST) 3: U.S. Naval Observatory (USNO) 4: International Bureau of Weights and Measures (BIPM) 5: European laboratories 6: Former Soviet Union (SU) 7: National Time Service Center (NTSC), China 15: Unknown
)

func (v NavTimeutcValid) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "ValidTOW")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "ValidWKN")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "ValidUTC")
	}
	v &^= 0x4
	if v&0xf0 != 0 {
		b = append(b, fmt.Sprintf("UtcStandard<%b>", (v&0xf0)>>4))
	}
	v &^= 0xf0
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-nav-velecef

// NavVelecef (Periodic/Polled) Velocity solution in ECEF
// Class/Id 0x01 0x11 (20 bytes)
// See important comments concerning validity of position given in section Navigation Output Filters.
type NavVelecef struct {
	ITOW_ms     uint32 // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	EcefVX_cm_s int32  // [cm/s] ECEF X velocity
	EcefVY_cm_s int32  // [cm/s] ECEF Y velocity
	EcefVZ_cm_s int32  // [cm/s] ECEF Z velocity
	SAcc_cm_s   uint32 // [cm/s] Speed accuracy estimate
}

func (NavVelecef) classID() uint16 { return 0x1101 }

// Message ubx-nav-velned

// NavVelned (Periodic/Polled) Velocity solution in NED frame
// Class/Id 0x01 0x12 (36 bytes)
// See important comments concerning validity of position given in section Navigation Output Filters.
type NavVelned struct {
	ITOW_ms       uint32 // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	VelN_cm_s     int32  // [cm/s] North velocity component
	VelE_cm_s     int32  // [cm/s] East velocity component
	VelD_cm_s     int32  // [cm/s] Down velocity component
	Speed_cm_s    uint32 // [cm/s] Speed (3-D)
	GSpeed_cm_s   uint32 // [cm/s] Ground speed (2-D)
	Heading_dege5 int32  // [1e-5 deg] Heading of motion 2-D
	SAcc_cm_s     uint32 // [cm/s] Speed accuracy Estimate
	CAcc_dege5    uint32 // [1e-5 deg] Course / Heading accuracy estimate
}

func (NavVelned) classID() uint16 { return 0x1201 }

// Message ubx-rxm-imes

// RxmImes (Periodic/Polled) Indoor Messaging System information
// Class/Id 0x02 0x61 (4 + N*44 bytes)
// This message shows the IMES stations the receiver is currently tracking, their data rate, the signal level, the Doppler (with respect to 1575.4282MHz) and what data (without protocol specific overhead) it has received from these stations so far. This message is sent out at the navigation rate the receiver is currently set to. Therefore it allows users to get an overview on the receiver's current state from the IMES perspective.
type RxmImes struct {
	NumTx     byte             `len:"Tx"` // Number of transmitters contained in the message
	Version   byte             // Message version (0x01 for this version)
	Reserved1 [2]byte          // Reserved
	Tx        []*RxmImesTxType // len: NumTx
}

func (RxmImes) classID() uint16 { return 0x6102 }

type RxmImesTxType struct {
	Reserved2     byte                // Reserved
	TxId          byte                // Transmitter identifier
	Reserved3     [3]byte             // Reserved
	Cno_dbhz      byte                // [dBHz] Carrier to Noise Ratio (Signal Strength)
	Reserved4     [2]byte             // Reserved
	Doppler_hzl12 int32               // [2^-12 Hz] Doppler frequency with respect to 1575. 4282MHz [IIIII.FFF Hz]
	Position1_1   RxmImesPosition1_1  // Position 1 Frame (part 1/2)
	Position1_2   RxmImesPosition1_2  // Position 1 Frame (part 2/2)
	Position2_1   RxmImesPosition2_1  // Position 2 Frame (part 1/3)
	Lat_deg       int32               // [180*2^-24 deg] Latitude, Position 2 Frame (part 2/3)
	Lon_deg       int32               // [360*2^-25 deg] Longitude, Position 2 Frame (part 3/3)
	ShortIdFrame  RxmImesShortIdFrame // Short ID Frame
	MediumIdLSB   uint32              // Medium ID LSB, Medium ID Frame (part 1/2)
	MediumId_2    RxmImesMediumId_2   // Medium ID Frame (part 2/2)
}

type RxmImesPosition1_1 uint32

const (
	RxmImesPos1Floor RxmImesPosition1_1 = 0xff       // Floor number [1.0 floor resolution] (Offset: -50 floor)
	RxmImesPos1Lat   RxmImesPosition1_1 = 0x7fffff00 // Latitude [deg * (180 / 2^23)]
)

type RxmImesPosition1_2 uint32

const (
	RxmImesPos1Lon   RxmImesPosition1_2 = 0xffffff  // Longitude [deg * (360 / 2^24)]
	RxmImesPos1Valid RxmImesPosition1_2 = 0x1000000 // Position 1 Frame valid
)

type RxmImesPosition2_1 uint32

const (
	RxmImesPos2Floor RxmImesPosition2_1 = 0x1ff    // Floor number [0.5 floor resolution] (Offset: -50 floor)
	RxmImesPos2Alt   RxmImesPosition2_1 = 0x1ffe00 // Altitude [m] (Offset: -95m)
	RxmImesPos2Acc   RxmImesPosition2_1 = 0x600000 // Accuracy Index (0:undef, 1:<7m, 2:<15m, 3:>15m)
	RxmImesPos2Valid RxmImesPosition2_1 = 0x800000 // Position 2 Frame valid
)

type RxmImesShortIdFrame uint32

const (
	RxmImesShortId       RxmImesShortIdFrame = 0xfff  // Short ID
	RxmImesShortValid    RxmImesShortIdFrame = 0x1000 // Short ID Frame valid
	RxmImesShortBoundary RxmImesShortIdFrame = 0x2000 // Boundary Bit
)

type RxmImesMediumId_2 uint32

const (
	RxmImesMediumIdMSB    RxmImesMediumId_2 = 0x1 // Medium ID MSB
	RxmImesMediumValid    RxmImesMediumId_2 = 0x2 // Medium ID Frame valid
	RxmImesMediumboundary RxmImesMediumId_2 = 0x4 // Boundary Bit
)

func (v RxmImesPosition1_1) String() string {
	var b []string
	if v&0xff != 0 {
		b = append(b, fmt.Sprintf("Pos1Floor<%b>", (v&0xff)>>0))
	}
	v &^= 0xff
	if v&0x7fffff00 != 0 {
		b = append(b, fmt.Sprintf("Pos1Lat<%b>", (v&0x7fffff00)>>8))
	}
	v &^= 0x7fffff00
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v RxmImesPosition1_2) String() string {
	var b []string
	if v&0xffffff != 0 {
		b = append(b, fmt.Sprintf("Pos1Lon<%b>", (v&0xffffff)>>0))
	}
	v &^= 0xffffff
	if v&0x1000000 != 0 {
		b = append(b, "Pos1Valid")
	}
	v &^= 0x1000000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v RxmImesPosition2_1) String() string {
	var b []string
	if v&0x1ff != 0 {
		b = append(b, fmt.Sprintf("Pos2Floor<%b>", (v&0x1ff)>>0))
	}
	v &^= 0x1ff
	if v&0x1ffe00 != 0 {
		b = append(b, fmt.Sprintf("Pos2Alt<%b>", (v&0x1ffe00)>>9))
	}
	v &^= 0x1ffe00
	if v&0x600000 != 0 {
		b = append(b, fmt.Sprintf("Pos2Acc<%b>", (v&0x600000)>>21))
	}
	v &^= 0x600000
	if v&0x800000 != 0 {
		b = append(b, "Pos2Valid")
	}
	v &^= 0x800000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v RxmImesShortIdFrame) String() string {
	var b []string
	if v&0xfff != 0 {
		b = append(b, fmt.Sprintf("ShortId<%b>", (v&0xfff)>>0))
	}
	v &^= 0xfff
	if v&0x1000 != 0 {
		b = append(b, "ShortValid")
	}
	v &^= 0x1000
	if v&0x2000 != 0 {
		b = append(b, "ShortBoundary")
	}
	v &^= 0x2000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v RxmImesMediumId_2) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "MediumIdMSB")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "MediumValid")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "Mediumboundary")
	}
	v &^= 0x4
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-rxm-measx

// RxmMeasx (Periodic/Polled) Satellite measurements for RRLP
// Class/Id 0x02 0x14 (44 + N*24 bytes)
// The message payload data is, where possible and appropriate, according to the Radio Resource LCS (Location Services) Protocol (RRLP) [1]. One exception is the satellite and GNSS IDs, which here are given according to the Satellite Numbering scheme. The correct satellites have to be selected and their satellite ID translated accordingly [1, tab. A.10.14] for use in a RRLP Measure Position Response Component. Similarly, the measurement reference time of week has to be forwarded correctly (modulo 14400000 for the 24 LSB GPS measurements variant, modulo 3600000 for the 22 LSB Galileo and Additional Navigation Satelllite Systems (GANSS) measurements variant) of the RRLP measure position response to the SMLC. Reference: [1] ETSI TS 144 031 V11.0.0 (2012-10), Digital cellular telecommunications system (Phase 2+), Location Services (LCS), Mobile Station (MS) - Serving Mobile Location Centre (SMLC), Radio Resource LCS Protocol (RRLP), (3GPP TS 44.031 version 11.0.0 Release 11).
type RxmMeasx struct {
	Version         byte              // Message version, currently 0x01
	Reserved1       [3]byte           // Reserved
	GpsTOW_ms       uint32            // [ms] GPS measurement reference time
	GloTOW_ms       uint32            // [ms] GLONASS measurement reference time
	BdsTOW_ms       uint32            // [ms] BeiDou measurement reference time
	Reserved2       [4]byte           // Reserved
	QzssTOW_ms      uint32            // [ms] QZSS measurement reference time
	GpsTOWacc_msl4  uint16            // [2^-4 ms] GPS measurement reference time accuracy (0xffff = > 4s)
	GloTOWacc_msl4  uint16            // [2^-4 ms] GLONASS measurement reference time accuracy (0xffff = > 4s)
	BdsTOWacc_msl4  uint16            // [2^-4 ms] BeiDou measurement reference time accuracy (0xffff = > 4s)
	Reserved3       [2]byte           // Reserved
	QzssTOWacc_msl4 uint16            // [2^-4 ms] QZSS measurement reference time accuracy (0xffff = > 4s)
	NumSV           byte              `len:"SV"` // Number of satellites in repeated block
	Flags           RxmMeasxFlags     // Flags
	Reserved4       [8]byte           // Reserved
	SV              []*RxmMeasxSVType // len: NumSV
}

func (RxmMeasx) classID() uint16 { return 0x1402 }

type RxmMeasxSVType struct {
	GnssId          byte    // GNSS ID (see Satellite Numbering)
	SvId            byte    // Satellite ID (see Satellite Numbering)
	CNo             byte    // carrier noise ratio (0..63)
	MpathIndic      byte    // multipath index (according to [1]) (0 = not measured, 1 = low, 2 = medium, 3 = high)
	DopplerMS_m_s   int32   // [0.04 m/s] Doppler measurement
	DopplerHz_hz    int32   // [0.2 Hz] Doppler measurement
	WholeChips      uint16  // whole value of the code phase measurement (0..1022 for GPS)
	FracChips       uint16  // fractional value of the code phase measurement (0..1023)
	CodePhase_msl21 uint32  // [2^-21 ms] Code phase
	IntCodePhase_ms byte    // [ms] Integer (part of the) code phase
	PseuRangeRMSErr byte    // pseudorange RMS error index (according to [1]) (0..63)
	Reserved5       [2]byte // Reserved
}

type RxmMeasxFlags byte

const (
	RxmMeasxTowSet RxmMeasxFlags = 0x3 // TOW set (0 = no, 1 or 2 = yes)
)

func (v RxmMeasxFlags) String() string {
	var b []string
	if v&0x3 != 0 {
		b = append(b, fmt.Sprintf("TowSet<%b>", (v&0x3)>>0))
	}
	v &^= 0x3
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-rxm-pmreq (2 versions)

// RxmPmreq (Command) Power management request
// Class/Id 0x02 0x41 (8 bytes)
// This message requests a power management related task of the receiver.
type RxmPmreq struct {
	Duration_ms uint32        // [ms] Duration of the requested task, set to zero for infinite duration. The maximum supported time is 12 days.
	Flags       RxmPmreqFlags // task flags
}

func (RxmPmreq) classID() uint16 { return 0x4102 }

type RxmPmreqFlags uint32

const (
	RxmPmreqBackup RxmPmreqFlags = 0x2 // The receiver goes into backup mode for a time period defined by duration, provided that it is not connected to USB
)

func (v RxmPmreqFlags) String() string {
	var b []string
	if v&0x2 != 0 {
		b = append(b, "Backup")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// RxmPmreq1 (Command) Power management request
// Class/Id 0x02 0x41 (16 bytes)
// This message requests a power management related task of the receiver.
type RxmPmreq1 struct {
	Version       byte                   // Message version (0x00 for this version)
	Reserved1     [3]byte                // Reserved
	Duration_ms   uint32                 // [ms] Duration of the requested task, set to zero for infinite duration. The maximum supported time is 12 days.
	Flags         RxmPmreq1Flags         // task flags
	WakeupSources RxmPmreq1WakeupSources // Configure pins to wake up the receiver. The receiver wakes up if there is either a falling or a rising edge on one of the configured pins.
}

func (RxmPmreq1) classID() uint16 { return 0x4102 }

type RxmPmreq1Flags uint32

const (
	RxmPmreq1Backup RxmPmreq1Flags = 0x2 // The receiver goes into backup mode for a time period defined by duration, provided that it is not connected to USB
	RxmPmreq1Force  RxmPmreq1Flags = 0x4 // Force receiver backup while USB is connected. USB interface will be disabled.
)

type RxmPmreq1WakeupSources uint32

const (
	RxmPmreq1Uartrx  RxmPmreq1WakeupSources = 0x8  // Wake up the receiver if there is an edge on the UART RX pin
	RxmPmreq1Extint0 RxmPmreq1WakeupSources = 0x20 // Wake up the receiver if there is an edge on the EXTINT0 pin
	RxmPmreq1Extint1 RxmPmreq1WakeupSources = 0x40 // Wake up the receiver if there is an edge on the EXTINT1 pin
	RxmPmreq1Spics   RxmPmreq1WakeupSources = 0x80 // Wake up the receiver if there is an edge on the SPI CS pin
)

func (v RxmPmreq1Flags) String() string {
	var b []string
	if v&0x2 != 0 {
		b = append(b, "Backup")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "Force")
	}
	v &^= 0x4
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v RxmPmreq1WakeupSources) String() string {
	var b []string
	if v&0x8 != 0 {
		b = append(b, "Uartrx")
	}
	v &^= 0x8
	if v&0x20 != 0 {
		b = append(b, "Extint0")
	}
	v &^= 0x20
	if v&0x40 != 0 {
		b = append(b, "Extint1")
	}
	v &^= 0x40
	if v&0x80 != 0 {
		b = append(b, "Spics")
	}
	v &^= 0x80
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-rxm-rawx

// RxmRawx (Periodic/Polled) Multi-GNSS raw measurements
// Class/Id 0x02 0x15 (16 + N*32 bytes)
// This message contains the information needed to be able to generate a RINEX 3 multi-GNSS observation file (see ftp://ftp.igs.org/pub/data/format/). This message contains pseudorange, Doppler, carrier phase, phase lock and signal quality information for GNSS satellites once signals have been synchronized. This message supports all active GNSS. The only difference between this version of the message and the previous version (UBX-RXM-RAWX-DATA0) is the addition of the version field.
type RxmRawx struct {
	RcvTow_s   float64            // [s] Measurement time of week in receiver local time approximately aligned to the GPS time system. The receiver local time of week, week number and leap second information can be used to translate the time to other time systems. More information about the difference in time systems can be found in the RINEX 3 format documentation. For a receiver operating in GLONASS only mode, UTC time can be determined by subtracting the leapS field from GPS time regardless of whether the GPS leap seconds are valid.
	Week_weeks uint16             // [weeks] GPS week number in receiver local time.
	LeapS_s    int8               // [s] GPS leap seconds (GPS-UTC). This field represents the receiver's best knowledge of the leap seconds offset. A flag is given in the recStat bitfield to indicate if the leap seconds are known.
	NumMeas    byte               `len:"Meas"` // Number of measurements to follow
	RecStat    RxmRawxRecStat     // Receiver tracking status bitfield
	Version    byte               // Message version (0x01 for this version)
	Reserved1  [2]byte            // Reserved
	Meas       []*RxmRawxMeasType // len: NumMeas
}

func (RxmRawx) classID() uint16 { return 0x1502 }

type RxmRawxMeasType struct {
	PrMes_m        float64        // [m] Pseudorange measurement [m]. GLONASS inter frequency channel delays are compensated with an internal calibration table.
	CpMes_cycles   float64        // [cycles] Carrier phase measurement [cycles]. The carrier phase initial ambiguity is initialized using an approximate value to make the magnitude of the phase close to the pseudorange measurement. Clock resets are applied to both phase and code measurements in accordance with the RINEX specification.
	DoMes_hz       float32        // [Hz] Doppler measurement (positive sign for approaching satellites) [Hz]
	GnssId         byte           // GNSS identifier (see Satellite Numbering for a list of identifiers)
	SvId           byte           // Satellite identifier (see Satellite Numbering)
	SigId          byte           // New style signal identifier (see Signal Identifiers).(not supported in protocol versions less than 27)
	FreqId         byte           // Only used for GLONASS: This is the frequency slot + 7 (range from 0 to 13)
	Locktime_ms    uint16         // [ms] Carrier phase locktime counter (maximum 64500ms)
	Cno_dbhz       byte           // [dBHz] Carrier-to-noise density ratio (signal strength) [dB-Hz]
	PrStdev_m      RxmRawxPrStdev // [0.01*2^n m] Estimated pseudorange measurement standard deviation
	CpStdev_cycles RxmRawxCpStdev // [0.004 cycles] Estimated carrier phase measurement standard deviation (note a raw value of 0x0F indicates the value is invalid)
	DoStdev_hz     RxmRawxDoStdev // [0.002*2^n Hz] Estimated Doppler measurement standard deviation.
	TrkStat        RxmRawxTrkStat // Tracking status bitfield (see graphic below )
	Reserved2      byte           // Reserved
}

type RxmRawxRecStat byte

const (
	RxmRawxLeapSec  RxmRawxRecStat = 0x1 // Leap seconds have been determined
	RxmRawxClkReset RxmRawxRecStat = 0x2 // Clock reset applied. Typically the receiver clock is changed in increments of integer milliseconds.
)

type RxmRawxPrStdev byte

const (
	RxmRawxPrStd RxmRawxPrStdev = 0xf // Estimated pseudorange standard deviation
)

type RxmRawxCpStdev byte

const (
	RxmRawxCpStd RxmRawxCpStdev = 0xf // Estimated carrier phase standard deviation
)

type RxmRawxDoStdev byte

const (
	RxmRawxDoStd RxmRawxDoStdev = 0xf // Estimated Doppler standard deviation
)

type RxmRawxTrkStat byte

const (
	RxmRawxPrValid    RxmRawxTrkStat = 0x1 // Pseudorange valid
	RxmRawxCpValid    RxmRawxTrkStat = 0x2 // Carrier phase valid
	RxmRawxHalfCyc    RxmRawxTrkStat = 0x4 // Half cycle valid
	RxmRawxSubHalfCyc RxmRawxTrkStat = 0x8 // Half cycle subtracted from phase
)

func (v RxmRawxRecStat) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "LeapSec")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "ClkReset")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v RxmRawxPrStdev) String() string {
	var b []string
	if v&0xf != 0 {
		b = append(b, fmt.Sprintf("PrStd<%b>", (v&0xf)>>0))
	}
	v &^= 0xf
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v RxmRawxCpStdev) String() string {
	var b []string
	if v&0xf != 0 {
		b = append(b, fmt.Sprintf("CpStd<%b>", (v&0xf)>>0))
	}
	v &^= 0xf
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v RxmRawxDoStdev) String() string {
	var b []string
	if v&0xf != 0 {
		b = append(b, fmt.Sprintf("DoStd<%b>", (v&0xf)>>0))
	}
	v &^= 0xf
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v RxmRawxTrkStat) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "PrValid")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "CpValid")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "HalfCyc")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "SubHalfCyc")
	}
	v &^= 0x8
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-rxm-rlm (2 versions)

// RxmRlm (Output) Galileo SAR short-RLM report
// Class/Id 0x02 0x59 (16 bytes)
// This message contains the contents of any Galileo Search and Rescue (SAR) Short Return Link Message detected by the receiver.
type RxmRlm struct {
	Version   byte    // Message version (0x00 for this version)
	Type      byte    // Message type (0x01 for Short-RLM)
	SvId      byte    // Identifier of transmitting satellite (see Satellite Numbering)
	Reserved1 byte    // Reserved
	Beacon    [8]byte // Beacon identifier (60 bits), with bytes ordered by earliest transmitted (most significant) first. Top four bits of first byte are zero.
	Message   byte    // Message code (4 bits)
	Params    [2]byte // Parameters (16 bits), with bytes ordered by earliest transmitted (most significant) first.
	Reserved2 byte    // Reserved
}

func (RxmRlm) classID() uint16 { return 0x5902 }

// RxmRlm1 (Output) Galileo SAR long-RLM report
// Class/Id 0x02 0x59 (28 bytes)
// This message contains the contents of any Galileo Search and Rescue (SAR) Long Return Link Message detected by the receiver.
type RxmRlm1 struct {
	Version   byte     // Message version (0x00 for this version)
	Type      byte     // Message type (0x02 for Long-RLM)
	SvId      byte     // Identifier of transmitting satellite (see Satellite Numbering)
	Reserved1 byte     // Reserved
	Beacon    [8]byte  // Beacon identifier (60 bits), with bytes ordered by earliest transmitted (most significant) first. Top four bits of first byte are zero.
	Message   byte     // Message code (4 bits)
	Params    [12]byte // Parameters (96 bits), with bytes ordered by earliest transmitted (most significant) first.
	Reserved2 [3]byte  // Reserved
}

func (RxmRlm1) classID() uint16 { return 0x5902 }

// Message ubx-rxm-rtcm

// RxmRtcm (Output) RTCM input status
// Class/Id 0x02 0x32 (8 bytes)
// This message shows info on a received RTCM input message. It is output upon successful parsing of an RTCM input message, irrespective of whether the RTCM message is supported or not by the receiver.
type RxmRtcm struct {
	Version    byte         // Message version (0x02 for this version)
	Flags      RxmRtcmFlags // RTCM input status flags
	SubType    uint16       // Message subtype, only applicable to u- blox proprietary RTCM message 4072 (not available on all products)
	RefStation uint16       // Reference station ID: For RTCM 2.3: Reference station ID of the received RTCM 2 input message. Valid range 0-1023. For RTCM 3.3: Reference station ID (DF003) of the received RTCM input message. Valid range 0-4095. Reported only for the standard RTCM messages that include the DF003 field and for the u- blox proprietary RTCM messages 4072.x. For all other messages, reports 0xFFFF.
	MsgType    uint16       // Message type
}

func (RxmRtcm) classID() uint16 { return 0x3202 }

type RxmRtcmFlags byte

const (
	RxmRtcmCrcFailed RxmRtcmFlags = 0x1 // 0 when RTCM message received and passed CRC check, 1 when failed, in which case refStation and msgType might be corrupted and misleading
	RxmRtcmMsgUsed   RxmRtcmFlags = 0x6 // 2 = RTCM message used successfully by the receiver, 1 = not used, 0 = do not know
)

func (v RxmRtcmFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "CrcFailed")
	}
	v &^= 0x1
	if v&0x6 != 0 {
		b = append(b, fmt.Sprintf("MsgUsed<%b>", (v&0x6)>>1))
	}
	v &^= 0x6
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-rxm-sfrbx

// RxmSfrbx (Output) Broadcast navigation data subframe
// Class/Id 0x02 0x13 (8 + N*4 bytes)
// This message reports a complete subframe of broadcast navigation data decoded from a single signal. The number of data words reported in each message depends on the nature of the signal. See the section on Broadcast Navigation Data for further details.
type RxmSfrbx struct {
	GnssId    byte                 // GNSS identifier (see Satellite Numbering)
	SvId      byte                 // Satellite identifier (see Satellite Numbering)
	Reserved1 byte                 // Reserved
	FreqId    byte                 // Only used for GLONASS: This is the frequency slot + 7 (range from 0 to 13)
	NumWords  byte                 `len:"Words"` // The number of data words contained in this message (up to 10, for currently supported signals)
	Chn       byte                 // The tracking channel number the message was received on
	Version   byte                 // Message version, (0x02 for this version)
	Reserved2 byte                 // Reserved
	Words     []*RxmSfrbxWordsType // len: NumWords
}

func (RxmSfrbx) classID() uint16 { return 0x1302 }

type RxmSfrbxWordsType struct {
	Dwrd uint32 // The data words
}

// Message ubx-rxm-svsi

// RxmSvsi (Periodic/Polled) SV status info
// Class/Id 0x02 0x20 (8 + N*6 bytes)
// Status of the receiver manager knowledge about GPS Orbit Validity This message has only been retained for backwards compatibility; users are recommended to use the UBX-NAV-ORB message in preference.
type RxmSvsi struct {
	ITOW_ms    uint32           // [ms] GPS time of week of the navigation epoch. See the description of iTOW for details.
	Week_weeks int16            // [weeks] GPS week number of the navigation epoch
	NumVis     byte             // Number of visible satellites
	NumSV      byte             `len:"SV"` // Number of per-SV data blocks following
	SV         []*RxmSvsiSVType // len: NumSV
}

func (RxmSvsi) classID() uint16 { return 0x2002 }

type RxmSvsiSVType struct {
	Svid   byte          // Satellite ID
	SvFlag RxmSvsiSvFlag // Information Flags
	Azim   int16         // Azimuth
	Elev   int8          // Elevation
	Age    RxmSvsiAge    // Age of Almanac and Ephemeris:
}

type RxmSvsiSvFlag byte

const (
	RxmSvsiUra      RxmSvsiSvFlag = 0xf  // Figure of Merit (URA) range 0..15
	RxmSvsiHealthy  RxmSvsiSvFlag = 0x10 // SV healthy flag
	RxmSvsiEphVal   RxmSvsiSvFlag = 0x20 // Ephemeris valid
	RxmSvsiAlmVal   RxmSvsiSvFlag = 0x40 // Almanac valid
	RxmSvsiNotAvail RxmSvsiSvFlag = 0x80 // SV not available
)

type RxmSvsiAge byte

const (
	RxmSvsiAlmAge RxmSvsiAge = 0xf  // Age of ALM in days offset by 4 i.e. the reference time may be in the future: ageOfAlm = (age & 0x0f) - 4
	RxmSvsiEphAge RxmSvsiAge = 0xf0 // Age of EPH in hours offset by 4. i.e. the reference time may be in the future: ageOfEph = ((age & 0xf0) >> 4) - 4
)

func (v RxmSvsiSvFlag) String() string {
	var b []string
	if v&0xf != 0 {
		b = append(b, fmt.Sprintf("Ura<%b>", (v&0xf)>>0))
	}
	v &^= 0xf
	if v&0x10 != 0 {
		b = append(b, "Healthy")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "EphVal")
	}
	v &^= 0x20
	if v&0x40 != 0 {
		b = append(b, "AlmVal")
	}
	v &^= 0x40
	if v&0x80 != 0 {
		b = append(b, "NotAvail")
	}
	v &^= 0x80
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v RxmSvsiAge) String() string {
	var b []string
	if v&0xf != 0 {
		b = append(b, fmt.Sprintf("AlmAge<%b>", (v&0xf)>>0))
	}
	v &^= 0xf
	if v&0xf0 != 0 {
		b = append(b, fmt.Sprintf("EphAge<%b>", (v&0xf0)>>4))
	}
	v &^= 0xf0
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-sec-uniqid

// SecUniqid (Output) Unique chip ID
// Class/Id 0x27 0x03 (9 bytes)
// This message is used to retrieve a unique chip identifier (40 bits, 5 bytes).
type SecUniqid struct {
	Version   byte    // Message version (0x01 for this version)
	Reserved1 [3]byte // Reserved
	UniqueId  [5]byte // Unique chip ID
}

func (SecUniqid) classID() uint16 { return 0x0327 }

// Message ubx-tim-dosc

// TimDosc (Output) Disciplined oscillator control
// Class/Id 0x0d 0x11 (8 bytes)
// The receiver sends this message when it is disciplining an external oscillator and the external oscillator is set up to be controlled via the host.
type TimDosc struct {
	Version   byte    // Message version (0x00 for this version)
	Reserved1 [3]byte // Reserved
	Value     uint32  // The raw value to be applied to the DAC controlling the external oscillator. The least significant bits should be written to the DAC, with the higher bits being ignored.
}

func (TimDosc) classID() uint16 { return 0x110d }

// Message ubx-tim-fchg

// TimFchg (Periodic/Polled) Oscillator frequency changed notification
// Class/Id 0x0d 0x16 (32 bytes)
// This message reports frequency changes commanded by the sync manager for the internal and external oscillator. It is output at the configured rate even if the sync manager decides not to command a frequency change.
type TimFchg struct {
	Version               byte    // Message version (0x00 for this version)
	Reserved1             [3]byte // Reserved
	ITOW_ms               uint32  // [ms] GPS time of week of the navigation epoch from which the sync manager obtains the GNSS specific data. Like for the NAV message, the iTOW can be used to group messages of a single sync manager run together (See the description of iTOW for details)
	IntDeltaFreq_ppbl8    int32   // [2^-8 ppb] Frequency increment of the internal oscillator
	IntDeltaFreqUnc_ppbl8 uint32  // [2^-8 ppb] Uncertainty of the internal oscillator frequency increment
	IntRaw                uint32  // Current raw DAC setting commanded to the internal oscillator
	ExtDeltaFreq_ppbl8    int32   // [2^-8 ppb] Frequency increment of the external oscillator
	ExtDeltaFreqUnc_ppbl8 uint32  // [2^-8 ppb] Uncertainty of the external oscillator frequency increment
	ExtRaw                uint32  // Current raw DAC setting commanded to the external oscillator
}

func (TimFchg) classID() uint16 { return 0x160d }

// Message ubx-tim-hoc

// TimHoc (Input) Host oscillator control
// Class/Id 0x0d 0x17 (8 bytes)
// This message can be sent by the host to force the receiver to bypass the disciplining algorithms in the SMGR and carry out the instructed changes to internal or external oscillator frequency. No checks are carried out on the size of the frequency change requested, so normal limits imposed by the SMGR are ignored. It is recommended that the disciplining of that oscillator is disabled before this message is sent (i.e. by clearing the enableInternal or enableExternal flag in the UBX-CFG-SMGR message), otherwise the autonomous disciplining processes may cancel the effect of the direct command. Note that the GNSS subsystem may temporarily lose track of some/all satellite signals if a large change of the internal oscillator is made.
type TimHoc struct {
	Version   byte        // Message version (0x00 for this version)
	OscId     byte        // Id of oscillator: 0: internal oscillator 1: external oscillator
	Flags     TimHocFlags // Flags
	Reserved1 byte        // Reserved
	Value     int32       // [2^-8 ppb/-] Required frequency offset or raw output, depending on the flags
}

func (TimHoc) classID() uint16 { return 0x170d }

type TimHocFlags byte

const (
	TimHocRaw        TimHocFlags = 0x1 // Type of value: 0: frequency offset 1: raw digital output
	TimHocDifference TimHocFlags = 0x2 // Nature of value: 0: absolute (i.e. relative to 0) 1: relative to current setting
)

func (v TimHocFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Raw")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "Difference")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-tim-smeas

// TimSmeas (Input/Output) Source measurement
// Class/Id 0x0d 0x13 (12 + N*24 bytes)
// Frequency and/or phase measurement of synchronization sources. The measurements are relative to the nominal frequency and nominal phase. The receiver reports the measurements on its sync sources using this message. Which measurements are reported can be configured using UBX-CFG-SMGR. The host may report offset of the receiver's outputs with this message as well. The receiver has to be configured using UBX-CFG-SMGR to enable the use of the external measurement messages. Otherwise the receiver will ignore them.
type TimSmeas struct {
	Version   byte                // Message version (0x00 for this version)
	NumMeas   byte                `len:"Meas"` // Number of measurements in repeated block
	Reserved1 [2]byte             // Reserved
	ITOW_ms   uint32              // [ms] Time of the week
	Reserved2 [4]byte             // Reserved
	Meas      []*TimSmeasMeasType // len: NumMeas
}

func (TimSmeas) classID() uint16 { return 0x130d }

type TimSmeasMeasType struct {
	SourceId             byte          // Index of source. SMEAS can provide six measurement sources. The first four sourceId values represent measurements made by the receiver and sent to the host. The first of these with a sourceId value of 0 is a measurement of the internal oscillator against the current receiver time-and-frequency estimate. The internal oscillator is being disciplined against that estimate and this result represents the current offset between the actual and desired internal oscillator states. The next three sourceId values represent frequency and time measurements made by the receiver against the internal oscillator. sourceId 1 represents the GNSS-derived frequency and time compared with the internal oscillator frequency and time. sourceId2 give measurements of a signal coming in on EXTINT0. sourceId 3 corresponds to a similar measurement on EXTINT1. The remaining two of these measurements (sourceId 4 and 5) are made by the host and sent to the receiver. A measurement with sourceId 4 is a measurement by the host of the internal oscillator and sourceId 5 indicates a host measurement of the external oscillator.
	Flags                TimSmeasFlags // Flags
	PhaseOffsetFrac_nsl8 int8          // [2^-8 ns] Sub-nanosecond phase offset; the total offset is the sum of phaseOffset and phaseOffsetFrac
	PhaseUncFrac_nsl8    byte          // [2^-8 ns] Sub-nanosecond phase uncertainty
	PhaseOffset_ns       int32         // [ns] Phase offset, positive if the source lags accurate phase and negative if the source is early
	PhaseUnc_ns          uint32        // [ns] Phase uncertainty (one standard deviation)
	Reserved3            [4]byte       // Reserved
	FreqOffset_ppbl8     int32         // [2^-8 ppb] Frequency offset, positive if the source frequency is too high, negative if the frequency is too low.
	FreqUnc_ppbl8        uint32        // [2^-8 ppb] Frequency uncertainty (one standard deviation)
}

type TimSmeasFlags byte

const (
	TimSmeasFreqValid  TimSmeasFlags = 0x1 // 1 = frequency measurement is valid
	TimSmeasPhaseValid TimSmeasFlags = 0x2 // 1 = phase measurement is valid
)

func (v TimSmeasFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "FreqValid")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "PhaseValid")
	}
	v &^= 0x2
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-tim-svin

// TimSvin (Periodic/Polled) Survey-in data
// Class/Id 0x0d 0x04 (28 bytes)
// This message contains information about survey-in parameters. For details about the Time mode see section Time mode configuration.
type TimSvin struct {
	Dur_s     uint32  // [s] Passed survey-in observation time
	MeanX_cm  int32   // [cm] Current survey-in mean position ECEF X coordinate
	MeanY_cm  int32   // [cm] Current survey-in mean position ECEF Y coordinate
	MeanZ_cm  int32   // [cm] Current survey-in mean position ECEF Z coordinate
	MeanV_mm2 uint32  // [mm^2] Current survey-in mean position 3D variance
	Obs       uint32  // Number of position observations used during survey-in
	Valid     byte    // Survey-in position validity flag, 1 = valid, otherwise 0
	Active    byte    // Survey-in in progress flag, 1 = in-progress, otherwise 0
	Reserved1 [2]byte // Reserved
}

func (TimSvin) classID() uint16 { return 0x040d }

// Message ubx-tim-tm2

// TimTm2 (Periodic/Polled) Time mark data
// Class/Id 0x0d 0x03 (28 bytes)
// This message contains information for high precision time stamping / pulse counting. The delay figures and timebase given in UBX-CFG-TP5 are also applied to the time results output in this message.
type TimTm2 struct {
	Ch           byte        // Channel (i.e. EXTINT) upon which the pulse was measured
	Flags        TimTm2Flags // Bitmask
	Count        uint16      // Rising edge counter
	WnR          uint16      // Week number of last rising edge
	WnF          uint16      // Week number of last falling edge
	TowMsR_ms    uint32      // [ms] Tow of rising edge
	TowSubMsR_ns uint32      // [ns] Millisecond fraction of tow of rising edge in nanoseconds
	TowMsF_ms    uint32      // [ms] Tow of falling edge
	TowSubMsF_ns uint32      // [ns] Millisecond fraction of tow of falling edge in nanoseconds
	AccEst_ns    uint32      // [ns] Accuracy estimate
}

func (TimTm2) classID() uint16 { return 0x030d }

type TimTm2Flags byte

const (
	TimTm2Mode           TimTm2Flags = 0x1  // 0=single 1=running
	TimTm2Run            TimTm2Flags = 0x2  // 0=armed 1=stopped
	TimTm2NewFallingEdge TimTm2Flags = 0x4  // New falling edge detected
	TimTm2TimeBase       TimTm2Flags = 0x18 // 0=Time base is Receiver time 1=Time base is GNSS time (the system according to the configuration in UBX-CFG-TP5 for tpIdx=0) 2=Time base is UTC (the variant according to the configuration in UBX-CFG-NAV5)
	TimTm2Utc            TimTm2Flags = 0x20 // 0=UTC not available 1=UTC available
	TimTm2Time           TimTm2Flags = 0x40 // 0=Time is not valid 1=Time is valid (Valid GNSS fix)
	TimTm2NewRisingEdge  TimTm2Flags = 0x80 // New rising edge detected
)

func (v TimTm2Flags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "Mode")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "Run")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "NewFallingEdge")
	}
	v &^= 0x4
	if v&0x18 != 0 {
		b = append(b, fmt.Sprintf("TimeBase<%b>", (v&0x18)>>3))
	}
	v &^= 0x18
	if v&0x20 != 0 {
		b = append(b, "Utc")
	}
	v &^= 0x20
	if v&0x40 != 0 {
		b = append(b, "Time")
	}
	v &^= 0x40
	if v&0x80 != 0 {
		b = append(b, "NewRisingEdge")
	}
	v &^= 0x80
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-tim-tos

// TimTos (Periodic) Time pulse time and frequency data
// Class/Id 0x0d 0x12 (56 bytes)
// This message contains information about the time pulse that has just happened and the state of the disciplined oscillators(s) at the time of the pulse. It gives the UTC and GNSS times and time uncertainty of the pulse together with frequency and frequency uncertainty of the disciplined oscillators. It also supplies leap second information.
type TimTos struct {
	Version                 byte        // Message version (0x00 for this version)
	GnssId                  byte        // GNSS system used for reporting GNSS time (see Satellite Numbering)
	Reserved1               [2]byte     // Reserved
	Flags                   TimTosFlags // Flags
	Year_y                  uint16      // [y] Year of UTC time
	Month_month             byte        // [month] Month of UTC time
	Day_d                   byte        // [d] Day of UTC time
	Hour_h                  byte        // [h] Hour of UTC time
	Minute_min              byte        // [min] Minute of UTC time
	Second_s                byte        // [s] Second of UTC time
	UtcStandard             byte        // UTC standard identifier: 0: unknown 3: UTC as operated by the U.S. Naval Observatory (USNO) 6: UTC as operated by the former Soviet Union 7: UTC as operated by the National Time Service Center (NTSC), China
	UtcOffset_ns            int32       // [ns] Time offset between the preceding pulse and UTC top of second
	UtcUncertainty_ns       uint32      // [ns] Uncertainty of utcOffset
	Week                    uint32      // GNSS week number
	TOW_s                   uint32      // [s] GNSS time of week
	GnssOffset_ns           int32       // [ns] Time offset between the preceding pulse and GNSS top of second
	GnssUncertainty_ns      uint32      // [ns] Uncertainty of gnssOffset
	IntOscOffset_ppbl8      int32       // [2^-8 ppb] Internal oscillator frequency offset
	IntOscUncertainty_ppbl8 uint32      // [2^-8 ppb] Internal oscillator frequency uncertainty
	ExtOscOffset_ppbl8      int32       // [2^-8 ppb] External oscillator frequency offset
	ExtOscUncertainty_ppbl8 uint32      // [2^-8 ppb] External oscillator frequency uncertainty
}

func (TimTos) classID() uint16 { return 0x120d }

type TimTosFlags uint32

const (
	TimTosLeapNow       TimTosFlags = 0x1    // 1 = currently in a leap second
	TimTosLeapSoon      TimTosFlags = 0x2    // 1 = leap second scheduled in current minute
	TimTosLeapPositive  TimTosFlags = 0x4    // 1 = positive leap second
	TimTosTimeInLimit   TimTosFlags = 0x8    // 1 = time pulse is within tolerance limit (UBX-CFG-SMGR timeTolerance field)
	TimTosIntOscInLimit TimTosFlags = 0x10   // 1 = internal oscillator is within tolerance limit (UBX-CFG-SMGR freqTolerance field)
	TimTosExtOscInLimit TimTosFlags = 0x20   // 1 = external oscillator is within tolerance limit (UBX-CFG-SMGR freqTolerance field)
	TimTosGnssTimeValid TimTosFlags = 0x40   // 1 = GNSS time is valid
	TimTosUTCTimeValid  TimTosFlags = 0x80   // 1 = UTC time is valid
	TimTosDiscSrc       TimTosFlags = 0x700  // Disciplining source identifier: 0: internal oscillator 1: GNSS 2: EXTINT0 3: EXTINT1 4: internal oscillator measured by the host 5: external oscillator measured by the host
	TimTosRaim          TimTosFlags = 0x800  // 1 = (T)RAIM system is currently active. Note this flag only reports the current state of the GNSS solution; it is not affected by whether or not the GNSS solution is being used to discipline the oscillator.
	TimTosCohPulse      TimTosFlags = 0x1000 // 1 = coherent pulse generation is currently in operation
	TimTosLockedPulse   TimTosFlags = 0x2000 // 1 = time pulse is locked
)

func (v TimTosFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "LeapNow")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "LeapSoon")
	}
	v &^= 0x2
	if v&0x4 != 0 {
		b = append(b, "LeapPositive")
	}
	v &^= 0x4
	if v&0x8 != 0 {
		b = append(b, "TimeInLimit")
	}
	v &^= 0x8
	if v&0x10 != 0 {
		b = append(b, "IntOscInLimit")
	}
	v &^= 0x10
	if v&0x20 != 0 {
		b = append(b, "ExtOscInLimit")
	}
	v &^= 0x20
	if v&0x40 != 0 {
		b = append(b, "GnssTimeValid")
	}
	v &^= 0x40
	if v&0x80 != 0 {
		b = append(b, "UTCTimeValid")
	}
	v &^= 0x80
	if v&0x700 != 0 {
		b = append(b, fmt.Sprintf("DiscSrc<%b>", (v&0x700)>>8))
	}
	v &^= 0x700
	if v&0x800 != 0 {
		b = append(b, "Raim")
	}
	v &^= 0x800
	if v&0x1000 != 0 {
		b = append(b, "CohPulse")
	}
	v &^= 0x1000
	if v&0x2000 != 0 {
		b = append(b, "LockedPulse")
	}
	v &^= 0x2000
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-tim-tp

// TimTp (Periodic/Polled) Time pulse time data
// Class/Id 0x0d 0x01 (16 bytes)
// This message contains information on the timing of the next pulse at the TIMEPULSE0 output. The recommended configuration when using this message is to set both the measurement rate (UBX-CFG-RATE) and the timepulse frequency (UBX-CFG-TP5) to 1 Hz. For more information see section Time pulse. TIMEPULSE0 and this message are not available from DR products using the dedicated I2C sensor interface, including NEO-M8L and NEO-M8U modules
type TimTp struct {
	TowMS_ms       uint32       // [ms] Time pulse time of week according to time base
	TowSubMS_msl32 uint32       // [2^-32 ms] Submillisecond part of towMS
	QErr_ps        int32        // [ps] Quantization error of time pulse
	Week_weeks     uint16       // [weeks] Time pulse week number according to time base
	Flags          TimTpFlags   // Flags
	RefInfo        TimTpRefInfo // Time reference information
}

func (TimTp) classID() uint16 { return 0x010d }

type TimTpFlags byte

const (
	TimTpTimeBase    TimTpFlags = 0x1  // 0 = Time base is GNSS 1 = Time base is UTC
	TimTpUtc         TimTpFlags = 0x2  // 0 = UTC not available 1 = UTC available
	TimTpRaim        TimTpFlags = 0xc  // (T)RAIM information 0 = Information not available 1 = Not active 2 = Active
	TimTpQErrInvalid TimTpFlags = 0x10 // 0 = Quantization error valid 1 = Quantization error invalid
)

type TimTpRefInfo byte

const (
	TimTpTimeRefGnss TimTpRefInfo = 0xf  // GNSS reference information. Only valid if time base is GNSS (timeBase=0). 0 = GPS 1 = GLONASS 2 = BeiDou 3 = Galileo 15 = Unknown
	TimTpUtcStandard TimTpRefInfo = 0xf0 // UTC standard identifier. Only valid if time base is UTC (timeBase=1). 0 = Information not available 1 = Communications Research Laboratory (CRL), Tokyo, Japan 2 = National Institute of Standards and Technology (NIST) 3 = U.S. Naval Observatory (USNO) 4 = International Bureau of Weights and Measures (BIPM) 5 = European laboratories 6 = Former Soviet Union (SU) 7 = National Time Service Center (NTSC), China 15 = Unknown
)

func (v TimTpFlags) String() string {
	var b []string
	if v&0x1 != 0 {
		b = append(b, "TimeBase")
	}
	v &^= 0x1
	if v&0x2 != 0 {
		b = append(b, "Utc")
	}
	v &^= 0x2
	if v&0xc != 0 {
		b = append(b, fmt.Sprintf("Raim<%b>", (v&0xc)>>2))
	}
	v &^= 0xc
	if v&0x10 != 0 {
		b = append(b, "QErrInvalid")
	}
	v &^= 0x10
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

func (v TimTpRefInfo) String() string {
	var b []string
	if v&0xf != 0 {
		b = append(b, fmt.Sprintf("TimeRefGnss<%b>", (v&0xf)>>0))
	}
	v &^= 0xf
	if v&0xf0 != 0 {
		b = append(b, fmt.Sprintf("UtcStandard<%b>", (v&0xf0)>>4))
	}
	v &^= 0xf0
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-tim-vcocal (3 versions)

// TimVcocal (Command) Stop calibration
// Class/Id 0x0d 0x15 (1 bytes)
// Stop all ongoing calibration (both oscillators are affected)
type TimVcocal struct {
	Type byte `stf:"0"` // Message type (0 for this message)
}

func (TimVcocal) classID() uint16 { return 0x150d }

// TimVcocal1 (Command) VCO calibration extended command
// Class/Id 0x0d 0x15 (12 bytes)
// Calibrate (measure) gain of the voltage controlled oscillator. The calibration is performed by varying the raw oscillator control values between the limits specified in raw0 and raw1. maxStepSize is the largest step change that can be used during the calibration process. The "raw values" are either PWM duty cycle values or DAC values depending on how the VCTCXO is connected to the system. The measured gain is the transfer function dRelativeFrequencyChange/dRaw (not dFrequency/dVoltage). The calibration process works as follows: Starting from the current raw output the control value is changed in the direction of raw0 in steps of size at most maxStepSize. Then the frequency is measured and the control value is changed towards raw1, again in steps of maxStepSize. When raw1 is reached, the frequency is again measured and the message version DATA0 is output containing the measured result. Normal operation then resumes. If the control value movement is less than maxStepSize then the transition will happen in one step - this will give fast calibration. Care must be taken when calibrating the internal oscillator against the GNSS source. In that case the changes applied to the oscillator frequency could be severe enough to lose satellite signal tracking, especially when signals are weak. If too many signals are lost, the GNSS system will lose its fix and be unable to measure the oscillator frequency - the calibration will then fail. In this case maxStepSize must be reasonably small. It is also important that only the chosen frequency source is enabled during the calibration process and that it remains stable throughout the calibration period; otherwise incorrect oscillator measurements will be made and this will lead to miscalibration and poor subsequent operation of the receiver.
type TimVcocal1 struct {
	Type                   byte    `stf:"2"` // Message type (2 for this message)
	Version                byte    // Message version (0x00 for this version)
	OscId                  byte    // Oscillator to be calibrated: 0: internal oscillator 1: external oscillator
	SrcId                  byte    // Reference source: 0: internal oscillator 1: GNSS 2: EXTINT0 3: EXTINT1 Option 0 should be used when calibrating the external oscillator. Options 1-3 should be used when calibrating the internal oscillator.
	Reserved1              [2]byte // Reserved
	Raw0                   uint16  // First value used for calibration
	Raw1                   uint16  // Second value used for calibration
	MaxStepSize_rawvalue_s uint16  // [rawvalue/s] Maximum step size to be used
}

func (TimVcocal1) classID() uint16 { return 0x150d }

// TimVcocal2 (Periodic/Polled) Results of the calibration
// Class/Id 0x0d 0x15 (12 bytes)
// This message is sent when the oscillator gain calibration process is finished (successful or unsuccessful). It notifies the user of the calibrated oscillator gain. If the oscillator gain calibration process was successful, this message will contain the measured gain (field gainVco) and its uncertainty (field gainUncertainty). The calibration process can however fail. In that case the two fields gainVco and gainUncertainty are set to zero.
type TimVcocal2 struct {
	Type            byte    `stf:"3"` // Message type (3 for this message)
	Version         byte    // Message version (0x00 for this version)
	OscId           byte    // Id of oscillator: 0: internal oscillator 1: external oscillator
	Reserved1       [3]byte // Reserved
	GainUncertainty uint16  // [2^-16 1/1] Relative gain uncertainty after calibration, 0 if calibration failed
	GainVco         int32   // [2^-16 ppb/raw LSB] Calibrated gain or 0 if calibration failed
}

func (TimVcocal2) classID() uint16 { return 0x150d }

// Message ubx-tim-vrfy

// TimVrfy (Periodic/Polled) Sourced time verification
// Class/Id 0x0d 0x06 (20 bytes)
// This message contains verification information about previous time received via assistance data or from RTC.
type TimVrfy struct {
	Itow_ms    int32        // [ms] integer millisecond tow received by source
	Frac_ns    int32        // [ns] sub-millisecond part of tow
	DeltaMs_ms int32        // [ms] integer milliseconds of delta time (current time minus sourced time)
	DeltaNs_ns int32        // [ns] Sub-millisecond part of delta time
	Wno_week   uint16       // [week] Week number
	Flags      TimVrfyFlags // Flags
	Reserved1  byte         // Reserved
}

func (TimVrfy) classID() uint16 { return 0x060d }

type TimVrfyFlags byte

const (
	TimVrfySrc TimVrfyFlags = 0x7 // Aiding time source 0 = no time aiding done 2 = source was RTC 3 = source was assistance data
)

func (v TimVrfyFlags) String() string {
	var b []string
	if v&0x7 != 0 {
		b = append(b, fmt.Sprintf("Src<%b>", (v&0x7)>>0))
	}
	v &^= 0x7
	if v != 0 {
		b = append(b, fmt.Sprintf("%x", uint(v)))
	}
	return strings.Join(b, ",")
}

// Message ubx-upd-sos (5 versions)

// UpdSos (Poll Request) Poll backup restore status
// Class/Id 0x09 0x14 (0 bytes)
// Sending this (empty) message to the receiver results in the receiver returning a System restored from backup message as defined below.
type UpdSos struct {
}

func (UpdSos) classID() uint16 { return 0x1409 }

// UpdSos1 (Command) Create backup in flash
// Class/Id 0x09 0x14 (4 bytes)
// The host can send this message in order to save part of the battery-backed memory (BBR) in a file in the flash file system. The feature is designed in order to emulate the presence of the backup battery even if it is not present; the host can issue the save on shutdown command before switching off the device supply. It is recommended to issue a GNSS stop command using UBX-CFG-RST before in order to keep the BBR memory content consistent.
type UpdSos1 struct {
	Cmd       byte    `stf:"0"` // Command (must be 0)
	Reserved1 [3]byte // Reserved
}

func (UpdSos1) classID() uint16 { return 0x1409 }

// UpdSos2 (Command) Clear backup in flash
// Class/Id 0x09 0x14 (4 bytes)
// The host can send this message in order to erase the backup file present in flash. It is recommended that the clear operation is issued after the host has received the notification that the memory has been restored after a reset. Alternatively the host can parse the startup string Restored data saved on shutdown or poll the UBX-UPD-SOS message for obtaining the status.
type UpdSos2 struct {
	Cmd       byte    `stf:"1"` // Command (must be 1)
	Reserved1 [3]byte // Reserved
}

func (UpdSos2) classID() uint16 { return 0x1409 }

// UpdSos3 (Output) Backup creation acknowledge
// Class/Id 0x09 0x14 (8 bytes)
// The message is sent from the device as confirmation of creation of a backup file in flash. The host can safely shut down the device after having received this message.
type UpdSos3 struct {
	Cmd       byte    `stf:"2"` // Command (must be 2)
	Reserved1 [3]byte // Reserved
	Response  byte    // 0 = Not acknowledged 1 = Acknowledged
	Reserved2 [3]byte // Reserved
}

func (UpdSos3) classID() uint16 { return 0x1409 }

// UpdSos4 (Output) System restored from backup
// Class/Id 0x09 0x14 (8 bytes)
// The message is sent from the device to notify the host the BBR has been restored from a backup file in the flash file sysetem. The host should clear the backup file after receiving this message. If the UBX-UPD-SOS message is polled, this message will be resent.
type UpdSos4 struct {
	Cmd       byte    `stf:"3"` // Command (must be 3)
	Reserved1 [3]byte // Reserved
	Response  byte    // 0 = Unknown 1 = Failed restoring from backup 2 = Restored from backup 3 = Not restored (no backup)
	Reserved2 [3]byte // Reserved
}

func (UpdSos4) classID() uint16 { return 0x1409 }

func mkMsg(classId, sz uint16, frame []byte) Message {
	switch classId {

	case 0x0105:
		return new(AckAck)

	case 0x0005:
		return new(AckNak)

	case 0x300b:

		switch sz {
		case 0:
			return new(AidAlm) // 0 0 [0]
		case 1:
			return new(AidAlm1) // 1 1 [0]
		case 8, 40:
			return new(AidAlm2) // 8 40 [0]
		}

	case 0x330b:

		switch sz {
		case 0:
			return new(AidAop) // 0 0 [0]
		case 1:
			return new(AidAop1) // 1 1 [0]
		case 68:
			return new(AidAop2) // 68 68 [0]
		}

	case 0x310b:

		switch sz {
		case 0:
			return new(AidEph) // 0 0 [0]
		case 1:
			return new(AidEph1) // 1 1 [0]
		case 8, 104:
			return new(AidEph2) // 8 104 [0]
		}

	case 0x020b:

		switch sz {
		case 0:
			return new(AidHui) // 0 0 [0]
		case 72:
			return new(AidHui1) // 72 72 [0]
		}

	case 0x010b:

		switch sz {
		case 0:
			return new(AidIni) // 0 0 [0]
		case 48:
			return new(AidIni1) // 48 48 [0]
		}

	case 0x1306:
		return new(CfgAnt)

	case 0x9306:
		return new(CfgBatch)

	case 0x0906:
		return new(CfgCfg)

	case 0x0606:

		switch sz {
		case 44:
			return new(CfgDat) // 44 44 [0]
		case 52:
			return new(CfgDat1) // 52 52 [0]
		}

	case 0x7006:
		return new(CfgDgnss)

	case 0x6106:
		return new(CfgDosc)

	case 0x4c06:
		return new(CfgEsfa)

	case 0x5606:
		return new(CfgEsfalg)

	case 0x4d06:
		return new(CfgEsfg)

	case 0x8206:
		return new(CfgEsfwt)

	case 0x6006:
		return new(CfgEsrc)

	case 0x6906:
		return new(CfgGeofence)

	case 0x3e06:
		return new(CfgGnss)

	case 0x5c06:
		return new(CfgHnr)

	case 0x0206:

		switch sz {
		case 1:
			return new(CfgInf1) // 1 1 [0]
		default:
			return new(CfgInf) // 0 [10]
		}

	case 0x3906:
		return new(CfgItfm)

	case 0x4706:
		return new(CfgLogfilter)

	case 0x0106:

		switch sz {
		case 2:
			return new(CfgMsg) // 2 2 [0]
		case 3:
			return new(CfgMsg1) // 3 3 [0]
		case 8:
			return new(CfgMsg2) // 8 8 [0]
		}

	case 0x2406:
		return new(CfgNav5)

	case 0x2306:

		switch sz {
		case 40:
			return new(CfgNavx5) // 40 40 [0]
		case 44:
			return new(CfgNavx51) // 44 44 [0]
		}

	case 0x1706:

		switch sz {
		case 4:
			return new(CfgNmea) // 4 4 [0]
		case 12:
			return new(CfgNmea1) // 12 12 [0]
		case 20:
			return new(CfgNmea2) // 20 20 [0]
		}

	case 0x1e06:
		return new(CfgOdo)

	case 0x3b06:

		switch sz {
		case 44:
			return new(CfgPm2) // 44 44 [0]
		case 48:
			return new(CfgPm21) // 48 48 [0]
		}

	case 0x8606:
		return new(CfgPms)

	case 0x0006:
		return mkCfgPrt(sz, frame)

	case 0x5706:
		return new(CfgPwr)

	case 0x0806:
		return new(CfgRate)

	case 0x3406:
		return new(CfgRinv)

	case 0x0406:
		return new(CfgRst)

	case 0x1106:
		return new(CfgRxm)

	case 0x1606:
		return new(CfgSbas)

	case 0x8806:
		return new(CfgSenif)

	case 0x8d06:
		return new(CfgSlas)

	case 0x6206:
		return new(CfgSmgr)

	case 0x6406:
		return new(CfgSpt)

	case 0x3d06:
		return new(CfgTmode2)

	case 0x7106:
		return new(CfgTmode3)

	case 0x3106:

		switch sz {
		case 0:
			return new(CfgTp5) // 0 0 [0]
		case 1:
			return new(CfgTp51) // 1 1 [0]
		case 32:
			return new(CfgTp52) // 32 32 [0]
		}

	case 0x5306:
		return new(CfgTxslot)

	case 0x1b06:
		return new(CfgUsb)

	case 0x1410:
		return new(EsfAlg)

	case 0x1510:
		return new(EsfIns)

	case 0x0210:
		return new(EsfMeas)

	case 0x0310:
		return new(EsfRaw)

	case 0x1010:
		return new(EsfStatus)

	case 0x0128:
		return new(HnrAtt)

	case 0x0228:
		return new(HnrIns)

	case 0x0028:
		return new(HnrPvt)

	case 0x0404:
		return new(InfDebug)

	case 0x0004:
		return new(InfError)

	case 0x0204:
		return new(InfNotice)

	case 0x0304:
		return new(InfTest)

	case 0x0104:
		return new(InfWarning)

	case 0x1121:
		return new(LogBatch)

	case 0x0721:
		return new(LogCreate)

	case 0x0321:
		return new(LogErase)

	case 0x0e21:

		switch sz {
		case 8:
			return new(LogFindtime) // 8 8 [0]
		case 12:
			return new(LogFindtime1) // 12 12 [0]
		}

	case 0x0821:

		switch sz {
		case 0:
			return new(LogInfo) // 0 0 [0]
		case 48:
			return new(LogInfo1) // 48 48 [0]
		}

	case 0x0921:
		return new(LogRetrieve)

	case 0x1021:
		return new(LogRetrievebatch)

	case 0x0b21:
		return new(LogRetrievepos)

	case 0x0f21:
		return new(LogRetrieveposextra)

	case 0x0d21:
		return new(LogRetrievestring)

	case 0x0421:
		return new(LogString)

	case 0x6013:
		return new(MgaAckData0)

	case 0x2013:
		return new(MgaAno)

	case 0x0313:

		switch sz {
		case 40:
			return new(MgaBdsAlm) // 40 40 [0]
		case 88:
			return new(MgaBdsEph1) // 88 88 [0]
		case 68:
			return new(MgaBdsHealth2) // 68 68 [0]
		case 16:
			return new(MgaBdsIono3) // 16 16 [0]
		case 20:
			return new(MgaBdsUtc4) // 20 20 [0]
		}

	case 0x8013:

		switch sz {
		case 0:
			return new(MgaDbd) // 0 0 [0]
		default:
			return new(MgaDbd1) // 12 [1]
		}

	case 0x2113:
		return mkMgaFlashAck(sz, frame)

	case 0x0213:

		switch sz {
		case 32:
			return new(MgaGalAlm) // 32 32 [0]
		case 76:
			return new(MgaGalEph1) // 76 76 [0]
		case 12:
			return new(MgaGalTimeoffset2) // 12 12 [0]
		case 20:
			return new(MgaGalUtc3) // 20 20 [0]
		}

	case 0x0613:

		switch sz {
		case 36:
			return new(MgaGloAlm) // 36 36 [0]
		case 48:
			return new(MgaGloEph1) // 48 48 [0]
		case 20:
			return new(MgaGloTimeoffset2) // 20 20 [0]
		}

	case 0x0013:

		switch sz {
		case 36:
			return new(MgaGpsAlm) // 36 36 [0]
		case 68:
			return new(MgaGpsEph1) // 68 68 [0]
		case 40:
			return new(MgaGpsHealth2) // 40 40 [0]
		case 16:
			return new(MgaGpsIono3) // 16 16 [0]
		case 20:
			return new(MgaGpsUtc4) // 20 20 [0]
		}

	case 0x4013:
		return mkMgaIniClkd(sz, frame)

	case 0x0513:

		switch sz {
		case 36:
			return new(MgaQzssAlm) // 36 36 [0]
		case 68:
			return new(MgaQzssEph1) // 68 68 [0]
		case 12:
			return new(MgaQzssHealth2) // 12 12 [0]
		}

	case 0x320a:
		return new(MonBatch)

	case 0x280a:
		return new(MonGnss)

	case 0x090a:
		return new(MonHw)

	case 0x0b0a:
		return new(MonHw2)

	case 0x020a:
		return new(MonIo)

	case 0x060a:
		return new(MonMsgpp)

	case 0x270a:

		switch sz {
		case 0:
			return new(MonPatch) // 0 0 [0]
		default:
			return new(MonPatch1) // 4 [16]
		}

	case 0x070a:
		return new(MonRxbuf)

	case 0x210a:
		return new(MonRxr)

	case 0x2e0a:
		return new(MonSmgr)

	case 0x2f0a:
		return new(MonSpt)

	case 0x080a:
		return new(MonTxbuf)

	case 0x040a:

		switch sz {
		case 0:
			return new(MonVer) // 0 0 [0]
		default:
			return new(MonVer1) // 40 [30]
		}

	case 0x6001:
		return new(NavAopstatus)

	case 0x0501:
		return new(NavAtt)

	case 0x2201:
		return new(NavClock)

	case 0x3601:
		return new(NavCov)

	case 0x3101:
		return new(NavDgps)

	case 0x0401:
		return new(NavDop)

	case 0x3d01:
		return new(NavEell)

	case 0x6101:
		return new(NavEoe)

	case 0x3901:
		return new(NavGeofence)

	case 0x1301:
		return new(NavHpposecef)

	case 0x1401:
		return new(NavHpposllh)

	case 0x2801:
		return new(NavNmi)

	case 0x0901:
		return new(NavOdo)

	case 0x3401:
		return new(NavOrb)

	case 0x0101:
		return new(NavPosecef)

	case 0x0201:
		return new(NavPosllh)

	case 0x0701:
		return new(NavPvt)

	case 0x3c01:
		return new(NavRelposned)

	case 0x1001:
		return new(NavResetodo)

	case 0x3501:
		return new(NavSat)

	case 0x3201:
		return new(NavSbas)

	case 0x4201:
		return new(NavSlas)

	case 0x0601:
		return new(NavSol)

	case 0x0301:
		return new(NavStatus)

	case 0x3b01:
		return new(NavSvin)

	case 0x3001:
		return new(NavSvinfo)

	case 0x2401:
		return new(NavTimebds)

	case 0x2501:
		return new(NavTimegal)

	case 0x2301:
		return new(NavTimeglo)

	case 0x2001:
		return new(NavTimegps)

	case 0x2601:
		return new(NavTimels)

	case 0x2101:
		return new(NavTimeutc)

	case 0x1101:
		return new(NavVelecef)

	case 0x1201:
		return new(NavVelned)

	case 0x6102:
		return new(RxmImes)

	case 0x1402:
		return new(RxmMeasx)

	case 0x4102:

		switch sz {
		case 8:
			return new(RxmPmreq) // 8 8 [0]
		case 16:
			return new(RxmPmreq1) // 16 16 [0]
		}

	case 0x1502:
		return new(RxmRawx)

	case 0x5902:

		switch sz {
		case 16:
			return new(RxmRlm) // 16 16 [0]
		case 28:
			return new(RxmRlm1) // 28 28 [0]
		}

	case 0x3202:
		return new(RxmRtcm)

	case 0x1302:
		return new(RxmSfrbx)

	case 0x2002:
		return new(RxmSvsi)

	case 0x0327:
		return new(SecUniqid)

	case 0x110d:
		return new(TimDosc)

	case 0x160d:
		return new(TimFchg)

	case 0x170d:
		return new(TimHoc)

	case 0x130d:
		return new(TimSmeas)

	case 0x040d:
		return new(TimSvin)

	case 0x030d:
		return new(TimTm2)

	case 0x120d:
		return new(TimTos)

	case 0x010d:
		return new(TimTp)

	case 0x150d:
		return mkTimVcocal(sz, frame)

	case 0x060d:
		return new(TimVrfy)

	case 0x1409:
		return mkUpdSos(sz, frame)

	}
	return nil
}

/** Please Provide implementations of the following functions to disambiguate


func mkCfgPrt(sz uint16, frame []byte) Message {
            return new(CfgPrt) // length: 1
            return new(CfgPrt1) // length: 20 subtype byte: default
            return new(CfgPrt2) // length: 20 subtype byte: 3
            return new(CfgPrt3) // length: 20 subtype byte: 4
            return new(CfgPrt4) // length: 20 subtype byte: 0
}
func mkMgaFlashAck(sz uint16, frame []byte) Message {
            return new(MgaFlashAck) // length: 6 subtype byte: 3
            return new(MgaFlashData1) // length: 6 [1] subtype byte: 1
            return new(MgaFlashStop2) // length: 2 subtype byte: 2
}
func mkMgaIniClkd(sz uint16, frame []byte) Message {
            return new(MgaIniClkd) // length: 12 subtype byte: 0x20
            return new(MgaIniEop1) // length: 72 subtype byte: 0x30
            return new(MgaIniFreq2) // length: 12 subtype byte: 0x21
            return new(MgaIniPos_llh3) // length: 20 subtype byte: 1
            return new(MgaIniPos_xyz4) // length: 20 subtype byte: 0
            return new(MgaIniTime_gnss5) // length: 24 subtype byte: 0x11
            return new(MgaIniTime_utc6) // length: 24 subtype byte: 0x10
}
func mkTimVcocal(sz uint16, frame []byte) Message {
            return new(TimVcocal) // length: 1 subtype byte: 0
            return new(TimVcocal1) // length: 12 subtype byte: 2
            return new(TimVcocal2) // length: 12 subtype byte: 3
}
func mkUpdSos(sz uint16, frame []byte) Message {
            return new(UpdSos) // length: 0
            return new(UpdSos1) // length: 4 subtype byte: 0
            return new(UpdSos2) // length: 4 subtype byte: 1
            return new(UpdSos3) // length: 8 subtype byte: 2
            return new(UpdSos4) // length: 8 subtype byte: 3
}

*/
