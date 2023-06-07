package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"mini-project-2/modul/actors"
	"mini-project-2/modul/customer"
)

func ConfigDatabase() (*gorm.DB, error) {
	dsn := "root:password@tcp(db:3306)/app?parseTime=true"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
}

func main() {
	db, err := ConfigDatabase()
	if err != nil {
		fmt.Println("Connect Database Error")
	}
	actorsHandler := actors.DefaultRequestHandler(db)
	customerHandler := customer.DefaultRequestHandler(db)
	r := gin.Default()

	r.POST("api/register/", actorsHandler.RegisActors)
	r.POST("api/login/", actorsHandler.LoginActors)
	r.GET("api/logout/", actorsHandler.Logout)
	Admin := r.Group("/api", actorsHandler.MiddlewareAdmin)
	{
		Admin.GET("customer/", customerHandler.ShowCustomer)
		Admin.GET("customer/:name", customerHandler.ShowCustomerByName)
		Admin.POST("customer/", customerHandler.AddCustomer)
		Admin.PUT("customer/:id", customerHandler.EditCustomer)
		Admin.DELETE("customer/:id", customerHandler.DeleteCustomer)
	}
	SuperAdmin := r.Group("/admin", actorsHandler.MiddlewareSuperAdmin)
	{
		SuperAdmin.GET("actors/", actorsHandler.ShowActors)
		SuperAdmin.GET("actors/:name", actorsHandler.ShowActorsByName)
		SuperAdmin.PUT("actors/:id", actorsHandler.EditActors)
		SuperAdmin.DELETE("actors/:id", actorsHandler.DeleteActors)

		SuperAdmin.GET("approval/", actorsHandler.ShowRegister)
		SuperAdmin.PUT("approval/:id", actorsHandler.ApprovalRegister)
	}

	r.Run(":8080")

}
