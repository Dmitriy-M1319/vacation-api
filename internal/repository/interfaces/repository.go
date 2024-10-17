package interfaces

import "github.com/Dmitriy-M1319/vacation-api/internal/models"

type EmployeeRepository interface {
	GetAll() []models.Employee
	GetById(id uint64) (models.Employee, error)
	Insert(e *models.Employee)
	Update(id uint64, e *models.Employee) error
	Delete(id uint64) error
}

type VacationRepository interface {
	GetAll() []models.Vacation
	GetById(id uint64) (models.Vacation, error)
	GetByEmployeeId(id uint64) ([]models.Vacation, error)
	Insert(e *models.Vacation)
	Update(id uint64, e *models.Vacation) error
	Delete(id uint64) error
}

type VacationNormRepository interface {
	GetAll() []models.Vacation
	GetById(id uint64) (models.Vacation, error)
	Insert(e *models.Vacation)
	Update(id uint64, e *models.Vacation) error
	Delete(id uint64) error
}
