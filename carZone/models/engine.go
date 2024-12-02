package models

import (
	"errors"

	"github.com/google/uuid"
)

type Engine struct {
	EngineID       uuid.UUID `json:"engine"`
	Displacement   int64     `json:"displacement"`
	NoOfCyclinders int64     `json:"noOfCyclinders"`
	CarRange       int64     `json:"carRange"`
}

type EngineRequest struct {
	Displacement   int64 `json:"displacement"`
	NoOfCyclinders int64 `json:"noOfCyclinders"`
	CarRange       int64 `json:"carRange"`
}

func ValidateEngineRequest(EngineReq EngineRequest) error {
	if err := ValidateDisplacement(EngineReq.Displacement); err != nil {
		return err
	}
	if err := ValidateNoOfCyclinders(EngineReq.NoOfCyclinders); err != nil {
		return err
	}
	if err := ValidateCarRange(EngineReq.CarRange); err != nil {
		return err
	}
	return nil
}

func ValidateDisplacement(displacement int64) error {
	if displacement <= 0 {
		return errors.New("displacement must be greater than Zero")
	}
	return nil
}

func ValidateNoOfCyclinders(NoOfCyclinders int64) error {
	if NoOfCyclinders <= 0 {
		return errors.New("NoOfCyclinders must be greater than Zero")
	}
	return nil
}

func ValidateCarRange(carRange int64) error {
	if carRange <= 0 {
		return errors.New("carRange must be greater than Zero")
	}
	return nil
}
