package httpapi

import (
	"errors"
	"net/http"

	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/transfers/domain"
)

func statusFor(err error) int {
	switch {
	case errors.Is(err, domain.ErrInvalidAmount):
		return http.StatusBadRequest
	case errors.Is(err, domain.ErrAccountNotFound):
		return http.StatusNotFound
	case errors.Is(err, domain.ErrInsufficientFunds):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
