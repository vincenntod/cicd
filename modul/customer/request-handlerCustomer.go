package customer

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type RequestHandler struct {
	ctrl *Controller
}

func DefaultRequestHandler(db *gorm.DB) *RequestHandler {
	return NewRequestHandler(
		NewController(
			NewUseCase(
				NewRepository(db),
			),
		),
	)
}
func NewRequestHandler(ctrl *Controller) *RequestHandler {
	return &RequestHandler{
		ctrl: ctrl,
	}
}
func (h RequestHandler) ShowCustomer(c *gin.Context) {
	res, err := h.ctrl.ShowCustomer(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, res)

}
func (h RequestHandler) ShowCustomerByName(c *gin.Context) {
	res, err := h.ctrl.ShowCustomerByName(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, res)

}
func (h RequestHandler) AddCustomer(c *gin.Context) {
	res, err := h.ctrl.AddCustomer(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h RequestHandler) EditCustomer(c *gin.Context) {
	res, err := h.ctrl.EditCustomer(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h RequestHandler) DeleteCustomer(c *gin.Context) {
	_, err := h.ctrl.DeleteCustomer(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Terhapus"})
}
