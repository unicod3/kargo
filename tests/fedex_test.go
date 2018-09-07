package kargoTest

import (
	"kargo"
	"testing"
)

func TestFedExPackageIsValid(t *testing.T) {
	p, _ := kargo.NewPackage("9632001960000000000400152152152158")
	fexpress := kargo.NewFedExExpress(p)
	fexpress.Validate()
	if fexpress.Package.IsValid != true {
		t.Errorf("Failed, expected: %t, got: %t.", true, fexpress.Package.IsValid)
	}
}
func TestFedExExpressCarrier(t *testing.T) {
	p, _ := kargo.NewPackage("9632001960000000000400152152152158")
	fexpress := kargo.NewFedExExpress(p)
	fexpress.Validate()
	if fexpress.Package.Carrier != fexpress.GetCarrierName() {
		t.Errorf("Failed, expected: %v, got: %v.", fexpress.GetCarrierName(), fexpress.Package.Carrier)
	}
}
