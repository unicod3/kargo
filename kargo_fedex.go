package kargo

type FedExGround96 struct {
	*Package
}

type FedExExpress struct {
	*Package
}

// NewFedExGround96 initializes a new FedExGround96 struct with package value
func NewFedExGround96(p *Package) *FedExGround96 {
	return &FedExGround96{Package: p}
}

// GetCarrierName Implements the CarrierFactory interface method
// Retuns the name of carrier
func (f *FedExGround96) GetCarrierName() string {
	return "FedEx"
}

// GetPackage Implements the CarrierFactory interface method
// Returns the package that carrier holds
func (f *FedExGround96) GetPackage() *Package {
	return f.Package
}

// Validate Implements the CarrierFactory interface method
// Checks whether is a package belongs to that carrier
func (f *FedExGround96) Validate() bool {
	//p.Carrier = "Fedex"
	return false
}

// NewFedExExpress initializes a new FedExExpress struct with package value
func NewFedExExpress(p *Package) *FedExExpress {
	return &FedExExpress{Package: p}
}

// GetCarrierName Implements the CarrierFactory interface method
// Retuns the name of carrier
func (f *FedExExpress) GetCarrierName() string {
	return "FedEx"
}

// GetPackage Implements the CarrierFactory interface method
// Returns the package that carrier holds
func (f *FedExExpress) GetPackage() *Package {
	return f.Package
}

// Validate Implements the CarrierFactory interface method
// Checks whether is a package belongs to that carrier
func (f *FedExExpress) Validate() bool {
	f.Package.Carrier = "Fedex"
	return false
}
