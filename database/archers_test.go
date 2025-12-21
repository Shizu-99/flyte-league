package database

import (
	"flyte-league/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBInsertArcher(t *testing.T) {
	archer := models.Archer{
		First_name: "Rin",
		Last_name:  "Shima",
		School:     "Motosu High School",
		Bowtype:    "Recurve",
		Age:        16,
	}
	tests := []struct {
		name           string
		archerToInsert *models.Archer
	}{
		{
			name:           "Successful Insertion",
			archerToInsert: &archer,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := OpenDatabase(":memory:"); err != nil {
				t.Fatal(err)
			}
			defer CloseDatabase()

			err := DBInsertArcher(test.archerToInsert)
			if err != nil {
				t.Error(err)
			}
			actualArcher, err := DBGetArcherByFullName("Rin", "Shima")
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, test.archerToInsert, actualArcher)
		})
	}
}

func TestDBGetArcherByFullName(t *testing.T) {
	archers := []*models.Archer{
		{
			First_name: "Rin",
			Last_name:  "Shima",
			School:     "Motosu High School",
			Bowtype:    "Recurve",
			Age:        16,
		},
		{
			First_name: "Nagi",
			Last_name:  "Arato",
			School:     "Maehiba Graduate School of Science",
			Bowtype:    "Compound",
			Age:        24,
		},
		{
			First_name: "Nagisa",
			Last_name:  "Shiota",
			School:     " Kunugigaoka Junior High School",
			Bowtype:    "Recurve",
			Age:        17,
		},
		{
			First_name: "Klaudia",
			Last_name:  "Valentz",
			School:     "",
			Bowtype:    "Barebow",
			Age:        21,
		},
	}
	tests := []struct {
		name             string
		archerToRetrieve *models.Archer
		expectErr        bool
	}{
		{
			name:             "Successful Retrieval",
			archerToRetrieve: archers[2],
			expectErr:        false,
		},
		{
			name: "Archer does not exist",
			archerToRetrieve: &models.Archer{
				First_name: "Nadeshiko",
				Last_name:  "Kagamihara",
				School:     "Motosu High School",
				Bowtype:    "Compound",
				Age:        16,
			},
			expectErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := OpenDatabase(":memory:"); err != nil {
				t.Fatal(err)
			}
			defer CloseDatabase()

			if err := DBInsertMultipleArchers(archers); err != nil {
				t.Fatal(err)
			}

			archerActual, err := DBGetArcherByFullName(test.archerToRetrieve.First_name, test.archerToRetrieve.Last_name)
			if test.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, test.archerToRetrieve, archerActual)
			}
		})
	}
}
