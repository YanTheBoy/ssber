package app

import (
	"context"
	"offering-service/internal/model"
	otel2 "offering-service/internal/otel"
)

func (s *FlatsService) GetFlat(ctx context.Context, id string) (*model.Flat, error) {
	ctx, span := otel2.GetTracer().Start(ctx, "serviceGetFlat")
	defer span.End()

	selectedFlat, err := s.storage.GetFlat(ctx, id)
	if err != nil {
		return nil, err
	}

	return selectedFlat, nil

}
