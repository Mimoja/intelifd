package intelifd

import (
	"bytes"
	"encoding/binary"
)

func parseVT(header *IFDHeader, bs []byte) ([]VSCCEntry, error){
	vsccs := []VSCCEntry{}

	reader := bytes.NewReader(bs[header.Flumap1.VTBA:])
	for i := uint32(0); i < header.Flumap1.VTL; i++ {
		var Vscc VSCCEntry
		err := binary.Read(reader, binary.LittleEndian, &Vscc)
		if err != nil {
			return vsccs, err
		}
		vsccs = append(vsccs, Vscc)
	}

	return vsccs, nil;
}

func isBitSet(val uint32, bit uint32) bool {
	return (val & (1 << bit)) != 0
}


func parseFLMSTR(flmstr uint32, ifdVersion uint32) MasterSectionEntry {
	var wr_shift uint32
	var rd_shift uint32

	FLMSTR_WR_SHIFT_V1 := uint32(24)
	FLMSTR_RD_SHIFT_V1 := uint32(16)

	FLMSTR_WR_SHIFT_V2 := uint32(20)
	FLMSTR_RD_SHIFT_V2 := uint32(8)

	if ifdVersion == 1 {
		wr_shift = FLMSTR_WR_SHIFT_V1
		rd_shift = FLMSTR_RD_SHIFT_V1
	} else {
		wr_shift = FLMSTR_WR_SHIFT_V2
		rd_shift = FLMSTR_RD_SHIFT_V2
	}

	entry := MasterSectionEntry{
		FlashDescriptorReadAccess:     isBitSet(flmstr, rd_shift+0),
		FlashDescriptorWriteAccess:    isBitSet(flmstr, wr_shift+0),
		HostCPUBIOSRegionReadAccess:   isBitSet(flmstr, rd_shift+1),
		HostCPUBIOSRegionWriteAccess:  isBitSet(flmstr, wr_shift+1),
		IntelMERegionReadAccess:       isBitSet(flmstr, rd_shift+2),
		IntelMERegionWriteAccess:      isBitSet(flmstr, wr_shift+2),
		GbERegionReadAccess:           isBitSet(flmstr, rd_shift+3),
		GbERegionWriteAccess:          isBitSet(flmstr, wr_shift+3),
		PlatformDataRegionReadAccess:  isBitSet(flmstr, rd_shift+4),
		PlatformDataRegionWriteAccess: isBitSet(flmstr, wr_shift+4),
		ECRegionReadAccess:            isBitSet(flmstr, rd_shift+8),
		ECRegionWriteAccess:           isBitSet(flmstr, wr_shift+8),
		RequesterID:                   uint16(flmstr&0xFFFF),
	}
	return entry
}

func parseFM(header *IFDHeader, bs []byte) (*FM, error){
	fm := FM{}
	reader := bytes.NewReader(bs[header.Flmap1.FMBA:])
	err := binary.Read(reader, binary.LittleEndian, &fm.Raw)

	if err != nil {
		return nil, err
	}

	fm.Master = MasterSection{
		BIOS:     parseFLMSTR(fm.Raw.Flmstr1, header.Version),
		ME:       parseFLMSTR(fm.Raw.Flmstr2, header.Version),
		ETHERNET: parseFLMSTR(fm.Raw.Flmstr3, header.Version),
		RESERVED: parseFLMSTR(fm.Raw.Flmstr4, header.Version),
		EC:       parseFLMSTR(fm.Raw.Flmstr5, header.Version),
	}

	return &fm, nil
}
