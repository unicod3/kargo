package kargoTest

import (
	"kargo"
	"testing"
)

func TestPackageTrackingNumber(t *testing.T) {
	expectedValue := "someTrackingNumber"
	p, _ := kargo.NewPackage(expectedValue)
	if p.TrackingNumber != expectedValue {
		t.Errorf("Failed, expected: %v, got: %v.", expectedValue, p.TrackingNumber)
	}
}

func TestPackageTrackingNumberWithSpaces(t *testing.T) {
	expectedValue := "1Z999AA10123456784"

	p, _ := kargo.NewPackage("  1Z9  99AA1012  345678 4")

	if p.TrackingNumber != expectedValue {
		t.Errorf("Failed, expected: %v, got: %v.", expectedValue, p.TrackingNumber)
	}
}
