package database

import "flyte-league/models"

func DBInsertArcher(archer *models.Archer) error {
	_, err := db.NamedExec(`INSERT OR IGNORE INTO archers (first_name, last_name, school, bowtype, age) VALUES (:first_name, :last_name, :school, :bowtype, :age)`, archer)
	if err != nil {
		return err
	}
	return nil
}

func DBInsertMultipleArchers(archers []*models.Archer) error {
	stmt, err := db.PrepareNamed(`INSERT OR IGNORE INTO archers (first_name, last_name, school, bowtype, age) VALUES (:first_name, :last_name, :school, :bowtype, :age)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, archer := range archers {
		_, err := stmt.Exec(archer)
		if err != nil {
			return err
		}
	}
	return nil
}

func DBGetArcherByFullName(first_name string, last_name string) (*models.Archer, error) {
	archer := &models.Archer{}
	err := db.Get(archer, `SELECT first_name, last_name, school, bowtype, age FROM archers WHERE first_name=$1 AND last_name=$2`, first_name, last_name)
	if err != nil {
		return nil, err
	}
	return archer, nil
}

func DBGetAllArchersInSchool(school string) ([]models.Archer, error) {
	archers := []models.Archer{}
	err := db.Select(archers, `SELECT first_name, last_name, school, bowtype, age FROM archers WHERE school=$1 ORDER BY last_name ASC`, school)
	if err != nil {
		return nil, err
	}
	return archers, nil
}

func DBGetArchersByfirst_name(name string) ([]models.Archer, error) {
	archers := []models.Archer{}
	err := db.Select(archers, `SELECT first_name, last_name, school, bowtype, age FROM archers WHERE first_name=$1 ORDER BY first_name ASC`, name)
	if err != nil {
		return nil, err
	}
	return archers, nil
}

func DBGetArchersBylast_name(name string) ([]models.Archer, error) {
	archers := []models.Archer{}
	err := db.Select(archers, `SELECT first_name, last_name, school, bowtype, age FROM archers WHERE last_name=$1 ORDER BY last_name ASC`, name)
	if err != nil {
		return nil, err
	}
	return archers, nil
}

func DBGetAllArchersByBowtype(bowtype string) ([]models.Archer, error) {
	archers := []models.Archer{}
	err := db.Select(archers, `SELECT first_name, last_name, school, bowtype, age FRROM archers WHERE bowtype=$1 ORDER BY last_name ASC`, bowtype)
	if err != nil {
		return nil, err
	}
	return archers, nil
}

func DBGetAllArchersByAge(upperAgeLimit int, lowerAgeLimit int) ([]models.Archer, error) {
	archers := []models.Archer{}
	err := db.Select(archers, `SELECT first_name, last_name, school, bowtype, age FRROM archers WHERE age<$1 AND age >$2 ORDER BY last_name ASC`, upperAgeLimit, lowerAgeLimit)
	if err != nil {
		return nil, err
	}
	return archers, nil
}
