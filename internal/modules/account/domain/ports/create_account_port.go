package ports

type CreateAccountRepository interface {
	NameValidation(name string) error
	EmailValidation(email string) error
	CPFValidation(cpf string) error
	PhoneNumberValidation(phoneNumber string) error
	GenerateID()
}
