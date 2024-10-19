package interfaces

import "github.com/Dmitriy-M1319/vacation-api/internal/models"

type EmployeeRepository interface {
	GetAll() ([]models.Employee, error)
	GetById(id uint64) (models.Employee, error)
	Insert(e *models.Employee) error
	Update(id uint64, e *models.Employee) error
	Delete(id uint64) error
}

type VacationRepository interface {
	GetAll() ([]models.Vacation, error)
	GetById(id uint64) (models.Vacation, error)
	GetByEmployeeId(id uint64) ([]models.Vacation, error)
	Insert(v *models.Vacation) error
	Update(id uint64, v *models.Vacation) error
	Delete(id uint64) error
}

type VacationNormRepository interface {
	GetAll() ([]models.VacationNorm, error)
	GetById(id uint64) (models.VacationNorm, error)
	Insert(n *models.VacationNorm) error
	Update(id uint64, n *models.VacationNorm) error
	Delete(id uint64) error
}
