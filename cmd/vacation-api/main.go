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
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var employeeRepo interfaces.EmployeeRepository
var vacationRepo interfaces.VacationRepository
var vacationNormRepo interfaces.VacationNormRepository

func runGrpcServer(port string) {
	sock, err := net.Listen("tcp", ":"+port)
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

func runGrpcGatewayRest(httpPort, grpcPort string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterVacationsServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts)
	if err != nil {
		panic(err)
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Разрешить все источники
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	log.Printf("server listening at %s", httpPort)
	if err := http.ListenAndServe(":"+httpPort, handler); err != nil {
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

	go runGrpcGatewayRest(os.Getenv("API_HTTP_PORT"), os.Getenv("API_GRPC_PORT"))
	runGrpcServer(os.Getenv("API_GRPC_PORT"))
}
