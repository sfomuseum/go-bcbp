# go-bcbp

Opinionated Go package for working with BCBP data both as raw strings and as PDF417 or Aztec barcodes.

## Documentation

Go documentation is incomplete.

## Motivation

This package is designed to marshal and unmarshal IATA BCBP data in to native Go structs.

It was written to account for errors parsing BCBP data using the [martinmroz/iata_bcbp](https://github.com/martinmroz/iata_bcbp) Rust package and to better understand the BCBP data format in general.

## Example

```
import (
       "fmt"

       "github.com/sfomuseum/go-bcbp"
)

func main() {

	raw := "M1DESMARAIS/LUC       EABC123 YULFRAAC 0834 326J001A0025 100"
	bcbp_data, _ := bcbp.Unmarshal(raw)

	// prints 1
	fmt.Println(len(bcbp_data.Legs))

	// prints YUL
	fmt.Println(bcbp_data.Legs[0].FromAirport)
```		     		

## Caveats and known-knowns

### Missing or invalid group separators for multi-leg strings

This code does not handle multi-leg BCBP data with invalid or missing group separators yet.

### Decoding PDF417 barcodes

This code does not support decoding PDF417 barcodes yet.

## See also

* https://github.com/zxing/zxing
* https://github.com/rxing-core/rxing
* https://github.com/martinmroz/iata_bcbp