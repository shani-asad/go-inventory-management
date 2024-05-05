package repository

import (
	"cats-social/model/database"
	"cats-social/model/dto"
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/lib/pq"
)

type CatRepository struct {
	db *sql.DB
}

func NewCatRepository(db *sql.DB) CatRepositoryInterface {
	return &CatRepository{db}
}

func (r *CatRepository) GetCatById(ctx context.Context ,id int) (err error) {
	query := `SELECT id FROM cats WHERE id = $1`

	_, err = r.db.QueryContext(ctx, query, id)
	fmt.Println(err)

	return err
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

func (r *CatRepository) GetCat(ctx context.Context, data dto.RequestGetCat) ([]dto.CatDetail, error) {
	// Initialize the base query
	query := "SELECT id, name, race, sex, age_in_month, description, image_urls, has_matched, created_at FROM cats WHERE 1=1"

	// Initialize the arguments slice
	args := make([]interface{}, 0)

	// Loop through the data array and append to the WHERE query from the queryParams
	if data.Id != "" {
			query += " AND id = $" + strconv.Itoa(len(args)+1)
			args = append(args, data.Id)
	}
	 
	if data.Sex != "" {
		if data.Sex == "male" || data.Sex == "female" {
			query += " AND sex = $" + strconv.Itoa(len(args)+1)
      args = append(args, data.Sex)
		}
  }
	if data.AgeInMonth != "" {
			// Parse the age filter string
			ageValue := data.AgeInMonth[1:]
			age, err := strconv.Atoi(ageValue)
			if err == nil {
				// Add the appropriate condition based on the prefix
					if strings.HasPrefix(data.AgeInMonth, "<") {
						// Less than condition
						query += " AND age_in_month < $" + strconv.Itoa(len(args)+1)
						args = append(args, age)
					} else if strings.HasPrefix(data.AgeInMonth, ">") {
							// Greater than condition
							query += " AND age_in_month > $" + strconv.Itoa(len(args)+1)
							args = append(args, age)
					} else {
							query += " AND age_in_month = $" + strconv.Itoa(len(args)+1)
							args = append(args, age)
					}
			}
  }
	if data.HasMatched {
		  query += " AND has_matched = $" + strconv.Itoa(len(args)+1)
      args = append(args, data.HasMatched)
  }
	if data.Owned {
		  query += " AND user_id = $" + strconv.Itoa(len(args)+1)
      args = append(args, data.Owned)
  }
	if data.Search != "" {
		  query += " AND (name LIKE '%' || $" + strconv.Itoa(len(args)+1) + " || '%' OR description LIKE '%' || $" + strconv.Itoa(len(args)+1) + " || '%')"
      args = append(args, data.Search)
  }

	// Add the ORDER BY clause to the query
	query += " ORDER BY created_at DESC"

	// Add the LIMIT clause if data.Limit is not nil
	if data.Limit != 0 {
			query += " LIMIT $" + strconv.Itoa(len(args)+1)
			args = append(args, data.Limit)
	}

	// Add the OFFSET clause if data.Offset is not nil
	if data.Offset != 0 {
			query += " OFFSET $" + strconv.Itoa(len(args)+1)
			args = append(args, data.Offset)
	}

	// Execute the query
	rows, err := r.db.QueryContext(ctx, query, args...)
	fmt.Println(err)
	if err != nil {
			return nil, err
	}
	defer rows.Close()

	// Initialize a slice to store the fetched cats
	var response []dto.CatDetail

	// Loop through the result set
	for rows.Next() {
			var cat dto.CatDetail
			// Scan each row into a Cat struct
			err := rows.Scan(&cat.Id, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.Description, pq.Array(&cat.ImageUrls), &cat.HasMatched, &cat.CreatedAt)
			fmt.Println(err)
			if err != nil {
					return nil, err
			}
			// Append the cat to the cats slice
			response = append(response, cat)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
			fmt.Println(err)
			return nil, err
	}

	// Return the fetched cats
	return response, err
}

func (r *CatRepository) UpdateCat(ctx context.Context, data database.Cat) (error) {
	query := `UPDATE cats SET name = $1, race = $2, sex = $3, age_in_month = $4, description = $5, image_urls = $6, updated_at = $7 WHERE id = $8`

	_, err := r.db.ExecContext(ctx, query, data.Name, data.Race, data.Sex, data.AgeInMonth, data.Description, pq.Array(data.ImageUrls), data.UpdatedAt, data.Id)

	if err != nil {
			return err
	}

	return nil
}

func (r *CatRepository) CheckHasMatch(ctx context.Context, id int) (bool, error) {
	query := `SELECT has_matched from cats WHERE id = $1`

    var hasMatched bool
    err := r.db.QueryRowContext(ctx, query, id).Scan(&hasMatched)

    if err != nil {
      return false, err
    }

    return hasMatched, nil
}

func (r *CatRepository) DeleteCat(ctx context.Context, id int) (error) {
	query := `DELETE FROM cats WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
			return err
	}

	return nil
}
