package handlers

import (
	"context"
	"offering-service/internal/model"
)

type FlatsManagerService interface {
	CreateFlat(ctx context.Context, model model.Flat) (*model.Flat, error)
	GetFlat(ctx context.Context, id string) (*model.Flat, error)
}
