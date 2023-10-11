package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/jackduh/goDomainDriven/aggregate"
	"github.com/jackduh/goDomainDriven/domain/patient"
)

// package memory is an in-memory implementation of the patient
type MemoryRepository struct {
	patients map[uuid.UUID]aggregate.Patient
	sync.Mutex
}

// factory function
func New() *MemoryRepository {
	return &MemoryRepository{
		patients: make(map[uuid.UUID]aggregate.Patient),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Patient, error) {
	if patient, ok := mr.patients[id]; ok {
		return patient, nil
	}

	return aggregate.Patient{}, patient.ErrPatientNotFound
}

func (mr *MemoryRepository) Add(p aggregate.Patient) error {
	if mr.patients == nil {
		mr.Lock()
		mr.patients = make(map[uuid.UUID]aggregate.Patient)
		mr.Unlock()
	}
	//Make sure patient isn't already in repo
	if _, ok := mr.patients[p.GetID()]; ok {
		return fmt.Errorf("patient already exists: %w", patient.ErrFailedToAddPatient)
	}
	mr.Lock()
	mr.patients[p.GetID()] = p
	mr.Unlock()
	return nil
}

func (mr *MemoryRepository) Update(p aggregate.Patient) error {
	if _, ok := mr.patients[p.GetID()]; !ok {
		return fmt.Errorf("patient does not exist: %w", patient.ErrFailedToUpdatePatient)
	}

	mr.Lock()
	mr.patients[p.GetID()] = p
	mr.Unlock()

	return nil
}
