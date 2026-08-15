package database

import "flyte-league/models"

func DBInsertArcher(archer *models.Archer) error {
	_, err := db.NamedExec(`INSERT OR IGNORE INTO archers (first_name, last_name) VALUES (:first_name, :last_name)`, archer)
	if err != nil {
		return err
	}
	return nil
}

func DBInsertMultipleArchers(archers []models.Archer) error {
	_, err := db.NamedExec(`INSERT OR IGNORE INTO archers (first_name, last_name) VALUES (:first_name, :last_name)`, archers)
	if err != nil {
		return err
	}
	return nil
}

func DBGetArcherByFullName(first_name string, last_name string) (*models.Archer, error) {
	archer := &models.Archer{}
	err := db.Get(archer, `SELECT first_name, last_name FROM archers WHERE first_name=$1 AND last_name=$2`, first_name, last_name)
	if err != nil {
		return nil, err
	}
	return archer, nil
}

func DBGetArchersByFirstName(name string) ([]models.Archer, error) {
	archers := []models.Archer{}
	err := db.Select(archers, `SELECT first_name, last_name FROM archers WHERE first_name LIKE '$1%' ORDER BY first_name ASC`, name)
	if err != nil {
		return nil, err
	}
	return archers, nil
}

func DBGetArchersBylast_name(name string) ([]models.Archer, error) {
	archers := []models.Archer{}
	err := db.Select(archers, `SELECT first_name, last_name FROM archers WHERE last_name LIKE '$1%' ORDER BY last_name ASC`, name)
	if err != nil {
		return nil, err
	}
	return archers, nil
}
