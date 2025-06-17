package service

import (
	"app/internal"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) Create(v internal.Vehicle) error {
	return s.rp.Create(v)
}

func (s *VehicleDefault) FindByColorAndYear(color string, year int) ([]internal.Vehicle, error) {

	return s.rp.FindByColorAndYear(color, year)
}

func (s *VehicleDefault) Delete(id int) error {

	return s.rp.Delete(id)
}

func (s *VehicleDefault) UpdateSpeed(id int, speed float64) error {

	return s.rp.UpdateSpeed(id, speed)
}

func (s *VehicleDefault) UpdateFuelType(id int, fuelType string) error {

	return s.rp.UpdateFuelType(id, fuelType)
}

func (s *VehicleDefault) FindByFuelType(fuelType string) ([]internal.Vehicle, error) {

	return s.rp.FindByFuelType(fuelType)
}

func (s *VehicleDefault) FindByTransmissionType(transmission string) ([]internal.Vehicle, error) {

	return s.rp.FindByTransmissionType(transmission)
}

func (s *VehicleDefault) CreateBatch(vehicles []internal.Vehicle) error {

	return s.rp.CreateBatch(vehicles)
}

func (s *VehicleDefault) FindByBrandAndBetweenYear(brand string, start, end int) ([]internal.Vehicle, error) {

	return s.rp.FindByBrandAndBetweenYear(brand, start, end)

}

func (s *VehicleDefault) FindById(id int) ([]internal.Vehicle, error) {

	return s.rp.FindById(id)

}

func (s *VehicleDefault) FindByBrandAverageSpeed(brand string) (float64, error) {
	return s.rp.FindByBrandAverageSpeed(brand)
}

func (s *VehicleDefault) FindByBrandAverageCapacity(brand string) (int, error) {
	return s.rp.FindByBrandAverageCapacity(brand)
}

func (s *VehicleDefault) FindByDimensions(lengthMin, lengthMax, widthMin, widthMax float64) ([]internal.Vehicle, error) {
	return s.rp.FindByDimensions(lengthMin, lengthMax, widthMin, widthMax)
}

func (s *VehicleDefault) FindByWeight(min, max float64) ([]internal.Vehicle, error) {
	return s.rp.FindByWeight(min, max)
}
