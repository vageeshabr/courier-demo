package models

type Order struct {
	Id             int
	ItemCount      int
	PackageDetails struct {
		SmallBags int
		LargeBags int
		Weight    float64
	}
	Address struct {
		L1      string
		L2      string
		State   string
		ZipCode string
		Country string
	}
}
