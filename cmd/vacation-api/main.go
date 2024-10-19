package main

import (
	"log"
	"os"

	"github.com/Dmitriy-M1319/vacation-api/internal/database"
	"github.com/Dmitriy-M1319/vacation-api/internal/repository"
	"github.com/Dmitriy-M1319/vacation-api/internal/repository/interfaces"
	"github.com/joho/godotenv"
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

	database.Close(conn)
}
