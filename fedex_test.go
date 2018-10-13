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

func TestFedExExpressPackage(t *testing.T) {
	trackingNumbers := [3]string {
		"9632001960000000000400152152192158", //WithNonValidNumber
		"96320019600000000004001521s2152151", //StringDigit
		"963200196000000000040015215215215s", //StringCheckDigit
	}

	for _, trackingNumber := range trackingNumbers {
		p, _ := NewPackage(trackingNumber)
		fexpress := NewFedExExpress(p)
		fexpress.Validate()
		if fexpress.Package.IsValid {
			t.Errorf("Failed for %s, expected: %t, got: %t.", trackingNumber, false, fexpress.Package.IsValid)
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

func TestFedExExpressCarrier(t *testing.T) {
	p, _ := NewPackage("9632001960000000000400152152152158")
	fexpress := NewFedExExpress(p)
	fexpress.Validate()
	if fexpress.Package.Carrier != fexpress.GetCarrierName() {
		t.Errorf("Failed, expected: %v, got: %v.", fexpress.GetCarrierName(), fexpress.Package.Carrier)
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
	trackingNumbers := [3]string {
		"9611019012345612945671", //WithNonValidNumber
		"96110190123456s2345671", //StringDigit
		"961101901234561234567s", //StringCheckDigit
	}

	for _, trackingNumber := range trackingNumbers {
		p, _ := NewPackage(trackingNumber)
		fground := NewFedExGround96(p)
		fground.Validate()
		if fground.Package.IsValid {
			t.Errorf("Failed for %s, expected: %t, got: %t.", trackingNumber, false, fground.Package.IsValid)
		}
	}
}

func TestFedExGround96PackageIsValid(t *testing.T) {
	p, _ := NewPackage("9611019012345612345671")
	fground := NewFedExGround96(p)
	fground.Validate()
	if !fground.Package.IsValid {
		t.Errorf("Failed, expected: %t, got: %t.", true, fground.Package.IsValid)
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

func TestFedExGround96(t *testing.T) {
	trackingNumbers := [2]string {
		"9611019012345612345640", //ZeroChecksum
		"9611019012345612345671", //Carrier
	}

	for _, trackingNumber := range trackingNumbers {
		p, _ := NewPackage(trackingNumber)
		fground := NewFedExGround96(p)
		fground.Validate()
		if fground.Package.Carrier != fground.GetCarrierName() {
			t.Errorf("Failed, expected: %v, got: %v.", fground.GetCarrierName(), fground.Package.Carrier)
		}
	}
}
