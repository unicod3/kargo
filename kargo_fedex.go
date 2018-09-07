package kargo

type FedExGround96 struct {
	*Package
}

type FedExExpress struct {
	*Package
}

func NewFedExGround96(p *Package) *FedExGround96 {
	return &FedExGround96{Package: p}
}

func (f *FedExGround96) GetPackage() *Package {
	return f.Package
}

func (f *FedExGround96) Validate() bool {
	//p.Carrier = "Fedex"
	return false
}

func NewFedExExpress(p *Package) *FedExExpress {
	return &FedExExpress{Package: p}
}

func (f *FedExExpress) GetPackage() *Package {
	return f.Package
}

func (f *FedExExpress) Validate() bool {
	f.Package.Carrier = "Fedex"
	return false
}
