package loader

import (
	"database/sql"
	"encoding/json"
	"errors"
	"os"

	"github.com/go-sql-driver/mysql"
)

type CustomerLoader struct {
	pathFile string
	db       *sql.DB
}

type jsonCustomer struct {
	Id        int    `json:"id"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Condition int    `json:"condition"`
}

func NewCustomerLoader(pathFile string, db *sql.DB) *CustomerLoader {
	return &CustomerLoader{
		pathFile: pathFile,
		db:       db,
	}
}

// Get customers from file and insert into the database
func (c *CustomerLoader) Load() (err error) {
	// open the file
	file, err := os.Open(c.pathFile)
	if err != nil {
		return
	}
	defer file.Close()

	// decode the file into a slice of customers
	var customers []jsonCustomer
	err = json.NewDecoder(file).Decode(&customers)
	if err != nil {
		return
	}

	// insert customer into the database
	for _, customer := range customers {
		_, err = c.db.Exec(
			"INSERT INTO customers VALUES (?, ?, ?, ?)",
			customer.Id, customer.FirstName, customer.LastName, customer.Condition,
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
