package order

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/octaviookumu/myproject/internal/platform/httpx"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /orders", h.placeOrder)
	mux.HandleFunc("GET /orders/{id}", h.getOrder)
}

func (h *Handler) placeOrder(w http.ResponseWriter, r *http.Request) {
	var req PlaceOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		httpx.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	order, err := h.svc.PlaceOrder(ctx, req)
	if err != nil {
		if errors.Is(err, ErrCustomerRequired) {
			httpx.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httpx.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.JSON(w, http.StatusCreated, order)
}

func (h *Handler) getOrder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	safeID, err := uuid.Parse(id)
	if err != nil {
		httpx.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := r.Context()

	order, orderErr := h.svc.GetOrder(ctx, safeID.String())
	if orderErr != nil {
		httpx.Error(w, orderErr.Error(), http.StatusInternalServerError)
		return
	}

	httpx.JSON(w, http.StatusOK, order)
}
