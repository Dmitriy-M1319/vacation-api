package main

import (
	"log"
	"net"
	"os"

	"github.com/Dmitriy-M1319/vacation-api/api/rpc"
	"github.com/Dmitriy-M1319/vacation-api/internal/database"
	"github.com/Dmitriy-M1319/vacation-api/internal/repository"
	"github.com/Dmitriy-M1319/vacation-api/internal/repository/interfaces"
	pb "github.com/Dmitriy-M1319/vacation-api/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var employeeRepo interfaces.EmployeeRepository
var vacationRepo interfaces.VacationRepository

// var vacationNormRepo interfaces.VacationNormRepository

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
	// vacationNormRepo, err = repository.NewPgVacationNormRepository(conn)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	lis, err := net.Listen("tcp", "[::1]:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	service := rpc.NewRpcService(employeeRepo, vacationRepo)

	pb.RegisterVacationsServiceServer(grpcServer, service)
	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalf("Error strating server: %v", err)
	}
}
