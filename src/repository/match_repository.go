package repository

import (
	"cats-social/model/database"
	"context"
	"database/sql"
	"cats-social/model/dto"
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

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (response database.User, err error) {
	err = r.db.QueryRowContext(ctx, "SELECT id, name, email, password FROM users WHERE email = $1", email).Scan(&response.Id, &response.Name, &response.Email, &response.Password)
	if err != nil {
		return
	}
	return
}

func (r *MatchRepository) GetMatch(ctx context.Context, userId string) (response dto.ResponseGetMatch, err error) {
	query := `
	select 
		m.message,
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
		u.created_at ASuserCreatedAt
	from matches m
	join cats userCat
		on m.user_cat_id = userCat.id
	join cats matchCat
		on m.match_cat_id = matchCat.id
	join users u
		on matchCat.user_id = u.id OR userCat.user_id = u.id
	where u.id = $1
	`

	/**
	INI CARA NYIMPAN KE ARRAY GIMANA CUY
	*/
	// _, err = r.db.QueryRowContext(
	// 	ctx,
	// 	query,
	// 	userId,
	// ).Scan()

	return 
}

