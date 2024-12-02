package models

import (
	"errors"

	"github.com/google/uuid"
)

type Engine struct {
	EngineID      uuid.UUID `json:"engine"`
	Displacement  int64     `json:"displacement"`
	NoOfCylinders int64     `json:"noOfcylinders"`
	CarRange      int64     `json:"carRange"`
}

type EngineRequest struct {
	Displacement  int64 `json:"displacement"`
	NoOfCylinders int64 `json:"noOfcylinders"`
	CarRange      int64 `json:"carRange"`
}

func ValidateEngineRequest(EngineReq EngineRequest) error {
	if err := ValidateDisplacement(EngineReq.Displacement); err != nil {
		return err
	}
	if err := ValidateNoOfCylinders(EngineReq.NoOfCylinders); err != nil {
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

func ValidateNoOfCylinders(NoOfCylinders int64) error {
	if NoOfCylinders <= 0 {
		return errors.New("NoOfCylinders must be greater than Zero")
	}
	return nil
}

func ValidateCarRange(carRange int64) error {
	if carRange <= 0 {
		return errors.New("carRange must be greater than Zero")
	}
	return nil
}
