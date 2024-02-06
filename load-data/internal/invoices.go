package loader

import (
	"database/sql"
	"encoding/json"
	"errors"
	"os"

	"github.com/go-sql-driver/mysql"
)

type InvoiceLoader struct {
	pathFile string
	db       *sql.DB
}

type jsonInvoice struct {
	Id         int     `json:"id"`
	DateTime   string  `json:"datetime"`
	CustomerId int     `json:"customer_id"`
	Total      float64 `json:"total"`
}

func NewInvoiceLoader(pathFile string, db *sql.DB) *InvoiceLoader {
	return &InvoiceLoader{
		pathFile: pathFile,
		db:       db,
	}
}

// Get invoices from file and insert into the database
func (i *InvoiceLoader) Load() (err error) {

	// open the file
	file, err := os.Open(i.pathFile)
	if err != nil {
		return
	}
	defer file.Close()

	// decode the file into a slice of invoice
	var invoices []jsonInvoice
	err = json.NewDecoder(file).Decode(&invoices)
	if err != nil {
		return
	}

	// insert invoice into the database
	for _, invoices := range invoices {
		_, err = i.db.Exec(
			"INSERT INTO invoices VALUES (?, ?, ?, ?)",
			invoices.Id, invoices.DateTime, invoices.CustomerId, invoices.Total,
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
