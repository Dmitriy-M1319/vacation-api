package main

import (
	"log"
	"net"
	"os"

	"github.com/Dmitriy-M1319/vacation-api/internal/api/rpc"
	"github.com/Dmitriy-M1319/vacation-api/internal/database"
	"github.com/Dmitriy-M1319/vacation-api/internal/repository"
	"github.com/Dmitriy-M1319/vacation-api/internal/repository/interfaces"
	pb "github.com/Dmitriy-M1319/vacation-api/pkg/vacation_api/v1"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var employeeRepo interfaces.EmployeeRepository
var vacationRepo interfaces.VacationRepository
var vacationNormRepo interfaces.VacationNormRepository

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

	// gRPC server
	lis, err := net.Listen("tcp", "[::1]:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	service := rpc.NewRpcService(employeeRepo, vacationRepo, vacationNormRepo)

	pb.RegisterVacationsServiceServer(grpcServer, service)
	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalf("Error strating server: %v", err)
	}

	// ctx := context.Background()
	// httpRouter := gin.Default()
	// httpRouter.Use(cors.Default())

	// employeeHandler := http.NewEmployeeHandlers(employeeRepo)
	// httpRouter.GET("/employees", employeeHandler.GetAll)
	// httpRouter.GET("/employees/:id", employeeHandler.GetById)
	// httpRouter.POST("/employees", employeeHandler.Insert)
	// httpRouter.PUT("/employees/:id", employeeHandler.Update)
	// httpRouter.DELETE("/employees/:id", employeeHandler.Delete)

	// vacationHandler := http.NewVacationHandlers(vacationRepo)
	// httpRouter.GET("/vacations", vacationHandler.GetAll)
	// httpRouter.GET("/vacations/:id", vacationHandler.GetById)
	// httpRouter.GET("/employees/:id/vacations", vacationHandler.GetByEmployeeId)
	// httpRouter.POST("/vacations", vacationHandler.Insert)
	// httpRouter.PUT("/vacations/:id", vacationHandler.Update)
	// httpRouter.DELETE("/vacations/:id", vacationHandler.Delete)

	// normHandler := http.NewVacationNormHandlers(vacationNormRepo)
	// httpRouter.GET("/norms", normHandler.GetAll)
	// httpRouter.GET("/norms/:id", normHandler.GetById)
	// httpRouter.POST("/norms", normHandler.Insert)
	// httpRouter.PUT("/norms/:id", normHandler.Update)
	// httpRouter.DELETE("/norms/:id", normHandler.Delete)

	// httpRouter.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// server := &http2.Server{
	// 	Addr:    "localhost:8081",
	// 	Handler: httpRouter.Handler(),
	// }

	// go func() {
	// 	err := server.ListenAndServe()
	// 	if err != http2.ErrServerClosed {
	// 		log.Fatal(err.Error())
	// 	}
	// }()

	// notifChan := make(chan os.Signal, 1)
	// signal.Notify(notifChan, syscall.SIGINT, syscall.SIGTERM)
	// _ = <-notifChan

	// closeCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	// defer cancel()

	// err = server.Shutdown(closeCtx)
	// if err != nil {
	// 	log.Fatalf("Failed to complete graceful shutdown: %v\n", err)
	// }
}
