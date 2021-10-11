type AccountRepository interface {
    Create(models.Account) error
    Update(models.Account) error
    Delete(int) error
    GetAll() ([]models.Account, error)
}
 
type exampleAccountRepository struct {

}

func NewAccountRepository() AccountRepository {
    return &exampleAccountRepository{}
}