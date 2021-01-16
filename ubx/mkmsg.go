package ubx

// implementations of mkMsg for ambiguous types that need to look inside the frame
// TODO generate as part of messages.go, now that we have the Subtype annotation in the xml

func mkCfgPrt(sz uint16, frame []byte) Message {
	if len(frame) < 2 {
		return new(CfgPrt) // 1
	}
	switch frame[0] {
	default:
		return new(CfgPrt1) // 20 uart
	case 3:
		return new(CfgPrt2) // 20  usb
	case 4:
		return new(CfgPrt3) // 20  spi
	case 0:
		return new(CfgPrt4) // 20  i2c
	}
	return nil
}

func mkMgaFlashAck(sz uint16, frame []byte) Message {
	if len(frame) < 1 {
		return nil
	}
	switch frame[0] {
	case 3:
		return new(MgaFlashAck) // 6
	case 1:
		return new(MgaFlashData1) // 6 [1]
	case 2:
		return new(MgaFlashStop2) // 2
	}
	return nil
}

func mkMgaIniClkd(sz uint16, frame []byte) Message {
	if len(frame) < 1 {
		return nil
	}
	switch frame[0] {
	case 0x20:
		return new(MgaIniClkd) // 12
	case 0x30:
		return new(MgaIniEop1) // 72
	case 0x21:
		return new(MgaIniFreq2) // 12
	case 0x01:
		return new(MgaIniPos_llh3) // 20
	case 0x00:
		return new(MgaIniPos_xyz4) // 20
	case 0x11:
		return new(MgaIniTime_gnss5) // 24
	case 0x10:
		return new(MgaIniTime_utc6) // 24
	}
	return nil
}

func mkTimVcocal(sz uint16, frame []byte) Message {
	if len(frame) < 1 {
		return nil
	}
	switch frame[0] {
	case 0:
		return new(TimVcocal) // 1
	case 2:
		return new(TimVcocal1) // 12
	case 3:
		return new(TimVcocal2) // 12
	}
	return nil
}

func mkUpdSos(sz uint16, frame []byte) Message {
	if len(frame) < 1 {
		return new(UpdSos)
	}
	switch frame[0] {
	case 0:
		return new(UpdSos1) // 4
	case 1:
		return new(UpdSos2) // 4
	case 2:
		return new(UpdSos3) // 8
	case 3:
		return new(UpdSos4) // 8
	}
	return nil
}
