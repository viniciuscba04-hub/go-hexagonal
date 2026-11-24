package ports

import "github.com/bendevjunior/hexagonal-architecture-go/internal/modules/account/domain"

type AccountRepository interface {
	Get(id string) (domain.Account, error)
	Save(accounts ...domain.Account) error
}
