package usecase

import (
	"context"
	"errors"
	"fmt"
	"inventory-management/helpers"
	"inventory-management/model/database"
	"inventory-management/model/dto"
	"inventory-management/src/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	iStaffRepository repository.StaffRepositoryInterface
	helper          helpers.HelperInterface
}

func NewAuthUsecase(
		iStaffRepository repository.StaffRepositoryInterface,
		helper helpers.HelperInterface,
	) AuthUsecaseInterface {
	return &AuthUsecase{iStaffRepository, helper}
}

func (u *AuthUsecase) Register(request dto.RequestCreateStaff) (token string, staffId int, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		fmt.Println("@ auth_usecase.go > Register >> GenerateFromPassword:", err)
		return "", 0, err
	}

	data := database.Staff{
		PhoneNumber: request.PhoneNumber,
		Password:    string(hash),
		Name:        request.Name,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	staffId, err = u.iStaffRepository.CreateStaff(context.TODO(), data)
	if err != nil {
		fmt.Println("@ auth_usecase.go > Register >> CreateStaff:", err)
		return "", 0, err
	}

	token, _ = u.helper.GenerateToken(staffId)

	return token, staffId, err
}

func (u *AuthUsecase) Login(request dto.RequestAuth) (token string, staff database.Staff, err error) {
	// check creds on database
	staffData, err := u.iStaffRepository.GetStaffByPhoneNumber(context.TODO(), request.PhoneNumber)
	if err != nil {
		return "", database.Staff{}, errors.New("staff not found")
	}

	// check the password
	isValid := u.verifyPassword(request.Password, staffData.Password)
	if !isValid {
		return "", database.Staff{}, errors.New("wrong password")
	}

	token, _ = u.helper.GenerateToken(staffData.Id)

	return token, staffData, nil
}

func (u *AuthUsecase) verifyPassword(password, passwordHash string) bool {
	byteHash := []byte(passwordHash)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(password))

	return err == nil
}

func (u *AuthUsecase) GetStaffByPhoneNumber(phoneNumber string) (bool, error) {
	_, err := u.iStaffRepository.GetStaffByPhoneNumber(context.TODO(), phoneNumber)
	if err != nil {
		return false, err
	}
	return true, nil
}
