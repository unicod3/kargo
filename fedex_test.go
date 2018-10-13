package kargo

import (
	"testing"
)

func TestFedExExpressPackageMatch(t *testing.T) {
	p, _ := NewPackage("9632001960000000000400152152152158")
	fexpress := NewFedExExpress(p)
	if !fexpress.Match() {
		t.Errorf("Failed, expected: %t, got: %t.", true, fexpress.Match())
	}
}

func TestFedExExpressPackageValidate(t *testing.T) {
	trackingNumbers := map[string]bool{
		"9632001960000000000400152152192158": false, // Not Valid Tracking Number
		"96320019600000000004001521s2152151": false, // Have Char Digit
		"963200196000000000040015215215215s": false, // Char Check Digit
		"9632001960000000000400152152152158": true,  // Valid Tracking Number
	}

	for trackingNumber, expected := range trackingNumbers {
		p, _ := NewPackage(trackingNumber)
		fexpress := NewFedExExpress(p)
		fexpress.Validate()
		if fexpress.Package.IsValid != expected {
			t.Errorf("Failed: %s, expected: %t, have: %t.", trackingNumber, expected, fexpress.Package.IsValid)
		}
	}
}

func TestFedExExpressPackageTrackingNumber(t *testing.T) {
	expectedValue := "152152152158"
	p, _ := NewPackage("9632001960000000000400152152152158")
	fexpress := NewFedExExpress(p)
	fexpress.Validate()
	if fexpress.Package.TrackingNumber != expectedValue {
		t.Errorf("Failed, expected: %v, got: %v.", expectedValue, fexpress.Package.TrackingNumber)
	}
}

func TestFedExGround96PackageMatch(t *testing.T) {
	p, _ := NewPackage("9611019012345612345671")
	fground := NewFedExGround96(p)
	if !fground.Match() {
		t.Errorf("Failed, expected: %t, got: %t.", true, fground.Match())
	}
}

func TestFedExGround96Package(t *testing.T) {
	trackingNumbers := map[string]bool{
		"9611019012345612945671": false, // Not Valid Tracking Number
		"96110190123456s2345671": false, // Have a Char Digit
		"961101901234561234567s": false, // Char Check Digit
		"9611019012345612345671": true,  // Valid Tracking Number
		"9611019012345612345640": true,  // Zero Check Digit
	}

	for trackingNumber, expected := range trackingNumbers {
		p, _ := NewPackage(trackingNumber)
		fground := NewFedExGround96(p)
		fground.Validate()
		if fground.Package.IsValid != expected {
			t.Errorf("Failed: %s, expected: %t, have: %t.", trackingNumber, expected, fground.Package.IsValid)
		}
	}
}

func TestFedExGround96PackageTrackingNumber(t *testing.T) {
	expectedValue := "012345612345671"
	p, _ := NewPackage("9611019012345612345671")
	fground := NewFedExGround96(p)
	fground.Validate()
	if fground.Package.TrackingNumber != expectedValue {
		t.Errorf("Failed, expected: %v, got: %v.", expectedValue, fground.Package.TrackingNumber)
	}
}
