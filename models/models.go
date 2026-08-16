package models

import "time"

type School struct {
	ID       int    `db:"school_id"`
	Name     string `db:"name"`
	Location string `db:"location"`
}

type Archer struct {
	ID        int    `db:"archer_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

type ArcherWithSchool struct {
	ID        int    `db:"archer_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	School    School `db:"school"`
}

type Event struct {
	Name        string    `db:"name"`
	Description string    `db:"description"`
	StartDate   time.Time `db:"start_date"`
	EndDate     time.Time `db:"end_date"`
}

type EventDay struct {
	Date      time.Time `db:"day_date"`
	DayNumber int       `db:"day_number"`
}

type Round struct {
	Name             string `db:"name"`
	MaxPossibleScore int    `db:"max_possible_score"`
}

type Session struct {
	Name      string    `db:"name"`
	StartTime time.Time `db:"start_time"`
}

type Score struct {
	Score int `db:"score"`
	Xs    int `db:"xs_count"`
	Tens  int `db:"tens_count"`
	hits  int `db:"hits"`
}

type Team struct {
	TeamName string `db:"name"`
	Division int    `db:"division"`
}

type Registration struct {
	Division string `db:"division"`
	Bowtype  string `db:"bow_type"`
}
