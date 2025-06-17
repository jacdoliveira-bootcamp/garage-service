package repository

import (
	"app/internal"
	"fmt"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

func (r *VehicleMap) Create(v internal.Vehicle) error {
	if _, exists := r.db[v.Id]; exists {
		return fmt.Errorf("vehicle with ID: %v, already exists", v.Id)
	}
	r.db[v.Id] = v
	return nil
}

func (r *VehicleMap) Delete(id int) error {
	if _, exists := r.db[id]; !exists {
		return fmt.Errorf("vehicle with ID: %v, not found", id)
	}
	delete(r.db, id)
	return nil
}

func (r *VehicleMap) UpdateSpeed(id int, speed float64) error {
	v, exists := r.db[id]
	if !exists {
		return fmt.Errorf("vehicle with ID: %v, does not found", id)
	}

	v.MaxSpeed = speed
	r.db[id] = v
	return nil
}

func (r *VehicleMap) UpdateFuelType(id int, fuelType string) error {
	v, exists := r.db[id]
	if !exists {
		return fmt.Errorf("vehicle with ID: %v, does not found", id)
	}

	v.FuelType = fuelType
	r.db[id] = v

	return nil
}

func (r *VehicleMap) FindByFuelType(fuelType string) ([]internal.Vehicle, error) {
	var result []internal.Vehicle
	for _, v := range r.db {
		if v.FuelType == fuelType {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("vehicle not found")
	}

	return result, nil
}

func (r *VehicleMap) FindByTransmissionType(transmission string) ([]internal.Vehicle, error) {
	var result []internal.Vehicle
	for _, v := range r.db {
		if v.Transmission == transmission {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("vehicle not found")
	}

	return result, nil
}

func (r *VehicleMap) FindByColorAndYear(color string, year int) ([]internal.Vehicle, error) {
	var result []internal.Vehicle

	for _, v := range r.db {
		if v.Color == color && v.FabricationYear == int(year) {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("vehicle not found")
	}

	return result, nil
}

func (r *VehicleMap) CreateBatch(vehicles []internal.Vehicle) error {
	for _, v := range vehicles {
		if _, exists := r.db[v.Id]; exists {
			return fmt.Errorf("vehicle with ID: %v already exists", v.Id)
		}
	}
	for _, v := range vehicles {
		r.db[v.Id] = v
	}
	return nil
}

func (r *VehicleMap) FindByBrandAndBetweenYear(brand string, start, end int) ([]internal.Vehicle, error) {
	var result []internal.Vehicle

	for _, v := range r.db {
		if v.Brand == brand && v.FabricationYear >= start && v.FabricationYear <= end {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("vehilce not found")
	}
	return result, nil
}
