package loader

import (
	"database/sql"
	"encoding/json"
	"errors"
	"os"

	"github.com/go-sql-driver/mysql"
)

// ProductAttributes is the struct that represents the attributes of a product.
type ProductAttributes struct {
	// Description is the description of the product.
	Description string
	// Price is the price of the product.
	Price float64
}

// Product is the struct that represents a product.
type Product struct {
	// Id is the unique identifier of the product.
	Id int
	// ProductAttributes is the attributes of the product.
	ProductAttributes
}

type ProductsLoader struct {
	pathFile string
	db       *sql.DB
}

func NewProductsLoader(pathFile string, db *sql.DB) *ProductsLoader {
	return &ProductsLoader{
		pathFile: pathFile,
		db:       db,
	}
}

// Get products the products from the json file.
func (p *ProductsLoader) Load() (err error) {
	// open the file
	file, err := os.Open(p.pathFile)
	if err != nil {
		return
	}
	defer file.Close()

	var products []Product
	err = json.NewDecoder(file).Decode(&products)
	if err != nil {
		return
	}

	// insert products into the database
	for _, product := range products {
		_, err = p.db.Exec(
			"INSERT INTO products VALUES (?, ?, ?)",
			product.Id, product.Description, product.Price,
		)

		if err != nil {
			var sqlError *mysql.MySQLError
			if errors.As(err, &sqlError) {
				switch sqlError.Number {
				case 1062:
					err = nil
					continue
				default:
					return
				}
			}
			return
		}
	}
	return
}
