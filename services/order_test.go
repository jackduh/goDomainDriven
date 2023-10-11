package services

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jackduh/goDomainDriven/aggregate"
)

func init_products(t *testing.T) []aggregate.Product {
	xray, err := aggregate.NewProduct("XRay", "X-Ray Service", 311.99)
	if err != nil {
		t.Fatal(err)
	}

	cast, err := aggregate.NewProduct("Cast", "Cast Service", 411.99)
	if err != nil {
		t.Fatal(err)
	}

	motrin, err := aggregate.NewProduct("motrin", "Pain Medicine", 41.99)
	if err != nil {
		t.Fatal(err)
	}

	return []aggregate.Product{
		xray, cast, motrin,
	}
}

func TestOrder_NewOrderService(t *testing.T) {

	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryPatientRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	p, err := aggregate.NewPatient("Jack")
	if err != nil {
		t.Error(err)
	}

	err = os.patients.Add(p)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(p.GetID(), order)

	if err != nil {
		t.Error(err)
	}
}
