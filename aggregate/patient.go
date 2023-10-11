package aggregate

// package aggregate holds our aggrets that combimes many objects to full object
import (
	"errors"

	"github.com/google/uuid"
	"github.com/jackduh/goDomainDriven/entity"
	"github.com/jackduh/goDomainDriven/valueobject"
)

var (
	ErrInvalidPerson = errors.New("a Patient needs to have a valid name")
)

type Patient struct {
	// person is root entity of Patient
	// all fields are lowercase, not accessible outside domain (no direct access)
	// no json (not up to aggregate on how data is formated)
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction //not a pointer as it cannot change
}

// NewPatient is a factory to create a new Patient aggregate
// validate the name is not empty
func NewPatient(name string) (Patient, error) {
	if name == "" {
		return Patient{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Patient{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

func (p *Patient) GetID() uuid.UUID {
	return p.person.ID
}

func (p *Patient) SetID(id uuid.UUID) {
	if p.person == nil {
		p.person = &entity.Person{}
	}
	p.person.ID = id
}

func (p Patient) SetName(name string) {
	if p.person == nil {
		p.person = &entity.Person{}
	}
	p.person.Name = name
}

func (p *Patient) GetName() string {
	return p.person.Name
}
