package service

import (
	"app/internal"
	"fmt"
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
	all, err := s.rp.FindAll()
	if err != nil {
		return nil, err
	}
	var result []internal.Vehicle
	for _, v := range all {
		if v.Color == color && v.FabricationYear == year {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("vehicle not found")
	}
	return result, nil
}

func (s *VehicleDefault) Delete(id int) error {

	return s.rp.Delete(id)
}

func (s *VehicleDefault) UpdateSpeed(id int, speed float64) error {

	return s.rp.UpdateSpeed(id, speed)
}
