package handler

import (
	"errors"
	"net/http"

	"app/internal"
	"app/platform/web/request"
	"app/platform/web/response"
)

// NewProductsDefault returns a new ProductsDefault
func NewProductsDefault(sv internal.ServiceProduct) *ProductsDefault {
	return &ProductsDefault{sv: sv}
}

// ProductsDefault is a struct that returns the product handlers
type ProductsDefault struct {
	// sv is the product's service
	sv internal.ServiceProduct
}

// ProductJSON is a struct that represents a product in JSON format
type ProductJSON struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// GetAll returns all products
func (h *ProductsDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		p, err := h.sv.FindAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error getting products")
			return
		}

		// response
		// - serialize
		pJSON := make([]ProductJSON, len(p))
		for ix, v := range p {
			pJSON[ix] = ProductJSON{
				Id:          v.Id,
				Description: v.Description,
				Price:       v.Price,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "products found",
			"data":    pJSON,
		})
	}
}

type ProductSoldJSON struct {
	Description string `json:"description"`
	Total       int    `json:"total"`
}

// GetTopFiveBestSeller returns the top five best seller products
func (h *ProductsDefault) GetTopFiveBestSeller() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//process
		p, err := h.sv.FindTopFiveBestSeller()
		if err != nil {
			if errors.Is(err, internal.ErrProductNotFound) {
				response.Error(w, http.StatusNotFound, "error getting top five best seller products")
				return
			}
			response.Error(w, http.StatusInternalServerError, "error getting top five best seller products")
			return
		}

		// serialize
		data := make([]ProductSoldJSON, 0, len(p))
		for _, product := range p {
			data = append(data, ProductSoldJSON{
				Description: product.Description,
				Total:       product.Total,
			})
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "top five best seller products found",
			"data":    data,
		})
	}
}

// RequestBodyProduct is a struct that represents the request body for a product
type RequestBodyProduct struct {
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// Create creates a new product
func (h *ProductsDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - body
		var reqBody RequestBodyProduct
		err := request.JSON(r, &reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "error parsing request body")
			return
		}

		// process
		// - deserialize
		p := internal.Product{
			ProductAttributes: internal.ProductAttributes{
				Description: reqBody.Description,
				Price:       reqBody.Price,
			},
		}
		// - save
		err = h.sv.Save(&p)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error creating product")
			return
		}

		// response
		// - serialize
		pr := ProductJSON{
			Id:          p.Id,
			Description: p.Description,
			Price:       p.Price,
		}
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "product created",
			"data":    pr,
		})
	}
}
