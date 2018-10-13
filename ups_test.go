package kargo

import (
	"testing"
)

func TestUPSPackageMatch(t *testing.T) {
	p, _ := NewPackage("1Z999AA10123456784")
	ups := NewUPS(p)
	if ups.Match() != true {
		t.Errorf("Failed, expected: %t, have: %t.", true, ups.Match())
	}
}

func TestUPSPackageValidate(t *testing.T) {
	trackingNumbers := map[string]bool{
		"1Z999AA10723456784": false, // Non Valid Tracking Number
		"1Z999AA1012345678s": false, // Char Check Digit
		"1Z999AA10123456784": true,  // Valid Tracking Number
		"1Z879E930346834440": true,  // Valid Zero Check Digit
	}

	for trackingNumber, expected := range trackingNumbers {
		p, _ := NewPackage(trackingNumber)
		ups := NewUPS(p)
		ups.Validate()
		if ups.Package.IsValid != expected {
			t.Errorf("Failed, expected: %t, have: %t.", expected, ups.Package.IsValid)
		}
	}
}
