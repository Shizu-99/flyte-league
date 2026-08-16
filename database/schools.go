package database

import (
	"database/sql"
	"flyte-league/models"
)

func DBInsertSchools(schools []*models.School) error {
	_, err := db.NamedExec(`INSERT INTO schools (name, location) VALUES (:name, :location)`, schools)
	if err != nil {
		return err
	}
	return nil
}

func DBGetAllSchools() ([]models.School, error) {
	schools := []models.School{}
	query := `SELECT name, location FROM schools`
	err := db.Select(&schools, query)
	if err != nil {
		return nil, err
	}
	if len(schools) == 0 {
		return nil, sql.ErrNoRows
	}
	return schools, nil
}

func DBGetSchoolByName(name string) (*models.School, error) {
	school := &models.School{}
	err := db.Get(school, `SELECT school_id, name, location FROM schools WHERE name=$1`, name)
	if err != nil {
		return nil, err
	}
	return school, nil
}
