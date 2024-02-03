package main

import (
	"github.com/TimeReapz/national-itmx/service"
	"github.com/labstack/echo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize Echo instance
	e := echo.New()

	// Initialize GORM instance with SQLite database
	db, err := gorm.Open(sqlite.Open("customers.db"), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal(err)
	}

	customerService := service.NewCustomerService(db)

	// Routes
	e.POST("/customers", customerService.CreateCustomer)
	e.PUT("/customers/:id", customerService.UpdateCustomer)
	e.DELETE("/customers/:id", customerService.DeleteCustomer)
	e.GET("/customers/:id", customerService.GetCustomer)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}