package handlers

import (
	"encoding/json"
	"github.com/segmentio/ksuid"
	"net/http"
	"rest-ws/models"
	"rest-ws/repository"
	"rest-ws/server"
)

type CreateProductRequest struct {
	Reference   string  `json:"reference"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type CreateProductResponse struct {
	Id        string `json:"id"`
	Reference string `json:"reference"`
	Name      string `json:"name"`
}

func CreateProductHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = CreateProductRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var product = models.Product{
			Id:          id.String(),
			Reference:   request.Reference,
			Name:        request.Name,
			Description: request.Description,
			Price:       request.Price,
			Quantity:    request.Quantity,
		}

		err = repository.InsertProduct(r.Context(), &product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var response = CreateProductResponse{
			Id:        product.Id,
			Reference: product.Reference,
			Name:      product.Name,
		}

		json.NewEncoder(w).Encode(response)
	}
}
