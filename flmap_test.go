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
			NR:   0,
			FRBA: 0x40,
			NC:   1,
			FCBA: 0x30,
		},
		Flmap1: FLMAP1{
			IFDEntry: IFDEntry{
				Raw: 0x58100208,
			},
			ISL:   0x58,
			FPSBA: 0x100,
			NM:    2,
			FMBA:  0x80,
		},
		Flmap2: FLMAP2{
			IFDEntry: IFDEntry{
				Raw: 0x00310330,
			},
			UNKNOWN1: 0,
			PSL:      0x3103,
			FMSBA:    0x300,
		},
		Flmap3: FLMAP3{
			IFDEntry: IFDEntry{
				Raw: 0xFFFFFFFF,
			},
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
	target.Flumap1.Raw = 0

	target.WriteMaps()

	assert.Equal(t, testIFDHeaderMapped, target)
}
