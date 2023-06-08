package customer

import "github.com/gin-gonic/gin"

type CustomerRepository interface {
	ShowCustomer(c *gin.Context) ([]Customer, error)
	ShowCustomerByName(c *gin.Context) ([]Customer, error)
	AddCustomer(c *gin.Context) (Customer, error)
	EditCustomer(c *gin.Context) (Customer, error)
	DeleteCustomer(c *gin.Context) (Customer, error)
}

type UseCase struct {
	repo CustomerRepository
}

func NewUseCase(repo CustomerRepository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) ShowCustomer(c *gin.Context) ([]Customer, error) {
	return u.repo.ShowCustomer(c)
}

func (u *UseCase) ShowCustomerByName(c *gin.Context) ([]Customer, error) {
	return u.repo.ShowCustomerByName(c)
}

func (u *UseCase) AddCustomer(c *gin.Context) (Customer, error) {
	return u.repo.AddCustomer(c)
}

func (u *UseCase) EditCustomer(c *gin.Context) (Customer, error) {
	return u.repo.EditCustomer(c)
}

func (u *UseCase) DeleteCustomer(c *gin.Context) (Customer, error) {
	return u.repo.DeleteCustomer(c)
}
