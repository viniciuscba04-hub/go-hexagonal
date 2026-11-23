package ports

import "github.com/bendevjunior/hexagonal-architecture-go/internal/modules/transfers/domain"

type AccountRepository interface {
	Get(id string) (domain.Account, error)
	Save(accounts ...domain.Account) error
}

type TransfersRepository interface {
	List(id string) ([]domain.Transfers, error)
	SaveTransfers(transfers ...domain.Transfers)
}
