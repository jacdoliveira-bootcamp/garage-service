package internal

type Dimensions struct {
	Height float64
	Length float64
	Width  float64
}

type VehicleAttributes struct {
	Brand           string
	Model           string
	Registration    string
	Color           string
	FabricationYear int
	Capacity        int
	MaxSpeed        float64
	FuelType        string
	Transmission    string
	Weight          float64
	Dimensions
}

type Vehicle struct {
	Id int
	VehicleAttributes
}
