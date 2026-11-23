package httpapi

import (
	"net/http"

	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/transfers/application"
	"github.com/bendevjunior/hexagonal-architecture-go/internal/platform/httpserver"
)

type Handler struct {
	service *application.TransferService
}

func NewHandler(service *application.TransferService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var input transferRequest
	if err := httpserver.ReadJSON(w, r, &input); err != nil {
		httpserver.Fail(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if err := h.service.Execute(input.FromID, input.ToID, input.Amount); err != nil {
		httpserver.Fail(w, statusFor(err), err.Error())
		return
	}

	httpserver.WriteJSON(w, http.StatusCreated, map[string]string{"message": "transfer completed"})
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	transfers, err := h.service.List(id)
	if err != nil {
		httpserver.Fail(w, statusFor(err), err.Error())
		return
	}
	httpserver.WriteJSON(w, http.StatusOK, transfers)
}
