package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/sfomuseum/go-bcbp/aztec"
)

/*

> go run cmd/decode/main.go ./fixtures/barcodes/20250216-UAL61-aztec.png | jq
{
  "format_code": "M",
  "number_of_legs": "1",
  "passenger_name": "DOE/JOHN",
  "electronic_ticket_indicator": "E",
  "operating_carrier_pnr": "XYZ123",
  "from_airport": "MEL",
  "to_airport": "SFO",
  "operating_carrier_designator": "UA",
  "flight_number": "61   ",
  "date_of_flight": "047",
  "compartment_code": "C",
  "seat_number": "12D",
  "checkin_sequence_number": "1 ",
  "passenger_status": "1"
}

*/

func main() {

	flag.Parse()

	for _, path := range flag.Args() {

		r, err := os.Open(path)

		if err != nil {
			log.Fatal(err)
		}

		defer r.Close()

		b, err := aztec.Unmarshal(r)

		if err != nil {
			log.Fatal(err)
		}

		enc := json.NewEncoder(os.Stdout)
		err = enc.Encode(b)

		if err != nil {
			log.Fatal(err)
		}
	}

}
