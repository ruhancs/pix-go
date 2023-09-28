package gateway

import (
	"context"

	"github.com/ruhanc/pix-go/internal/domain/entity"
)

type PixKeyRepositoryInterface interface {
	Register(ctx context.Context,pixKey *entity.PixKey) (*entity.PixKey,error)
	FindKeyByID(ctx context.Context,key, id string) (*entity.PixKey,error)
	AddBank(ctx context.Context,bank *entity.Bank) error
	AddAccount(ctx context.Context,account *entity.Account) error
	FindAccount(ctx context.Context,id string) (*entity.Account,error)
}