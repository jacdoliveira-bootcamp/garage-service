package internal

type VehicleRepository interface {
	FindAll() (v map[int]Vehicle, err error)
	Create(v Vehicle) error
}
