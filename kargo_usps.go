package kargo

import (
	"regexp"
	"strings"
)

// USPS is a carrier that have package struct
type USPS struct {
	*Package
}

// NewUSPS initialize a new USPS struct with package value
func NewUSPS(p *Package) *USPS {
	return &USPS{Package: p}
}

// GetCarrierName Implements the CarrierFactory interface method
// Retuns the name of carrier struct with package value
func (u *USPS) GetCarrierName() string {
	return "USPS"
}

// Match Implements the CarrierFactory interface method
// Retuns the name of carrier struct with package value
func (u *USPS) Match() bool {

	//Regex to Match in this function
	//Source: https://tools.usps.com/go/TrackConfirmAction_input
	//### LONG
	// USPS Tracking®						9400 1000 0000 0000 0000 00
	//										94001 [0-9]{17}
	// Priority Mail®						9205 5000 0000 0000 0000 00
	//										92055 [0-9]{17}
	// Certified Mail®						9407 3000 0000 0000 0000 00
	//										94073 [0-9]{17}
	// Collect On Delivery Hold For Pickup	9303 3000 0000 0000 0000 00
	//										93033 [0-9]{17}
	// Registered Mail™						9208 8000 0000 0000 0000 00
	//										92088 [0-9]{17}
	// Signature Confirmation™				9202 1000 0000 0000 0000 00
	//										92021 [0-9]{17}
	// Priority Mail Express®				9270 1000 0000 0000 0000 00
	//										92701 [0-9]{17}
	//### MEDIUM (With Tail)
	// Priority Mail Express® (again)		EA 000 000 000 US
	//										EA [0-9]{9} US
	// Priority Mail Express International®	EC 000 000 000 US
	//										EC [0-9]{9} US
	// Priority Mail International®			CP 000 000 000 US
	//										CP [0-9]{9} US
	//### SHORT
	// Global Express Guaranteed®			82 000 000 00
	//										82 [0-9]{8}

	//NOTE: I'm sure there's other codes (like the tail 'US' out there for different Nations)
	//So, I'm putting everything into the below slices so that it'll be easier to add later.
	frontValueLong := []string{"94001", "92055", "93033", "92088", "92021", "92701"}
	frontValueMedium := []string{"EA", "EC", "CP"}
	frontValueShort := []string{"82"}
	backValue := []string{"US"}

	patternLong := "^(" + strings.Join(frontValueLong, "|") + ")[0-9]{17}$"
	patternMediumWithTail := "^(" + strings.Join(frontValueMedium, "|") + ")[0-9]{9}(" + strings.Join(backValue, "|") + ")&"
	patternShort := "^(" + strings.Join(frontValueShort, "|") + ")[0-9]{8}&"

	long, _ := regexp.MatchString(patternLong, u.Package.TrackingNumber)
	medium, _ := regexp.MatchString(patternMediumWithTail, u.Package.TrackingNumber)
	short, _ := regexp.MatchString(patternShort, u.Package.TrackingNumber)

	if (long || medium || short) == false {
		return false
	}
	u.Package.Carrier = u.GetCarrierName()
	return true
}

// GetPackage Implements the CarrierFactory interface method
// Returns the package that carrier holds
func (u *USPS) GetPackage() *Package {
	return u.Package
}

// Validate Implements the CarrierFactory interface method
// Checks whether is a package belongs to that carrier
func (u *USPS) Validate() bool {
	if u.Match() == false {
		return false
	}
	u.Package.Carrier = u.GetCarrierName()
	u.Package.IsValid = true
	return true
}
