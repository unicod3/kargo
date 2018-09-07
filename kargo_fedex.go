package kargo

type FedEx struct {
	*Package
}

func (f *FedEx) Validate(p *Package) bool {
	//p.Carrier = "Fedex"
	return false
}
