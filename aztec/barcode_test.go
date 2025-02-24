package aztec

import (
	"context"
	_ "fmt"
	"os"
	"testing"

	"github.com/sfomuseum/go-bcbp"
)

func TestBarcodeMarshal(t *testing.T) {

	tests := []string{
		"M1DOE/JOHN            EXYZ123 MELSFOUA 61   047C012D0001 100",
	}

	ctx := context.Background()

	bc, err := bcbp.NewBarcode(ctx, "aztec://")

	if err != nil {
		t.Fatalf("Failed to create new barcode, %v", err)
	}

	for _, data := range tests {

		bcbp_data, err := bcbp.Unmarshal(data)

		if err != nil {
			t.Fatalf("Failed to unmarshal data '%s', %v", data, err)
		}

		wr, err := os.CreateTemp("", "example.*.png")

		if err != nil {
			t.Fatalf("Failed to create temp file, %v", err)
		}

		defer os.Remove(wr.Name()) // clean up

		err = bc.Encode(bcbp_data, wr)

		if err != nil {
			t.Fatalf("Failed to marshal data, %v", err)
		}

		err = wr.Close()

		if err != nil {
			t.Fatalf("Failed to close %s after writing, %v", wr.Name(), err)
		}
	}

}

func TestBarcodeUnmarshal(t *testing.T) {

	tests := []string{
		"M1DOE/JOHN            EXYZ123 MELSFOUA 61   047C012D0001 100",
	}

	ctx := context.Background()

	bc, err := bcbp.NewBarcode(ctx, "aztec://")

	if err != nil {
		t.Fatalf("Failed to create new barcode, %v", err)
	}

	for _, data := range tests {

		bcbp_data, err := bcbp.Unmarshal(data)

		if err != nil {
			t.Fatalf("Failed to unmarshal data '%s', %v", data, err)
		}

		wr, err := os.CreateTemp("", "example.*.png")

		if err != nil {
			t.Fatalf("Failed to create temp file, %v", err)
		}

		defer os.Remove(wr.Name()) // clean up

		err = bc.Encode(bcbp_data, wr)

		if err != nil {
			t.Fatalf("Failed to marshal data, %v", err)
		}

		err = wr.Close()

		if err != nil {
			t.Fatalf("Failed to close %s after writing, %v", wr.Name(), err)
		}

		r, err := os.Open(wr.Name())

		if err != nil {
			t.Fatalf("Failed to open %s for writing, %v", wr.Name(), err)
		}

		defer r.Close()

		bcbp_data, err = bc.Decode(r)

		if err != nil {
			t.Fatalf("Failed to unmarshal %s, %v", wr.Name(), err)
		}

		if bcbp_data.String() != data {
			t.Fatalf("Unexpected barcode data '%s'. Expected '%s'", bcbp_data.String(), data)
		}
	}

}
