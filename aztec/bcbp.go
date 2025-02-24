package aztec

// ISO/IEC 24778:2008

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"io"

	"github.com/boombuler/barcode"
	bc_aztec "github.com/boombuler/barcode/aztec"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/aztec"
	"github.com/sfomuseum/go-bcbp"
)

func Marshal(b *bcbp.BCBP, wr io.Writer) error {

	data := []byte(b.String())

	bc, err := bc_aztec.Encode(data, 33, 0)

	if err != nil {
		return fmt.Errorf("Failed to encode barcode, %w", err)
	}

	bc, err = barcode.Scale(bc, 300, 300)

	if err != nil {
		return fmt.Errorf("Failed to scale barcode, %w", err)
	}

	return png.Encode(wr, bc)
}

func Unmarshal(r io.Reader) (*bcbp.BCBP, error) {

	im, _, err := image.Decode(r)

	if err != nil {
		return nil, fmt.Errorf("Failed to decode image, %w", err)
	}

	bm, err := gozxing.NewBinaryBitmapFromImage(im)

	if err != nil {
		return nil, fmt.Errorf("Failed to create bitmap from image, %w", err)
	}

	az_r := aztec.NewAztecReader()

	rsp, err := az_r.DecodeWithoutHints(bm)

	if err != nil {
		return nil, fmt.Errorf("Failed to decode barcode, %w", err)
	}

	return bcbp.Unmarshal(rsp.GetText())
}
