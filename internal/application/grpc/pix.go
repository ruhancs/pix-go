package grpc

import (
	"context"

	"github.com/ruhanc/pix-go/internal/application/grpc/pb"
	"github.com/ruhanc/pix-go/internal/application/usecase"
)

type PixGrpcService struct {
	pb.UnimplementedPixServiceServer
	PixUseCase usecase.PixUseCase 
}

func (p *PixGrpcService) RegisterPixKey(ctx context.Context, in *pb.PixRegistration) (*pb.PixKeyCreatedResult, error) {
	key,err := p.PixUseCase.RegisterKey(ctx,in.Key,in.Kind,in.AccountID)
	if err != nil {
		return &pb.PixKeyCreatedResult{
			Status: "not created",
			Error: err.Error(),
		},err
	}

	return &pb.PixKeyCreatedResult{
		Id: key.ID,
		Status: "created",
	}, nil
}

func (p *PixGrpcService) Find(ctx context.Context, in *pb.PixKey) (*pb.PixKeyInfo, error) {
	key,err := p.PixUseCase.FindKey(ctx,in.Id)
	if err != nil {
		return &pb.PixKeyInfo{},err
	}

	return &pb.PixKeyInfo{
		Id: key.ID,
		Key: key.Key,
		Kind: key.Kind,
		Account: &pb.Account{
			OwnerName: key.Account.OwnerName,
			AccountID: key.Account.ID,
			AccountNumber: key.Account.Number,
			BankID: key.Account.Bank.ID,
			BankName: key.Account.Bank.Name,
			CreatedAt: key.Account.CreatedAt.String(),
		},
		CreatedAt: key.CreatedAt.String(),
	}, nil
}

func NewPixGRPCService(usecase usecase.PixUseCase) *PixGrpcService {
	return &PixGrpcService{
		PixUseCase: usecase,
	}
}