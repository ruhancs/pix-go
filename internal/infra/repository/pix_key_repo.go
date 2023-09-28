package repository

import (
	"context"
	"database/sql"

	"github.com/ruhanc/pix-go/internal/domain/entity"
	"github.com/ruhanc/pix-go/internal/infra/db"
)

type PixKeyRepository struct {
	DB *sql.DB
	Queries *db.Queries
}

func NewPixKeyRepository(database *sql.DB) *PixKeyRepository {
	return &PixKeyRepository{
		DB: database,
		Queries: db.New(database),
	}
}

func (p *PixKeyRepository) AddBank(ctx context.Context,bank *entity.Bank) error {
	err := p.Queries.CreateBank(ctx, db.CreateBankParams{
		ID: bank.ID,
		Code: bank.Code,
		Name: bank.Name,
		CreatedAt: bank.CreatedAt,
	})

	if err != nil {
		return err
	}

	return nil
}

func (p *PixKeyRepository) AddAccount(ctx context.Context, account *entity.Account) error {
	err := p.Queries.CreateAccount(ctx, db.CreateAccountParams{
		ID: account.ID,
		OwnerName: account.OwnerName,
		Balance: 0,
		Number: account.Number,
		BankID: account.Bank.ID,
	})

	if err != nil {
		return err
	}

	return nil
}

func (p *PixKeyRepository) Register(ctx context.Context, pixKey *entity.PixKey) (*entity.PixKey,error) {
	err := p.Queries.CreatePixKey(ctx,db.CreatePixKeyParams{
		ID: pixKey.ID,
		Kind: pixKey.Kind,
		Key: pixKey.Key,
		Status: pixKey.Status,
		AccountID: pixKey.Account.ID,
		CreatedAt: pixKey.CreatedAt,
	})

	if err != nil {
		return nil,err
	}

	return pixKey,nil
}

func (p *PixKeyRepository) FindKeyByID(ctx context.Context,key, id string) (*entity.PixKey,error) {
	pixKey := &entity.PixKey{}
	accountEntity := &entity.Account{}
	bankEntity := &entity.Bank{}

	res,err := p.Queries.FindPixKeyByID(ctx,id)
	if err != nil {
		return nil, err
	}
	pixKey.ID = res.ID
	pixKey.CreatedAt = res.CreatedAt
	pixKey.Key = res.Key
	pixKey.Kind = res.Kind
	pixKey.Status = res.Status

	account,err := p.Queries.FindAccountByID(ctx,res.AccountID)
	if err != nil {
		return nil,err
	}
	bank,err := p.Queries.FindBankByID(ctx,account.BankID)
	if err != nil {
		return nil,err
	}
	bankEntity.ID = bank.ID
	bankEntity.Code = bank.Code
	bankEntity.Name = bank.Name
	bankEntity.CreatedAt = bank.CreatedAt
	accountEntity.ID = account.ID
	accountEntity.Bank = bankEntity
	accountEntity.CreatedAt = account.CreatedAt
	accountEntity.OwnerName = account.OwnerName
	accountEntity.Number = account.Number

	pixKey.Account = accountEntity
	
	
	return pixKey,nil
}

func (p *PixKeyRepository) FindAccount(ctx context.Context,id string) (*entity.Account,error) {
	accountEntity := &entity.Account{}
	bankEntity := &entity.Bank{}

	account,err := p.Queries.FindAccountByID(ctx,id)
	if err != nil {
		return nil,err
	}
	bank,err := p.Queries.FindBankByID(ctx,account.BankID)
	if err != nil {
		return nil,err
	}
	bankEntity.ID=bank.ID
	bankEntity.CreatedAt = bank.CreatedAt
	bankEntity.Code = bank.Code
	bankEntity.Name = bank.Name

	accountEntity.ID = account.ID
	accountEntity.CreatedAt = account.CreatedAt
	accountEntity.OwnerName = account.OwnerName
	accountEntity.Number = account.Number
	accountEntity.Bank = bankEntity

	return accountEntity,nil
}

