package services

import (
	"github.com/Dmitriy-M1319/vacation-api/internal/models"
	"github.com/Dmitriy-M1319/vacation-api/internal/repository"
)

type VacationNormService struct {
	repo repository.SqlVacationNormRepository
}

func NewVacationNormService(r repository.SqlVacationNormRepository) *VacationNormService {
	service := VacationNormService{repo: r}
	return &service
}

func (service *VacationNormService) GetAllVacationNorms() ([]models.VacationNormDto, error) {
	norms, err := service.repo.GetAll()
	result := make([]models.VacationNormDto, len(norms))
	for i, n := range norms {
		result[i] = models.VacationNormDto{VacationNorm: n}
	}
	return result, err
}

func (service *VacationNormService) GetVacationNormById(id uint64) (models.VacationNormDto, error) {
	norm, err := service.repo.GetById(id)
	return models.VacationNormDto{VacationNorm: norm}, err
}

func (service *VacationNormService) CreateNewVacationNorm(body models.VacationNormBodyDto) (models.VacationNormDto, error) {
	result := models.VacationNorm{Month: body.Month, VacationsCount: body.VacationsCount}
	err := service.repo.Insert(&result)
	return models.VacationNormDto{VacationNorm: result}, err
}

func (service *VacationNormService) UpdateExistingVacationNorm(id uint64, body models.VacationNormBodyDto) (models.VacationNormDto, error) {
	result := models.VacationNorm{Month: body.Month, VacationsCount: body.VacationsCount}
	err := service.repo.Update(id, &result)
	return models.VacationNormDto{VacationNorm: result}, err
}

func (service *VacationNormService) DeleteExistingVacationNorm(id uint64) error {
	err := service.repo.Delete(id)
	return err
}
