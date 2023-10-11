package services

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackduh/goDomainDriven/aggregate"
	"github.com/jackduh/goDomainDriven/domain/patient"
	"github.com/jackduh/goDomainDriven/domain/patient/memory"
	"github.com/jackduh/goDomainDriven/domain/patient/mongo"
	"github.com/jackduh/goDomainDriven/domain/product"
	prodmem "github.com/jackduh/goDomainDriven/domain/product/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	patients patient.PatientRepository
	products product.ProductRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	//loop through all the config and apply
	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

// applies patient repository to order service
func WithPatientRepository(pt patient.PatientRepository) OrderConfiguration {
	//return a function that matches the order configuration
	return func(os *OrderService) error {
		os.patients = pt
		return nil
	}
}

func WithMemoryPatientRepository() OrderConfiguration {
	cr := memory.New()
	return WithPatientRepository(cr)
}

func WithMongoPatientRepository(ctx context.Context, connStr string) OrderConfiguration {
	return func(os *OrderService) error {
		pr, err := mongo.New(ctx, connStr)
		if err != nil {
			return err
		}
		os.patients = pr
		return nil
	}
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()

		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}

		os.products = pr
		return nil
	}
}
func (o *OrderService) CreateOrder(patientID uuid.UUID, productsIDs []uuid.UUID) (float64, error) {
	//fetch patient
	p, err := o.patients.Get(patientID)
	if err != nil {
		return 0, err
	}
	//Get each Product Repository
	var products []aggregate.Product
	var total float64

	for _, id := range productsIDs {
		prod, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, prod)
		total += prod.GetPrice()
	}

	log.Printf("Patient: %s has ordered %d products", p.GetID(), len(products))
	return total, nil
}
