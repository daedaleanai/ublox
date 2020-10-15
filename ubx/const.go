package ubx

//go:generate stringer -output=strings.go -type=ClassID

// A ClassID encodes Class and Message ID constants as per p.146..152 into a single 16 bit unsigned value.
type ClassID uint16

const (
	// Ack/Nak Messages: Acknowledge or Reject messages to UBX-CFG input messages
	ACK_ACK ClassID = 0x0501
	ACK_NAK ClassID = 0x0500

	// AssistNow Aiding Messages: Ephemeris, Almanac, other A-GPS data input
	AID_ALM ClassID = 0x0b30
	AID_AOP ClassID = 0x0b33
	AID_EPH ClassID = 0x0b31
	AID_HUI ClassID = 0x0b02
	AID_INI ClassID = 0x0b01

	// Configuration Input Messages: Configure the receiver.
	CFG_ANT       ClassID = 0x0613
	CFG_BATCH     ClassID = 0x0693
	CFG_CFG       ClassID = 0x0609
	CFG_DAT       ClassID = 0x0606
	CFG_DGNSS     ClassID = 0x0670
	CFG_DOSC      ClassID = 0x0661
	CFG_ESRC      ClassID = 0x0660
	CFG_GEOFENCE  ClassID = 0x0669
	CFG_GNSS      ClassID = 0x063e
	CFG_HNR       ClassID = 0x065c
	CFG_INF       ClassID = 0x0602
	CFG_ITFM      ClassID = 0x0639
	CFG_LOGFILTER ClassID = 0x0647
	CFG_MSG       ClassID = 0x0601
	CFG_NAV5      ClassID = 0x0624
	CFG_NAVX5     ClassID = 0x0623
	CFG_NMEA      ClassID = 0x0617
	CFG_ODO       ClassID = 0x061e
	CFG_PM2       ClassID = 0x063b
	CFG_PMS       ClassID = 0x0686
	CFG_PRT       ClassID = 0x0600
	CFG_PWR       ClassID = 0x0657
	CFG_RATE      ClassID = 0x0608
	CFG_RINV      ClassID = 0x0634
	CFG_RST       ClassID = 0x0604
	CFG_RXM       ClassID = 0x0611
	CFG_SBAS      ClassID = 0x0616
	CFG_SLAS      ClassID = 0x068d
	CFG_SMGR      ClassID = 0x0662
	CFG_TMODE2    ClassID = 0x063d
	CFG_TMODE3    ClassID = 0x0671
	CFG_TP5       ClassID = 0x0631
	CFG_TXSLOT    ClassID = 0x0653
	CFG_USB       ClassID = 0x061b

	// External Sensor Fusion Messages: External Sensor Measurements and Status Information
	ESF_INS    ClassID = 0x1015
	ESF_MEAS   ClassID = 0x1002
	ESF_RAW    ClassID = 0x1003
	ESF_STATUS ClassID = 0x1010

	// High Rate Navigation Results Messages: High rate time, position, speed, heading
	HNR_INS ClassID = 0x2802
	HNR_PVT ClassID = 0x2800

	// Information Messages: Printf-Style Messages, with IDs such as Error, Warning, Notice
	INF_DEBUG   ClassID = 0x0404
	INF_ERROR   ClassID = 0x0400
	INF_NOTICE  ClassID = 0x0402
	INF_TEST    ClassID = 0x0403
	INF_WARNING ClassID = 0x0401

	// Logging Messages: Log creation, deletion, info and retrieval
	LOG_BATCH            ClassID = 0x2111
	LOG_CREATE           ClassID = 0x2107
	LOG_ERASE            ClassID = 0x2103
	LOG_FINDTIME         ClassID = 0x210e
	LOG_INFO             ClassID = 0x2108
	LOG_RETRIEVE         ClassID = 0x2109
	LOG_RETRIEVEBATCH    ClassID = 0x2110
	LOG_RETRIEVEPOSEXTRA ClassID = 0x210f
	LOG_RETRIEVEPOS      ClassID = 0x210b
	LOG_RETRIEVESTRING   ClassID = 0x210d
	LOG_STRING           ClassID = 0x2104

	// Multiple GNSS Assistance Messages: Assistance data for various GNSS
	MGA_ACK_DATA0      ClassID = 0x1360
	MGA_ANO            ClassID = 0x1320
	MGA_BDS_ALM        ClassID = 0x1303
	MGA_BDS_EPH        ClassID = 0x1303
	MGA_BDS_HEALTH     ClassID = 0x1303
	MGA_BDS_IONO       ClassID = 0x1303
	MGA_BDS_UTC        ClassID = 0x1303
	MGA_DBD            ClassID = 0x1380
	MGA_FLASH_ACK      ClassID = 0x1321
	MGA_FLASH_DATA     ClassID = 0x1321
	MGA_FLASH_STOP     ClassID = 0x1321
	MGA_GAL_ALM        ClassID = 0x1302
	MGA_GAL_EPH        ClassID = 0x1302
	MGA_GAL_TIMEOFFSET ClassID = 0x1302
	MGA_GAL_UTC        ClassID = 0x1302
	MGA_GLO_ALM        ClassID = 0x1306
	MGA_GLO_EPH        ClassID = 0x1306
	MGA_GLO_TIMEOFFSET ClassID = 0x1306
	MGA_GPS_ALM        ClassID = 0x1300
	MGA_GPS_EPH        ClassID = 0x1300
	MGA_GPS_HEALTH     ClassID = 0x1300
	MGA_GPS_IONO       ClassID = 0x1300
	MGA_GPS_UTC        ClassID = 0x1300
	MGA_INI_CLKD       ClassID = 0x1340
	MGA_INI_EOP        ClassID = 0x1340
	MGA_INI_FREQ       ClassID = 0x1340
	MGA_INI_POS_LLH    ClassID = 0x1340
	MGA_INI_POS_XYZ    ClassID = 0x1340
	MGA_INI_TIME_GNSS  ClassID = 0x1340
	MGA_INI_TIME_UTC   ClassID = 0x1340
	MGA_QZSS_ALM       ClassID = 0x1305
	MGA_QZSS_EPH       ClassID = 0x1305
	MGA_QZSS_HEALTH    ClassID = 0x1305

	// Monitoring Messages: Communication Status, CPU Load, Stack Usage, Task Status
	MON_BATCH ClassID = 0x0a32
	MON_GNSS  ClassID = 0x0a28
	MON_HW    ClassID = 0x0a09
	MON_HW2   ClassID = 0x0a0b
	MON_IO    ClassID = 0x0a02
	MON_MSGPP ClassID = 0x0a06
	MON_PATCH ClassID = 0x0a27
	MON_RXBUF ClassID = 0x0a07
	MON_RXR   ClassID = 0x0a21
	MON_SMGR  ClassID = 0x0a2e
	MON_TXBUF ClassID = 0x0a08
	MON_VER   ClassID = 0x0a04

	// Navigation Results Messages: Position, Speed, Time, Acceleration, Heading, DOP, SVs used
	NAV_AOPSTATUS ClassID = 0x0160
	NAV_ATT       ClassID = 0x0105
	NAV_CLOCK     ClassID = 0x0122
	NAV_DGPS      ClassID = 0x0131
	NAV_DOP       ClassID = 0x0104
	NAV_EOE       ClassID = 0x0161
	NAV_GEOFENCE  ClassID = 0x0139
	NAV_HPPOSECEF ClassID = 0x0113
	NAV_HPPOSLLH  ClassID = 0x0114
	NAV_NMI       ClassID = 0x0128
	NAV_ODO       ClassID = 0x0109
	NAV_ORB       ClassID = 0x0134
	NAV_POSECEF   ClassID = 0x0101
	NAV_POSLLH    ClassID = 0x0102
	NAV_PVT       ClassID = 0x0107
	NAV_RELPOSNED ClassID = 0x013c
	NAV_RESETODO  ClassID = 0x0110
	NAV_SAT       ClassID = 0x0135
	NAV_SBAS      ClassID = 0x0132
	NAV_SLAS      ClassID = 0x0142
	NAV_SOL       ClassID = 0x0106
	NAV_STATUS    ClassID = 0x0103
	NAV_SVIN      ClassID = 0x013b
	NAV_SVINFO    ClassID = 0x0130
	NAV_TIMEBDS   ClassID = 0x0124
	NAV_TIMEGAL   ClassID = 0x0125
	NAV_TIMEGLO   ClassID = 0x0123
	NAV_TIMEGPS   ClassID = 0x0120
	NAV_TIMELS    ClassID = 0x0126
	NAV_TIMEUTC   ClassID = 0x0121
	NAV_VELECEF   ClassID = 0x0111
	NAV_VELNED    ClassID = 0x0112

	// Receiver Manager Messages: Satellite Status, RTC Status
	RXM_IMES  ClassID = 0x0261
	RXM_MEASX ClassID = 0x0214
	RXM_PMREQ ClassID = 0x0241
	RXM_RAWX  ClassID = 0x0215
	RXM_RLM   ClassID = 0x0259
	RXM_RTCM  ClassID = 0x0232
	RXM_SFRBX ClassID = 0x0213
	RXM_SVSI  ClassID = 0x0220

	// Security Feature Messages
	SEC_UNIQID ClassID = 0x2703

	// Timing Messages: Time Pulse Output, Time Mark Results
	TIM_DOSC   ClassID = 0x0d11
	TIM_FCHG   ClassID = 0x0d16
	TIM_HOC    ClassID = 0x0d17
	TIM_SMEAS  ClassID = 0x0d13
	TIM_SVIN   ClassID = 0x0d04
	TIM_TM2    ClassID = 0x0d03
	TIM_TOS    ClassID = 0x0d12
	TIM_TP     ClassID = 0x0d01
	TIM_VCOCAL ClassID = 0x0d15
	TIM_VRFY   ClassID = 0x0d06

	// Firmware Update Messages: Memory/Flash erase/write, Reboot, Flash identification, etc.
	UPD_SOS ClassID = 0x0914
)

func (c ClassID) Class() uint8 { return uint8(c >> 8) }
func (c ClassID) ID() uint8    { return uint8(c) }
