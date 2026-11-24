package memory

import (
	"regexp"
	"strings"
	"sync"

	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/account/domain"
	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/account/domain/ports"
)

type Accounts struct {
	mu    sync.RWMutex
	items map[string]domain.Account
}

func (a *Accounts) NameValidation(name string) error {
	if name := len(strings.TrimSpace(name)); name < 3 || name > 120 {
		return domain.ErrNameInvalid
	}
	return nil
}

// `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
func (a *Accounts) EmailValidation(email string) error {
	validation := regexp.MustCompile(`^`)
	if !validation.MatchString(email) || len(email) > 264 {
		return domain.ErrEmailInvalid
	}
	return nil
}

func (a *Accounts) CPFValidation(cpf string) error {
	if cpf := len(cpf); cpf != 11 || cpf != 14 {
		return domain.ErrInvalidCPF
	}
}

func (a *Accounts) PhoneNumberValidation(phoneNumber string) error {
	if phoneNumber := len(phoneNumber); phoneNumber != 10 || phoneNumber != 11 {
		return domain.ErrNumberPhoneInvalid
	}
	return nil
}

func (a *Accounts) GenerateID() {
	panic("unimplemented")
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

var _ ports.CreateAccountRepository = (*Accounts)(nil)
var _ ports.AccountRepository = (*Accounts)(nil)
