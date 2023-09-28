package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base `valid:"required"`
	OwnerName string `json:"owner_name" valid:"notnull"`
	Bank *Bank `valid:"-"`
	Number string `json:"number" valid:"notnull"`
	PixKeys []*PixKey `valid:"-"`
}

func NewAccount(ownerName, number string, bank *Bank) (*Account,error) {
	account := &Account{
		OwnerName: ownerName,
		Number: number,
		Bank: bank,
	}
	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	if err := account.Validate(); err != nil {
		return nil,err
	}

	return account,nil
}

func (b *Account) Validate() error {
	_, err := govalidator.ValidateStruct(b)
	if err != nil {
		return err
	}

	return nil
}