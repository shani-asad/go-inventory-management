package repository

import (
	"cats-social/model/database"
	"context"
	"database/sql"

	"github.com/lib/pq"
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

func (r *CatRepository) CreateCat(ctx context.Context, data database.Cat) (int64, error) {
	var id int64

	query := `INSERT INTO cats (name, user_id, race, sex, age_in_month, description, image_urls, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	err := r.db.QueryRowContext(
			ctx,
			query,
			data.Name,
			data.UserId,
			data.Race,
			data.Sex,
			data.AgeInMonth,
			data.Description,
			pq.Array(data.ImageUrls),
			data.CreatedAt,
			data.UpdatedAt,
	).Scan(&id)

	if err != nil {
			return 0, err
	}

	return id, nil
}

