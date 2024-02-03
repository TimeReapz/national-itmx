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
	// Auto Migrate Customer struct to database
	err = db.AutoMigrate(&service.Customer{})
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Initialize mock data
	createInitialData(db)

	customerService := service.NewCustomerService(db)

	// Routes
	e.POST("/customers", customerService.CreateCustomer)
	e.PUT("/customers/:id", customerService.UpdateCustomer)
	e.DELETE("/customers/:id", customerService.DeleteCustomer)
	e.GET("/customers/:id", customerService.GetCustomer)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func createInitialData(db *gorm.DB) {
	customers := []service.Customer{
		{Name: "Tony Stark", Age: 30},
		{Name: "Black Widow", Age: 25},
		{Name: "Scarlet Witch", Age: 35},
	}
	for _, customer := range customers {
		db.Create(&customer)
	}
}
