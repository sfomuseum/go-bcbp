package aztec

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/aztec"
	"github.com/sfomuseum/go-bcbp"
)

func Parse(r io.Reader) (*bcbp.BCBP, error) {

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

	return bcbp.Parse(rsp.GetText())
}
