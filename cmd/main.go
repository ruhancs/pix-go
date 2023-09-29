package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/ruhanc/pix-go/internal/application/grpc"
	_ "github.com/lib/pq" //driver postgres
)


func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env")
	}

	conn, err := sql.Open(os.Getenv("DB_TYPE"), os.Getenv("DSN"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	grpc.StartGrpcServer(conn, 50051)
}