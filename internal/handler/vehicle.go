package handler

import (
	"app/internal"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
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
			if strings.Contains(err.Error(), "identifier of the existing vehicle") {
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
			"message": "vehicle created successfully",
		})
	}
}

func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		color := chi.URLParam(r, "color")
		yearStr := chi.URLParam(r, "year")

		year, err := strconv.Atoi(yearStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "invalid formated year"})
			return
		}

		vehicles, err := h.sv.FindByColorAndYear(color, year)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]string{"error": "vehicle not found"})
		}

		var data []VehicleJSON
		for _, value := range vehicles {
			data = append(data, VehicleJSON{
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
			})
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) DeleteById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": "invalid ID",
			})
			return
		}
		err = h.sv.Delete(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"error": err.Error(),
			})
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func (h *VehicleDefault) PutUpdateSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": "invalid ID",
			})
			return
		}

		var body struct {
			MaxSpeed float64 `json:"max_speed"`
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
			return
		}

		err = h.sv.UpdateSpeed(id, body.MaxSpeed)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "max speed update sucessfully",
		})

	}
}

func (h *VehicleDefault) UpdateFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": "invalid ID",
			})
			return
		}

		var body struct {
			FuelType string `json:"fuel_type"`
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
			return
		}

		err = h.sv.UpdateFuelType(id, body.FuelType)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
			return
		}
		response.JSON(w, http.StatusOK, map[string]string{
			"message": "sucess update fueltype",
		})
	}
}

func (h *VehicleDefault) GetByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fuelType := chi.URLParam(r, "type")

		vehicles, err := h.sv.FindByFuelType(fuelType)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]string{"error": "vehicle not found"})
			return
		}

		var data []VehicleJSON
		for _, value := range vehicles {
			data = append(data, VehicleJSON{
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
			})
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) GetByTransmissionType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transmission := chi.URLParam(r, "type")

		vehicles, err := h.sv.FindByTransmissionType(transmission)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]string{"error": "vehicle not found"})
			return
		}

		var data []VehicleJSON
		for _, value := range vehicles {
			data = append(data, VehicleJSON{
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
			})
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) PostCreateBatch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req []VehicleJSON
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"error": "invalid JSON",
			})
			return
		}
		var vehicles []internal.Vehicle
		for _, item := range req {
			vehicles = append(vehicles, internal.Vehicle{
				Id: item.ID,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           item.Brand,
					Model:           item.Model,
					Registration:    item.Registration,
					Color:           item.Color,
					FabricationYear: item.FabricationYear,
					Capacity:        item.Capacity,
					MaxSpeed:        item.MaxSpeed,
					FuelType:        item.FuelType,
					Transmission:    item.Transmission,
					Weight:          item.Weight,
					Dimensions: internal.Dimensions{
						Height: item.Height,
						Length: item.Length,
						Width:  item.Width,
					},
				},
			})
		}
		err := h.sv.CreateBatch(vehicles)
		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				response.JSON(w, http.StatusConflict, map[string]string{
					"error": err.Error(),
				})
			} else {
				response.JSON(w, http.StatusBadRequest, map[string]string{
					"error": err.Error(),
				})
			}
			return
		}

		response.JSON(w, http.StatusCreated, map[string]string{
			"message": "vehicles created sucessfuly",
		})
	}

}

func (h *VehicleDefault) GetByBrandAndBetweenYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")
		startStr := chi.URLParam(r, "start_year")
		endStr := chi.URLParam(r, "end_year")

		start, err := strconv.Atoi(startStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"error": "invalid start year",
			})
		}

		end, err := strconv.Atoi(endStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"error": "invalid end year",
			})
		}

		vehicles, err := h.sv.FindByBrandAndBetweenYear(brand, start, end)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
			return
		}

		var data []VehicleJSON
		for _, v := range vehicles {
			data = append(data, VehicleJSON{
				ID:              v.Id,
				Brand:           v.Brand,
				Model:           v.Model,
				Registration:    v.Registration,
				Color:           v.Color,
				FabricationYear: v.FabricationYear,
				Capacity:        v.Capacity,
				MaxSpeed:        v.MaxSpeed,
				FuelType:        v.FuelType,
				Transmission:    v.Transmission,
				Weight:          v.Weight,
				Height:          v.Dimensions.Height,
				Length:          v.Dimensions.Length,
				Width:           v.Dimensions.Width,
			})
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "sucess",
			"data":    data,
		})

	}
}

