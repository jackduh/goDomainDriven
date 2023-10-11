package patient

//repository pattern to be able to change repository backend
import (
	"errors"

	"github.com/google/uuid"
	"github.com/jackduh/goDomainDriven/aggregate"
)

var (
	ErrPatientNotFound       = errors.New("the patient was not found in the repositorys")
	ErrFailedToAddPatient    = errors.New("failed to add the patient")
	ErrFailedToUpdatePatient = errors.New("failed to update the patient")
)

type PatientRepository interface {
	Get(uuid.UUID) (aggregate.Patient, error)
	Add(aggregate.Patient) error
	Update(aggregate.Patient) error
}
