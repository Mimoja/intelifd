package intelifd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testIFDHeaderMapped = IFDHeader{
		Offset:   0x10,
		Flvalsig: 0xff0a55a,
		Flmap0: FLMAP0{
			IFDEntry: IFDEntry{
				Raw: 0x00040003,
			},
			NumberOfRegions:         0,
			RegionBase:              0x40,
			NumberOfFlashComponents: 0,
			ComponentBase:           0x30,
		},
		Flmap1: FLMAP1{
			IFDEntry: IFDEntry{
				Raw: 0x58100208,
			},
			NumberOfPCHStraps: 0x58,
			PCHStrapsBase:     0x100,
			NumberOfMasters:   2,
			MasterBase:        0x80,
		},
		Flmap2: FLMAP2{
			IFDEntry: IFDEntry{
				Raw: 0x00310330,
			},
			IccTableBase:            0x310,
			NumberOfIccTableEntries: 0x00,
			NumberOfProcessorStraps: 0x3,
			ProcessorStrapsBase:     0x300,
		},
		Flmap3: FLMAP3{
			IFDEntry: IFDEntry{
				Raw: 0xFFFFFFFF,
			},
			NumberOfDMITableEntries: 0xFF,
			DMITableBase:            0xFF0,
			Reserved0:               0xFF,
			Reserved1:               0xFF,
		},
		Flumap1: FLUMAP1{
			IFDEntry: IFDEntry{
				Raw: 0x000002df,
			},
			VTL:  2,
			VTBA: 0xdf0,
		},
	}
)

func TestIFDHeader_ParseMaps(t *testing.T) {
	target := testIFDHeader

	target.ParseMaps()

	assert.Equal(t, testIFDHeaderMapped, target)
}

func TestIFDHeader_WriteMaps(t *testing.T) {
	target := testIFDHeaderMapped

	target.Flmap0.Raw = 0
	target.Flmap1.Raw = 0
	target.Flmap2.Raw = 0
	target.Flmap3.Raw = 0
	target.Flumap1.Raw = 0

	target.WriteMaps()

	assert.Equal(t, testIFDHeaderMapped, target)
}
