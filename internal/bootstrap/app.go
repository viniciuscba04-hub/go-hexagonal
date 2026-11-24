package bootstrap

import (
	"net/http"

	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/transfers/adapters/persistence/memory"
	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/transfers/application"
	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/transfers/domain"
	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/transfers/entrypoints/httpapi"
)

func Bootstrap() http.Handler {
	accounts := memory.NewAccounts(
		domain.Account{ID: "a", Balance: 5000},
		domain.Account{ID: "b", Balance: 1000},
		domain.Account{ID: "c", Balance: 500},
	)
	transfers := memory.NewTransfers()
	transfer := application.NewTransferService(accounts, transfers)
	handler := httpapi.NewHandler(transfer)

	routeHandlers := []RouteRegister{handler}
	return NewRouter(routeHandlers...)
}
