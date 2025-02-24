package pdf417

import (
	"context"
	"io"

	"github.com/sfomuseum/go-bcbp"
)

type PDF417Barcode struct {
	bcbp.Barcode
}

func init() {
	err := bcbp.RegisterBarcode(context.Background(), "pdf417", NewPDF417Barcode)
	if err != nil {
		panic(err)
	}
}

func NewPDF417Barcode(ctx context.Context, uri string) (bcbp.Barcode, error) {

	bc := &PDF417Barcode{}
	return bc, nil
}

func (a *PDF417Barcode) Decode(r io.Reader) (*bcbp.BCBP, error) {
	return Unmarshal(r)
}

func (a *PDF417Barcode) Encode(b *bcbp.BCBP, wr io.Writer) error {
	return Marshal(b, wr)
}
