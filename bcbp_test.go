package bcbp

import (
	"fmt"
	"testing"
)

func TestParseBCBP(t *testing.T) {

	tests := map[string]*BCBP{
		/*
			"M1DOE/JOHN            EXYZ123 MELSFOUA 61   047C012D0001 100": &BCBP{
				FormatCode:                 "M",
				PassengerName:              "DOE/JOHN",
				OperatingCarrierPNR:        "XYZ123",
				FromAirport:                "MEL",
				ToAirport:                  "SFO",
				OperatingCarrierDesignator: "UA",
				FlightNumber:               "61   ", // Note the trailing white space, TBD...
			},
		*/
		"M1DESMARAIS/LUC       EABC123 YULFRAAC 0834 226F001A0025 1": &BCBP{
			FormatCode:                 "M",
			PassengerName:              "DESMARAIS/LUC",
			OperatingCarrierPNR:        "ABC123",
			FromAirport:                "YUL",
			ToAirport:                  "FRA",
			OperatingCarrierDesignator: "AC",
			FlightNumber:               "834 ", // Note the trailing white space, TBD...
		},
		"M1EWING/SHAUN MR      E1A11A1 BNESYDQF 0551 107Y026J0037 0": &BCBP{
			FormatCode:                 "M",
			PassengerName:              "EWING/SHAUN MR",
			OperatingCarrierPNR:        "1A11A1",
			FromAirport:                "BNE",
			ToAirport:                  "SYD",
			OperatingCarrierDesignator: "QF",
			FlightNumber:               "551 ", // Note the trailing white space, TBD...
		},
	}

	for str_bcbp, expected := range tests {

		fmt.Println(len(str_bcbp), str_bcbp)
		b, err := Parse(str_bcbp)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str_bcbp, err)
		}

		if b.PassengerName != expected.PassengerName {
			t.Fatalf("Invalid passenger name. Expected '%s', got '%s'", expected.PassengerName, b.PassengerName)
		}

		if b.OperatingCarrierPNR != expected.OperatingCarrierPNR {
			t.Fatalf("Invalid PNR. Expected '%s', got '%s'", expected.OperatingCarrierPNR, b.OperatingCarrierPNR)
		}

		if b.FromAirport != expected.FromAirport {
			t.Fatalf("Invalid origin airport. Expected '%s', got '%s'", expected.FromAirport, b.FromAirport)
		}

		if b.ToAirport != expected.ToAirport {
			t.Fatalf("Invalid destination airport. Expected '%s', got '%s'", expected.ToAirport, b.ToAirport)
		}

		if b.OperatingCarrierDesignator != expected.OperatingCarrierDesignator {
			t.Fatalf("Invalid airline. Expected '%s', got '%s'", expected.OperatingCarrierDesignator, b.OperatingCarrierDesignator)
		}

		if b.FlightNumber != expected.FlightNumber {
			t.Fatalf("Invalid flight number. Expected '%s', got '%s'", expected.FlightNumber, b.FlightNumber)
		}

		if b.String() != str_bcbp {
			t.Fatalf("String value of parsed BCBP string ('%s') does not match: '%s'", str_bcbp, b.String())
		}
	}
}
