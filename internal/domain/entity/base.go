package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// executa toda vez que é executada, valida todos fieldes required de base
func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Base struct {
	ID        string    `json:"id" valid:"required,uuid"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
}
