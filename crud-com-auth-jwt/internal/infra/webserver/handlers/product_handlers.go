package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Marlliton/go/crud-com-auth-jwt/internal/dto"
	"github.com/Marlliton/go/crud-com-auth-jwt/internal/entity"
	"github.com/Marlliton/go/crud-com-auth-jwt/internal/infra/database"
	"github.com/Marlliton/go/crud-com-auth-jwt/internal/infra/webserver/error_response"
	entityPKG "github.com/Marlliton/go/crud-com-auth-jwt/pkg/entity"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Create products godoc
// @Summary Create products
// @Description Create products
// @Tags products
// @Accept json
// @Produce json
// @Param request body dto.CreateProductInput true "products request"
// @Success 201
// @Failure 401 string no token found
// @Failure 400 {object} error_response.ErrorResponse
// @Failure 500 {object} error_response.ErrorResponse
// @Router /products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var produtcInput dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&produtcInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error_response.ErrorResponse{
			Success: false,
			Errors:  err,
		})
		return
	}
	p, err := entity.NewProduct(produtcInput.Name, produtcInput.Price) // NOTE: não fazer isso, user casos de uso
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error_response.ErrorResponse{
			Success: false,
			Errors:  err,
		})
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error_response.ErrorResponse{
			Success: false,
			Errors:  err,
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error_response.ErrorResponse{
			Success: false,
			Message: "ID is required",
		})
		return
	}

	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error_response.ErrorResponse{
			Success: false,
			Errors:  err,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error_response.ErrorResponse{
			Success: false,
			Message: "ID is required",
		})
		return
	}

	var productUpdated *entity.Product
	err := json.NewDecoder(r.Body).Decode(&productUpdated)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error_response.ErrorResponse{
			Success: false,
			Errors:  err,
		})
		return
	}

	productUpdated.ID, err = entityPKG.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error_response.ErrorResponse{
			Success: false,
			Errors:  err,
		})
		return
	}

	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error_response.ErrorResponse{
			Success: false,
			Errors:  err,
		})
		return
	}

	err = h.ProductDB.Update(productUpdated)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error_response.ErrorResponse{
			Success: false,
			Errors:  err,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error_response.ErrorResponse{
			Success: false,
			Message: "ID is required",
		})
		return
	}

	_, err := entityPKG.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error_response.ErrorResponse{
			Success: false,
			Errors:  err,
		})
		return
	}

	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error_response.ErrorResponse{
			Success: false,
			Errors:  err,
		})
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error_response.ErrorResponse{
			Success: false,
			Errors:  err,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 0
	}
	sort := r.URL.Query().Get("sort")

	products, err := h.ProductDB.FindAll(page, limit, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error_response.ErrorResponse{
			Success: false,
			Errors:  err,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
