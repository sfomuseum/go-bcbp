package bcbp

import (
	_ "fmt"
	"log/slog"
	"strings"
)

/*

Mandatory Fields (Fixed Length, 60 Characters)
These fields must be present in every BCBP string.

Field No.	Field Name	Offset Position	Length (Chars)	Notes
1	Format Code	0	1	Usually "M"
2	Number of Legs Encoded	1	1	"1" = Single-leg, "2" = Multi-leg
3	Passenger Name	2	20	Left-justified, trailing spaces
4	Electronic Ticket Indicator	22	1	"E" for e-ticket, " " for paper
5	Operating Carrier PNR Code	23	7	Left-justified, trailing spaces
6	From Airport Code	30	3	IATA 3-letter airport code
7	To Airport Code	33	3	IATA 3-letter airport code
8	Operating Carrier Designator	36	3	IATA airline code, left-justified
9	Flight Number	39	5	Right-justified, leading zeros
10	Date of Flight (Julian Date)	44	3	Format: DDD (001-366)
11	Compartment Code	47	1	Cabin class (e.g., Y = Economy)
12	Seat Number	48	4	Right-justified, leading zeros
13	Check-in Sequence Number	52	5	Right-justified, leading zeros
14	Passenger Status	57	1	"0" = Not checked in, "1" = Checked in

*/

// M|1|DESMARAIS/LUC       |E|ABC123 |YUL|FRA|AC |0834 |226|F|001A|0025 |1

const FORMAT_CODE int = 1
const NUMBER_OF_LEGS int = 1
const PASSENGER_NAME int = 20
const ELECTRONIC_TICKET_INDICATOR int = 1
const OPERATING_CARRIER_PNR int = 7
const DEPARTURE_AIRPORT int = 3
const ARRIVAL_AIRPORT int = 3
const OPERATING_CARRIER_DESIGNATOR int = 3
const FLIGHT_NUMBER int = 5
const FLIGHT_DATE int = 3
const COMPARTMENT_CODE int = 1
const SEAT_NUMBER int = 4
const CHECK_IN_SEQUENCE_NUMBER int = 5
const PASSENGER_STATUS int = 1

type BCBP struct {
	FormatCode                 string `json:"format_code"`
	NumberOfLegs               string `json:"number_of_legs"`
	PassengerName              string `json:"passenger_name"`
	ElectronicTicketIndicator  string `json:"electronic_ticket_indicator"`
	OperatingCarrierPNR        string `json:"operating_carrier_pnr"`
	FromAirport                string `json:"from_airport"`
	ToAirport                  string `json:"to_airport"`
	OperatingCarrierDesignator string `json:"operating_carrier_designator"`
	FlightNumber               string `json:"flight_number"`
	DateOfFlight               string `json:"date_of_flight"`
	CompartmentCode            string `json:"compartment_code"`
	SeatNumber                 string `json:"seat_number"`
	CheckInSequenceNumber      string `json:"checkin_sequence_number"`
	PassengerStatus            string `json:"passenger_status"`
}

func (b *BCBP) String() string {

	parts := []string{
		b.FormatCode,
		b.NumberOfLegs,
		rightPad(b.PassengerName, " ", PASSENGER_NAME),
		b.ElectronicTicketIndicator,
		rightPad(b.OperatingCarrierPNR, " ", OPERATING_CARRIER_PNR),
		b.FromAirport,
		b.ToAirport,
		rightPad(b.OperatingCarrierDesignator, " ", OPERATING_CARRIER_DESIGNATOR),
		leftPad(b.FlightNumber, "0", FLIGHT_NUMBER),
		b.DateOfFlight,
		b.CompartmentCode,
		leftPad(b.SeatNumber, "0", SEAT_NUMBER),
		leftPad(b.CheckInSequenceNumber, "0", CHECK_IN_SEQUENCE_NUMBER),
		b.PassengerStatus,
	}

	return strings.Join(parts, "")
}

// M1EWING/SHAUN MR      E1A11A1 BNESYDQF   551107Y 26J   370
// M1EWING/SHAUN MR      E1A11A1 BNESYDQF 551  107Y26J 37   0

func Parse(bcbp string) (*BCBP, error) {

	bcbpData := &BCBP{
		FormatCode:                 getField(bcbp, 0, FORMAT_CODE),
		NumberOfLegs:               getField(bcbp, 1, NUMBER_OF_LEGS),
		PassengerName:              strings.TrimSpace(getField(bcbp, 2, PASSENGER_NAME)),
		ElectronicTicketIndicator:  getField(bcbp, 22, ELECTRONIC_TICKET_INDICATOR),
		OperatingCarrierPNR:        strings.TrimSpace(getField(bcbp, 23, OPERATING_CARRIER_PNR)),
		FromAirport:                strings.TrimSpace(getField(bcbp, 30, DEPARTURE_AIRPORT)),
		ToAirport:                  strings.TrimSpace(getField(bcbp, 33, ARRIVAL_AIRPORT)),
		OperatingCarrierDesignator: strings.TrimSpace(getField(bcbp, 36, OPERATING_CARRIER_DESIGNATOR)),
		FlightNumber:               strings.TrimLeft(getField(bcbp, 39, FLIGHT_NUMBER), "0"),
		DateOfFlight:               getField(bcbp, 44, FLIGHT_DATE),
		CompartmentCode:            getField(bcbp, 47, COMPARTMENT_CODE),
		SeatNumber:                 strings.TrimLeft(getField(bcbp, 48, SEAT_NUMBER), "0"),
		CheckInSequenceNumber:      strings.TrimLeft(getField(bcbp, 52, CHECK_IN_SEQUENCE_NUMBER), "0"),
		PassengerStatus:            getField(bcbp, 57, PASSENGER_STATUS),
	}

	return bcbpData, nil
}

func leftPad(raw string, pad string, length int) string {

	for len(raw) < length {
		raw = pad + raw
	}

	return raw
}

func rightPad(raw string, pad string, length int) string {

	for len(raw) < length {
		raw = raw + pad
	}

	return raw
}

func getField(raw string, offset int, length int) string {
	v := raw[offset : offset+length]
	slog.Info("field", "raw", raw, "offset", offset, "length", length, "v", v)
	return v
}
