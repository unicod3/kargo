package kargo

// Identify investigates the carriers and checks if
// the tracking number is valid for one of them
func Identify(trackingNumber string) (*Package, error) {
	p, err := NewPackage(trackingNumber)
	if err != nil {
		return nil, err
	}

	carriers := []CarrierFactory{NewUPS(p), NewFedExGround96(p), NewFedExExpress(p)}
	for _, carrier := range carriers {
		if carrier.Validate() {
			return carrier.GetPackage(), nil
		}
	}
	return p, nil
}
