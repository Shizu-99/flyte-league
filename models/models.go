package models

type Archer struct {
	Firstname string `db:"firstname"`
	Lastname  string `db:"lastname"`
	School    string `db:"school"`
	Bowtype   string `db:"bowtype"`
	Age       int    `db:"age"`
}
