package handler

import (
	"log"
	"math"
	"net/http"

	"app/internal"
	"app/platform/web/request"
	"app/platform/web/response"
)

// NewCustomersDefault returns a new CustomersDefault
func NewCustomersDefault(sv internal.ServiceCustomer) *CustomersDefault {
	return &CustomersDefault{sv: sv}
}

// CustomersDefault is a struct that returns the customer handlers
type CustomersDefault struct {
	// sv is the customer's service
	sv internal.ServiceCustomer
}

// CustomerJSON is a struct that represents a customer in JSON format
type CustomerJSON struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Condition int    `json:"condition"`
}

// GetAll returns all customers
func (h *CustomersDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		c, err := h.sv.FindAll()
		if err != nil {
			log.Println(err)
			response.Error(w, http.StatusInternalServerError, "error getting customers")
			return
		}

		// response
		// - serialize
		csJSON := make([]CustomerJSON, len(c))
		for ix, v := range c {
			csJSON[ix] = CustomerJSON{
				Id:        v.Id,
				FirstName: v.FirstName,
				LastName:  v.LastName,
				Condition: v.Condition,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "customers found",
			"data":    csJSON,
		})
	}
}

type CustomerTotalAmountGroupByConditionJSON struct {
	Condition   string  `json:"condition"`
	TotalAmount float64 `json:"total_amount"`
}

func (h *CustomersDefault) GetTotalAmountGroupByCondition() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//request
		// ...

		// process
		c, err := h.sv.FindTotalAmountGroupByCondition()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error getting total amount of customers group by condition"+err.Error())
			return
		}
		// response
		var data []CustomerTotalAmountGroupByConditionJSON
		for _, customer := range c {
			data = append(data, CustomerTotalAmountGroupByConditionJSON{
				Condition:   customer.Condition,
				TotalAmount: math.Round(customer.TotalAmount*100) / 100,
			})
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "total amount of customers group by condition found",
			"data":    data,
		})

	}
}

type CustomerActiveWithHighestAmountSpent struct {
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	TotalAmount float64 `json:"total_amount"`
}

func (h *CustomersDefault) GetActiveWithHighestAmountSpent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// process
		c, err := h.sv.FindActiveWithHighestAmountSpent()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error getting total amount of customers group by condition")
			return
		}
		// response
		var data []CustomerActiveWithHighestAmountSpent
		for _, customer := range c {
			data = append(data, CustomerActiveWithHighestAmountSpent{
				FirstName:   customer.FirstName,
				LastName:    customer.LastName,
				TotalAmount: math.Round(customer.TotalAmount*100) / 100,
			})
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "active customer with highest amount spent found",
			"data":    data,
		})

	}
}

// RequestBodyCustomer is a struct that represents the request body for a customer
type RequestBodyCustomer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Condition int    `json:"condition"`
}

// Create creates a new customer
func (h *CustomersDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - body
		var reqBody RequestBodyCustomer
		err := request.JSON(r, &reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "error deserializing request body")
			return
		}

		// process
		// - deserialize
		c := internal.Customer{
			CustomerAttributes: internal.CustomerAttributes{
				FirstName: reqBody.FirstName,
				LastName:  reqBody.LastName,
				Condition: reqBody.Condition,
			},
		}
		// - save
		err = h.sv.Save(&c)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error saving customer")
			return
		}

		// response
		// - serialize
		cs := CustomerJSON{
			Id:        c.Id,
			FirstName: c.FirstName,
			LastName:  c.LastName,
			Condition: c.Condition,
		}
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "customer created",
			"data":    cs,
		})
	}
}
