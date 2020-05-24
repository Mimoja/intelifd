package intelifd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testVTEntries = []VSCCEntry{
		{Jid: 0x1840ef, Vscc: 0x20252025},
		{Jid: 0x123456, Vscc: 0x42424242},
	}

	testVTBytes = []byte{
		0xef, 0x40, 0x18, 0x00, 0x25, 0x20, 0x25, 0x20,
		0x56, 0x34, 0x12, 0x00, 0x42, 0x42, 0x42, 0x42,
	}

	testFMBytes = []byte{
		0x00, 0x0f, 0xa0, 0x00,
		0x00, 0x0d, 0x40, 0x00,
		0x00, 0x09, 0x80, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x01, 0x01, 0x10,
	}

	testFM = FM{
		Raw: struct {
			Flmstr1 uint32
			Flmstr2 uint32
			Flmstr3 uint32
			Flmstr4 uint32
			Flmstr5 uint32
		}{
			Flmstr1: 0xa00f00,
			Flmstr2: 0x400d00,
			Flmstr3: 0x800900,
			Flmstr4: 0x000000,
			Flmstr5: 0x10010100,
		},
		Master: MasterSection{
			BIOS: MasterSectionEntry{
				FlashDescriptorReadAccess:     true,
				FlashDescriptorWriteAccess:    false,
				HostCPUBIOSRegionReadAccess:   true,
				HostCPUBIOSRegionWriteAccess:  true,
				IntelMERegionReadAccess:       true,
				IntelMERegionWriteAccess:      false,
				GbERegionReadAccess:           true,
				GbERegionWriteAccess:          true,
				PlatformDataRegionReadAccess:  false,
				PlatformDataRegionWriteAccess: false,
				ECRegionReadAccess:            false,
				ECRegionWriteAccess:           false,
				RequesterID:                   0xF00,
			},
			ME: MasterSectionEntry{
				FlashDescriptorReadAccess:     true,
				FlashDescriptorWriteAccess:    false,
				HostCPUBIOSRegionReadAccess:   false,
				HostCPUBIOSRegionWriteAccess:  false,
				IntelMERegionReadAccess:       true,
				IntelMERegionWriteAccess:      true,
				GbERegionReadAccess:           true,
				GbERegionWriteAccess:          false,
				PlatformDataRegionReadAccess:  false,
				PlatformDataRegionWriteAccess: false,
				ECRegionReadAccess:            false,
				ECRegionWriteAccess:           false,
				RequesterID:                   0xD00,
			},
			ETHERNET: MasterSectionEntry{
				FlashDescriptorReadAccess:     true,
				FlashDescriptorWriteAccess:    false,
				HostCPUBIOSRegionReadAccess:   false,
				HostCPUBIOSRegionWriteAccess:  false,
				IntelMERegionReadAccess:       false,
				IntelMERegionWriteAccess:      false,
				GbERegionReadAccess:           true,
				GbERegionWriteAccess:          true,
				PlatformDataRegionReadAccess:  false,
				PlatformDataRegionWriteAccess: false,
				ECRegionReadAccess:            false,
				ECRegionWriteAccess:           false,
				RequesterID:                   0x900,
			},
			RESERVED: MasterSectionEntry{
				FlashDescriptorReadAccess:     false,
				FlashDescriptorWriteAccess:    false,
				HostCPUBIOSRegionReadAccess:   false,
				HostCPUBIOSRegionWriteAccess:  false,
				IntelMERegionReadAccess:       false,
				IntelMERegionWriteAccess:      false,
				GbERegionReadAccess:           false,
				GbERegionWriteAccess:          false,
				PlatformDataRegionReadAccess:  false,
				PlatformDataRegionWriteAccess: false,
				ECRegionReadAccess:            false,
				ECRegionWriteAccess:           false,
				RequesterID:                   0,
			},
			EC: MasterSectionEntry{
				FlashDescriptorReadAccess:     true,
				FlashDescriptorWriteAccess:    false,
				HostCPUBIOSRegionReadAccess:   false,
				HostCPUBIOSRegionWriteAccess:  false,
				IntelMERegionReadAccess:       false,
				IntelMERegionWriteAccess:      false,
				GbERegionReadAccess:           false,
				GbERegionWriteAccess:          false,
				PlatformDataRegionReadAccess:  false,
				PlatformDataRegionWriteAccess: false,
				ECRegionReadAccess:            true,
				ECRegionWriteAccess:           true,
				RequesterID:                   0x100,
			},
		},
	}
)

func mockImage() []byte {
	firmware := make([]byte, testImage16MB)
	copy(firmware[:], testIFDBytes)
	copy(firmware[0x10+flumapPosition:], testFlumap1)
	copy(firmware[testIFDHeaderMapped.Flumap1.VTBA:], testVTBytes)
	copy(firmware[testIFDHeaderMapped.Flmap1.FMBA:], testFMBytes)

	return firmware
}

func TestParseVT(t *testing.T) {

	bs := mockImage()

	entries, err := parseVT(&testIFDHeaderMapped, bs)
	assert.Nil(t, err)
	assert.Equal(t, testVTEntries, entries)
}

func TestParseFM(t *testing.T) {

	bs := mockImage()

	fm, err := parseFM(&testIFDHeaderMapped, bs)
	assert.Nil(t, err)
	assert.Equal(t, testFM, *fm)
}
