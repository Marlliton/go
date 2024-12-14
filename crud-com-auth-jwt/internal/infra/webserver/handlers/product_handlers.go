package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Marlliton/go/crud-com-auth-jwt/internal/dto"
	"github.com/Marlliton/go/crud-com-auth-jwt/internal/entity"
	"github.com/Marlliton/go/crud-com-auth-jwt/internal/infra/database"
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

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var produtcInput dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&produtcInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(produtcInput.Name, produtcInput.Price) // NOTE: não fazer isso, user casos de uso
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var productUpdated *entity.Product
	err := json.NewDecoder(r.Body).Decode(&productUpdated)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	productUpdated.ID, err = entityPKG.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Update(productUpdated)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := entityPKG.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
