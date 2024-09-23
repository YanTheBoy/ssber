package db

import (
	"context"
	"offering-service/internal/model"
)

func (s *FlatsStorage) CreateFlat(ctx context.Context, f model.Flat) (*model.Flat, error) {
	query := `INSERT INTO offering 
    	(name, description, price, created_by) VALUES ($1, $2, $3, $4) RETURNING id,name,description,price,created_at`
	var createdFlat model.Flat
	err := s.db.QueryRowContext(ctx, query, f.Name, f.Description, f.Price, f.CreatedBy).Scan(&createdFlat.ID,
		&createdFlat.Name,&createdFlat.Description,&createdFlat.Price,&createdFlat.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &model.Flat{
		ID:   createdFlat.ID,
		Name: createdFlat.Name,
		Description: createdFlat.Description,
		Price: createdFlat.Price,
		CreatedAt: createdFlat.CreatedAt,
	}, err
}

func (s *FlatsStorage) GetFlat(ctx context.Context, id string) (*model.Flat, error) {
	var flat model.Flat
	sql := `SELECT id, name, description, price, created_by, created_at FROM offering WHERE id = $1`
	err := s.db.QueryRowContext(ctx, sql, id).Scan(&flat)
	if err != nil {
		return nil, err
	}

	return &flat, nil
}
