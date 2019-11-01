package intelifd

type (
	IFDEntry struct {
		Raw uint32
	}

	FLMAP0 struct {
		IFDEntry
		NR   uint32
		FRBA uint32
		NC   uint32
		FCBA uint32
	}

	FLMAP1 struct {
		IFDEntry
		ISL   uint32
		FPSBA uint32
		NM    uint32
		FMBA  uint32
	}

	FLMAP2 struct {
		IFDEntry
		UNKNOWN1 uint32
		PSL      uint32
		FMSBA    uint32
	}

	FLMAP3 struct {
		IFDEntry
	}

	FLUMAP1 struct {
		IFDEntry
		VTL  uint32
		VTBA uint32
	}
)

func (ifd *IFDHeader) ParseMaps() {

	ifd.Flmap0.NR = ((ifd.Flmap0.Raw >> 24) & 0x07)
	ifd.Flmap0.FRBA = ((ifd.Flmap0.Raw >> 16) & 0xFF) << 4
	ifd.Flmap0.NC = (ifd.Flmap0.Raw>>8)&0x03 + 1
	ifd.Flmap0.FCBA = (ifd.Flmap0.Raw & 0xFF) << 4

	ifd.Flmap1.ISL = (ifd.Flmap1.Raw >> 24) & 0xFF
	ifd.Flmap1.FPSBA = ((ifd.Flmap1.Raw >> 16) & 0xFF) << 4
	ifd.Flmap1.NM = (ifd.Flmap1.Raw >> 8) & 0x03
	ifd.Flmap1.FMBA = (ifd.Flmap1.Raw & 0xFF) << 4

	ifd.Flmap2.UNKNOWN1 = (ifd.Flmap2.Raw >> 24) & 0xFF
	ifd.Flmap2.PSL = (ifd.Flmap2.Raw >> 8) & 0xFFFF
	ifd.Flmap2.FMSBA = (ifd.Flmap2.Raw & 0xFF) << 4

	ifd.Flumap1.VTL = (ifd.Flumap1.Raw >> 8) & 0xFF
	ifd.Flumap1.VTBA = (ifd.Flumap1.Raw & 0xFF) << 4

}

func (ifd *IFDHeader) WriteMaps() {
	ifd.Flmap0.Raw = ifd.Flmap0.NR<<24 + (ifd.Flmap0.FRBA>>4)<<16 + (((ifd.Flmap0.NC - 1) >> 4) << 8) + (ifd.Flmap0.FCBA >> 4)
	ifd.Flmap1.Raw = ifd.Flmap1.ISL<<24 + (ifd.Flmap1.FPSBA>>4)<<16 + ifd.Flmap1.NM<<8 + ifd.Flmap1.FMBA>>4
	ifd.Flmap2.Raw = ifd.Flmap2.UNKNOWN1<<24 + (ifd.Flmap2.PSL << 8) + ifd.Flmap2.FMSBA>>4

	ifd.Flumap1.Raw = ifd.Flumap1.VTL<<8 + ifd.Flumap1.VTBA>>4

}
