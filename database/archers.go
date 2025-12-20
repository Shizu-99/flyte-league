package database

import "flyte-league/models"

func DBInsertArcher(archer *models.Archer) error {
	_, err := db.NamedExec(`INSERT OR IGNORE INTO archers (firstname, lastname, school, bowtype, age) VALUES (:firstname, :lastname, :school, :bowtype, :age)`, archer)
	if err != nil {
		return err
	}
	return nil
}

func DBInsertMultipleArchers(archers []*models.Archer) error {
	stmt, err := db.PrepareNamed(`INSERT OR IGNORE INTO archers (firstname, lastname, school, bowtype, age) VALUES (:firstname, :lastname, :school, :bowtype, :age)`)
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

func DBGetArcherByFullName(firstname string, lastname string) (*models.Archer, error) {
	archer := &models.Archer{}
	err := db.Get(archer, `SELECT firstname, lastname, school, bowtype, age FROM archers WHERE firstname=$1 AND lastname=$2`, firstname, lastname)
	if err != nil {
		return nil, err
	}
	return archer, nil
}

func DBGetAllArchersInSchool(school string) ([]models.Archer, error) {
	archers := []models.Archer{}
	err := db.Select(archers, `SELECT firstname, lastname, school, bowtype, age FRROM archers WHERE school=$1 ORDER BY lastname ASC`, school)
	if err != nil {
		return nil, err
	}
	return archers, nil
}

func DBGetArchersByFirstName(name string) ([]models.Archer, error) {
	archers := []models.Archer{}
	err := db.Select(archers, `SELECT firstname, lastname, school, bowtype, age FROM archers WHERE firstname=$1 ORDER BY firstname ASC`, name)
	if err != nil {
		return nil, err
	}
	return archers, nil
}

func DBGetArchersByLastName(name string) ([]models.Archer, error) {
	archers := []models.Archer{}
	err := db.Select(archers, `SELECT firstname, lastname, school, bowtype, age FROM archers WHERE lastname=$1 ORDER BY lastname ASC`, name)
	if err != nil {
		return nil, err
	}
	return archers, nil
}

func DBGetAllArchersByBowtype(bowtype string) ([]models.Archer, error) {
	archers := []models.Archer{}
	err := db.Select(archers, `SELECT firstname, lastname, school, bowtype, age FRROM archers WHERE bowtype=$1 ORDER BY lastname ASC`, bowtype)
	if err != nil {
		return nil, err
	}
	return archers, nil
}

func DBGetAllArchersByAge(upperAgeLimit int, lowerAgeLimit int) ([]models.Archer, error) {
	archers := []models.Archer{}
	err := db.Select(archers, `SELECT firstname, lastname, school, bowtype, age FRROM archers WHERE age<$1 AND age >$2 ORDER BY lastname ASC`, upperAgeLimit, lowerAgeLimit)
	if err != nil {
		return nil, err
	}
	return archers, nil
}
