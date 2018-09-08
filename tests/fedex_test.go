package kargoTest

import (
	"kargo"
	"testing"
)

func TestFedExExpressPackageIsValid(t *testing.T) {
	p, _ := kargo.NewPackage("9632001960000000000400152152152158")
	fexpress := kargo.NewFedExExpress(p)
	fexpress.Validate()
	if fexpress.Package.IsValid != true {
		t.Errorf("Failed, expected: %t, got: %t.", true, fexpress.Package.IsValid)
	}
}

func TestFedExExpressPackageTrackingNumber(t *testing.T) {
	expectedValue := "152152152158"
	p, _ := kargo.NewPackage("9632001960000000000400152152152158")
	fexpress := kargo.NewFedExExpress(p)
	fexpress.Validate()
	if fexpress.Package.TrackingNumber != expectedValue {
		t.Errorf("Failed, expected: %v, got: %v.", expectedValue, fexpress.Package.TrackingNumber)
	}
}

func TestFedExExpressCarrier(t *testing.T) {
	p, _ := kargo.NewPackage("9632001960000000000400152152152158")
	fexpress := kargo.NewFedExExpress(p)
	fexpress.Validate()
	if fexpress.Package.Carrier != fexpress.GetCarrierName() {
		t.Errorf("Failed, expected: %v, got: %v.", fexpress.GetCarrierName(), fexpress.Package.Carrier)
	}
}

func TestFedExGround96PackageIsValid(t *testing.T) {
	p, _ := kargo.NewPackage("9611019012345612345671")
	fground := kargo.NewFedExGround96(p)
	fground.Validate()
	if fground.Package.IsValid != true {
		t.Errorf("Failed, expected: %t, got: %t.", true, fground.Package.IsValid)
	}
}

func TestFedExGround96PackageTrackingNumber(t *testing.T) {
	expectedValue := "012345612345671"
	p, _ := kargo.NewPackage("9611019012345612345671")
	fground := kargo.NewFedExGround96(p)
	fground.Validate()
	if fground.Package.TrackingNumber != expectedValue {
		t.Errorf("Failed, expected: %v, got: %v.", expectedValue, fground.Package.TrackingNumber)
	}
}

func TestFedExGround96ValidateZeroChecksum(t *testing.T) {
	p, _ := kargo.NewPackage("9611019012345612345640")
	fground := kargo.NewFedExGround96(p)
	fground.Validate()
	if fground.Package.Carrier != fground.GetCarrierName() {
		t.Errorf("Failed, expected: %v, got: %v.", fground.GetCarrierName(), fground.Package.Carrier)
	}
}

func TestFedExGround96Carrier(t *testing.T) {
	p, _ := kargo.NewPackage("9611019012345612345671")
	fground := kargo.NewFedExGround96(p)
	fground.Validate()
	if fground.Package.Carrier != fground.GetCarrierName() {
		t.Errorf("Failed, expected: %v, got: %v.", fground.GetCarrierName(), fground.Package.Carrier)
	}
}
