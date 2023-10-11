package product

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jackduh/goDomainDriven/aggregate"
)

var (
	ErrProductNotFound       = errors.New("the product was not found in the repositorys")
	ErrProductAlreadyExists  = errors.New("there is already such a product")
	ErrFailedToAddProduct    = errors.New("failed to add the product")
	ErrFailedToUpdateProduct = errors.New("failed to update the product")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
