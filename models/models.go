package models

import "time"

type Archer struct {
	First_name string `db:"first_name"`
	Last_name  string `db:"last_name"`
	School     string `db:"school"`
	Bowtype    string `db:"bowtype"`
	Age        int    `db:"age"`
}

type Round struct {
	Round_num  int       `db:"round_num"`
	Session    int       `db:"session"`
	Date       time.Time `db:"date"`
	Event_name string    `db:"event_name"`
}

type Score struct {
	Archer Archer `db:"archer"`
	Event  string `db:"event"`
	Score  int    `db:"score"`
	Xs     int    `db:"xs"`
	tens   int    `db:"tens"`
	hits   int    `db:"hits"`
}

type Team struct {
	Team_name string `db:"team_name"`
	School    string `db:"school"`
	Division  int    `db:"division"`
	Bowtype   string `db:"bowtype"`
}
