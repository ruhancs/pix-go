package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateAccount(t *testing.T) {
	assert := assert.New(t)

	newBank,_ := NewBank("456","BB")
	account, err := NewAccount("test","123456",newBank)

	assert.Empty(err)
	assert.Equal(account.OwnerName, "test")
	assert.Equal(account.Bank.ID,newBank.ID)
}