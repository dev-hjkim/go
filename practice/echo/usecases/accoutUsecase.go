type AccountUsecase interface {
    CreateAccount(models.Account) error
    UpdateAccount(models.Account) error
    DeleteAccount(int) error
    GetAllAccounts() []models.Account
}

type exampleAccountUsecase struct {
    ar repositories.AccountRepository
}

func NewAccountUsecase(ar repositories.AccountRepository) AccountUsecase {
    return &exampleAccountUsecase{ar}
}

func (s *exampleAccountUsecase) UpdateAccount(account models.Account) error {
    return s.ar.Update(account)
}