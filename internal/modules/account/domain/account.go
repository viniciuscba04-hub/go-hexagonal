package domain

type Account struct {
	ID          string
	Name        string
	Email       string
	PhoneNumber string
	CPF         string
	Balance     int64
}

func (a *Account) CreateAccount(
	ID string,
	Name string,
	Email string,
	PhoneNumber string,
	CPF string,
) {
	a.ID = ID
	a.Name = Name
	a.Email = Email
	a.PhoneNumber = PhoneNumber
	a.CPF = CPF
	a.Balance = 0
}
