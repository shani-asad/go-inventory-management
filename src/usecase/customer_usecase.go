// usecase.go

package usecase

import (
	"context"
	"inventory-management/model/database"
	"inventory-management/model/dto"
	"inventory-management/src/repository"
	"time"
)

type CustomerUsecase struct {
	iCustomerRepository repository.CustomerRepositoryInterface
}

func NewCustomerUsecase(
	iCustomerRepository repository.CustomerRepositoryInterface) CustomerUsecaseInterface {
	return &CustomerUsecase{iCustomerRepository}
}

func (uc *CustomerUsecase) RegisterCustomer(request dto.RegisterCustomerRequest) (string, error) {
	data := database.Customer{
		PhoneNumber: request.PhoneNumber,
		Name: request.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	 id, err := uc.iCustomerRepository.RegisterCustomer(context.TODO(), data)
	 if err != nil {
			 return "", err
	 }

	return id, err
}

func (uc *CustomerUsecase) SearchCustomers(request dto.SearchCustomersRequest) ([]dto.CustomerDTO, error) {
    customers, err := uc.iCustomerRepository.SearchCustomers(context.TODO(), request)
    if err != nil {
        return nil, err
    }

    return customers, nil
}

func (u *CustomerUsecase) GetCustomerByPhoneNumber(phoneNumber string) (bool, error) {
	_, err := u.iCustomerRepository.GetCustomerByPhoneNumber(context.TODO(), phoneNumber)
	if err != nil {
		return false, err
	}
	return true, nil
}
