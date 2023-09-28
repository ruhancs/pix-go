package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Bank struct {
	Base `valid:"required"`
	Code string `json:"code" valid:"notnull"`
	Name string `json:"name" valid:"notnull"`
	Accounts []*Account `valid:"-"`
}

func NewBank(code, name string) (*Bank, error) {
	bank := &Bank{}
	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()
	bank.Code = code
	bank.Name = name

	if err := bank.Validate(); err != nil {
		return nil, err
	}

	return bank, nil
}

func (b *Bank) Validate() error {
	_, err := govalidator.ValidateStruct(b)
	if err != nil {
		return err
	}

	return nil
}
