package service

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type Customer struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Age  uint   `gorm:"not null" json:"age"`
}

type CustomerService interface {
	CreateCustomer(c echo.Context) error
	UpdateCustomer(c echo.Context) error
	DeleteCustomer(c echo.Context) error
	GetCustomer(c echo.Context) error
}
type customerService struct {
	DB *gorm.DB
}

func NewCustomerService(db *gorm.DB) CustomerService {
	return customerService{
		DB: db,
	}
}

func (s customerService) CreateCustomer(c echo.Context) error {
	customer := new(Customer)
	if err := c.Bind(customer); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	s.DB.Create(&customer)
	return c.JSON(http.StatusCreated, customer)
}

func (s customerService) UpdateCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	customer := new(Customer)
	if err := s.DB.First(&customer, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Customer not found")
	}
	if err := c.Bind(customer); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	s.DB.Save(&customer)
	return c.JSON(http.StatusOK, customer)
}

func (s customerService) DeleteCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	customer := new(Customer)
	if err := s.DB.First(&customer, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Customer not found")
	}
	s.DB.Delete(&customer)
	return c.NoContent(http.StatusNoContent)
}

func (s customerService) GetCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	customer := new(Customer)
	if err := s.DB.First(&customer, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Customer not found")
	}
	return c.JSON(http.StatusOK, customer)
}
