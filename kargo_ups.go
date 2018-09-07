package kargo

import (
	"strconv"
)

// UPS is a carrier that have package struct
type UPS struct {
	*Package
}

// GetCarrierName retuns the name of carrier struct with package value
func (u *UPS) GetCarrierName() string {
	return "UPS"
}

// NewUPS initialize a new UPS struct with package value
func NewUPS(p *Package) *UPS {
	return &UPS{Package: p}
}

// GetPackage Implements the CarrierFactory interface method
// Returns the package that carrier holds
func (u *UPS) GetPackage() *Package {
	return u.Package
}

// Validate Implements the CarrierFactory interface method
// Checks whether is a package belongs to that carrier
func (u *UPS) Validate() bool {
	chars := u.Package.TrackingNumber[2 : len(u.Package.TrackingNumber)-1]
	checkDigit, err := strconv.Atoi(u.Package.TrackingNumber[len(u.Package.TrackingNumber)-1:])
	if err != nil {
		return false
	}

	var odd, even int = 0, 0
	for i, char := range chars {

		t := (string(char))
		num, err := strconv.Atoi(t)
		if err != nil {
			num = int(char-3) % 10
		}

		if i%2 == 0 {
			even += num
			continue
		}

		odd += num

	}

	check := ((odd * 2) + even) % 10
	if check != 0 {
		check = 10 - check
	}

	if check != checkDigit {
		return false
	}

	u.Package.Carrier = u.GetCarrierName()
	u.Package.IsValid = true
	return true
}
