package models

type Archer struct {
	Name    string `db:"name"`
	School  string `db:"school"`
	Bowtype string `db:"bowtype"`
	Age     int    `db:"age"`
}
