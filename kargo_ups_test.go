package kargo

import (
	"testing"
)

func TestUPSPackageMatch(t *testing.T) {
	p, _ := NewPackage("1Z999AA10123456784")
	ups := NewUPS(p)
	if ups.Match() != true {
		t.Errorf("Failed, expected: %t, want: %t.", true, ups.Match())
	}
}

func TestUPSPackageWithNonValidNumber(t *testing.T) {
	p, _ := NewPackage("1Z999AA10723456784")
	ups := NewUPS(p)
	ups.Validate()
	if ups.Package.IsValid != false {
		t.Errorf("Failed, expected: %t, want: %t.", false, ups.Package.IsValid)
	}
}

func TestUPSPackageStringCheckDigit(t *testing.T) {
	p, _ := NewPackage("1Z999AA1012345678s")
	ups := NewUPS(p)
	ups.Validate()
	if ups.Package.IsValid != false {
		t.Errorf("Failed, expected: %t, want: %t.", true, ups.Package.IsValid)
	}
}

func TestUPSPackageIsValid(t *testing.T) {
	p, _ := NewPackage("1Z999AA10123456784")
	ups := NewUPS(p)
	ups.Validate()
	if ups.Package.IsValid != true {
		t.Errorf("Failed, expected: %t, want: %t.", true, ups.Package.IsValid)
	}
}

func TestUPSValidateZeroChecksum(t *testing.T) {
	p, _ := NewPackage("1Z999AA10123456784")
	ups := NewUPS(p)
	ups.Validate()
	if ups.Package.IsValid != true {
		t.Errorf("Failed, expected: %t, want: %t.", true, ups.Package.IsValid)
	}
}
