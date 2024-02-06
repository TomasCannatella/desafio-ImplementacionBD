package loader

import (
	"database/sql"
	"encoding/json"
	"errors"
	"os"

	"github.com/go-sql-driver/mysql"
)

type SaleLoader struct {
	pathFile string
	db       *sql.DB
}

type jsonSale struct {
	Id        int `json:"id"`
	ProductId int `json:"product_id"`
	InvoiceId int `json:"invoice_id"`
	Quantity  int `json:"quantity"`
}

func NewSaleLoader(pathFile string, db *sql.DB) *SaleLoader {
	return &SaleLoader{
		pathFile: pathFile,
		db:       db,
	}
}

// Get sale the sales from the json file.
func (s *SaleLoader) Load() (err error) {

	// open the file
	file, err := os.Open(s.pathFile)
	if err != nil {
		return
	}
	defer file.Close()

	// decode the file into a slice of sale
	var sales []jsonSale
	err = json.NewDecoder(file).Decode(&sales)
	if err != nil {
		return
	}

	// insert sale into the database
	for _, sale := range sales {
		_, err = s.db.Exec(
			"INSERT INTO sales VALUES (?, ?, ?, ?)",
			sale.Id, sale.ProductId, sale.InvoiceId, sale.Quantity,
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
