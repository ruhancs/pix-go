package gateway

import "github.com/ruhanc/pix-go/internal/domain/entity"

type PixKeyRepositoryInterface interface {
	Register(pixKey *entity.PixKey) (*entity.PixKey,error)
	FindKeyByKind(key, kind string) (*entity.PixKey,error)
	AddBank(bank *entity.Bank) error
	AddAccount(account *entity.Account) error
	FindAccount(id string) (*entity.Account,error)
}