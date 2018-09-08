package kargo

import (
	"errors"
	"strings"
	"unicode"
)

// CarrierFactory handles the validation of tracking number for carrier
type CarrierFactory interface {
	GetCarrierName() string
	GetPackage() *Package
	Match() bool
	Validate() bool
}

// Package is the base struct of all carriers
type Package struct {
	Carrier        string
	TrackingNumber string
	IsValid        bool
}

// NewPackage initializes a new Package struct with a Tracking Number value
func NewPackage(trackingNumber string) (*Package, error) {
	if len(trackingNumber) == 0 {
		return nil, errors.New("Tracking Number can not be empty!")
	}
	t := strings.TrimSpace(stripSpaces(trackingNumber))

	return &Package{TrackingNumber: t, Carrier: "Unknown", IsValid: false}, nil
}

//stripSpaces removes spaces from given string
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
