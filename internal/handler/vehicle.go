package handler

import (
	"app/internal"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bootcamp-go/web/response"
)

type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

type VehicleDefault struct {
	sv internal.VehicleService
}

func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) PostCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req VehicleJSON

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
			return
		}
		v := internal.Vehicle{
			Id: req.ID,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           req.Brand,
				Model:           req.Model,
				Registration:    req.Registration,
				Color:           req.Color,
				FabricationYear: req.FabricationYear,
				Capacity:        req.Capacity,
				MaxSpeed:        req.MaxSpeed,
				FuelType:        req.FuelType,
				Transmission:    req.Transmission,
				Weight:          req.Weight,
				Dimensions: internal.Dimensions{
					Height: req.Height,
					Length: req.Length,
					Width:  req.Width,
				},
			},
		}

		err := h.sv.Create(v)
		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				response.JSON(w, http.StatusConflict, map[string]string{
					"error": err.Error(),
				})
				return
			}
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}
		response.JSON(w, http.StatusCreated, map[string]string{
			"message": "vehicle created",
		})
	}
}
