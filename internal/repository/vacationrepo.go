package repository

import (
	"fmt"

	"github.com/Dmitriy-M1319/vacation-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type SqlVacationRepository struct {
	connection *sqlx.DB
	lastId     uint64
}

func NewPgVacationRepository(c *sqlx.DB) (*SqlVacationRepository, error) {
	rep := SqlVacationRepository{connection: c}

	//seed
	query := `CREATE TABLE IF NOT EXISTS vacations(
	id serial primary key,
	emp_id integer REFERENCES employees(id),
	start_date varchar(255),
	end_date varchar(255),
	days_count integer
	)`

	_, err := c.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("failed to init vacation repository: %s", err.Error())
	}

	var maxId int
	err = c.Get(&maxId, "SELECT (CASE WHEN MAX(id) IS NULL THEN 0 ELSE MAX(id) END) AS maxId FROM vacations")
	if err != nil {
		return nil, fmt.Errorf("failed to init vacation repository: %s", err.Error())
	}

	rep.lastId = uint64(maxId)
	return &rep, nil
}

func (repo SqlVacationRepository) GetAll() ([]models.Vacation, error) {
	var result []models.Vacation
	err := repo.connection.Select(&result, "SELECT * FROM vacations ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("failed to get vacations: %s", err.Error())
	}
	return result, nil
}

func (repo SqlVacationRepository) GetById(id uint64) (models.Vacation, error) {
	var result models.Vacation = models.Vacation{}
	err := repo.connection.Get(&result, "SELECT * FROM vacations WHERE id = $1", id)
	if err != nil {
		return result, fmt.Errorf("failed to get vacation: %s", err.Error())
	}
	return result, nil
}

func (repo SqlVacationRepository) GetByEmployeeId(id uint64) ([]models.Vacation, error) {
	var result []models.Vacation
	err := repo.connection.Select(&result, "SELECT * FROM vacations WHERE emp_id=$1 ORDER BY id", id)
	if err != nil {
		return nil, fmt.Errorf("failed to get vacations for employee: %s", err.Error())
	}
	return result, nil
}

func (repo *SqlVacationRepository) Insert(v *models.Vacation) error {
	var newId uint64
	err := repo.connection.Get(&newId,
		"INSERT INTO vacations(emp_id, start_date, end_date, days_count) VALUES($1, $2, $3, $4) RETURNING id",
		v.EmployeeID, v.StartDate, v.EndDate, v.DaysCount)
	if err != nil {
		return fmt.Errorf("failed to insert new vacation")
	}

	v.ID = newId
	return nil
}

func (repo SqlVacationRepository) Update(id uint64, v *models.Vacation) error {
	_, err := repo.GetById(id)
	if err != nil {
		return err
	}

	_, err = repo.connection.Exec("UPDATE vacations SET emp_id=$1, start_date=$2, end_date=$3 days_count=$4 WHERE id=$5",
		v.EmployeeID, v.StartDate, v.EndDate, v.DaysCount, id)
	if err != nil {
		return fmt.Errorf("failed to update vacation with id = %d", id)
	}
	return nil
}

func (repo SqlVacationRepository) Delete(id uint64) error {
	_, err := repo.GetById(id)
	if err != nil {
		return err
	}
	_, err = repo.connection.Exec("DELETE FROM vacations WHERE id=$1",
		id)
	if err != nil {
		return fmt.Errorf("failed to delete vacation with id = %d", id)
	}
	return nil
}
