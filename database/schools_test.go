package database

import (
	"testing"

	"flyte-league/models"

	"github.com/stretchr/testify/assert"
)

func TestDBInsertSchool(t *testing.T) {
	school := models.School{
		ID:       1,
		Name:     "Motosu High School",
		Location: "Japan",
	}
	tests := []struct {
		name           string
		schoolToInsert *models.School
	}{
		{
			name:           "Successful Insertion",
			schoolToInsert: &school,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := OpenDatabase(":memory:"); err != nil {
				t.Fatal(err)
			}
			defer CloseDatabase()

			err := DBInsertSchool(test.schoolToInsert)
			if err != nil {
				t.Error(err)
			}
			actualSchool, err := DBGetSchoolByName("Motosu High School")
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, test.schoolToInsert, actualSchool)
		})
	}
}
