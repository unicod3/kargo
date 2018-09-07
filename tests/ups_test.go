package kargoTest

import (
	"kargo"
	"testing"
)

func TestUPSTrackingNumber(t *testing.T) {
	expectedValue := "1Z999AA10123456784"

	p, _ := kargo.NewPackage(expectedValue)
	ups := kargo.NewUPS(p)

	if ups.Package.TrackingNumber != expectedValue {
		t.Errorf("got: %s, want: %s.", ups.Package.TrackingNumber, expectedValue)
	}
}
func TestUPSTrackingNumberWithSpaces(t *testing.T) {
	expectedValue := "1Z999AA10123456784"

	p, _ := kargo.NewPackage("1Z9  99AA1012  345678 4")
	ups := kargo.NewUPS(p)

	if ups.Package.TrackingNumber != expectedValue {
		t.Errorf("got: %s, want: %s.", ups.TrackingNumber, expectedValue)
	}
}

func TestUPSPackageIsValid(t *testing.T) {
	p, _ := kargo.NewPackage("1Z999AA10123456784")
	ups := kargo.NewUPS(p)
	ups.Validate()
	if ups.Package.IsValid != true {
		t.Errorf("got: %t, want: %t.", ups.Package.IsValid, true)
	}
}
func TestUPSValidateZeroChecksum(t *testing.T) {
	p, _ := kargo.NewPackage("1Z999AA10123456784")
	ups := kargo.NewUPS(p)
	ups.Validate()
	if ups.Package.IsValid != true {
		t.Errorf("got: %t, want: %t.", ups.Package.IsValid, true)
	}
}
