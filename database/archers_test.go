package database

import (
	"flyte-league/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBInsertArcher(t *testing.T) {
	archer := models.Archer{
		Firstname: "Rin",
		Lastname:  "Shima",
		School:    "Motosu High School",
		Bowtype:   "Recurve",
		Age:       16,
	}
	tests := []struct {
		name           string
		archerToInsert *models.Archer
		expectErr      bool
	}{
		{
			name:           "Successful Insertion",
			archerToInsert: &archer,
			expectErr:      false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := OpenDatabase(":memory:"); err != nil {
				t.Fatal(err)
			}
			defer CloseDatabase()

			err := DBInsertArcher(test.archerToInsert)
			if assert.Error(t, err) && !test.expectErr {
				t.Fatal("expected nil, but got error: ", err)
			} else if !assert.Error(t, err) && !test.expectErr {
				actualArcher, err := DBGetArcherByFullName("Rin", "Shima")
				if err != nil {
					t.Error(err)
				}
				assert.Equal(t, test.archerToInsert, actualArcher)
			}
		})
	}
}
