package usecase

import (
	"context"

	"github.com/ruhanc/pix-go/internal/domain/entity"
	"github.com/ruhanc/pix-go/internal/domain/gateway"
)

type PixUseCase struct {
	PixKeyRepository gateway.PixKeyRepositoryInterface
}

func (usecase *PixUseCase) RegisterKey(ctx context.Context, key, kind, accountID string) (*entity.PixKey, error) {
	account, err := usecase.PixKeyRepository.FindAccount(ctx, accountID)
	if err != nil {
		return nil, err
	}

	pixKey, err := entity.NewPixKey(kind, key, account)
	if err != nil {
		return nil, err
	}

	registeredPixKey, err := usecase.PixKeyRepository.Register(ctx, pixKey)
	if err != nil {
		return nil, err
	}

	return registeredPixKey, nil
}

func (usecase *PixUseCase) FindKey(ctx context.Context, id string) (*entity.PixKey, error) {
	pixKey, err := usecase.PixKeyRepository.FindKeyByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
