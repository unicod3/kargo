package kargo

import (
	"strconv"
)

type UPS struct {
	*Package
}

func (u *UPS) Validate(p *Package) bool {
	chars := p.TrackingNumber[2 : len(p.TrackingNumber)-1]
	checkDigit, err := strconv.Atoi(p.TrackingNumber[len(p.TrackingNumber)-1:])
	if err != nil {
		return false
	}

	var odd, even int = 0, 0
	for i, char := range chars {

		t := (string(char))
		num, err := strconv.Atoi(t)
		if err != nil {
			num = int(char-3) % 10
		}

		if i%2 == 0 {
			even += num
			continue
		}

		odd += num

	}

	check := ((odd * 2) + even) % 10
	if check != 0 {
		check = 10 - check
	}

	if check != checkDigit {
		return false
	}

	p.Carrier = "UPS"
	p.IsValid = true
	return true
}
