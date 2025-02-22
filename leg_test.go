package bcbp

import (
	"fmt"
	"testing"
)

func TestParseLeg(t *testing.T) {

	tests := map[string]*Leg{
		"M1DOE/JOHN            EXYZ123 MELSFOUA 61   047C012D0001 100": &Leg{
			FormatCode:                 "M",
			PassengerName:              "DOE/JOHN",
			OperatingCarrierPNR:        "XYZ123",
			FromAirport:                "MEL",
			ToAirport:                  "SFO",
			OperatingCarrierDesignator: "UA",
			FlightNumber:               "61   ", // Note the trailing white space, TBD...
		},
		"M1DESMARAIS/LUC       EABC123 YULFRAAC 0834 226F001A0025 100": &Leg{
			FormatCode:                 "M",
			PassengerName:              "DESMARAIS/LUC",
			OperatingCarrierPNR:        "ABC123",
			FromAirport:                "YUL",
			ToAirport:                  "FRA",
			OperatingCarrierDesignator: "AC",
			FlightNumber:               "834 ", // Note the trailing white space, TBD...
		},
		"M1EWING/SHAUN MR      E1A11A1 BNESYDQF 0551 107Y026J0037 000": &Leg{
			FormatCode:                 "M",
			PassengerName:              "EWING/SHAUN MR",
			OperatingCarrierPNR:        "1A11A1",
			FromAirport:                "BNE",
			ToAirport:                  "SYD",
			OperatingCarrierDesignator: "QF",
			FlightNumber:               "551 ", // Note the trailing white space, TBD...
		},
	}

	for str_leg, expected := range tests {

		fmt.Println(len(str_leg), str_leg)

		leg, err := ParseLeg(str_leg)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str_leg, err)
		}

		if leg.PassengerName != expected.PassengerName {
			t.Fatalf("Invalid passenger name. Expected '%s', got '%s'", expected.PassengerName, leg.PassengerName)
		}

		if leg.OperatingCarrierPNR != expected.OperatingCarrierPNR {
			t.Fatalf("Invalid PNR. Expected '%s', got '%s'", expected.OperatingCarrierPNR, leg.OperatingCarrierPNR)
		}

		if leg.FromAirport != expected.FromAirport {
			t.Fatalf("Invalid origin airport. Expected '%s', got '%s'", expected.FromAirport, leg.FromAirport)
		}

		if leg.ToAirport != expected.ToAirport {
			t.Fatalf("Invalid destination airport. Expected '%s', got '%s'", expected.ToAirport, leg.ToAirport)
		}

		if leg.OperatingCarrierDesignator != expected.OperatingCarrierDesignator {
			t.Fatalf("Invalid airline. Expected '%s', got '%s'", expected.OperatingCarrierDesignator, leg.OperatingCarrierDesignator)
		}

		if leg.FlightNumber != expected.FlightNumber {
			t.Fatalf("Invalid flight number. Expected '%s', got '%s'", expected.FlightNumber, leg.FlightNumber)
		}

		if leg.String() != str_leg {
			t.Fatalf("String value of parsed Leg string ('%s') does not match: '%s'", str_leg, leg.String())
		}
	}
}
