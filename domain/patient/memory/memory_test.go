package memory

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/jackduh/goDomainDriven/aggregate"
	"github.com/jackduh/goDomainDriven/domain/patient"
)

func TestMemory_GetPatient(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	p, err := aggregate.NewPatient("jack")
	if err != nil {
		t.Fatal(err)
	}

	id := p.GetID()

	repo := MemoryRepository{
		patients: map[uuid.UUID]aggregate.Patient{
			id: p,
		},
	}

	testCases := []testCase{
		{
			name:        "no patient by id",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: patient.ErrPatientNotFound,
		},
		{
			name:        "patient by id",
			id:          id,
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemory_AddPatient(t *testing.T) {
	type testCase struct {
		name        string
		patient     string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add Patient",
			patient:     "Jack",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := MemoryRepository{
				patients: map[uuid.UUID]aggregate.Patient{},
			}

			p, err := aggregate.NewPatient(tc.patient)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(p)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(p.GetID())
			if err != nil {
				t.Fatal(err)
			}
			if found.GetID() != p.GetID() {
				t.Errorf("Expected %v, got %v", p.GetID(), found.GetID())
			}
		})
	}
}
