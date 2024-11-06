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

	query := `CREATE TABLE IF NOT EXISTS employees(
	id serial primary key,
	first_name varchar(255),
	last_name varchar(255),
	patronymic varchar(255)
	)`

	_, err := c.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("failed to init employee repository: %s", err.Error())
	}

	var maxId int
	err = c.Get(&maxId, "SELECT (CASE WHEN MAX(id) IS NULL THEN 0 ELSE MAX(id) END) AS maxId FROM employees")
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

func (repo *SqlEmployeeRepository) Insert(e *models.Employee) error {
	var newId uint64
	err := repo.connection.Get(&newId,
		"INSERT INTO employees(first_name, last_name, patronymic) VALUES($1, $2, $3) RETURNING id",
		e.FirstName, e.LastName, e.Patronymic)
	if err != nil {
		return fmt.Errorf("failed to insert new employee")
	}

	e.ID = newId
	return nil
}

func (repo SqlEmployeeRepository) Update(id uint64, e *models.Employee) error {
	_, err := repo.GetById(id)
	if err != nil {
		return err
	}

	_, err = repo.connection.Exec("UPDATE employees SET first_name=$1, last_name=$2, patronymic=$3 WHERE id=$4",
		e.FirstName, e.LastName, e.Patronymic, id)
	if err != nil {
		return fmt.Errorf("failed to update employee with id = %d", id)
	}
	return nil
}

func (repo SqlEmployeeRepository) Delete(id uint64) error {
	_, err := repo.GetById(id)
	if err != nil {
		return err
	}
	_, err = repo.connection.Exec("DELETE FROM employees WHERE id=$1",
		id)
	if err != nil {
		return fmt.Errorf("failed to delete employee with id = %d", id)
	}
	return nil
}
