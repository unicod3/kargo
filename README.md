# ![Logo](https://png.icons8.com/color/32/000000/snail.png)Kargo 
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/unicod3/kargo) [![Build Status](https://travis-ci.org/unicod3/kargo.svg?branch=master)](https://travis-ci.org/unicod3/kargo) [![Coverage Status](https://coveralls.io/repos/github/unicod3/kargo/badge.svg)](https://coveralls.io/github/unicod3/kargo)

Kargo determines carrier of the tracking number barcode and checks if that tracking number barcode is valid by matching the format and calculating checksum validity. Feel free to create pull requests and issues.

Package supports  
- [x] UPS
- [x] FedEx Express
- [x] FedEx Ground "96"
- [x] USPS IMpb
- [x] USPS UPU S10
- [x] USPS 20

But more on the way
- [ ] DHL

## Install

```bash
go get github.com/unicod3/kargo
```

## Usage and Examples

Kargo's usage is pretty straight forward, supply a tracking number and get the Package struct

```go
package main

import (
	"fmt"
	"github.com/unicod3/kargo"
)

func main() {
	pkg, err := kargo.Identify("1Z999AA10123456784")
	if err != nil {
		fmt.Println(err)
	}

	var trackingNumber string = pkg.TrackingNumber
	var carrier string = pkg.Carrier
	var isValid bool = pkg.IsValid
	fmt.Printf("Tracking Number: %v, Carrier: %v, Is Valid: %t",
		trackingNumber, carrier, isValid)
}

```
Output: `Tracking Number: 1Z999AA10123456784, Carrier: UPS, Is Valid: true`

For more examples please check out the test files.

## Running the tests

You can run tests with standard test tool of Go:

```bash
go test -v 
```

## Changelog

Version 1.1.2
   * Added support for USPS; IMpb, UPU S10, 20 character barcodes
   * Removed the unnecessary prefix from filenames
   * Improved the tests

Version 1.0.1
   * Fix stringutils package dependency issue

Version 1.0.0
   * Initial release


## License

GNU General Public License v3.0 - see [LICENSE](LICENSE) for more details
