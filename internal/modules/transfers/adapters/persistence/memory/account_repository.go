package memory

import (
	"sync"

	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/transfers/domain"
	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/transfers/domain/ports"
)

type Accounts struct {
	mu    sync.RWMutex
	items map[string]domain.Account
}

func NewAccounts(initial ...domain.Account) *Accounts {
	items := make(map[string]domain.Account, len(initial))

	for _, account := range initial {
		items[account.ID] = account
	}

	return &Accounts{items: items}
}

func (a *Accounts) Get(id string) (domain.Account, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	account, ok := a.items[id]
	if !ok {
		return domain.Account{}, domain.ErrAccountNotFound
	}
	return account, nil
}

func (a *Accounts) Save(accounts ...domain.Account) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	for _, account := range accounts {
		if _, ok := a.items[account.ID]; !ok {
			return domain.ErrAccountNotFound
		}
	}
	for _, account := range accounts {
		a.items[account.ID] = account
	}
	return nil
}

var _ ports.AccountRepository = (*Accounts)(nil)
