package pdf417

// https://en.wikipedia.org/wiki/PDF417
// https://github.com/sparkfish/pdf417decoder

import (
	"fmt"
	_ "image"
	_ "image/jpeg"
	"image/png"
	"io"

	"github.com/boombuler/barcode"
	bc_pdf417 "github.com/boombuler/barcode/pdf417"
	"github.com/sfomuseum/go-bcbp"
)

func Marshal(b *bcbp.BCBP, wr io.Writer) error {

	data := b.String()

	bc, err := bc_pdf417.Encode(data, 4)

	if err != nil {
		return fmt.Errorf("Failed to encode barcode, %w", err)
	}

	bc, err = barcode.Scale(bc, 300, 100)

	if err != nil {
		return fmt.Errorf("Failed to scale barcode, %w", err)
	}

	return png.Encode(wr, bc)
}

func Unmarshal(r io.Reader) (*bcbp.BCBP, error) {

	return nil, fmt.Errorf("Not implemented")

	/*
		im, _, err := image.Decode(r)

		if err != nil {
			return nil, err
		}
	*/
}