func (h *VehicleDefault) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"error": "invalid ID",
			})
		}
		vehicles, err := h.sv.FindById(id)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}
		var data []VehicleJSON
		for _, v := range vehicles {
			data = append(data, VehicleJSON{
				ID:              v.Id,
				Brand:           v.Brand,
				Model:           v.Model,
				Registration:    v.Registration,
				Color:           v.Color,
				FabricationYear: v.FabricationYear,
				Capacity:        v.Capacity,
				MaxSpeed:        v.MaxSpeed,
				FuelType:        v.FuelType,
				Transmission:    v.Transmission,
				Weight:          v.Weight,
				Height:          v.Dimensions.Height,
				Length:          v.Dimensions.Length,
				Width:           v.Dimensions.Width,
			})
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "sucess",
			"data":    data,
		})

	}
}

func (h *VehicleDefault) GetByBrandAverageSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")

		avg, err := h.sv.FindByBrandAverageSpeed(brand)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message":           "sucess",
			"avarage_max_speed": avg,
		})
	}
}

func (h *VehicleDefault) GetByBrandAverageCapacity() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")

		avg, err := h.sv.FindByBrandAverageCapacity(brand)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message":              "sucess",
			"avarage_max_capacity": avg,
		})
	}
}

func (h *VehicleDefault) GetByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// length=10-200
		length := strings.Split(r.URL.Query().Get("length"), "-")
		width := strings.Split(r.URL.Query().Get("width"), "-")

		if len(length) != 2 || len(width) != 2 {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "invalid format"})
			return
		}

		lengthMin, _ := strconv.ParseFloat(length[0], 64)
		lengthMax, _ := strconv.ParseFloat(length[1], 64)
		widthMin, _ := strconv.ParseFloat(width[0], 64)
		widthMax, _ := strconv.ParseFloat(width[1], 64)

		vehicles, err := h.sv.FindByDimensions(lengthMin, lengthMax, widthMin, widthMax)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}

		var data []VehicleJSON
		for _, v := range vehicles {
			data = append(data, VehicleJSON{
				ID:              v.Id,
				Brand:           v.Brand,
				Model:           v.Model,
				Registration:    v.Registration,
				Color:           v.Color,
				FabricationYear: v.FabricationYear,
				Capacity:        v.Capacity,
				MaxSpeed:        v.MaxSpeed,
				FuelType:        v.FuelType,
				Transmission:    v.Transmission,
				Weight:          v.Weight,
				Height:          v.Height,
				Length:          v.Length,
				Width:           v.Width,
			})
		}

		response.JSON(w, http.StatusOK, map[string]any{"message": "success", "data": data})
	}
}

func (h *VehicleDefault) GetByWeightRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		minStr := r.URL.Query().Get("min")
		maxStr := r.URL.Query().Get("max")

		min, _ := strconv.ParseFloat(minStr, 64)
		max, _ := strconv.ParseFloat(maxStr, 64)

		vehicles, err := h.sv.FindByWeight(min, max)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}

		var data []VehicleJSON
		for _, v := range vehicles {
			data = append(data, VehicleJSON{
				ID:              v.Id,
				Brand:           v.Brand,
				Model:           v.Model,
				Registration:    v.Registration,
				Color:           v.Color,
				FabricationYear: v.FabricationYear,
				Capacity:        v.Capacity,
				MaxSpeed:        v.MaxSpeed,
				FuelType:        v.FuelType,
				Transmission:    v.Transmission,
				Weight:          v.Weight,
				Height:          v.Height,
				Length:          v.Length,
				Width:           v.Width,
			})
		}

		response.JSON(w, http.StatusOK, map[string]any{"message": "success", "data": data})
	}
}

func (h *VehicleDefault) GetByColor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		color := chi.URLParam(r, "color")

		vehicles, err := h.sv.FindByColor(color)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]string{"error": "vehicle not found"})
		}

		var data []VehicleJSON
		for _, value := range vehicles {
			data = append(data, VehicleJSON{
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
			})
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})

	}
}
