package entity

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransactionPending = "pending"
	TransactionCompleted = "completed"
	TransactionError = "error"
	TransactionConfirmed = "confirmed"
)

type Transaction struct {
	Base `valid:"required"`
	AccountFrom *Account `valid:"-"`
	Amount float64 `json:"amount" valid:"notnull"`
	PixKeyTo *PixKey `json:"pix_key" valid:"-"`
	Status string `json:"status" valid:"notnull"`
	Description string `json:"description" valid:"notnull"`
	CancelDescription string `json:"cancel_description" valid:"-"`
}

func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction,error) {
	t := &Transaction{
		AccountFrom: accountFrom,
		Amount: amount,
		PixKeyTo: pixKeyTo,
		Status: TransactionPending,
		Description: description,
	}
	t.ID = uuid.NewV4().String()
	t.CreatedAt = time.Now()

	if err := t.Validate(); err != nil {
		return nil,err
	}

	return t,nil
}

func (t *Transaction) Validate() error {
	_, err := govalidator.ValidateStruct(t)

	if t.Amount <= 0 {
		return errors.New("the amount must be greather than 0")
	}

	if t.Status != TransactionPending && t.Status != TransactionCompleted && t.Status != TransactionError && t.Status != TransactionConfirmed{
		return errors.New("invalid transaction status")
	}

	if t.PixKeyTo.Account.ID == t.AccountFrom.ID {
		return errors.New("invalid transaction to same account")
	}

	if err != nil {
		return err
	}

	return nil
}

func (t *Transaction) Complete() error {
	t.Status = TransactionCompleted
	if err := t.Validate(); err != nil {
		return err
	}
	return nil
}

func (t *Transaction) Confirm() error {
	t.Status = TransactionConfirmed
	if err := t.Validate(); err != nil {
		return err
	}
	return nil
}

func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionError
	t.Description = description
	if err := t.Validate(); err != nil {
		return err
	}
	return nil
}