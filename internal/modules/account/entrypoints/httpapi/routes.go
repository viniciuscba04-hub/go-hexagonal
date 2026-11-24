package httpapi

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /transfers", h.create)
	mux.HandleFunc("GET /transfers/list/{id}", h.list)
}
