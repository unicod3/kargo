package kargo

type UPS struct {
	*Package
}

func (u *UPS) Validate(p *Package) bool {
	p.Carrier = "UPS"
	return false
}
