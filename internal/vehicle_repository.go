package internal

type VehicleRepository interface {
	FindAll() (v map[int]Vehicle, err error)
	Create(v Vehicle) error
	FindByColorAndYear(color string, year int) ([]Vehicle, error)
	Delete(id int) error
	UpdateSpeed(id int, speed float64) error
	UpdateFuelType(id int, fuelType string) error
	FindByFuelType(fuelType string) ([]Vehicle, error)
	FindByTransmissionType(transmission string) ([]Vehicle, error)
	CreateBatch([]Vehicle) error
}
