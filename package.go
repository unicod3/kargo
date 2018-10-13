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
	p := &Package{TrackingNumber: "", Carrier: "Unknown", IsValid: false}
	if len(trackingNumber) == 0 {
		return p, errors.New("Tracking Number can not be empty!")
	}
	t := strings.TrimSpace(stripSpaces(trackingNumber))
	p.TrackingNumber = t
	return p, nil
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

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
