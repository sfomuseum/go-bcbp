package bcbp

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"sort"
	"strings"

	"github.com/aaronland/go-roster"
)

type Barcode interface {
	Encode(*BCBP, io.Writer) error
	Decode(io.Reader) (*BCBP, error)
}

var barcode_roster roster.Roster

// BarcodeInitializationFunc is a function defined by individual barcode package and used to create
// an instance of that barcode
type BarcodeInitializationFunc func(ctx context.Context, uri string) (Barcode, error)

// RegisterBarcode registers 'scheme' as a key pointing to 'init_func' in an internal lookup table
// used to create new `Barcode` instances by the `NewBarcode` method.
func RegisterBarcode(ctx context.Context, scheme string, init_func BarcodeInitializationFunc) error {

	err := ensureBarcodeRoster()

	if err != nil {
		return err
	}

	return barcode_roster.Register(ctx, scheme, init_func)
}

func ensureBarcodeRoster() error {

	if barcode_roster == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		barcode_roster = r
	}

	return nil
}

// NewBarcode returns a new `Barcode` instance configured by 'uri'. The value of 'uri' is parsed
// as a `url.URL` and its scheme is used as the key for a corresponding `BarcodeInitializationFunc`
// function used to instantiate the new `Barcode`. It is assumed that the scheme (and initialization
// function) have been registered by the `RegisterBarcode` method.
func NewBarcode(ctx context.Context, uri string) (Barcode, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := barcode_roster.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	init_func := i.(BarcodeInitializationFunc)
	return init_func(ctx, uri)
}

// BarcodeSchemes returns the list of schemes that have been registered.
func BarcodeSchemes() []string {

	ctx := context.Background()
	schemes := []string{}

	err := ensureBarcodeRoster()

	if err != nil {
		return schemes
	}

	for _, dr := range barcode_roster.Drivers(ctx) {
		scheme := fmt.Sprintf("%s://", strings.ToLower(dr))
		schemes = append(schemes, scheme)
	}

	sort.Strings(schemes)
	return schemes
}
