package kargo

import (
	"regexp"
	"strconv"
)

// UPS is a carrier that have package struct
type UPS struct {
	*Package
}

// NewUPS initialize a new UPS struct with package value
func NewUPS(p *Package) *UPS {
	return &UPS{Package: p}
}

// GetCarrierName Implements the CarrierFactory interface method
// Retuns the name of carrier struct with package value
func (u *UPS) GetCarrierName() string {
	return "UPS"
}

// Match Implements the CarrierFactory interface method
// Retuns the name of carrier struct with package value
func (u *UPS) Match() bool {
	if m, _ := regexp.MatchString(`^1Z[A-Z0-9]{16}$`, u.Package.TrackingNumber); m == false {
		return false
	}
	u.Package.Carrier = u.GetCarrierName()
	return true
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
