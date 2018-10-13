package kargo

import (
	"regexp"
	"strconv"
)

// USPSIMpb is a carrier that have package struct
type USPSIMpb struct {
	*Package
}

// USPS10 is a carrier that have package struct
type USPSS10 struct {
	*Package
}

// USPS20 is a carrier that have package struct
type USPS20 struct {
	*Package
}

// NewUSPSIMpb initialize a new USPSIMpb struct with package value
func NewUSPSIMpb(p *Package) *USPSIMpb {
	return &USPSIMpb{Package: p}
}

// NewUSPSS10 initialize a new NewUSPSS10 struct with package value
func NewUSPSS10(p *Package) *USPSS10 {
	return &USPSS10{Package: p}
}

// NewUSPS20 initialize a new USPS20 struct with package value
func NewUSPS20(p *Package) *USPS20 {
	return &USPS20{Package: p}
}

// GetCarrierName Implements the CarrierFactory interface method
// Retuns the name of carrier struct with package value
func (u *USPSIMpb) GetCarrierName() string {
	return "USPS"
}

// GetCarrierName Implements the CarrierFactory interface method
// Retuns the name of carrier struct with package value
func (u *USPSS10) GetCarrierName() string {
	return "USPS"
}

// GetCarrierName Implements the CarrierFactory interface method
// Retuns the name of carrier struct with package value
func (u *USPS20) GetCarrierName() string {
	return "USPS"
}

// Match Implements the CarrierFactory interface method
// Retuns the name of carrier struct with package value
func (u *USPSIMpb) Match() bool {
	regexIMpb := `^(?:420(?:\d{9}|\d{5}))?(9[1-5]\d{20,24})$`
	if m, _ := regexp.MatchString(regexIMpb, u.Package.TrackingNumber); m == false {
		return false
	}
	u.Package.Carrier = u.GetCarrierName()
	return true
}

// Match Implements the CarrierFactory interface method
// Retuns the name of carrier struct with package value
func (u *USPSS10) Match() bool {
	if m, _ := regexp.MatchString(`^[A-Z]{2}\d{9}[A-Z]{2}$`, u.Package.TrackingNumber); m == false {
		return false
	}
	u.Package.Carrier = u.GetCarrierName()
	return true
}

// Match Implements the CarrierFactory interface method
// Retuns the name of carrier struct with package value
func (u *USPS20) Match() bool {
	if m, _ := regexp.MatchString(`^(71|73|77|81)\d{18}$`, u.Package.TrackingNumber); m == false {
		return false
	}
	u.Package.Carrier = u.GetCarrierName()
	return true
}

// GetPackage Implements the CarrierFactory interface method
// Returns the package that carrier holds
func (u *USPSIMpb) GetPackage() *Package {
	return u.Package
}

// GetPackage Implements the CarrierFactory interface method
// Returns the package that carrier holds
func (u *USPSS10) GetPackage() *Package {
	return u.Package
}

// GetPackage Implements the CarrierFactory interface method
// Returns the package that carrier holds
func (u *USPS20) GetPackage() *Package {
	return u.Package
}

// formatTrackingNumber formats the usps IMpb barcode tracking number to
// normal tracking number
func (u *USPSIMpb) formatTrackingNumber() {
	regexpIMpb := `^(?:420(?:\d{9}|\d{5}))?(9[1-5]\d{20,24})$`
	re := regexp.MustCompile(regexpIMpb)
	subs := re.FindStringSubmatch(u.Package.TrackingNumber)
	if len(subs) > 1 {
		u.Package.TrackingNumber = subs[1]
	}
}

// Validate Implements the CarrierFactory interface method
// Checks whether is a package belongs to that carrier
func (u *USPSIMpb) Validate() bool {
	u.formatTrackingNumber()
	chars := u.Package.TrackingNumber[:len(u.Package.TrackingNumber)-1]
	checkDigit, err := strconv.Atoi(u.Package.TrackingNumber[len(u.Package.TrackingNumber)-1:])
	if err != nil {
		return false
	}

	odd, even := 0, 0
	reversed := Reverse(chars)
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

	u.Package.Carrier = u.GetCarrierName()
	u.Package.IsValid = true
	return true
}

// Validate Implements the CarrierFactory interface method
// Checks whether is a package belongs to that carrier
func (u *USPSS10) Validate() bool {
	chars := u.Package.TrackingNumber[2:10]
	checkDigit, err := strconv.Atoi(u.Package.TrackingNumber[10:11])
	if err != nil {
		return false
	}
	total := 0
	factors := [8]int{8, 6, 4, 2, 3, 5, 9, 7}
	for i, char := range chars {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		total += num * factors[i]
	}
	check := 11 - (total % 11)
	if check == 10 {
		check = 0
	}
	if check == 11 {
		check = 5
	}
	if check != checkDigit {
		return false
	}

	u.Package.Carrier = u.GetCarrierName()
	u.Package.IsValid = true
	return true
}

// Validate Implements the CarrierFactory interface method
// Checks whether is a package belongs to that carrier
func (u *USPS20) Validate() bool {
	chars := u.Package.TrackingNumber[:len(u.Package.TrackingNumber)-1]
	checkDigit, err := strconv.Atoi(u.Package.TrackingNumber[len(u.Package.TrackingNumber)-1:])
	if err != nil {
		return false
	}

	odd, even := 0, 0
	reversed := Reverse(chars)
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
	u.Package.Carrier = u.GetCarrierName()
	u.Package.IsValid = true
	return true
}
