package repository

import (
	"fmt"

	"github.com/Dmitriy-M1319/vacation-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type SqlVacationNormRepository struct {
	connection *sqlx.DB
	lastId     uint64
}

func NewPgVacationNormRepository(c *sqlx.DB) (*SqlVacationNormRepository, error) {
	rep := SqlVacationNormRepository{connection: c}

	//seed
	query := `CREATE TABLE IF NOT EXISTS vacation_norms(
	id serial primary key,
	month varchar(20),
	vacations_count integer
	)`

	_, err := c.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("failed to init vacation norm repository: %s", err.Error())
	}

	var maxId int
	err = c.Get(&maxId, "SELECT (CASE WHEN MAX(id) IS NULL THEN 0 ELSE MAX(id) END) AS maxId FROM vacation_norms")
	if err != nil {
		return nil, fmt.Errorf("failed to init vacation norm repository: %s", err.Error())
	}

	rep.lastId = uint64(maxId)
	return &rep, nil
}

func (repo SqlVacationNormRepository) GetAll() ([]models.VacationNorm, error) {
	var result []models.VacationNorm
	err := repo.connection.Select(&result, "SELECT * FROM vacation_norms ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("failed to get vacation norms: %s", err.Error())
	}
	return result, nil
}

func (repo SqlVacationNormRepository) GetById(id uint64) (models.VacationNorm, error) {
	var result models.VacationNorm = models.VacationNorm{}
	err := repo.connection.Get(&result, "SELECT * FROM vacation_norms WHERE id = $1", id)
	if err != nil {
		return result, fmt.Errorf("failed to get vacation norms: %s", err.Error())
	}
	return result, nil
}

func (repo *SqlVacationNormRepository) Insert(v *models.VacationNorm) error {
	var newId uint64
	err := repo.connection.Get(&newId,
		"INSERT INTO vacation_norms(month, vacations_count) VALUES($1, $2) RETURNING id",
		v.Month, v.VacationsCount)
	if err != nil {
		return fmt.Errorf("failed to insert new vacation norm")
	}

	v.ID = newId
	return nil
}

func (repo SqlVacationNormRepository) Update(id uint64, v *models.VacationNorm) error {
	_, err := repo.GetById(id)
	if err != nil {
		return err
	}

	_, err = repo.connection.Exec("UPDATE vacation_norms SET month=$1, vacations_count=$2 WHERE id=$3",
		v.Month, v.VacationsCount, id)
	if err != nil {
		return fmt.Errorf("failed to update vacation norm with id = %d", id)
	}
	return nil
}

func (repo SqlVacationNormRepository) Delete(id uint64) error {
	_, err := repo.GetById(id)
	if err != nil {
		return err
	}
	_, err = repo.connection.Exec("DELETE FROM vacation_norms WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("failed to delete vacation norm with id = %d", id)
	}
	return nil
}
