package service

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestDatabase() *gorm.DB {
	// Open a test database connection
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}

	_ = db.AutoMigrate(Customer{})

	customers := []Customer{
		{Name: "Test1", Age: 30},
	}
	for _, customer := range customers {
		db.Create(&customer)
	}

	return db
}

func TestCreateCustomer(t *testing.T) {
	db := setupTestDatabase()
	service := NewCustomerService(db)

	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		reqBody := []byte(`{"name":"Test2","age":30}`)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := service.CreateCustomer(c)
		actual := strings.TrimSpace(rec.Body.String())
		expected := `{"id":2,"name":"Test2","age":30}`

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expected, actual)
	})

	t.Run("ErrorBind", func(t *testing.T) {
		reqBody := []byte(``)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		_ = service.CreateCustomer(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestUpdateCustomer(t *testing.T) {
	db := setupTestDatabase()
	service := NewCustomerService(db)

	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		reqBody := []byte(`{"name":"Nick","age":34}`)
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		err := service.UpdateCustomer(c)
		actual := strings.TrimSpace(rec.Body.String())
		expected := `{"id":1,"name":"Nick","age":34}`

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, actual)
	})

	t.Run("CustomerNotFound", func(t *testing.T) {
		reqBody := []byte(`{"name":"test99","age":30}`)
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("99")

		_ = service.UpdateCustomer(c)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("ErrorBind", func(t *testing.T) {
		reqBody := []byte(``)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		_ = service.UpdateCustomer(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestDeleteCustomer(t *testing.T) {
	db := setupTestDatabase()
	service := NewCustomerService(db)

	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		err := service.DeleteCustomer(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})

	t.Run("CustomerNotFound", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("99")

		_ = service.DeleteCustomer(c)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestGetCustomer(t *testing.T) {
	db := setupTestDatabase()
	service := NewCustomerService(db)

	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		err := service.GetCustomer(c)
		actual := strings.TrimSpace(rec.Body.String())
		expected := `{"id":1,"name":"Test1","age":30}`

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, actual)
	})

	t.Run("CustomerNotFound", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("99")

		_ = service.GetCustomer(c)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}
