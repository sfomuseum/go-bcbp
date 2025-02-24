package aztec

import (
	"context"
	"io"

	"github.com/sfomuseum/go-bcbp"
)

type AztecBarcode struct {
	bcbp.Barcode
}

func init() {
	err := bcbp.RegisterBarcode(context.Background(), "aztec", NewAztecBarcode)
	if err != nil {
		panic(err)
	}
}

func NewAztecBarcode(ctx context.Context, uri string) (bcbp.Barcode, error) {

	bc := &AztecBarcode{}
	return bc, nil
}

func (a *AztecBarcode) Decode(r io.Reader) (*bcbp.BCBP, error) {
	return Unmarshal(r)
}

func (a *AztecBarcode) Encode(b *bcbp.BCBP, wr io.Writer) error {
	return Marshal(b, wr)
}
