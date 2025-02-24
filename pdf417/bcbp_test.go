package pdf417

import (
	_ "fmt"
	"os"
	"testing"

	"github.com/sfomuseum/go-bcbp"
)

func TestMarshal(t *testing.T) {

	tests := []string{
		"M1DOE/JOHN            EXYZ123 MELSFOUA 61   047C012D0001 100",
	}

	for _, data := range tests {

		b, err := bcbp.Unmarshal(data)

		if err != nil {
			t.Fatalf("Failed to unmarshal data '%s', %v", data, err)
		}

		wr, err := os.CreateTemp("", "example.*.png")

		if err != nil {
			t.Fatalf("Failed to create temp file, %v", err)
		}

		defer os.Remove(wr.Name()) // clean up

		err = Marshal(b, wr)

		if err != nil {
			t.Fatalf("Failed to marshal data, %v", err)
		}

		err = wr.Close()

		if err != nil {
			t.Fatalf("Failed to close %s after writing, %v", wr.Name(), err)
		}
	}

}

func TestUnmarshal(t *testing.T) {

	t.Skip()

	tests := []string{
		"M1DOE/JOHN            EXYZ123 MELSFOUA 61   047C012D0001 100",
	}

	for _, data := range tests {

		b, err := bcbp.Unmarshal(data)

		if err != nil {
			t.Fatalf("Failed to unmarshal data '%s', %v", data, err)
		}

		wr, err := os.CreateTemp("", "example.*.png")

		if err != nil {
			t.Fatalf("Failed to create temp file, %v", err)
		}

		defer os.Remove(wr.Name()) // clean up

		err = Marshal(b, wr)

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

		b2, err := Unmarshal(r)

		if err != nil {
			t.Fatalf("Failed to unmarshal %s, %v", wr.Name(), err)
		}

		if b2.String() != data {
			t.Fatalf("Unexpected barcode data '%s'. Expected '%s'", b2.String(), data)
		}
	}

}
