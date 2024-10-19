package services

import (
	"github.com/Dmitriy-M1319/vacation-api/internal/models"
	"github.com/Dmitriy-M1319/vacation-api/internal/repository"
)

type VacationService struct {
	repo repository.SqlVacationRepository
}

func NewVacationService(r repository.SqlVacationRepository) *VacationService {
	service := VacationService{repo: r}
	return &service
}

func (service *VacationService) GetAllVacations() ([]models.VacationDto, error) {
	vacations, err := service.repo.GetAll()
	result := make([]models.VacationDto, len(vacations))
	for i, v := range vacations {
		result[i] = models.VacationDto{Vacation: v}
	}
	return result, err
}

func (service *VacationService) GetVacationById(id uint64) (models.VacationDto, error) {
	vac, err := service.repo.GetById(id)
	return models.VacationDto{Vacation: vac}, err
}

func (service *VacationService) CreateNewVacation(body models.VacationBodyDto) (models.VacationDto, error) {
	result := models.Vacation{EmployeeID: body.EmployeeID, StartDate: body.StartDate, EndDate: body.EndDate, DaysCount: body.DaysCount}
	err := service.repo.Insert(&result)
	return models.VacationDto{Vacation: result}, err
}

func (service *VacationService) UpdateExistingVacation(id uint64, body models.VacationBodyDto) (models.VacationDto, error) {
	result := models.Vacation{EmployeeID: body.EmployeeID, StartDate: body.StartDate, EndDate: body.EndDate, DaysCount: body.DaysCount}
	err := service.repo.Update(id, &result)
	return models.VacationDto{Vacation: result}, err
}

func (service *VacationService) DeleteExistingVacation(id uint64) error {
	err := service.repo.Delete(id)
	return err
}
