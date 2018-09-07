package kargo

import "errors"

// CarrierFactory handles the validation of tracking number for carrier
type CarrierFactory interface {
	Validate(*Package) bool
}

// Package is the base struct of all carriers
type Package struct {
	Carrier        string
	TrackingNumber string
}

// Identify investigates the carriers and checks if
// the tracking number is valid for one of them
func Identify(trackingNumber string) (*Package, error) {
	if len(trackingNumber) == 0 {
		return nil, errors.New("Tracking Number can not be empty!")
	}

	p := &Package{TrackingNumber: trackingNumber}
	carriers := []CarrierFactory{&UPS{}, &FedEx{}}
	for _, carrier := range carriers {
		if carrier.Validate(p) {
			return p, nil
		}
	}
	return nil, errors.New("Tracking number could not be identified!")
}
