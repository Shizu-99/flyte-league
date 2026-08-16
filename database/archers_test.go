package database

import (
	"database/sql"
	"testing"

	"flyte-league/models"

	"github.com/stretchr/testify/assert"
)

func TestDBInsertArchers(t *testing.T) {
	archers := []*models.Archer{
		{
			FirstName: "Rin",
			LastName:  "Shima",
			//School: models.School{
			//	Name:     "Motosu High School",
			//	Location: "Japan",
			//},
		},
		{
			FirstName: "Nagi",
			LastName:  "Arato",
			//School: models.School{
			//	Name:     "Maehiba Graduate School of Science",
			//	Location: "Japan",
			//},
		},
		{
			FirstName: "Nagisa",
			LastName:  "Shiota",
			//School: models.School{
			//	Name:     "Kunugigaoka Junior High School",
			//	Location: "Japan",
			//},
		},
		{
			FirstName: "Klaudia",
			LastName:  "Valentz",
			//School:     models.School{},
		},
	}
	tests := []struct {
		name            string
		archersToInsert []*models.Archer
		expected        *models.Archer
		expectErr       bool
		expectedErr     error
	}{
		{
			name:            "Successful Insertion",
			archersToInsert: archers,
			expected:        archers[0],
			expectErr:       false,
			expectedErr:     nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := OpenDatabase(":memory:"); err != nil {
				t.Fatal(err)
			}
			defer CloseDatabase()

			err := DBInsertArchers(test.archersToInsert)
			if test.expectErr {
				assert.Error(t, err)
			} else {
				if assert.NoError(t, err) {
					actualArcher, err := DBGetArcherByFullName("Rin", "Shima")
					if err != nil {
						t.Error(err)
					}
					assert.Equal(t, test.expected, actualArcher)
				}
			}
		})
	}
}

func TestDBGetArcherByFullName(t *testing.T) {
	archers := []*models.Archer{
		{
			FirstName: "Rin",
			LastName:  "Shima",
		},
		{
			FirstName: "Nagi",
			LastName:  "Arato",
		},
		{
			FirstName: "Nagisa",
			LastName:  "Shiota",
		},
		{
			FirstName: "Klaudia",
			LastName:  "Valentz",
		},
	}
	tests := []struct {
		name        string
		expected    *models.Archer
		expectErr   bool
		expectedErr error
	}{
		{
			name:        "Successful Retrieval",
			expected:    archers[2],
			expectErr:   false,
			expectedErr: nil,
		},
		{
			name: "Archer does not exist",
			expected: &models.Archer{
				FirstName: "Nadeshiko",
				LastName:  "Kagamihara",
			},
			expectErr:   true,
			expectedErr: sql.ErrNoRows,
		},
		{
			name: "Incorrect first name",
			expected: &models.Archer{
				FirstName: "Nagsia",
				LastName:  "Shiota",
			},
			expectErr:   true,
			expectedErr: sql.ErrNoRows,
		},
		{
			name: "Incorrect last name",
			expected: &models.Archer{
				FirstName: "Nagisa",
				LastName:  "Shoita",
			},
			expectErr:   true,
			expectedErr: sql.ErrNoRows,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := OpenDatabase(":memory:"); err != nil {
				t.Fatal(err)
			}
			defer CloseDatabase()

			if err := DBInsertArchers(archers); err != nil {
				t.Fatal(err)
			}

			archerActual, err := DBGetArcherByFullName(test.expected.FirstName, test.expected.LastName)
			if test.expectErr {
				if assert.Error(t, err) {
					assert.ErrorIs(t, err, test.expectedErr)
				}
			} else {
				if assert.NoError(t, err) {
					assert.Equal(t, test.expected, archerActual)
				}
			}
		})
	}
}

func TestDBGetArchersByFirstName(t *testing.T) {
	archers := []*models.Archer{
		{
			FirstName: "Rin",
			LastName:  "Shima",
		},
		{
			FirstName: "Nagi",
			LastName:  "Arato",
		},
		{
			FirstName: "Nagisa",
			LastName:  "Shiota",
		},
		{
			FirstName: "Klaudia",
			LastName:  "Valentz",
		},
	}
	tests := []struct {
		name              string
		firstNameFragment string
		expected          []models.Archer
		expectErr         bool
		expectedErr       error
	}{
		{
			name:              "Multiple people with substring start",
			firstNameFragment: "Nagi",
			expected: []models.Archer{
				{
					FirstName: "Nagi",
					LastName:  "Arato",
				},
				{
					FirstName: "Nagisa",
					LastName:  "Shiota",
				},
			},
			expectErr:   false,
			expectedErr: nil,
		},
		{
			name:              "No person with substring start",
			firstNameFragment: "Eri",
			expected:          nil,
			expectErr:         true,
			expectedErr:       sql.ErrNoRows,
		},
		{
			name:              "One person with substring start",
			firstNameFragment: "Kl",
			expected: []models.Archer{
				{
					FirstName: "Klaudia",
					LastName:  "Valentz",
				},
			},
			expectErr:   false,
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := OpenDatabase(":memory:"); err != nil {
				t.Fatal(err)
			}
			defer CloseDatabase()

			if err := DBInsertArchers(archers); err != nil {
				t.Fatal(err)
			}

			actualArchers, err := DBGetArchersByFirstName(test.firstNameFragment)
			if test.expectErr {
				if assert.Error(t, err) {
					assert.ErrorIs(t, err, test.expectedErr)
				}
			} else {
				if assert.NoError(t, err) {
					assert.ElementsMatch(t, actualArchers, test.expected)
				}
			}
		})
	}
}

func TestDBGetArchersByLastName(t *testing.T) {
	archers := []*models.Archer{
		{
			FirstName: "Rin",
			LastName:  "Shima",
		},
		{
			FirstName: "Nagi",
			LastName:  "Arato",
		},
		{
			FirstName: "Nagisa",
			LastName:  "Shiota",
		},
		{
			FirstName: "Klaudia",
			LastName:  "Valentz",
		},
	}
	tests := []struct {
		name             string
		lastNameFragment string
		expected         []models.Archer
		expectErr        bool
		expectedErr      error
	}{
		{
			name:             "Multiple people with substring start",
			lastNameFragment: "Shi",
			expected: []models.Archer{
				{
					FirstName: "Rin",
					LastName:  "Shima",
				},
				{
					FirstName: "Nagisa",
					LastName:  "Shiota",
				},
			},
			expectErr:   false,
			expectedErr: nil,
		},
		{
			name:             "No person with substring start",
			lastNameFragment: "Eri",
			expected:         nil,
			expectErr:        true,
			expectedErr:      sql.ErrNoRows,
		},
		{
			name:             "One person with substring start",
			lastNameFragment: "Val",
			expected: []models.Archer{
				{
					FirstName: "Klaudia",
					LastName:  "Valentz",
				},
			},
			expectErr:   false,
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := OpenDatabase(":memory:"); err != nil {
				t.Fatal(err)
			}
			defer CloseDatabase()

			if err := DBInsertArchers(archers); err != nil {
				t.Fatal(err)
			}

			actualArchers, err := DBGetArchersByLastName(test.lastNameFragment)
			if test.expectErr {
				if assert.Error(t, err) {
					assert.ErrorIs(t, err, test.expectedErr)
				}
			} else {
				if assert.NoError(t, err) {
					assert.ElementsMatch(t, actualArchers, test.expected)
				}
			}
		})
	}
}
