package kargoTest

import (
	"kargo"
	"testing"
)

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
