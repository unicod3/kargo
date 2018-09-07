package kargo

import (
	"errors"
	"strings"
	"unicode"
)

// CarrierFactory handles the validation of tracking number for carrier
type CarrierFactory interface {
	Validate() bool
	GetPackage() *Package
}

// Package is the base struct of all carriers
type Package struct {
	Carrier        string
	TrackingNumber string
	IsValid        bool
}

func NewPackage(trackingNumber string) (*Package, error) {
	if len(trackingNumber) == 0 {
		return nil, errors.New("Tracking Number can not be empty!")
	}
	t := strings.TrimSpace(stripSpaces(trackingNumber))

	return &Package{TrackingNumber: t, IsValid: false}, nil
}

// Identify investigates the carriers and checks if
// the tracking number is valid for one of them
func Identify(trackingNumber string) (*Package, error) {
	if len(trackingNumber) == 0 {
		return nil, errors.New("Tracking Number can not be empty!")
	}

	t := strings.TrimSpace(stripSpaces(trackingNumber))

	p, _ := NewPackage(t)
	carriers := []CarrierFactory{NewUPS(p), NewFedExGround96(p), NewFedExExpress(p)}
	for _, carrier := range carriers {
		if carrier.Validate() {
			return carrier.GetPackage(), nil
		}
	}
	return p, nil
}

func stripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			// if the character is a space, drop it
			return -1
		}
		// else keep it in the string
		return r
	}, str)
}
