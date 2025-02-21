package aztec

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"io"

	bc_aztec "github.com/boombuler/barcode/aztec"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/aztec"
	"github.com/sfomuseum/go-bcbp"
)

/*

> go run cmd/encode/main.go -data 'M1DOE/JOHN            EXYZ123 MELSFOUA 61   047C012D0001 1'
[asc][asc@SD-931-4][8:38:04] /Users/asc/sfomuseum/go-bcbp                                                                                                                                                 > open barcode.png
[asc][asc@SD-931-4][8:38:06] /Users/asc/sfomuseum/go-bcbp                                                                                                                                                 > go run cmd/decode/main.go ./barcode.png
2025/02/21 08:38:18 Failed to decode barcode, NotFoundException: NotFoundException: NotFoundException: nbCenterLayers = 3
exit status 1

*/

func Marshal(b *bcbp.BCBP, wr io.Writer) error {

	data := []byte(b.String())

	bc, err := bc_aztec.Encode(data, 33, 15)

	if err != nil {
		return fmt.Errorf("Failed to encode barcode, %w", err)
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

	return bcbp.Parse(rsp.GetText())
}
