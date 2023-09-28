package entity

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKey struct {
	Base `valid:"required"`
	Kind string `json:"kind" valid:"notnull"`
	Key string `json:"key" valid:"notnull"`
	AccountID string `json:"account_id" valid:"notnull"`
	Account *Account `json:"account" valid:"-"`
	Status string `json:"status" valid:"notnull"`
}

func NewPixKey(kind,key string, account *Account) (*PixKey,error) {
	pixKey := &PixKey{
		Kind: kind,
		Key: key,
		Status: "active",
		Account: account,
	}
	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	if err := pixKey.Validate(); err != nil {
		return nil,err
	}

	return pixKey,nil
}

func (p *PixKey) Validate() error {
	_, err := govalidator.ValidateStruct(p)

	if p.Key != "email" && p.Key != "cpf"{
		return errors.New("invalid key type")
	}
	
	if p.Status != "active" && p.Status != "inactive"{
		return errors.New("invalid key type")
	}

	if err != nil {
		return err
	}

	return nil
}