package bcbp

import (
	"fmt"
	"testing"
)

func TestParseBCBP(t *testing.T) {

	tests := map[string]*BCBP{
		"M1DESMARAIS/LUC       EABC123 YULFRAAC 0834 226F001A0025 1": &BCBP{},
		// "M1DESMARAIS/LUC       EABC123YULFRAAC 00834326J001A00251": &BCBP{},
		// "M1EWING/SHAUN MR      E1A11A1BNESYDQF 00551107Y026J00370": &BCBP{},
	}

	for str_bcbp, _ := range tests {

		fmt.Println(len(str_bcbp), str_bcbp)
		b, err := Parse(str_bcbp)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str_bcbp, err)
		}

		if b.String() != str_bcbp {
			t.Fatalf("String value of parsed BCBP string ('%s') does not match: '%s'", str_bcbp, b.String())
		}
	}
}
