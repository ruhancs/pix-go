package factory

import (
	"database/sql"

	"github.com/ruhanc/pix-go/internal/application/usecase"
	"github.com/ruhanc/pix-go/internal/infra/repository"
)

func TransactionUseCaseFactory(database *sql.DB) usecase.TransactionUseCase {
	pixRepository := repository.NewPixKeyRepository(database)
	transactionRepository := repository.NewTransactionRepository(database)

	usecase := usecase.TransactionUseCase{
		PixRepository: pixRepository,
		TransactionRepository: transactionRepository,
	}

	return usecase
}