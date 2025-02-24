package main

import (
	"flag"
	"log"
	"os"

	"github.com/sfomuseum/go-bcbp"
	"github.com/sfomuseum/go-bcbp/aztec"
)

func main() {

	var path string
	var data string

	flag.StringVar(&path, "path", "barcode.png", "")
	flag.StringVar(&data, "data", "", "")

	flag.Parse()

	b, err := bcbp.Unmarshal(data)

	if err != nil {
		log.Fatalf("Failed to unmarshal data, %v", err)
	}

	wr, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Fatalf("Failed to open %s for writing, %v", path, err)
	}

	err = aztec.Marshal(b, wr)

	if err != nil {
		log.Fatalf("Failed to marshal data, %v", err)
	}

	err = wr.Close()

	if err != nil {
		log.Fatalf("Failed to close %s after writing, %v", path, err)
	}

}
