package user

import (
	"encoding/json"
	"net/http"

	"github.com/octaviookumu/myproject/internal/platform/httpx"
)

// transport layer - http. calls the service

type Handler struct {
	// we add the dependencies here
	svc *Service // dep is a service, a pointer to the service struct
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

// methods
func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	usersList, err := h.svc.List(ctx)
	if err != nil {
		httpx.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httpx.JSON(w, http.StatusOK, usersList)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")
	ctx := r.Context()

	user, err := h.svc.Get(ctx, userId)
	if err != nil {
		httpx.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.JSON(w, http.StatusOK, user)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	decodeError := json.NewDecoder(r.Body).Decode(&req) // JSON -> Go (Decode)

	if decodeError != nil {
		httpx.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	newUser, err := h.svc.Create(ctx, req)
	if err != nil {
		httpx.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.JSON(w, http.StatusCreated, newUser)
}
