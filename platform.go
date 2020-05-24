package intelifd

import "log"

const (
	CHIPSET_ICH_UNKNOWN = iota
	CHIPSET_ICH
	CHIPSET_ICH2345
	CHIPSET_ICH6
	CHIPSET_POULSBO      /* SCH U* */
	CHIPSET_TUNNEL_CREEK /* Atom E6xx */
	CHIPSET_CENTERTON    /* Atom S1220 S1240 S1260 */
	CHIPSET_ICH7
	CHIPSET_ICH8
	CHIPSET_ICH9
	CHIPSET_ICH10
	CHIPSET_5_SERIES_IBEX_PEAK
	CHIPSET_6_SERIES_COUGAR_POINT
	CHIPSET_7_SERIES_PANTHER_POINT
	CHIPSET_8_SERIES_LYNX_POINT
	CHIPSET_BAYTRAIL
	/* Actually all with Silvermont architecture:
	 * Bay Trail, Avoton/Rangeley
	 */
	CHIPSET_8_SERIES_LYNX_POINT_LP
	CHIPSET_8_SERIES_WELLSBURG
	CHIPSET_9_SERIES_WILDCAT_POINT
	CHIPSET_9_SERIES_WILDCAT_POINT_LP
	CHIPSET_100_SERIES_SUNRISE_POINT
	/* also 6th/7th gen Core i/o (LP)
	 * variants
	 */
	CHIPSET_C620_SERIES_LEWISBURG
)

const (
	PLATFORM_APL = iota
	PLATFORM_CNL
	PLATFORM_GLK
	PLATFORM_ICL
	PLATFORM_SKLKBL
)

/**
 * Stolen from flashrom
 */
func (ifd *IFDHeader) GuessChipset() int {

	if ifd.Flmap2.NumberOfProcessorStraps>>8 == 0x00 {
		if ifd.Flmap2.NumberOfProcessorStraps&0xFF == 0 && ifd.Flmap1.NumberOfPCHStraps <= 2 {
			return CHIPSET_ICH8
		} else if ifd.Flmap1.NumberOfPCHStraps <= 2 {
			return CHIPSET_ICH9
		} else if ifd.Flmap1.NumberOfPCHStraps <= 10 {
			return CHIPSET_ICH10
		} else if ifd.Flmap1.NumberOfPCHStraps <= 16 {
			return CHIPSET_5_SERIES_IBEX_PEAK
		}
		log.Printf("Peculiar firmware descriptor, assuming Ibex Peak compatibility.\n")
		return CHIPSET_5_SERIES_IBEX_PEAK
	} else if ifd.Flmap2.NumberOfProcessorStraps>>8 < 0x31 && (ifd.Flmap2.Raw&0xff) < 0x30 {
		if ifd.Flmap2.NumberOfProcessorStraps&0xFF == 0 && ifd.Flmap1.NumberOfPCHStraps <= 17 {
			return CHIPSET_BAYTRAIL
		} else if ifd.Flmap2.NumberOfProcessorStraps&0xFF <= 1 && ifd.Flmap1.NumberOfPCHStraps <= 18 {
			return CHIPSET_6_SERIES_COUGAR_POINT
		} else if ifd.Flmap2.NumberOfProcessorStraps&0xFF <= 1 && ifd.Flmap1.NumberOfPCHStraps <= 21 {
			return CHIPSET_8_SERIES_LYNX_POINT
		}
		log.Printf("Peculiar firmware descriptor, assuming Wildcat Point compatibility.\n")
		return CHIPSET_9_SERIES_WILDCAT_POINT
	} else if ifd.Flmap1.NumberOfMasters == 6 {
		return CHIPSET_C620_SERIES_LEWISBURG
	} else {
		return CHIPSET_100_SERIES_SUNRISE_POINT
	}

}
