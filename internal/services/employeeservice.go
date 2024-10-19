package services

import (
	"github.com/Dmitriy-M1319/vacation-api/internal/models"
	"github.com/Dmitriy-M1319/vacation-api/internal/repository"
)

type EmployeeService struct {
	repo repository.SqlEmployeeRepository
}

func NewEmployeeService(r repository.SqlEmployeeRepository) *EmployeeService {
	service := EmployeeService{repo: r}
	return &service
}

func (service *EmployeeService) GetAllEmployees() ([]models.EmployeeDto, error) {
	employees, err := service.repo.GetAll()
	result := make([]models.EmployeeDto, len(employees))
	for i, e := range employees {
		result[i] = models.EmployeeDto{Employee: e}
	}
	return result, err
}

func (service *EmployeeService) GetEmployeeById(id uint64) (models.EmployeeDto, error) {
	emp, err := service.repo.GetById(id)
	return models.EmployeeDto{Employee: emp}, err
}

func (service *EmployeeService) CreateNewEmployee(body models.EmployeeBodyDto) (models.EmployeeDto, error) {
	result := models.Employee{FirstName: body.FirstName, LastName: body.LastName, Patronymic: body.Patronymic}
	err := service.repo.Insert(&result)
	return models.EmployeeDto{Employee: result}, err
}

func (service *EmployeeService) UpdateExistingEmployee(id uint64, body models.EmployeeBodyDto) (models.EmployeeDto, error) {
	result := models.Employee{FirstName: body.FirstName, LastName: body.LastName, Patronymic: body.Patronymic}
	err := service.repo.Update(id, &result)
	return models.EmployeeDto{Employee: result}, err
}

func (service *EmployeeService) DeleteExistingEmployee(id uint64) error {
	err := service.repo.Delete(id)
	return err
}
