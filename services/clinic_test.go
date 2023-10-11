package services

import (
	//"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jackduh/goDomainDriven/aggregate"
)

func Test_Clinic(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryPatientRepository(),
		//WithMongoPatientRepository(context.Background(), "mongodb://localhost:27017"),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	clinic, err := NewClinic(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	patient, err := aggregate.NewPatient("Jack")
	if err != nil {
		t.Fatal(err)
	}

	if err = os.patients.Add(patient); err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = clinic.Order(patient.GetID(), order)

	if err != nil {
		t.Fatal(err)
	}
}
