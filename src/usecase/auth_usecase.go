package usecase

import (
	"cats-social/helpers"
	"cats-social/model/database"
	"cats-social/model/dto"
	"cats-social/src/repository"
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	iUserRepository repository.UserRepositoryInterface
	helper          helpers.HelperInterface
}

func NewAuthUsecase(
	iUserRepository repository.UserRepositoryInterface,
	helper helpers.HelperInterface) AuthUsecaseInterface {
	return &AuthUsecase{iUserRepository, helper}
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

func (u *AuthUsecase) Login(request dto.RequestAuth) (token string, err error) {
	// check creds on database
	userData, err := u.iUserRepository.GetUserByEmail(context.TODO(), request.Email)
	if err != nil {
		return token, errors.New("wrong credentials")
	}

	fmt.Println(userData)

	// check the password
	isValid := u.verifyPassword(request.Password, userData.Password)
	if !isValid {
		return token, errors.New("wrong credentials")
	}

	token, _ = u.helper.GenerateToken(userData.Id)

	return token, nil
}

func (u *AuthUsecase) verifyPassword(password, passwordHash string) bool {
	byteHash := []byte(passwordHash)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(password))

	return err == nil
}
