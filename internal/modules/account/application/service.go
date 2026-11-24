package application

import (
	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/transfers/domain"
	"github.com/bendevjunior/hexagonal-architecture-go/internal/modules/transfers/domain/ports"
)

type TransferService struct {
	accounts  ports.AccountRepository
	transfers ports.TransfersRepository
}

func NewTransferService(
	accounts ports.AccountRepository,
	transfers ports.TransfersRepository,
) *TransferService {
	return &TransferService{
		accounts:  accounts,
		transfers: transfers,
	}
}

func (s *TransferService) Execute(fromID string, toID string, amount int64) error {
	from, err := s.accounts.Get(fromID)

	if err != nil {
		return err
	}

	to, err := s.accounts.Get(toID)
	if err != nil {
		return err
	}

	if err := from.Withdraw(amount); err != nil {
		return err
	}

	to.Deposit(amount)

	result := s.accounts.Save(from, to)
	if result != nil {
		return domain.ErrAccountNotFound
	}

	transfer_register := domain.Transfers{
		FromID: fromID,
		ToID:   toID,
		Amount: amount,
	}
	s.transfers.SaveTransfers(transfer_register)

	return result
}

func (s *TransferService) List(id string) ([]domain.Transfers, error) {
	if _, err := s.accounts.Get(id); err != nil {
		return nil, err
	}

	return s.transfers.List(id)
}
