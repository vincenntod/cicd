package customer

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
func (r Repository) ShowCustomer(c *gin.Context) ([]Customer, error) {
	var customers []Customer

	err := r.db.Find(&customers).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return customers, err
	}

	resp, err := http.Get("https://reqres.in/api/users?page=2")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return customers, err
	}
	defer resp.Body.Close()

	var responseData struct {
		Data []Customer `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return customers, err
	}

	for _, customer := range responseData.Data {
		if err := r.db.Where("id != ?", customer.Id).Create(&customer).Error; err != nil {

		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil ditambahkan dari API", "data": &customer})
		}
	}

	return customers, err
}
func (r Repository) ShowCustomerByName(c *gin.Context) ([]Customer, error) {
	var customer []Customer
	name := c.Param("name")
	err := r.db.Where("first_name LIKE ?", name+"%").Or("last_name LIKE ?", name+"%").Or("email LIKE ?", name+"%").Find(&customer).Error

	return customer, err

}
func (r Repository) AddCustomer(c *gin.Context) (Customer, error) {
	var customer Customer
	if err := c.BindJSON(&customer); err != nil {
		c.JSON(500, gin.H{"error": err})
		return customer, err
	}
	err := r.db.Create(&customer).Error
	return customer, err
}
func (r Repository) EditCustomer(c *gin.Context) (Customer, error) {
	var customer Customer
	id := c.Param("id")
	if err := c.BindJSON(&customer); err != nil {
		c.JSON(500, gin.H{"error": err})
		return customer, err
	}
	err := r.db.Where("id=?", id).Updates(&customer).Error
	return customer, err
}
func (r Repository) DeleteCustomer(c *gin.Context) (Customer, error) {
	var customer Customer
	id := c.Param("id")
	if err := c.BindJSON(&customer); err != nil {
		c.JSON(500, gin.H{"error": err})
		return customer, err
	}
	err := r.db.Delete(&customer, id).Error
	return customer, err
}
