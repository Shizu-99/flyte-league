package database

import (
	"database/sql"
	"flyte-league/models"
)

func DBInsertArchers(archers []*models.Archer) error {
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

// DBGetArchersByFirstName returns a slice of Archer structs
// from the database. The function returns all archers
// whose first name starts with the given substring.
//
// If there are no archers whose first name starts with a given substring
// the function returns sql.ErrNoRows.
func DBGetArchersByFirstName(name string) ([]models.Archer, error) {
	archers := []models.Archer{}
	query := `SELECT first_name, last_name FROM archers WHERE first_name LIKE ? ORDER BY first_name ASC`
	searchPattern := name + "%"
	err := db.Select(&archers, query, searchPattern)
	if err != nil {
		return nil, err
	}
	if len(archers) == 0 {
		return nil, sql.ErrNoRows
	}
	return archers, nil
}

// DBGetArchersByLastName returns a slice of Archer structs
// from the database. The function returns all archers
// whose last name starts with the given substring.
//
// If there are no archers whose last name starts with a given substring
// the function returns sql.ErrNoRows.
func DBGetArchersByLastName(name string) ([]models.Archer, error) {
	archers := []models.Archer{}
	query := `SELECT first_name, last_name FROM archers WHERE last_name LIKE ? ORDER BY last_name ASC`
	searchPattern := name + "%"
	err := db.Select(&archers, query, searchPattern)
	if err != nil {
		return nil, err
	}
	if len(archers) == 0 {
		return nil, sql.ErrNoRows
	}
	return archers, nil
}
