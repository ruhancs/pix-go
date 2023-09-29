package gateway

import (
	"context"

	"github.com/ruhanc/pix-go/internal/domain/entity"
)

type TransactionRepositoryInterface interface {
	Register(ctx context.Context, transaction *entity.Transaction) error
	Save(ctx context.Context,transaction *entity.Transaction) error
	Find(ctx context.Context,id string) (*entity.Transaction,error)
}