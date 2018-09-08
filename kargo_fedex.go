package kargo

import (
	"strconv"
	"strings"

	"github.com/golang/example/stringutil"
)

type FedExGround96 struct {
	*Package
}

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

// Validate Implements the CarrierFactory interface method
// Checks whether is a package belongs to that carrier
func (f *FedExGround96) Validate() bool {
	//p.Carrier = "Fedex"
	return false
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

func (f *FedExExpress) formatTrackingNumber() {
	left := strings.TrimLeft(f.Package.TrackingNumber[20:22], "0")
	right := f.Package.TrackingNumber[22:len(f.Package.TrackingNumber)]
	f.Package.TrackingNumber = left + right
}

// Validate Implements the CarrierFactory interface method
// Checks whether is a package belongs to that carrier
func (f *FedExExpress) Validate() bool {

	f.formatTrackingNumber()
	chars, checkDigit := f.Package.TrackingNumber[:len(f.Package.TrackingNumber)-1],
		f.Package.TrackingNumber[len(f.Package.TrackingNumber)-1:]

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

	numCheckDigit, err := strconv.Atoi(string(checkDigit))
	if err != nil {
		return false
	}

	if total%11%10 != numCheckDigit {
		return false
	}

	f.Package.Carrier = f.GetCarrierName()
	f.Package.IsValid = true
	return true
}
