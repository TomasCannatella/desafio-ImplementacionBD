package handler_test

import (
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func init() {
	cfg := mysql.Config{
		User:   "fantasy_products",
		Passwd: "fantasy_products",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "fantasy_products_test",
	}

	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}

func TestGetTopFiveBestSeller(t *testing.T) {
	// arrange
	db, err := sql.Open("txdb", "fantasy_products_test")
	require.NoError(t, err)
	// - mock
	err = func(db *sql.DB) error {
		// - create table
		_, err := db.Exec(`
			INSERT INTO products (id,description,price) VALUES
				(1, 'product 1', 10),
				(2, 'product 2', 20),
				(3, 'product 3', 30),
				(4, 'product 4', 40),
				(5, 'product 5', 50);
			`)
		if err != nil {
			return err
		}
		_, err = db.Exec("INSERT INTO customers (id, first_name,last_name,`condition`) VALUES (1, 'John', 'Doe', 0), (2, 'Jane', 'Doe', 1);")
		if err != nil {
			return err
		}

		_, err = db.Exec("INSERT INTO invoices (id,`datetime`,customer_id,total) VALUES (1, '2021-01-01 00:00:00', 1, 100.00),(2, '2021-01-01 00:00:00', 1, 100.00);")
		if err != nil {
			return err
		}

		_, err = db.Exec(`
			INSERT INTO sales (id,quantity,invoice_id,product_id) VALUES
				(1,1, 1, 1),
				(2,2, 1, 2),
				(3,3, 1,3),
				(4,4, 1, 4),
				(5,5, 1, 5),
				(6,1, 1, 1),
				(7,2, 1 , 2),
				(8,3, 1, 3),
				(9,4, 1, 4),
				(10,5, 1, 5);
		`)
		if err != nil {
			return err
		}
		return nil
	}(db)

	if err != nil {
		t.Fatal(err)
	}
	// repository
	rp := repository.NewProductsMySQL(db)
	// service
	sv := service.NewProductsDefault(rp)
	// handler
	h := handler.NewProductsDefault(sv)
	// - request
	req := httptest.NewRequest("GET", "/products/topFiveBestSeller", nil)
	// - response
	res := httptest.NewRecorder()
	// act
	h.GetTopFiveBestSeller()(res, req)

	// assert
	expectedCode := http.StatusOK
	expectedBody := `{"data":[
		{
		 "description": "product 5", 
		 "total": 10
		}, 
		{
			"description": "product 4", 
			"total": 8
		}, 
		{
			"description": "product 3", 
			"total": 6
		}, 
		{
			"description": "product 2", 
			"total": 4
		}, 
		{
			"description": "product 1", 
			"total": 2
		}
		], 
		"message" : "top five best seller products found"}`

	expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

	require.Equal(t, expectedCode, res.Code)
	require.JSONEq(t, expectedBody, res.Body.String())
	require.Equal(t, expectedHeader, res.Header())
}
