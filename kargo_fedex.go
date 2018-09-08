package kargo

import (
	"strconv"
	"strings"

	"github.com/golang/example/stringutil"
)

// FedExGround96 defines a struct for Fedex Ground packages
type FedExGround96 struct {
	*Package
}

// FedExExpress defines a struct for Fedex Express packages
type FedExExpress struct {
	*Package
}

// NewFedExGround96 initializes a new FedExGround96 struct with package value
func NewFedExGround96(p *Package) *FedExGround96 {
	return &FedExGround96{Package: p}
}

// GetCarrierName Implements the CarrierFactory interface method
// Retuns the name of carrier
func (f *FedExGround96) GetCarrierName() string {
	return "FedEx"
}

// GetPackage Implements the CarrierFactory interface method
// Returns the package that carrier holds
func (f *FedExGround96) GetPackage() *Package {
	return f.Package
}

// formatTrackingNumber formats the fedex barcode tracking number to
// normal tracking number
func (f *FedExGround96) formatTrackingNumber() {
	f.Package.TrackingNumber = f.Package.TrackingNumber[7:]
}

// Validate Implements the CarrierFactory interface method
// Checks whether is a package belongs to that carrier
func (f *FedExGround96) Validate() bool {

	f.formatTrackingNumber()
	chars := f.Package.TrackingNumber[:len(f.Package.TrackingNumber)-1]
	checkDigit, err := strconv.Atoi(f.Package.TrackingNumber[len(f.Package.TrackingNumber)-1:])
	if err != nil {
		return false
	}

	odd, even := 0, 0
	reversed := stringutil.Reverse(chars)
	for i, char := range reversed {
		t := (string(char))
		num, err := strconv.Atoi(t)
		if err != nil {
			return false
		}

		if i%2 == 0 {
			even += num
			continue
		}
		odd += num
	}
	check := ((even * 3) + odd) % 10
	if check != 0 {
		check = 10 - check
	}

	if check != checkDigit {
		return false
	}

	f.Package.Carrier = f.GetCarrierName()
	f.Package.IsValid = true
	return true
}

// NewFedExExpress initializes a new FedExExpress struct with package value
func NewFedExExpress(p *Package) *FedExExpress {
	return &FedExExpress{Package: p}
}

// GetCarrierName Implements the CarrierFactory interface method
// Retuns the name of carrier
func (f *FedExExpress) GetCarrierName() string {
	return "FedEx"
}

// GetPackage Implements the CarrierFactory interface method
// Returns the package that carrier holds
func (f *FedExExpress) GetPackage() *Package {
	return f.Package
}

// formatTrackingNumber formats the fedex barcode tracking number to
// normal tracking number
func (f *FedExExpress) formatTrackingNumber() {
	left := strings.TrimLeft(f.Package.TrackingNumber[20:22], "0")
	right := f.Package.TrackingNumber[22:]
	f.Package.TrackingNumber = left + right
}

// Validate Implements the CarrierFactory interface method
// Checks whether is a package belongs to that carrier
func (f *FedExExpress) Validate() bool {

	f.formatTrackingNumber()
	chars := f.Package.TrackingNumber[:len(f.Package.TrackingNumber)-1]
	checkDigit, err := strconv.Atoi(f.Package.TrackingNumber[len(f.Package.TrackingNumber)-1:])
	if err != nil {
		return false
	}

	total := 0
	factors := [3]int{1, 3, 7}
	reversed := stringutil.Reverse(chars)
	for i, char := range reversed {

		num, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		total += num * factors[i%3]
	}

	if total%11%10 != checkDigit {
		return false
	}

	f.Package.Carrier = f.GetCarrierName()
	f.Package.IsValid = true
	return true
}
