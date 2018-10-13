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

func TestReverse(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	} {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
