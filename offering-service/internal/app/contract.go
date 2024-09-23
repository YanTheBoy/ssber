package app

import (
	"context"
	"offering-service/internal/model"
)

type FlatsStorage interface {
	CreateFlat(ctx context.Context, f model.Flat) (*model.Flat, error)
	GetFlat(ctx context.Context, id string) (*model.Flat, error)
}
