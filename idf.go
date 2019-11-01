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
	}
	IFDHeader struct {
		Offset   uint32
		Flvalsig uint32
		Flmap0   uint32
		Flmap1   uint32
		Flmap2   uint32
		Flmap3   uint32
		Flumap1  uint32
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

func ParseIFD(firmwareBytes []byte) (*IFD, error) {
	return nil, nil
}

func ParseIFDHeader(firmwareBytes []byte, offset uint32) (*IFDHeader, error) {
	if len(firmwareBytes) < int(offset+flumapPosition+4) {
		return nil, fmt.Errorf("image to small")
	}

	return &IFDHeader{
		Offset:   offset,
		Flvalsig: binary.LittleEndian.Uint32(firmwareBytes[offset: ]),
		Flmap0:   binary.LittleEndian.Uint32(firmwareBytes[offset+4: ]),
		Flmap1:   binary.LittleEndian.Uint32(firmwareBytes[offset+8: ]),
		Flmap2:   binary.LittleEndian.Uint32(firmwareBytes[offset+12: ]),
		Flmap3:   binary.LittleEndian.Uint32(firmwareBytes[offset+16: ]),
		Flumap1:  binary.BigEndian.Uint32(firmwareBytes[offset+flumapPosition : offset+flumapPosition+4]),
	}, nil
}
