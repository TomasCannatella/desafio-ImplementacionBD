package main

import (
	"database/sql"
	internal "load-data/internal"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_HOST"),
		DBName: os.Getenv("DB_NAME"),
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	products := internal.NewProductsLoader("./db/products.json", db)
	err = products.Load()
	if err != nil {
		panic(err)
	}

	invoices := internal.NewInvoiceLoader("./db/invoices.json", db)
	err = invoices.Load()
	if err != nil {
		panic(err)
	}

	customers := internal.NewCustomerLoader("./db/customers.json", db)
	err = customers.Load()
	if err != nil {
		panic(err)
	}

	sales := internal.NewSaleLoader("./db/sales.json", db)
	err = sales.Load()
	if err != nil {
		panic(err)
	}

}
