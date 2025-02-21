package aztec

import (
	"io"

	"github.com/sfomuseum/go-bcbp"
)

type AztecBarcode struct {
	bcbp.Barcode
}

func (a *AztecBarcode) Decode(r io.Reader) (*bcbp.BCBP, error) {
	return Unmarshal(r)
}

func (a *AztecBarcode) Encode(b *bcbp.BCBP, wr io.Writer) error {
	return Marshal(b, wr)
}
