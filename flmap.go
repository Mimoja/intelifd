package intelifd

type (
	IFDEntry struct {
		Raw uint32
	}

	FLMAP0 struct {
		IFDEntry
		NumberOfRegions         uint32
		RegionBase              uint32
		NumberOfFlashComponents uint32
		ComponentBase           uint32
	}

	FLMAP1 struct {
		IFDEntry
		NumberOfPCHStraps uint32
		PCHStrapsBase     uint32
		NumberOfMasters   uint32
		MasterBase        uint32
	}

	FLMAP2 struct {
		IFDEntry
		NumberOfIccTableEntries uint32
		IccTableBase            uint32
		NumberOfProcessorStraps uint32
		ProcessorStrapsBase     uint32
	}

	FLMAP3 struct {
		IFDEntry
		NumberOfDMITableEntries uint32
		DMITableBase            uint32
		Reserved0               uint32
		Reserved1               uint32
	}

	FLUMAP1 struct {
		IFDEntry
		VTL  uint32
		VTBA uint32
	}
)

func (ifd *IFDHeader) ParseMaps() (*IFDHeader){

	ifd.Flmap0.ComponentBase = (ifd.Flmap0.Raw & 0xFF) << 4
	ifd.Flmap0.NumberOfFlashComponents = (ifd.Flmap0.Raw>>8) & 0xFF
	ifd.Flmap0.RegionBase = ((ifd.Flmap0.Raw >> 16) & 0xFF) << 4
	ifd.Flmap0.NumberOfRegions = ((ifd.Flmap0.Raw >> 24) & 0x07)

	ifd.Flmap1.MasterBase = (ifd.Flmap1.Raw & 0xFF) << 4
	ifd.Flmap1.NumberOfMasters = (ifd.Flmap1.Raw >> 8) & 0xFF
	ifd.Flmap1.PCHStrapsBase = ((ifd.Flmap1.Raw >> 16) & 0xFF) << 4
	ifd.Flmap1.NumberOfPCHStraps = (ifd.Flmap1.Raw >> 24) & 0xFF

	ifd.Flmap2.ProcessorStrapsBase = (ifd.Flmap2.Raw & 0xFF) << 4
	ifd.Flmap2.NumberOfProcessorStraps = (ifd.Flmap2.Raw >> 8) & 0xFF
	ifd.Flmap2.IccTableBase = ((ifd.Flmap2.Raw >> 16) & 0xFF) << 4
	ifd.Flmap2.NumberOfIccTableEntries = (ifd.Flmap2.Raw >> 24) & 0xFF

	ifd.Flmap3.DMITableBase = (ifd.Flmap3.Raw & 0xFF) << 4
	ifd.Flmap3.NumberOfDMITableEntries = (ifd.Flmap3.Raw >> 8) & 0xFF
	ifd.Flmap3.Reserved0 = (ifd.Flmap3.Raw >> 16) & 0xFF
	ifd.Flmap3.Reserved1 = (ifd.Flmap3.Raw >> 24) & 0xFF

	ifd.Flumap1.VTL = (ifd.Flumap1.Raw >> 8) & 0xFF
	ifd.Flumap1.VTBA = (ifd.Flumap1.Raw & 0xFF) << 4

	return ifd;
}

func (ifd *IFDHeader) WriteMaps() {
	ifd.Flmap0.Raw = ifd.Flmap0.NumberOfRegions << 24 + (ifd.Flmap0.RegionBase >> 4) << 16 + (((ifd.Flmap0.NumberOfFlashComponents) >> 4) << 8) + (ifd.Flmap0.ComponentBase >> 4)
	ifd.Flmap1.Raw = ifd.Flmap1.NumberOfPCHStraps << 24 + (ifd.Flmap1.PCHStrapsBase >> 4) << 16 + ifd.Flmap1.NumberOfMasters << 8 + ifd.Flmap1.MasterBase >> 4
	ifd.Flmap2.Raw = ifd.Flmap2.NumberOfIccTableEntries << 24 + (ifd.Flmap2.IccTableBase >> 4) << 16 + ifd.Flmap2.NumberOfProcessorStraps << 8 + ifd.Flmap2.ProcessorStrapsBase >> 4
	ifd.Flmap3.Raw = ifd.Flmap3.Reserved1 << 24 + ifd.Flmap3.Reserved0 << 16 + ifd.Flmap3.NumberOfDMITableEntries << 8 + (ifd.Flmap3.DMITableBase >> 4)
	ifd.Flumap1.Raw = ifd.Flumap1.VTL << 8 + ifd.Flumap1.VTBA >> 4

}
