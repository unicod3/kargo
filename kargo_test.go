package kargo

import (
	"testing"
)

func TestIdentify(t *testing.T) {
	trackingNumbers := map[string]string{
		"1Z999AA10123456784":                 "UPS",   // UPS
		"9632001960000000000400152152152158": "FedEx", // FedexExpress
		"9611019012345612345640":             "FedEx", // FedexGround96
		"420221539101026837331000039521":     "USPS",  // USPSIMpb
		"EF123456785US":                      "USPS",  // USPSS10
		"71123456789123456787":               "USPS",  // USPS20
		"":            "Unknown", // Unknown: Empty Tracking Number
		"bla bla bla": "Unknown", // Unknown: Some other value
	}

	for trackingNumber, carrier := range trackingNumbers {
		pkg, _ := Identify(trackingNumber)
		if pkg.Carrier != carrier {
			t.Errorf("Failed, expected: %v, got: %v.", carrier, pkg.TrackingNumber)
		}
	}

}

func TestKargoIdentifyNotValid(t *testing.T) {
	expected := false
	pkg, _ := Identify("1Z399AA10123456784")

	if pkg.IsValid != expected {
		t.Errorf("Failed, expected: %v, got: %v.", expected, pkg.TrackingNumber)
	}
}
