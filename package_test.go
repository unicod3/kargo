package kargo

import (
	"testing"
)

func TestPackageTrackingNumber(t *testing.T) {
	expectedValue := "someTrackingNumber"
	p, _ := NewPackage(expectedValue)
	if p.TrackingNumber != expectedValue {
		t.Errorf("Failed, expected: %v, got: %v.", expectedValue, p.TrackingNumber)
	}
}

func TestPackageTrackingNumberWithSpaces(t *testing.T) {
	expectedValue := "someTrackingNumber"

	p, _ := NewPackage("  some Tracking Number")

	if p.TrackingNumber != expectedValue {
		t.Errorf("Failed, expected: %v, got: %v.", expectedValue, p.TrackingNumber)
	}
}
