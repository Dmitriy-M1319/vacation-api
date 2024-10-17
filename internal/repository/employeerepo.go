package repository

import (
	"fmt"

	"github.com/Dmitriy-M1319/vacation-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type SqlEmployeeRepository struct {
	connection *sqlx.DB
	lastId     uint64
}

func NewPgEmployeeRepository(c *sqlx.DB) (*SqlEmployeeRepository, error) {
	rep := SqlEmployeeRepository{connection: c}

	var maxId int
	err := c.Get(&maxId, "SELECT MAX(id) AS maxId FROM employees")
	if err != nil {
		return nil, fmt.Errorf("failed to init employee repository: %s", err.Error())
	}

	rep.lastId = uint64(maxId)
	return &rep, nil
}

func (repo SqlEmployeeRepository) GetAll() ([]models.Employee, error) {
	var result []models.Employee
	err := repo.connection.Select(&result, "SELECT * FROM employees ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("failed to get employees: %s", err.Error())
	}
	return result, nil
}

func (repo SqlEmployeeRepository) GetById(id uint64) (models.Employee, error) {
	var result models.Employee = models.Employee{}
	err := repo.connection.Get(&result, "SELECT * FROM employees WHERE id = $1", id)
	if err != nil {
		return result, fmt.Errorf("failed to get employees: %s", err.Error())
	}
	return result, nil
}

func (repo *SqlEmployeeRepository) Insert(e *models.Employee) {

}

func (repo SqlEmployeeRepository) Update(id uint64, e *models.Employee) error {

}

func (repo SqlEmployeeRepository) Delete(id uint64) error {

}
