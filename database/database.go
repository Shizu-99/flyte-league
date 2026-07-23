package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

const SCHEMA = `
PRAGMA journal_mode = WAL;
PRAGMA busy_timeout = 5000;
PRAGMA foreign_keys = ON;

-- -----------------------------------------------------------------------------
-- 1. Core Entities: Schools & Archers
-- -----------------------------------------------------------------------------

CREATE TABLE schools (
    school_id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    location TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE archers (
    archer_id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    default_school_id INTEGER NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (default_school_id) REFERENCES schools(school_id) ON DELETE SET NULL
);

-- -----------------------------------------------------------------------------
-- 2. Event Hierarchy: Events, Event Days, Rounds, and Sessions
-- -----------------------------------------------------------------------------

CREATE TABLE events (
    event_id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    CHECK (end_date >= start_date)
);

CREATE TABLE event_days (
    event_day_id INTEGER PRIMARY KEY AUTOINCREMENT,
    event_id INTEGER NOT NULL,
    day_date DATE NOT NULL,
    day_number INTEGER NOT NULL CHECK (day_number > 0),
    FOREIGN KEY (event_id) REFERENCES events(event_id) ON DELETE CASCADE,
    UNIQUE (event_id, day_date)
);

CREATE TABLE rounds (
    round_id INTEGER PRIMARY KEY AUTOINCREMENT,
    event_id INTEGER NOT NULL,
    name TEXT NOT NULL, -- e.g., "70m Qualification Round", "Elimination Round 1"
    max_possible_score INTEGER CHECK (max_possible_score > 0),
    FOREIGN KEY (event_id) REFERENCES events(event_id) ON DELETE CASCADE
);

CREATE TABLE sessions (
    session_id INTEGER PRIMARY KEY AUTOINCREMENT,
    event_id INTEGER NOT NULL,
    event_day_id INTEGER NULL,
    round_id INTEGER NULL,
    name TEXT NOT NULL, -- e.g., "Morning Session", "Flight 1"
    start_time DATETIME,
    FOREIGN KEY (event_id) REFERENCES events(event_id) ON DELETE CASCADE,
    FOREIGN KEY (event_day_id) REFERENCES event_days(event_day_id) ON DELETE SET NULL,
    FOREIGN KEY (round_id) REFERENCES rounds(round_id) ON DELETE SET NULL
);

-- -----------------------------------------------------------------------------
-- 3. Teams & Event Registrations (Participation & Metadata)
-- -----------------------------------------------------------------------------

CREATE TABLE teams (
    team_id INTEGER PRIMARY KEY AUTOINCREMENT,
    event_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    division TEXT, -- e.g., "Varsity", "Open Mixed"
    school_id INTEGER NULL,
    FOREIGN KEY (event_id) REFERENCES events(event_id) ON DELETE CASCADE,
    FOREIGN KEY (school_id) REFERENCES schools(school_id) ON DELETE SET NULL
);

CREATE TABLE event_registrations (
    registration_id INTEGER PRIMARY KEY AUTOINCREMENT,
    event_id INTEGER NOT NULL,
    archer_id INTEGER NOT NULL,
    team_id INTEGER NULL,
    school_id INTEGER NULL, -- Dynamic override per competition
    division TEXT NOT NULL, -- e.g., "U18", "Open", "50+"
    bow_type TEXT NOT NULL, -- e.g., "Recurve", "Compound", "Barebow"
    registered_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (event_id) REFERENCES events(event_id) ON DELETE CASCADE,
    FOREIGN KEY (archer_id) REFERENCES archers(archer_id) ON DELETE CASCADE,
    FOREIGN KEY (team_id) REFERENCES teams(team_id) ON DELETE SET NULL,
    FOREIGN KEY (school_id) REFERENCES schools(school_id) ON DELETE SET NULL,
    UNIQUE (event_id, archer_id)
);

-- -----------------------------------------------------------------------------
-- 4. Scoring Core
-- -----------------------------------------------------------------------------

CREATE TABLE scores (
    score_id INTEGER PRIMARY KEY AUTOINCREMENT,
    registration_id INTEGER NOT NULL,
    round_id INTEGER NULL,
    session_id INTEGER NULL,
    score INTEGER NOT NULL DEFAULT 0 CHECK (score >= 0),
    tens_count INTEGER NOT NULL DEFAULT 0 CHECK (tens_count >= 0),
    xs_count INTEGER NOT NULL DEFAULT 0 CHECK (xs_count >= 0),
	hits INTEGER NOT NULL DEFAULT 0 CHECK (hits >= 0),
    recorded_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (registration_id) REFERENCES event_registrations(registration_id) ON DELETE CASCADE,
    FOREIGN KEY (round_id) REFERENCES rounds(round_id) ON DELETE CASCADE,
    FOREIGN KEY (session_id) REFERENCES sessions(session_id) ON DELETE CASCADE,
    -- Archery Domain Integrity: X's are inner bullseyes and inherently subset of 10's
    CHECK (tens_count >= xs_count)
);

-- -----------------------------------------------------------------------------
-- 5. Performance Indexes
-- -----------------------------------------------------------------------------

CREATE INDEX idx_archers_name ON archers(last_name, first_name);
CREATE INDEX idx_event_days_event ON event_days(event_id);
CREATE INDEX idx_rounds_event ON rounds(event_id);
CREATE INDEX idx_sessions_event_day ON sessions(event_day_id);
CREATE INDEX idx_teams_event ON teams(event_id);

CREATE INDEX idx_registrations_event ON event_registrations(event_id);
CREATE INDEX idx_registrations_archer ON event_registrations(archer_id);
CREATE INDEX idx_registrations_team ON event_registrations(team_id);
CREATE INDEX idx_registrations_category ON event_registrations(event_id, division, bow_type);

CREATE INDEX idx_scores_registration ON scores(registration_id);
CREATE INDEX idx_scores_round ON scores(round_id);
CREATE INDEX idx_scores_session ON scores(session_id);
`

func OpenDatabase(dbPath string) error {
	var err error
	db, err = sqlx.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	db.MustExec(SCHEMA)
	return nil
}

func CloseDatabase() {
	if db != nil {
		db.Close()
	}
}
