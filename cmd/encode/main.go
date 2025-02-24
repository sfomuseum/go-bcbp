package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/sfomuseum/go-bcbp"
	_ "github.com/sfomuseum/go-bcbp/aztec"
	_ "github.com/sfomuseum/go-bcbp/pdf417"
)

func main() {

	var barcode_uri string
	var path string
	var data string

	flag.StringVar(&barcode_uri, "barcode-uri", "aztec://", "...")
	flag.StringVar(&path, "path", "barcode.png", "...")
	flag.StringVar(&data, "data", "", "...")

	flag.Parse()

	ctx := context.Background()

	bc, err := bcbp.NewBarcode(ctx, barcode_uri)

	if err != nil {
		log.Fatalf("Failed to create barcode, %v", err)
	}

	b, err := bcbp.Unmarshal(data)

	if err != nil {
		log.Fatalf("Failed to unmarshal data, %v", err)
	}

	wr, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Fatalf("Failed to open %s for writing, %v", path, err)
	}

	err = bc.Encode(b, wr)

	if err != nil {
		log.Fatalf("Failed to marshal data, %v", err)
	}

	err = wr.Close()

	if err != nil {
		log.Fatalf("Failed to close %s after writing, %v", path, err)
	}

}
