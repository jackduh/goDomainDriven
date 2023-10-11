package services

import (
	"log"

	"github.com/google/uuid"
)

type ClinicConfiguration func(os *Clinic) error

type Clinic struct {
	//Take Orders
	OrderService *OrderService
	//Billing
	BillingService interface{}
}

func NewClinic(cfgs ...ClinicConfiguration) (*Clinic, error) {
	c := &Clinic{}

	for _, cfg := range cfgs {
		if err := cfg(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func WithOrderService(os *OrderService) ClinicConfiguration {
	return func(c *Clinic) error {
		c.OrderService = os
		return nil
	}
}

func (c *Clinic) Order(patient uuid.UUID, products []uuid.UUID) error {
	price, err := c.OrderService.CreateOrder(patient, products)

	if err != nil {
		return err
	}

	log.Printf("\nBill the patient: %0.0f\n", price)

	return nil
}
