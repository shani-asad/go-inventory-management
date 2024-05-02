package usecase

import (
	"cats-social/model/database"
	"cats-social/model/dto"
	"cats-social/src/repository"
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	iUserRepository repository.UserRepositoryInterface
}

func NewAuthUsecase(iUserRepository repository.UserRepositoryInterface) AuthUsecaseInterface {
	return &AuthUsecase{iUserRepository}
}

func (u *AuthUsecase) Register(request dto.RequestCreateUser) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	data := database.User{
		Email:     request.Email,
		Password:  string(hash),
		Name:      request.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = u.iUserRepository.CreateUser(context.TODO(), data)
	return err
}

func (u *AuthUsecase) Login(request dto.RequestAuth) error {
	// check creds on database
	data, err := u.iUserRepository.GetUserByEmail(context.TODO(), request.Email)
	if err != nil {
		return err
	}

	fmt.Println(data)

	// check the password

	return nil
}
