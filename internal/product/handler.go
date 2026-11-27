package product

import (
	"net/http"

	"github.com/octaviookumu/myproject/internal/platform/httpx"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /products", h.list)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	products, err := h.svc.List(ctx)
	if err != nil {
		httpx.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.JSON(w, http.StatusOK, products)
}
