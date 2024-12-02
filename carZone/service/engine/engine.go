package service

import (
	"context"

	"github.com/Quanghh2233/MyApp/models"
	"github.com/Quanghh2233/MyApp/store"
)

type EngineService struct {
	store store.EngineStoreInterface
}

func NewEngineService(store store.EngineStoreInterface) *EngineService {
	return &EngineService{
		store: store,
	}
}

func (s *EngineService) GetEngineByID(ctx context.Context, id string) (*models.Engine, error) {
	engine, err := s.store.EngineByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &engine, nil
}

func (s *EngineService) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (*models.Engine, error) {
	if err := models.ValidateEngineRequest(*engineReq); err != nil {
		return nil, err
	}

	createdEngine, err := s.store.EngineCreated(ctx, engineReq)
	if err != nil {
		return nil, err
	}
	return &createdEngine, nil
}

func (s *EngineService) UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (*models.Engine, error) {
	if err := models.ValidateEngineRequest(*engineReq); err != nil {
		return nil, err
	}

	updatedEngine, err := s.store.EngineUpdate(ctx, id, engineReq)
	if err != nil {
		return nil, err
	}
	return &updatedEngine, nil
}

func (s *EngineService) DeleteEngine(ctx context.Context, id string) (*models.Engine, error) {
	deletedEngine, err := s.store.DeleteEngine(ctx, id)
	if err != nil {
		return nil, err
	}
	return &deletedEngine, nil
}