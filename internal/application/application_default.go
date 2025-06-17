package application

import (
	"app/internal/handler"
	"app/internal/loader"
	"app/internal/repository"
	"app/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
	// LoaderFilePath is the path to the file that contains the vehicles
	LoaderFilePath string
}

// NewServerChi is a function that returns a new instance of ServerChi
func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	// default values
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = cfg.LoaderFilePath
		}
	}

	return &ServerChi{
		serverAddress:  defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePath is the path to the file that contains the vehicles
	loaderFilePath string
}

// Run is a method that runs the application
func (a *ServerChi) Run() (err error) {
	// dependencies
	// - loader
	ld := loader.NewVehicleJSONFile(a.loaderFilePath)
	db, err := ld.Load()
	if err != nil {
		return
	}
	// - repository
	rp := repository.NewVehicleMap(db)
	// - service
	sv := service.NewVehicleDefault(rp)
	// - handler
	hd := handler.NewVehicleDefault(sv)
	// router
	rt := chi.NewRouter()
	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)
	// - endpoints
	rt.Route("/vehicles", func(rt chi.Router) {
		rt.Get("/", hd.GetAll())
		rt.Post("/", hd.PostCreate())
		rt.Get("/color/{color}/year/{year}", hd.GetByColorAndYear())
		rt.Delete("/{id}", hd.DeleteById())
		rt.Put("/{id}/update_speed", hd.PutUpdateSpeed())
		rt.Put("/{id}/update_fuel", hd.UpdateFuelType())
		rt.Get("/fuel_type/{type}", hd.GetByFuelType())
		rt.Get("/transmission/{type}", hd.GetByTransmissionType())
		rt.Post("/batch", hd.PostCreateBatch())
		rt.Get("/brand/{brand}/between/{start_year}/{end_year}", hd.GetByBrandAndBetweenYear())
		rt.Get("/id/{id}", hd.GetById())
		rt.Get("/avarage_speed/brand/{brand}", hd.GetByBrandAverageSpeed())
		rt.Get("/avarage_capacity/brand/{brand}", hd.GetByBrandAverageCapacity())

	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
