package repository

import (
	"context"
	"database/sql"

	"github.com/ruhanc/pix-go/internal/domain/entity"
	"github.com/ruhanc/pix-go/internal/infra/db"
)

type TransactionRepository struct{
	DB *sql.DB
	Queries *db.Queries
}

func NewTransactionRepository(database *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: database,
		Queries: db.New(database),
	}
}

func (repo *TransactionRepository) Register(ctx context.Context, transaction *entity.Transaction) error {
	err := repo.Queries.CreateTransaction(ctx,db.CreateTransactionParams{
		ID: transaction.ID,
		CreatedAt: transaction.CreatedAt,
		AccountFromID: transaction.AccountFrom.ID,
		PixKeyID: transaction.PixKeyTo.ID,
		Amount: int64(transaction.Amount),
		Status: transaction.Status,
		Description: transaction.Description,
	})
	if err != nil {
		return err
	}
	return nil
}

func (repo *TransactionRepository) Save(ctx context.Context,transaction *entity.Transaction) error {
	_,err := repo.Queries.UpdateTransactionStatus(ctx, db.UpdateTransactionStatusParams{
		ID: transaction.ID,
		Status: transaction.Status,
	})
	if err != nil {
		return err
	}
	return nil
}

func (repo *TransactionRepository) Find(ctx context.Context,id string) (*entity.Transaction,error) {
	transactionEntity := &entity.Transaction{}
	accountEntity := &entity.Account{}
	bankEntity := &entity.Bank{}
	pixKeyEntity := &entity.PixKey{}
	transaction,err := repo.Queries.FindTransactionByID(ctx, id)
	if err != nil {
		return nil,err
	}

	account,err := repo.Queries.FindAccountByID(ctx,transaction.AccountFromID)
	if err != nil {
		return nil,err
	}
	
	bank,err := repo.Queries.FindBankByID(ctx,account.BankID)
	if err != nil {
		return nil,err
	}
	bankEntity.ID = bank.ID
	bankEntity.CreatedAt = bank.CreatedAt
	bankEntity.Name = bank.Name
	bankEntity.Code = bank.Code

	accountEntity.ID = account.ID
	accountEntity.Bank = bankEntity
	accountEntity.OwnerName = account.OwnerName
	accountEntity.Number = account.Number
	accountEntity.CreatedAt = account.CreatedAt
	
	pixKey,err := repo.Queries.FindPixKeyByID(ctx,transaction.PixKeyID)
	if err != nil {
		return nil,err
	}
	pixKeyEntity.ID = pixKey.ID
	pixKeyEntity.CreatedAt = pixKey.CreatedAt
	pixKeyEntity.Key = pixKey.Key
	pixKeyEntity.Kind = pixKey.Kind
	pixKeyEntity.Status = pixKey.Status
	pixKeyEntity.Account = accountEntity

	transactionEntity.ID = transaction.ID
	transactionEntity.CreatedAt = transaction.CreatedAt
	transactionEntity.Amount = float64(transaction.Amount)
	transactionEntity.Description = transaction.Description
	transactionEntity.Status = transaction.Status
	transactionEntity.PixKeyTo = pixKeyEntity
	transactionEntity.AccountFrom = accountEntity

	return transactionEntity,nil
}