package usecase

import (
	"context"
	"inventory-management/model/dto"
	"inventory-management/src/repository"
)

type TransactionUsecase struct {
	iTransactionRepository repository.TransactionRepositoryInterface 
}

func NewTransactionUsecase(
		iTransactionRepository repository.TransactionRepositoryInterface,
	) TransactionUsecaseInterface {
	return &TransactionUsecase{iTransactionRepository}
}

func (u *TransactionUsecase) GetTransactions(request dto.GetTransactionRequest) (transactions []dto.TransactionData, err error) {
	params := dto.GetTransactionRequest{
		Limit:    validateLimit(request.Limit),
		Offset:   validateOffset(request.Offset),
		CustomerId: request.CustomerId,
		CreatedAt: request.CreatedAt,
	}

	response, err := u.iTransactionRepository.GetTransactions(context.TODO(), params)
	
	
	return response, err

}
