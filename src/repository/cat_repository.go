package repository

import (
	"cats-social/model/database"
	"context"
	"database/sql"
)

type CatRepository struct {
	db *sql.DB
}

func NewCatRepository(db *sql.DB) CatRepositoryInterface {
	return &CatRepository{db}
}

func (r *CatRepository) GetCatById(id int) (response database.Cat, err error) {
	return response, err
}

func (r *CatRepository) CreateCat(ctx context.Context, data database.Cat) (err error) {	
	query := `INSERT INTO cats (name, race, sex, age_in_month, description, image_urls, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	_, err = r.db.ExecContext(
		ctx,
		query,
		data.Name,
		data.Race,
		data.Sex,
		data.AgeInMonth,
		data.Description,
		data.ImageUrls,
		data.CreatedAt,
		data.UpdatedAt,
	)

	return err
}
