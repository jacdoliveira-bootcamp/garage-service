package internal

type VehicleService interface {
	FindAll() (v map[int]Vehicle, err error)
	Create(v Vehicle) error
	FindByColorAndYear(color string, year int) ([]Vehicle, error)
	Delete(id int) error
	UpdateSpeed(id int, speed float64) error
}
