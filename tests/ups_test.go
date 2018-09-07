package kargoTest

import (
	"kargo"
	"testing"
)

func TestUPSTrackingNumber(t *testing.T) {
	expectedValue := "1Z999AA10123456784"
	result, _ := kargo.Identify(expectedValue)

	if result.TrackingNumber != expectedValue {
		t.Errorf("Tracking Number is wrong, got: %s, want: %s.", result.TrackingNumber, expectedValue)
	}
}

func TestUPSTrackingNumberWithSpaces(t *testing.T) {
	expectedValue := "1Z999AA10123456784"
	result, _ := kargo.Identify("1Z9  99AA1012  345678 4")
	if result.TrackingNumber != expectedValue {
		t.Errorf("Tracking Number is wrong, got: %s, want: %s.", result.TrackingNumber, expectedValue)
	}
}

func TestUPSPackageIsValid(t *testing.T) {
	result, _ := kargo.Identify("1Z999AA10123456784")
	if result.IsValid != true {
		t.Errorf("Result is wrong, got: %t, want: %t.", result.IsValid, true)
	}
}
