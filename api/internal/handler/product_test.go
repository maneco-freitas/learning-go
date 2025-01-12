package handler

import (
	"api/internal/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) FindAll() ([]model.Product, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Product), args.Error(1)
}

func (m *MockProductRepository) FindByID(id string) (*model.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Product), args.Error(1)
}

func (m *MockProductRepository) Create(product *model.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Update(id string, product *model.Product) error {
	args := m.Called(id, product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(id uint) error { // uint conforme seu código
	args := m.Called(id)
	return args.Error(0)
}

func setupTest() (*gin.Engine, *ProductHandler, *MockProductRepository) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	mockRepo := new(MockProductRepository)
	handler := NewProductHandler(mockRepo)
	return router, handler, mockRepo
}

func TestGetAll(t *testing.T) {
	// Test Setup
	router, handler, mockRepo := setupTest()
	// Test Given
	router.GET("/products", handler.GetAll)
	mockProducts := []model.Product{
		{Name: "Product 1", Price: 10},
		{Name: "Product 2", Price: 20},
	}
	mockRepo.On("FindAll").Return(mockProducts, nil)
	// Test When
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/products", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	// Test Then
	var response []model.Product
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)
	mockRepo.AssertExpectations(t)
}

func TestGetById(t *testing.T) {
	// Test Setup
	router, handler, mockRepo := setupTest()
	// Test Given
	router.GET("/products/:id", handler.GetByID)
	mockProduct := model.Product{Name: "Product 1", Price: 10}
	mockRepo.On("FindByID", "1").Return(&mockProduct, nil)
	// Test When
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/products/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	// Test Then
	var response model.Product
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, mockProduct, response)
	mockRepo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	// Test Setup
	router, handler, mockRepo := setupTest()
	// Test Given
	router.POST("/products", handler.Create)
	mockProduct := model.Product{Name: "Product 1", Price: 10}
	mockRepo.On("Create", mock.AnythingOfType("*model.Product")).Return(nil)
	// Test When
	w := httptest.NewRecorder()
	jsonBody, _ := json.Marshal(mockProduct)
	req := httptest.NewRequest("POST", "/products", bytes.NewBuffer(jsonBody))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	// Test Then
	var response model.Product
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, mockProduct, response)
	mockRepo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	// Test Setup
	router, handler, mockRepo := setupTest()
	// Test Given
	router.PUT("/products/:id", handler.Update)
	mockProduct := model.Product{Name: "Product 1", Price: 10}
	mockRepo.On("Update", "1", mock.AnythingOfType("*model.Product")).Return(nil)
	// Test When
	w := httptest.NewRecorder()
	jsonBody, _ := json.Marshal(mockProduct)
	req := httptest.NewRequest("PUT", "/products/1", bytes.NewBuffer(jsonBody))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	// Test Then
	var response model.Product
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, mockProduct, response)
	mockRepo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	// Test Setup
	router, handler, mockRepo := setupTest()
	// Test Given
	router.DELETE("/products/:id", handler.Delete)
	mockRepo.On("Delete", uint(1)).Return(nil)
	// Test When
	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/products/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
	// Test Then
	mockRepo.AssertExpectations(t)
}
