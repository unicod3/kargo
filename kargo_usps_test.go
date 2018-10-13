package kargo

import (
	"testing"
)

func TestUSPSPackageMatch(t *testing.T) {
	p, _ := NewPackage("9400 1000 1337 0000 0000 00")
	usps := NewUSPS(p)
	if usps.Match() != true {
		t.Errorf("Failed, expected: %t, want: %t.", true, usps.Match())
	}
}

func TestUSPSPackageWithNonValidNumber(t *testing.T) {
	p, _ := NewPackage("1Z999AA10723456784")
	usps := NewUSPS(p)
	usps.Validate()
	if usps.Package.IsValid != false {
		t.Errorf("Failed, expected: %t, want: %t.", false, usps.Package.IsValid)
	}
}

func TestUSPSPackageStringCheckDigit(t *testing.T) {
	p, _ := NewPackage("9400 1000 leet 0000 0000 00")
	usps := NewUSPS(p)
	usps.Validate()
	if usps.Package.IsValid != false {
		t.Errorf("Failed, expected: %t, want: %t.", true, usps.Package.IsValid)
	}
}

func TestUSPSPackageIsValid(t *testing.T) {
	p, _ := NewPackage("9400 1000 1337 0000 0000 00")
	usps := NewUSPS(p)
	usps.Validate()
	if usps.Package.IsValid != true {
		t.Errorf("Failed, expected: %t, want: %t.", true, usps.Package.IsValid)
	}
}
