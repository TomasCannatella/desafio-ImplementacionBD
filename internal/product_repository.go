package internal

import "errors"

var (
	// ErrProductNotFound is the error returned when the product is not found.
	ErrProductNotFound = errors.New("repository: product not found")
)

// RepositoryProduct is the interface that wraps the basic methods that a product repository must have.
type RepositoryProduct interface {
	// FindAll returns all products saved in the database.
	FindAll() (p []Product, err error)
	// Save saves a product into the database.
	Save(p *Product) (err error)
	//FindTopFiveBestSeller returns the top five best seller products.
	FindTopFiveBestSeller() (p []ProductSold, err error)
}
