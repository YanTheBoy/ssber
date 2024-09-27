package app

import (
	"context"
	"offering-service/internal/model"
	otel2 "offering-service/internal/otel"
)

func (s *FlatsService) CreateFlat(ctx context.Context, f model.Flat) (*model.Flat, error) {
	// TODO add validation

	ctx, span := otel2.GetTracer().Start(ctx, "serviceCreateFlat")
	defer span.End()

	createdFlat, err := s.storage.CreateFlat(ctx, f)
	if err != nil {
		return nil, err
	}

	return createdFlat, nil

}
