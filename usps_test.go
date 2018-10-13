package kargo

import (
	"testing"
)

// Tests for USPSIMpb
func TestUSPSIMpbPackageMatch(t *testing.T) {
	p, _ := NewPackage("420221539101026837331000039521")
	usps := NewUSPSIMpb(p)
	if usps.Match() != true {
		t.Errorf("Failed, expected: %t, have: %t.", true, usps.Match())
	}
}

func TestUSPSIMpbPackageValidate(t *testing.T) {
	trackingNumbers := map[string]bool{
		"9212391234567812745670":             false, // Non Valid Tracking Number
		"420973179200920559276775601584255r": false, // Char Check Digit
		"420973179200r205592767756015842559": false, // Have a Char
		"4209731792009205592767756015842559": false, // Non Valid Check Digit
		"4209731792009205592767756015842558": true,  // Valid Nine Digit Zip
		"9212391234567812345670":             true,  // Valid Zero Check Digit
	}

	for trackingNumber, expected := range trackingNumbers {
		p, _ := NewPackage(trackingNumber)
		u := NewUSPSIMpb(p)
		u.Validate()
		if u.Package.IsValid != expected {
			t.Errorf("Failed: %s, expected: %t, have: %t.", trackingNumber, expected, u.Package.IsValid)
		}
	}
}

func TestUSPSIMpbTrackingNumberFormatting(t *testing.T) {
	expected_value := "9101026837331000039521"
	p, _ := NewPackage("420221539101026837331000039521")
	usps := NewUSPSIMpb(p)
	usps.Validate()
	if usps.Package.TrackingNumber != expected_value {
		t.Errorf("Failed, expected: %v, have: %v.",
			expected_value, usps.Package.TrackingNumber)
	}
}

// Tests for USPSS10
func TestUSPSS10PackageMatch(t *testing.T) {
	p, _ := NewPackage("EF123456785US")
	usps := NewUSPSS10(p)
	if usps.Match() != true {
		t.Errorf("Failed, expected: %t, have: %t.", true, usps.Match())
	}
}

func TestUSPSS10PackageValidate(t *testing.T) {
	trackingNumbers := map[string]bool{
		"EF12345678US":  false, // Non Valid Tracking Number
		"EF12r456786US": false, // Have a Char
		"EF123456786US": false, // Non Valid Check Digit
		"EF123456785US": true,  // Valid Tracking Number
		"RZ030057180PH": true,  // Valid Check Digit
		"VA456789015KG": true,  // Valid Check Digit With Remainder 11
	}

	for trackingNumber, expected := range trackingNumbers {
		p, _ := NewPackage(trackingNumber)
		u := NewUSPSS10(p)
		u.Validate()
		if u.Package.IsValid != expected {
			t.Errorf("Failed: %s, expected: %t, have: %t.", trackingNumber, expected, u.Package.IsValid)
		}
	}
}

// Tests for USPS20
func TestUSPS20PackageMatch(t *testing.T) {
	p, _ := NewPackage("71123456789123456787")
	usps := NewUSPS20(p)
	if usps.Match() != true {
		t.Errorf("Failed, expected: %t, have: %t.", true, usps.Match())
	}
}

func TestUSPS20PackageValidate(t *testing.T) {
	trackingNumbers := map[string]bool{
		"72123456789123456787": false, // Non Valid Tracking Number
		"71123456r89123456787": false, // Have a Char
		"7112345678912345678r": false, // Char Check Digit
		"71123456789123456787": true,  // Valid Tracking Number
	}

	for trackingNumber, expected := range trackingNumbers {
		p, _ := NewPackage(trackingNumber)
		u := NewUSPS20(p)
		u.Validate()
		if u.Package.IsValid != expected {
			t.Errorf("Failed: %s, expected: %t, have: %t.", trackingNumber, expected, u.Package.IsValid)
		}
	}
}
