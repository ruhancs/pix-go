package grpc

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/ruhanc/pix-go/internal/application/grpc/pb"
	"github.com/ruhanc/pix-go/internal/application/usecase"
	"github.com/ruhanc/pix-go/internal/infra/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer (database *sql.DB, port int) {
	grpcServer := grpc.NewServer()

	//debug do grpc
	reflection.Register(grpcServer)

	pixRepository := repository.NewPixKeyRepository(database)
	pixUseCase := usecase.PixUseCase{PixKeyRepository: pixRepository}
	pixGRPCService := NewPixGRPCService(pixUseCase)

	//adicionar o pixGRPCService ao servidor grpc
	pb.RegisterPixServiceServer(grpcServer,pixGRPCService)

	address := fmt.Sprintf("0.0.0.0:%d",port)
	listener,err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start grpc ", err)
	}
	
	log.Printf("GRPC server has been started on port %d", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc ", err)
	}

}