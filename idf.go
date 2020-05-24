package intelifd

import (
	"encoding/binary"
	"fmt"
)

const IFD_HEADER = 0x5aa5f00f
const flumapPosition = 4096 - 256 - 4

type (
	IFD struct {
		Header IFDHeader
		FM     *FM
		VSCCs  []VSCCEntry
	}
	IFDHeader struct {
		Offset   uint32
		Version uint32
		Flvalsig uint32
		Flmap0   FLMAP0
		Flmap1   FLMAP1
		Flmap2   FLMAP2
		Flmap3   FLMAP3
		Flumap1  FLUMAP1
	}
	/*
	// FLMAP0
		ComponentBase      uint8
		NumberOfFlashChips uint8
		RegionBase         uint8
		NumberOfRegions    uint8
		// FLMAP1
		MasterBase        uint8
		NumberOfMasters   uint8
		PCHStrapsBase     uint8
		NumberOfPchStraps uint8
		// FLMAP2
		ProcStrapsBase          uint8
		NumberOfProcStraps      uint8
		IccTableBase            uint8
		NumberOfIccTableEntries uint8
		// FLMAP3
		DmiTableBase            uint8
		NumberOfDmiTableEntries uint8
		Reserved0               uint8
		Reserved1               uint8
	 */
	VSCCEntry struct {
		Jid  uint32
		Vscc uint32
	}
	FM struct {
		Raw struct {
			Flmstr1 uint32
			Flmstr2 uint32
			Flmstr3 uint32
			Flmstr4 uint32
			Flmstr5 uint32
		}
		Master MasterSection
	}

	MasterSection struct {
		BIOS     MasterSectionEntry
		ME       MasterSectionEntry
		ETHERNET MasterSectionEntry
		RESERVED MasterSectionEntry
		EC       MasterSectionEntry
	}

	MasterSectionEntry struct {
		FlashDescriptorReadAccess     bool
		FlashDescriptorWriteAccess    bool
		HostCPUBIOSRegionReadAccess   bool
		HostCPUBIOSRegionWriteAccess  bool
		IntelMERegionReadAccess       bool
		IntelMERegionWriteAccess      bool
		GbERegionReadAccess           bool
		GbERegionWriteAccess          bool
		PlatformDataRegionReadAccess  bool
		PlatformDataRegionWriteAccess bool
		ECRegionReadAccess            bool
		ECRegionWriteAccess           bool
		RequesterID uint16
	}
)

func FindFD(firmwareBytes []byte) (uint32, error) {

	for i := 0; i < len(firmwareBytes)-4; i += 4 {
		if binary.BigEndian.Uint32(firmwareBytes[i:i+4]) == IFD_HEADER {
			return uint32(i), nil
		}
	}

	return 0, fmt.Errorf("could not find IFD Header 0x%X", IFD_HEADER)
}

func ParseIFD(firmwareBytes []byte, offset uint32) (*IFD, error) {
	ifd := IFD{}
	header, e := ParseIFDHeader(firmwareBytes, offset)
	if e != nil {
		return nil, e
	}
	ifd.Header = *header.ParseMaps()
	//FIXME guess Header Version correctly!
	ifd.Header.Version = 1;

	ifd.FM, _ = parseFM(header, firmwareBytes)

	ifd.VSCCs, _ = parseVT(header, firmwareBytes)

	return &ifd, nil
}

func ParseIFDHeader(firmwareBytes []byte, offset uint32) (*IFDHeader, error) {
	if len(firmwareBytes) < int(offset+flumapPosition+4) {
		return nil, fmt.Errorf("image to small")
	}

	ifd := IFDHeader{
		Offset:   offset,
		Flvalsig: binary.LittleEndian.Uint32(firmwareBytes[offset:]),
	}
	ifd.Flmap0.Raw = binary.LittleEndian.Uint32(firmwareBytes[offset+4:])
	ifd.Flmap1.Raw = binary.LittleEndian.Uint32(firmwareBytes[offset+8:])
	ifd.Flmap2.Raw = binary.LittleEndian.Uint32(firmwareBytes[offset+12:])
	ifd.Flmap3.Raw = binary.LittleEndian.Uint32(firmwareBytes[offset+16:])
	ifd.Flumap1.Raw = binary.BigEndian.Uint32(firmwareBytes[offset+flumapPosition : offset+flumapPosition+4])

	return &ifd, nil
}
