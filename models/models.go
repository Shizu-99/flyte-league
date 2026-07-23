package models

import "time"

type School struct {
	ID       int    `db:"school_id"`
	Name     string `db:"name"`
	Location string `db:"location"`
}

type Archer struct {
	ID         int    `db:"archer_id"`
	First_name string `db:"first_name"`
	Last_name  string `db:"last_name"`
}

type Event struct {
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Start_date  time.Time `db:"start_date"`
	End_date    time.Time `db:"end_date"`
}

type EventDay struct {
	Date       time.Time `db:"day_date"`
	Day_number int       `db:"day_number"`
}

type Round struct {
	Name               string `db:"name"`
	Max_possible_score int    `db:"max_possible_score"`
}

type Session struct {
	Name       string    `db:"name"`
	Start_time time.Time `db:"start_time"`
}

type Score struct {
	Score int `db:"score"`
	Xs    int `db:"xs_count"`
	Tens  int `db:"tens_count"`
	hits  int `db:"hits"`
}

type Team struct {
	Team_name string `db:"name"`
	Division  int    `db:"division"`
}

type Registration struct {
	Division string `db:"division"`
	Bowtype  string `db:"bow_type"`
}
