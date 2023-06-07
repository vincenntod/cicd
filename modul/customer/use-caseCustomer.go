package customer

import "github.com/gin-gonic/gin"

type UseCase struct {
	repo *Repository
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
func (u UseCase) ShowCustomer(c *gin.Context) ([]Customer, error) {
	return u.repo.ShowCustomer(c)
}
func (u UseCase) ShowCustomerByName(c *gin.Context) ([]Customer, error) {
	return u.repo.ShowCustomerByName(c)
}
func (u UseCase) AddCustomer(c *gin.Context) (Customer, error) {
	return u.repo.AddCustomer(c)
}
func (u UseCase) EditCustomer(c *gin.Context) (Customer, error) {
	return u.repo.EditCustomer(c)
}
func (u UseCase) DeleteCustomer(c *gin.Context) (Customer, error) {
	return u.repo.DeleteCustomer(c)
}
