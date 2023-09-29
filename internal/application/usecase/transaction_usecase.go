package usecase

import (
	"context"

	"github.com/ruhanc/pix-go/internal/domain/entity"
	"github.com/ruhanc/pix-go/internal/domain/gateway"
)

type TransactionUseCase struct {
	TransactionRepository gateway.TransactionRepositoryInterface
	PixRepository         gateway.PixKeyRepositoryInterface
}

func (usecase *TransactionUseCase) RegisterTransaction(ctx context.Context, accountID, pixKeyTo, Description string, amount float64) (*entity.Transaction,error) {
	account,err := usecase.PixRepository.FindAccount(ctx,accountID)
	if err != nil {
		return nil,err
	}
	
	pixKey,err := usecase.PixRepository.FindKeyByID(ctx,pixKeyTo)
	if err != nil {
		return nil,err
	}

	transaction,err := entity.NewTransaction(account,amount,pixKey,Description)
	if err != nil {
		return nil,err
	}

	err = usecase.TransactionRepository.Register(ctx,transaction)
	if err != nil {
		return nil,err
	}

	return transaction,nil
}

func (usecase *TransactionUseCase) ConfirmTransaction(ctx context.Context, transactionID string) (*entity.Transaction,error) {
	transaction,err := usecase.TransactionRepository.Find(ctx,transactionID)
	if err != nil {
		return nil,err
	}

	transaction.Status = entity.TransactionConfirmed
	err = usecase.TransactionRepository.Save(ctx,transaction)
	if err != nil {
		return nil,err
	}

	return transaction,nil
}

func (usecase *TransactionUseCase) CompleteTransaction(ctx context.Context, transactionID string) (*entity.Transaction,error) {
	transaction,err := usecase.TransactionRepository.Find(ctx,transactionID)
	if err != nil {
		return nil,err
	}

	transaction.Status = entity.TransactionCompleted
	err = usecase.TransactionRepository.Save(ctx,transaction)
	if err != nil {
		return nil,err
	}

	return transaction,nil
}
