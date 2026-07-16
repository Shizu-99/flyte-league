package models

import "time"

type Archer struct {
	First_name string `db:"first_name"`
	Last_name  string `db:"last_name"`
	School     string `db:"school"`
	Bowtype    string `db:"bowtype"`
	Age        int    `db:"age"`
}

type Event struct {
	Event_name string    `db:"event_name"`
	Date       time.Time `db:"date"`
	Round_num  int       `db:"round_num"`
	Session    int       `db:"session"`
}

type Score struct {
	Archer Archer `db:"archer"`
	Event  Event  `db:"event"`
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
