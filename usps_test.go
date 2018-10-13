package kargo

import (
	"testing"
)

// Tests for usps IMpb
func TestUSPSIMpbPackageMatch(t *testing.T) {
	p, _ := NewPackage("420221539101026837331000039521")
	usps := NewUSPSIMpb(p)
	if usps.Match() != true {
		t.Errorf("Failed, expected: %t, have: %t.", true, usps.Match())
	}
}

func TestUSPSIMpbPackageValidNoZip(t *testing.T) {
	p, _ := NewPackage("9212391234567812345670")
	usps := NewUSPSIMpb(p)
	if usps.Match() != true {
		t.Errorf("Failed, expected: %t, have: %t.", false, usps.Package.IsValid)
	}
}

func TestUSPSIMpbPackageValidNineDigitZip(t *testing.T) {
	p, _ := NewPackage("4209731792009205592767756015842558")
	usps := NewUSPSIMpb(p)
	if usps.Match() != true {
		t.Errorf("Failed, expected: %t, have: %t.", false, usps.Package.IsValid)
	}
}

func TestUSPSIMpbPackageWithNonValidNumber(t *testing.T) {
	p, _ := NewPackage("1Z999AA10723456784")
	usps := NewUSPSIMpb(p)
	usps.Validate()
	if usps.Package.IsValid != false {
		t.Errorf("Failed, expected: %t, have: %t.", false, usps.Package.IsValid)
	}
}

func TestUSPSIMpbPackageWithNonValidCharCheckDigit(t *testing.T) {
	p, _ := NewPackage("420973179200920559276775601584255r")
	usps := NewUSPSIMpb(p)
	usps.Validate()
	if usps.Package.IsValid != false {
		t.Errorf("Failed, expected: %t, have: %t.", false, usps.Package.IsValid)
	}
}

func TestUSPSIMpbPackageWithNonValidCheckDigit(t *testing.T) {
	p, _ := NewPackage("4209731792009205592767756015842559")
	usps := NewUSPSIMpb(p)
	usps.Validate()
	if usps.Package.IsValid != false {
		t.Errorf("Failed, expected: %t, have: %t.", false, usps.Package.IsValid)
	}
}

func TestUSPSIMpbPackageZeroCheckDigit(t *testing.T) {
	p, _ := NewPackage("9212391234567812345670")
	usps := NewUSPSIMpb(p)
	usps.Validate()
	if usps.Package.IsValid != true {
		t.Errorf("Failed, expected: %t, have: %t.", true, usps.Package.IsValid)
	}
}

func TestUSPSIMpbPackageIsValid(t *testing.T) {
	expected_value := "9101026837331000039521"
	p, _ := NewPackage("420221539101026837331000039521")
	usps := NewUSPSIMpb(p)
	usps.Validate()
	if usps.Package.TrackingNumber != expected_value {
		t.Errorf("Failed, expected: %v, have: %v.",
			expected_value, usps.Package.TrackingNumber)
	}
}

// Tests for usps s10
func TestUSPSS10PackageMatch(t *testing.T) {
	p, _ := NewPackage("EF123456785US")
	usps := NewUSPSS10(p)
	if usps.Match() != true {
		t.Errorf("Failed, expected: %t, have: %t.", true, usps.Match())
	}
}

func TestUSPSS10PackageWithNonValidChar(t *testing.T) {
	p, _ := NewPackage("EF12r456786US")
	usps := NewUSPSS10(p)
	usps.Validate()
	if usps.Package.IsValid != false {
		t.Errorf("Failed, expected: %t, have: %t.", false, usps.Package.IsValid)
	}
}

func TestUSPSS10PackageWithNonValidCheckDigit(t *testing.T) {
	p, _ := NewPackage("EF123456786US")
	usps := NewUSPSS10(p)
	usps.Validate()
	if usps.Package.IsValid != false {
		t.Errorf("Failed, expected: %t, have: %t.", false, usps.Package.IsValid)
	}
}
func TestUSPSS10PackageWithNonValidNumber(t *testing.T) {
	p, _ := NewPackage("EF12345678US") // 8 digit
	usps := NewUSPSS10(p)
	usps.Validate()
	if usps.Package.IsValid != false {
		t.Errorf("Failed, expected: %t, have: %t.", false, usps.Package.IsValid)
	}
}

func TestUSPSS10PackageValidCheckDigit(t *testing.T) {
	p, _ := NewPackage("RZ030057180PH")
	usps := NewUSPSS10(p)
	usps.Validate()
	if usps.Package.IsValid != true {
		t.Errorf("Failed, expected: %t, have: %t.", true, usps.Package.IsValid)
	}
}
func TestUSPSS10PackageValidCheckDigitRemainder11(t *testing.T) {
	p, _ := NewPackage("VA456789015KG")
	usps := NewUSPSS10(p)
	usps.Validate()
	if usps.Package.IsValid != true {
		t.Errorf("Failed, expected: %t, have: %t.", true, usps.Package.IsValid)
	}
}

func TestUSPSS10PackageIsValid(t *testing.T) {
	p, _ := NewPackage("EF123456785US")
	usps := NewUSPSS10(p)
	usps.Validate()
	if usps.Package.IsValid != true {
		t.Errorf("Failed, expected: %t, have: %t.", true, usps.Package.IsValid)
	}
}

// Tests for usps 20
func TestUSPS20PackageMatch(t *testing.T) {
	p, _ := NewPackage("71123456789123456787")
	usps := NewUSPS20(p)
	if usps.Match() != true {
		t.Errorf("Failed, expected: %t, have: %t.", true, usps.Match())
	}
}

func TestUSPS20PackageWithNonValidNumber(t *testing.T) {
	p, _ := NewPackage("72123456789123456787")
	usps := NewUSPS20(p)
	usps.Validate()
	if usps.Package.IsValid != false {
		t.Errorf("Failed, expected: %t, have: %t.", false, usps.Package.IsValid)
	}
}

func TestUSPS20PackageWithNonValidChar(t *testing.T) {
	p, _ := NewPackage("71123456r89123456787")
	usps := NewUSPS20(p)
	usps.Validate()
	if usps.Package.IsValid != false {
		t.Errorf("Failed, expected: %t, have: %t.", true, usps.Package.IsValid)
	}
}
func TestUSPS20PackageWrongCheckDigit(t *testing.T) {
	p, _ := NewPackage("7112345678912345678r")
	usps := NewUSPS20(p)
	usps.Validate()
	if usps.Package.IsValid != false {
		t.Errorf("Failed, expected: %t, have: %t.", true, usps.Package.IsValid)
	}
}

func TestUSPS20PackageIsValid(t *testing.T) {
	p, _ := NewPackage("71123456789123456787")
	usps := NewUSPS20(p)
	usps.Validate()
	if usps.Package.IsValid != true {
		t.Errorf("Failed, expected: %t, have: %t.", true, usps.Package.IsValid)
	}
}
