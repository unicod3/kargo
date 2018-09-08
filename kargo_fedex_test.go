package kargo

import (
	"testing"
)

func TestFedExExpressPackageMatch(t *testing.T) {
	p, _ := NewPackage("9632001960000000000400152152152158")
	fexpress := NewFedExExpress(p)
	if fexpress.Match() != true {
		t.Errorf("Failed, expected: %t, got: %t.", true, fexpress.Match())
	}
}

func TestFedExExpressPackageWithNonValidNumber(t *testing.T) {
	p, _ := NewPackage("9632001960000000000400152152192158")
	fexpress := NewFedExExpress(p)
	fexpress.Validate()
	if fexpress.Package.IsValid != false {
		t.Errorf("Failed, expected: %t, got: %t.", false, fexpress.Package.IsValid)
	}
}

func TestFedExExpressPackageStringDigit(t *testing.T) {
	p, _ := NewPackage("96320019600000000004001521s2152151")
	fexpress := NewFedExExpress(p)
	fexpress.Validate()
	if fexpress.Package.IsValid != false {
		t.Errorf("Failed, expected: %t, got: %t.", true, fexpress.Package.IsValid)
	}
}

func TestFedExExpressPackageStringCheckDigit(t *testing.T) {
	p, _ := NewPackage("963200196000000000040015215215215s")
	fexpress := NewFedExExpress(p)
	fexpress.Validate()
	if fexpress.Package.IsValid != false {
		t.Errorf("Failed, expected: %t, got: %t.", true, fexpress.Package.IsValid)
	}
}

func TestFedExExpressPackageIsValid(t *testing.T) {
	p, _ := NewPackage("9632001960000000000400152152152158")
	fexpress := NewFedExExpress(p)
	fexpress.Validate()
	if fexpress.Package.IsValid != true {
		t.Errorf("Failed, expected: %t, got: %t.", true, fexpress.Package.IsValid)
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
	if fground.Match() != true {
		t.Errorf("Failed, expected: %t, got: %t.", true, fground.Match())
	}
}

func TestFedExGround96PackageWithNonValidNumber(t *testing.T) {
	p, _ := NewPackage("9611019012345612945671")
	fground := NewFedExGround96(p)
	fground.Validate()
	if fground.Package.IsValid != false {
		t.Errorf("Failed, expected: %t, got: %t.", false, fground.Package.IsValid)
	}
}

func TestFedExGround96PackageStringDigit(t *testing.T) {
	p, _ := NewPackage("96110190123456s2345671")
	fground := NewFedExGround96(p)
	fground.Validate()
	if fground.Package.IsValid != false {
		t.Errorf("Failed, expected: %t, got: %t.", true, fground.Package.IsValid)
	}
}

func TestFedExGround96PackageStringCheckDigit(t *testing.T) {
	p, _ := NewPackage("961101901234561234567s")
	fground := NewFedExGround96(p)
	fground.Validate()
	if fground.Package.IsValid != false {
		t.Errorf("Failed, expected: %t, got: %t.", true, fground.Package.IsValid)
	}
}

func TestFedExGround96PackageIsValid(t *testing.T) {
	p, _ := NewPackage("9611019012345612345671")
	fground := NewFedExGround96(p)
	fground.Validate()
	if fground.Package.IsValid != true {
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

func TestFedExGround96ValidateZeroChecksum(t *testing.T) {
	p, _ := NewPackage("9611019012345612345640")
	fground := NewFedExGround96(p)
	fground.Validate()
	if fground.Package.Carrier != fground.GetCarrierName() {
		t.Errorf("Failed, expected: %v, got: %v.", fground.GetCarrierName(), fground.Package.Carrier)
	}
}

func TestFedExGround96Carrier(t *testing.T) {
	p, _ := NewPackage("9611019012345612345671")
	fground := NewFedExGround96(p)
	fground.Validate()
	if fground.Package.Carrier != fground.GetCarrierName() {
		t.Errorf("Failed, expected: %v, got: %v.", fground.GetCarrierName(), fground.Package.Carrier)
	}
}
