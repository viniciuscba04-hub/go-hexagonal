package memory

import (
	"sync"

	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/transfers/domain"
	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/transfers/domain/ports"
)

type Transfers struct {
	mu    sync.RWMutex
	items map[string][]domain.Transfers
}

func NewTransfers() *Transfers {
	return &Transfers{
		items: make(map[string][]domain.Transfers),
	}
}

func (t *Transfers) SaveTransfers(transfers ...domain.Transfers) {
	t.mu.Lock()
	defer t.mu.Unlock()

	for _, transfer := range transfers {
		t.items[transfer.FromID] = append(t.items[transfer.FromID], transfer)
		t.items[transfer.ToID] = append(t.items[transfer.ToID], transfer)
	}
}

func (t *Transfers) List(id string) ([]domain.Transfers, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if transfers, ok := t.items[id]; !ok {
		return nil, domain.ErrAccountNotFound
	} else {
		return transfers, nil
	}
}

var _ ports.TransfersRepository = (*Transfers)(nil)
