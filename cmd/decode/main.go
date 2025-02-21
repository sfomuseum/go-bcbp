package main

import (
	"encoding/json"
	"flag"
	"fmt"
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

	var format string

	flag.StringVar(&format, "format", "string", "")

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

		switch format {
		case "json":

			enc := json.NewEncoder(os.Stdout)
			err = enc.Encode(b)

			if err != nil {
				log.Fatal(err)
			}

		default:
			fmt.Println(b.String())
		}
	}

}
