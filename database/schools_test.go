package database

import (
	"database/sql"
	"testing"

	"flyte-league/models"

	"github.com/stretchr/testify/assert"
)

func TestDBInsertSchools(t *testing.T) {
	schools := []*models.School{
		{
			Name: "Motosu High School",
			Location: sql.NullString{
				String: "Japan",
				Valid:  true,
			},
		},
		{
			Name: "Maehiba Graduate School of Science",
			Location: sql.NullString{
				String: "Japan",
				Valid:  true,
			},
		},
		{
			Name: "Kunugigaoka Junior High School",
			Location: sql.NullString{
				String: "Japan",
				Valid:  true,
			},
		},
		{
			Name: "Motosu High School",
			Location: sql.NullString{
				String: "",
				Valid:  false,
			},
		},
	}
	expectedSchools := []models.School{
		{
			Name: "Motosu High School",
			Location: sql.NullString{
				String: "Japan",
				Valid:  true,
			},
		},
		{
			Name: "Maehiba Graduate School of Science",
			Location: sql.NullString{
				String: "Japan",
				Valid:  true,
			},
		},
		{
			Name: "Kunugigaoka Junior High School",
			Location: sql.NullString{
				String: "Japan",
				Valid:  true,
			},
		},
		{
			Name: "Motosu High School",
			Location: sql.NullString{
				String: "",
				Valid:  false,
			},
		},
		{
			Name: "Raelion Academy",
			Location: sql.NullString{
				String: "Elydes",
				Valid:  true,
			},
		},
	}
	tests := []struct {
		name           string
		schoolToInsert []*models.School
		expected       []models.School
		expectErr      bool
		expectedErr    error
	}{
		{
			name: "Successful Insertion",
			schoolToInsert: []*models.School{
				{
					Name: "Raelion Academy",
					Location: sql.NullString{
						String: "Elydes",
						Valid:  true,
					},
				},
			},
			expected:    expectedSchools,
			expectErr:   false,
			expectedErr: nil,
		},
		{
			name:           "Non-unique name and location",
			schoolToInsert: schools[:1],
			expected:       nil,
			expectErr:      true,
			expectedErr:    nil,
		},
		{
			name:           "Non-unique name with NULL location",
			schoolToInsert: schools[3:],
			expected:       nil,
			expectErr:      true,
			expectedErr:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := OpenDatabase(":memory:"); err != nil {
				t.Fatal(err)
			}
			defer CloseDatabase()

			err := DBInsertSchools(schools)
			if err != nil {
				t.Error(err)
			}

			err = DBInsertSchools(test.schoolToInsert)
			if test.expectErr {
				if assert.Error(t, err) {
					if test.expectedErr != nil {
						assert.ErrorIs(t, err, test.expectedErr)
					} else {
						assert.ErrorContains(t, err, "UNIQUE constraint failed:")
					}
				}
			} else {
				if assert.NoError(t, err) {
					actualSchools, err := DBGetAllSchools()
					if err != nil {
						t.Error(err)
					}
					assert.ElementsMatch(t, expectedSchools, actualSchools)
				}
			}
		})
	}
}
