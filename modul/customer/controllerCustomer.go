package customer

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Controller struct {
	useCase *UseCase
}
type ReadResponse struct {
	Data []CustomerResponse `json:"data"`
}
type CustomerResponse struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewController(useCase *UseCase) *Controller {
	return &Controller{
		useCase: useCase,
	}
}
func (c Controller) ShowCustomer(cxt *gin.Context) (*ReadResponse, error) {
	customers, err := c.useCase.ShowCustomer(cxt)
	if err != nil {
		return nil, err
	}

	res := &ReadResponse{}
	for _, customer := range customers {
		item := CustomerResponse{
			Id:        customer.Id,
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
			CreatedAt: customer.CreatedAt,
			UpdatedAt: customer.UpdatedAt,
		}
		res.Data = append(res.Data, item)
	}

	return res, nil
}
func (c Controller) ShowCustomerByName(ctx *gin.Context) (*ReadResponse, error) {
	customers, err := c.useCase.ShowCustomerByName(ctx)
	if err != nil {
		return nil, err
	}

	res := &ReadResponse{}
	for _, customer := range customers {
		item := CustomerResponse{
			Id:        customer.Id,
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
			CreatedAt: customer.CreatedAt,
			UpdatedAt: customer.UpdatedAt,
		}
		res.Data = append(res.Data, item)
	}

	return res, nil
}
func (c Controller) AddCustomer(ctx *gin.Context) (Customer, error) {
	customer, err := c.useCase.AddCustomer(ctx)
	if err != nil {
		return customer, err
	}
	return customer, nil
}
func (c Controller) EditCustomer(ctx *gin.Context) (Customer, error) {
	customer, err := c.useCase.EditCustomer(ctx)
	if err != nil {
		return customer, err
	}

	return customer, nil
}
func (c Controller) DeleteCustomer(ctx *gin.Context) (Customer, error) {
	customer, err := c.useCase.DeleteCustomer(ctx)
	if err != nil {
		return customer, err
	}

	return customer, nil
}
