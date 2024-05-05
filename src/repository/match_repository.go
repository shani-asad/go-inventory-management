package repository

import (
	"cats-social/model/database"
	"cats-social/model/dto"
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type MatchRepository struct {
	db *sql.DB
}

func NewMatchRepository(db *sql.DB) MatchRepositoryInterface {
	return &MatchRepository{db}
}

func (r *MatchRepository) CreateMatch(ctx context.Context, data database.Match) (err error) {
	query := `
	INSERT INTO matches (match_cat_id, user_cat_id, message, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	_, err = r.db.ExecContext(
		ctx,
		query,
		data.MatchCatId,
		data.UserCatId,
		data.Message,
		data.CreatedAt,
		data.UpdatedAt,
	)

	return err
}

func (r *MatchRepository) GetMatch(ctx context.Context, userId int) (response []dto.ResponseGetMatch, err error) {
	query := `
	select 
		m.message AS message,
		m.created_at	AS matchCreatedAt,
		userCat.id	AS userCatId,
		userCat.name	AS userCatName,
		userCat.race	AS userCatRace,
		userCat.sex	AS userCatSex,
		userCat.description	AS userCatDescription,
		userCat.age_in_month	AS userCatAgeInMonth,
		userCat.image_urls	AS userCatImageUrls,
		userCat.created_at	AS userCatCreatedAt,
		matchCat.id	AS matchCatId,
		matchCat.name	AS matchCatName,
		matchCat.race	AS matchCatRace,
		matchCat.sex	AS matchCatSex,
		matchCat.description	AS matchCatDescription,
		matchCat.age_in_month	AS matchCatAgeInMonth,
		matchCat.image_urls	AS matchCatImageUrls,
		matchCat.created_at	AS matchCatCreatedAt,
		u.name	AS userName,
		u.email	AS userEmail,
		u.created_at AS userCreatedAt
	from matches m
	join cats userCat
		on m.user_cat_id = userCat.id
	join cats matchCat
		on m.match_cat_id = matchCat.id
	join users u
		on matchCat.user_id = u.id OR userCat.user_id = u.id
	where u.id = $1
	`

	rows, err := r.db.QueryContext(
		ctx,
		query,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var match dto.ResponseGetMatch
		err := rows.Scan(
			&match.Message,
			&match.CreatedAt,
			&match.UserCatDetail.Id,
			&match.UserCatDetail.Name,
			&match.UserCatDetail.Race,
			&match.UserCatDetail.Sex,
			&match.UserCatDetail.Description,
			&match.UserCatDetail.AgeInMonth,
			pq.Array(&match.UserCatDetail.ImageUrls),
			&match.UserCatDetail.CreatedAt,
			&match.MatchCatDetail.Id,
			&match.MatchCatDetail.Name,
			&match.MatchCatDetail.Race,
			&match.MatchCatDetail.Sex,
			&match.MatchCatDetail.Description,
			&match.MatchCatDetail.AgeInMonth,
			pq.Array(&match.MatchCatDetail.ImageUrls),
			&match.MatchCatDetail.CreatedAt,
			&match.IssuedBy.Name,
			&match.IssuedBy.Email,
			&match.IssuedBy.CreatedAt,
		)

		if err != nil {
			fmt.Println("XXXXXXXXXXXXXXXXXXXX",err)
			return nil, err
		}
		response = append(response, match)
	}

	return response, err
}

