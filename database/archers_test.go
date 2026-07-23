package database

import (
	"testing"

	"flyte-league/models"

	"github.com/stretchr/testify/assert"
)

func TestDBInsertArcher(t *testing.T) {
	archer := models.Archer{
		First_name: "Rin",
		Last_name:  "Shima",
		//School: models.School{
		//	Name:     "Motosu High School",
		//	Location: "Japan",
		//},
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
		},
		{
			First_name: "Nagi",
			Last_name:  "Arato",
			//School: models.School{
			//	Name:     "Maehiba Graduate School of Science",
			//	Location: "Japan",
			//},
		},
		{
			First_name: "Nagisa",
			Last_name:  "Shiota",
			//School: models.School{
			//	Name:     "Kunugigaoka Junior High School",
			//	Location: "Japan",
			//},
		},
		{
			First_name: "Klaudia",
			Last_name:  "Valentz",
			//School:     models.School{},
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
