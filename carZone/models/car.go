package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Year      string    `json:"year"`
	Brand     string    `json:"brand"`
	FuelType  string    `json:"fueltype"`
	Engine    Engine    `json:"engine"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CarRequest struct {
	Name     string  `json:"name"`
	Year     string  `json:"year"`
	Brand    string  `json:"brand"`
	FuelType string  `json:"fueltype"`
	Engine   Engine  `json:"engine"`
	Price    float64 `json:"price"`
}

func ValidateRequest(carReq CarRequest) error {
	if err := ValidateName(carReq.Name); err != nil {
		return err
	}
	if err := ValidateYear(carReq.Year); err != nil {
		return err
	}
	if err := ValidateBrand(carReq.Brand); err != nil {
		return err
	}
	if err := ValidateFuelType(carReq.FuelType); err != nil {
		return err
	}
	if err := ValidateEngine(carReq.Engine); err != nil {
		return err
	}
	if err := ValidatePrice(carReq.Price); err != nil {
		return err
	}
	return nil
}

func ValidateName(name string) error {
	if name == "" {
		return errors.New("Name is Required")
	}
	return nil
}

func ValidateYear(year string) error {
	if year == "" {
		return errors.New("Year is Required")
	}
	_, err := strconv.Atoi(year)
	if err != nil {
		return errors.New("year must be a valid number")
	}
	currentYear := time.Now().Year()
	yearInt, _ := strconv.Atoi(year)
	if yearInt <= 1886 || yearInt > currentYear {
		return errors.New("Year must be between 1886 and current year")
	}
	return nil
}

func ValidateBrand(brand string) error {
	if brand == "" {
		return errors.New("Brand is Required")
	}
	return nil
}

func ValidateFuelType(fuelType string) error {
	validateFuelType := []string{"Petrol", "Diesel", "Electric", "Hybrid"}
	for _, validType := range validateFuelType {
		if fuelType == validType {
			return nil
		}
	}
	return errors.New("FuelType must be one of: Petrol, Diesel, Electric, Hybrid")
}

func ValidateEngine(engine Engine) error {
	if engine.EngineID == uuid.Nil {
		return errors.New("EngineId is Required")
	}
	if engine.Displacement <= 0 {
		return errors.New("displacement must be greater than Zero")
	}
	if engine.NoOfCylinders <= 0 {
		return errors.New("NoOfCylinders must be greater than Zero")
	}
	if engine.CarRange <= 0 {
		return errors.New("CarRange must be greater than Zero")
	}
	return nil
}

func ValidatePrice(price float64) error {
	if price <= 0 {
		return errors.New("CarRange must be greater than Zero")
	}
	return nil
}
