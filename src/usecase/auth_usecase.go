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
	helper helpers.HelperInterface) AuthUsecaseInterface {
	return &AuthUsecase{iStaffRepository, helper}
}

func (u *AuthUsecase) Register(request dto.RequestCreateStaff) (token string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	data := database.Staff{
		PhoneNumber:     request.PhoneNumber,
		Password:  string(hash),
		Name:      request.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	u.iStaffRepository.CreateStaff(context.TODO(), data)

	staffData, err := u.iStaffRepository.GetStaffByPhoneNumber(context.TODO(), request.PhoneNumber)

	fmt.Println("staffData:::::: ", staffData)

	token, _ = u.helper.GenerateToken(staffData.Id)

	return token, err
}

func (u *AuthUsecase) Login(request dto.RequestAuth) (token string, staff database.Staff, err error) {
	// check creds on database
	staffData, err := u.iStaffRepository.GetStaffByPhoneNumber(context.TODO(), request.PhoneNumber)
	if err != nil {
		return "", database.Staff{}, errors.New("staff not found")
	}

	fmt.Println(staffData)

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