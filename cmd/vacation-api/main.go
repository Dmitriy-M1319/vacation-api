package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/Dmitriy-M1319/vacation-api/internal/api/rpc"
	"github.com/Dmitriy-M1319/vacation-api/internal/database"
	"github.com/Dmitriy-M1319/vacation-api/internal/repository"
	"github.com/Dmitriy-M1319/vacation-api/internal/repository/interfaces"
	pb "github.com/Dmitriy-M1319/vacation-api/pkg/vacation_api/v1"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var employeeRepo interfaces.EmployeeRepository
var vacationRepo interfaces.VacationRepository
var vacationNormRepo interfaces.VacationNormRepository

func runGrpcServer() {
	sock, err := net.Listen("tcp", ":12201")
	if err != nil {
		log.Fatalf("failed to listen port: %v", err)
	}

	grpcServer := grpc.NewServer()
	service := rpc.NewRpcService(employeeRepo, vacationRepo, vacationNormRepo)

	pb.RegisterVacationsServiceServer(grpcServer, service)
	err = grpcServer.Serve(sock)

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func runGrpcGatewayRest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterVacationsServiceHandlerFromEndpoint(ctx, mux, "localhost:12201", opts)
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := database.NewConnection(os.Getenv("POSTGRES_IP"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DATABASE"))
	defer database.Close(conn)
	if err != nil {
		log.Fatal(err)
	}

	employeeRepo, err = repository.NewPgEmployeeRepository(conn)
	if err != nil {
		log.Fatal(err)
	}
	vacationRepo, err = repository.NewPgVacationRepository(conn)
	if err != nil {
		log.Fatal(err)
	}
	vacationNormRepo, err = repository.NewPgVacationNormRepository(conn)
	if err != nil {
		log.Fatal(err)
	}

	go runGrpcGatewayRest()
	runGrpcServer()
}
