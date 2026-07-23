package database

import "flyte-league/models"

func DBInsertSchool(school *models.School) error {
	_, err := db.NamedExec(`INSERT OR IGNORE INTO schools (name, location) VALUES (:name, :location)`, school)
	if err != nil {
		return err
	}
	return nil
}

func DBGetSchoolByName(name string) (*models.School, error) {
	school := &models.School{}
	err := db.Get(school, `SELECT school_id, name, location FROM schools WHERE name=$1`, name)
	if err != nil {
		return nil, err
	}
	return school, nil
}
