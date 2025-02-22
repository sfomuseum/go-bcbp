package main

import (
	"flag"
	"log"
	"os"

	"github.com/sfomuseum/go-bcbp"
	"github.com/sfomuseum/go-bcbp/aztec"
)

func main() {

	var data string

	flag.StringVar(&data, "data", "", "")

	flag.Parse()

	b, err := bcbp.Unmarshal(data)

	if err != nil {
		log.Fatal(err)
	}

	wr, err := os.OpenFile("barcode.png", os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}

	err = aztec.Marshal(b, wr)

	if err != nil {
		log.Fatal(err)
	}

	err = wr.Close()

	if err != nil {
		log.Fatal(err)
	}

}
