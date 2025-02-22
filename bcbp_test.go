package bcbp

import (
	"testing"
)

func TestUnmarshal(t *testing.T) {

	// https://github.com/KDE/kitinerary/blob/master/autotests/bcbpparsertest.cpp#L37

	tests := map[string]int{
		"M1DESMARAIS/LUC       EABC123 YULFRAAC 0834 326J001A0025 100": 1,
		"M1DESMARAIS/LUC       EAB12C3 YULFRAAC 0834 326J003A0027 167>5321WW1325BAC 0014123456002001412346700100141234789012A0141234567890 1AC AC 1234567890123    4PCYLX58Z^108ABCDEFGH":                                                                                                                    1,
		"M1GRANDMAIRE/MELANIE  EABC123 GVACDGAF 0123 339C002F0025 130>5002A0571234567890  AF AF 1234567890123456    Y^108ABCDEFGH":                                                                                                                                                                           1,
		"M2DESMARAIS/LUC       EAB12C3 YULFRAAC 0834 326J003A0027 167>5321WW1325BAC 0014123456002001412346700100141234789012A0141234567890 1AC AC 1234567890123    4PCYLX58Z" + string(GROUP_SEPARATOR) + "EDEF456 FRAGVALH 3664 327C012C0002 12E2A0140987654321 1AC AC 1234567890123    3PCNWQ^108ABCDEFGH": 2,
		"M2GRANDMAIRE/MELANIE  EABC123 GVACDGAF 0123 339C002F0025 130>5002A0571234567890  AF AF 1234567890123456" + string(GROUP_SEPARATOR) + "    YDEF456 CDGDTWNW 0049 339F001A0002 12C2A012098765432101                       2PC ^108ABCDEFGH":                                                           2,
	}

	for raw, expected_count := range tests {

		b, err := Unmarshal(raw)

		if err != nil {
			t.Fatalf("Failed to unmarshal BCBP (%s), %v", raw, err)
		}

		if len(b.Legs) != expected_count {
			t.Fatalf("Unexpected count for (%s). Expected %d but got %d", raw, expected_count, len(b.Legs))
		}
	}

}
