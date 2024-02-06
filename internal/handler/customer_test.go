package handler_test

import (
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTotalAmountGroupByCondition(t *testing.T) {
	t.Run("should return total amount of customers group by condition", func(t *testing.T) {
		db, err := sql.Open("txdb", "fantasy_products_test")
		if err != nil {
			t.Fatal(err)
		}
		err = func(db *sql.DB) error {
			// - create table
			// - insert data
			_, err := db.Exec("INSERT INTO customers (id, first_name,last_name,`condition`) VALUES (1, 'John', 'Doe', 0), (2, 'Jane', 'Doe', 1)")
			if err != nil {
				return err
			}

			_, err = db.Exec("INSERT INTO invoices (id,datetime,customer_id,total) VALUES (1, '2021-01-01 00:00:00', 1, 100.00),(2, '2021-01-01 00:00:00', 2, 100.00)")
			if err != nil {
				return err
			}
			// - return nil
			return nil
		}(db)
		if err != nil {
			t.Fatal(err)
		}

		// repository
		rp := repository.NewCustomersMySQL(db)
		// service
		sv := service.NewCustomersDefault(rp)
		// handler
		hd := handler.NewCustomersDefault(sv)

		req := httptest.NewRequest("GET", "/customers/totalAmountGroupByCondition", nil)
		res := httptest.NewRecorder()

		// act
		hd.GetTotalAmountGroupByCondition()(res, req)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"data":[{"condition":"Inactivo","total_amount":100.00},{"condition":"Activo","total_amount":100.00}],"message":"total amount of customers group by condition found"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})
}

func TestActiveWithHighestAmountSpent(t *testing.T) {
	t.Run("should return active customer with highest amount spent", func(t *testing.T) {
		db, err := sql.Open("txdb", "fantasy_products_test")
		require.NoError(t, err)
		err = func(db *sql.DB) error {
			// - create table
			// - insert data
			_, err := db.Exec("INSERT INTO customers (id, first_name,last_name,`condition`) VALUES (1, 'John', 'Doe', 0), (2, 'Jane', 'Doe', 1), (3, 'John', 'Doe', 1), (4, 'John', 'Doe', 1), (5, 'John', 'Doe', 1)")
			if err != nil {
				return err
			}

			_, err = db.Exec(`INSERT INTO invoices (id,datetime,customer_id,total) 
							  VALUES (1, '2021-01-01 00:00:00', 1, 100.00),
							  		 (2, '2021-01-01 00:00:00', 2, 100.00),
									 (3, '2021-01-01 00:00:00', 3, 100.00),
									 (4, '2021-01-01 00:00:00', 4, 100.00),
									 (5, '2021-01-01 00:00:00', 5, 100.00),
									 (6, '2021-01-01 00:00:00', 1, 100.00),
									 (7, '2021-01-01 00:00:00', 2, 100.00),
									 (8, '2021-01-01 00:00:00', 3, 100.00),
									 (9, '2021-01-01 00:00:00', 4, 100.00),
									 (10, '2021-01-01 00:00:00', 5, 100.00)`)
			if err != nil {
				return err
			}
			// - return nil
			return nil
		}(db)
		require.NoError(t, err)

		// repository
		rp := repository.NewCustomersMySQL(db)
		// service
		sv := service.NewCustomersDefault(rp)
		// handler
		hd := handler.NewCustomersDefault(sv)

		req := httptest.NewRequest("GET", "/customers/activeWithHighestAmountSpent", nil)
		res := httptest.NewRecorder()

		// act
		hd.GetActiveWithHighestAmountSpent()(res, req)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"data":[
								{
									"first_name":"John",
									"last_name":"Doe",
									"total_amount":200.00
								},
								{
									"first_name":"Jane",
									"last_name":"Doe",
									"total_amount":200.00
								},
								{
									"first_name":"John",
									"last_name":"Doe",
									"total_amount":200.00
								},
								{
									"first_name":"John",
									"last_name":"Doe",
									"total_amount":200.00
								},
								{
									"first_name":"John",
									"last_name":"Doe",
									"total_amount":200.00
								}		
								],
								"message":"active customer with highest amount spent found"}
		`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())

	})
}
