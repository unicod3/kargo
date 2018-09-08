package kargo

import (
	"testing"
)

func TestKargoIdentifyFedExExpressTrackingNumber(t *testing.T) {
	expected := "FedEx"
	pkg, _ := Identify("9632001960000000000400152152152158")

	if pkg.Carrier != expected {
		t.Errorf("Failed, expected: %v, want: %v.", expected, pkg.TrackingNumber)
	}
}

func TestKargoIdentifyFedExGround96TrackingNumber(t *testing.T) {
	expected := "FedEx"
	pkg, _ := Identify("9611019012345612345640")

	if pkg.Carrier != expected {
		t.Errorf("Failed, expected: %v, want: %v.", expected, pkg.TrackingNumber)
	}
}

func TestKargoIdentifyUPSTrackingNumber(t *testing.T) {
	expected := "UPS"
	pkg, _ := Identify("1Z999AA10123456784")

	if pkg.Carrier != expected {
		t.Errorf("Failed, expected: %v, want: %v.", expected, pkg.TrackingNumber)
	}
}

func TestKargoIdentifyReachPackageWithEmptyTrackingNumber(t *testing.T) {
	expected := "Unknown"
	pkg, err := Identify("")

	if pkg.Carrier != expected && err != nil {
		t.Errorf("Failed, expected: %v, want: %v.", expected, pkg.Carrier)
	}
}

func TestKargoIdentifyUnknownCarrier(t *testing.T) {
	expected := "Unknown"
	pkg, _ := Identify("WrongTrackingNumber")

	if pkg.Carrier != expected {
		t.Errorf("Failed, expected: %v, want: %v.", expected, pkg.TrackingNumber)
	}
}

func TestKargoIdentifyCarrier(t *testing.T) {
	expected := "UPS"
	pkg, _ := Identify("1Z999AA10123456784")

	if pkg.Carrier != expected {
		t.Errorf("Failed, expected: %v, want: %v.", expected, pkg.TrackingNumber)
	}
}

func TestKargoIdentifyNotValid(t *testing.T) {
	expected := false
	pkg, _ := Identify("1Z399AA10123456784")

	if pkg.IsValid != expected {
		t.Errorf("Failed, expected: %v, want: %v.", expected, pkg.TrackingNumber)
	}
}

func TestKargoIdentifyValid(t *testing.T) {
	expected := true
	pkg, _ := Identify("1Z999AA10123456784")

	if pkg.IsValid != expected {
		t.Errorf("Failed, expected: %v, want: %v.", expected, pkg.TrackingNumber)
	}
}

func TestKargoIdentifyEmptyTracking(t *testing.T) {
	expected := "Tracking Number can not be empty!"
	pkg, err := Identify("")

	if err == nil {
		t.Errorf("Failed, expected: %v, want: %v.", expected, pkg.TrackingNumber)
	}
}
