package repository

import (
	"api/internal/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTest() (ProductRepositoryInterface, *gorm.DB) {

	mockDB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Cria o repository com o DB mockado
	return NewProductRepository(mockDB), mockDB
}

func TestProductRepository_FindAll(t *testing.T) {
	// Test Setup
	repo, mockDB := setupTest()
	migrateError := mockDB.AutoMigrate(&model.Product{})
	if migrateError != nil {
		t.Fatal("failed to migrate database")
	}

	// Test Given
	mockProducts := []model.Product{
		{Name: "Product 1", Price: 10},
		{Name: "Product 2", Price: 20},
	}
	for _, p := range mockProducts {
		mockDB.Create(&p)
	}

	// Test When
	products, err := repo.FindAll()

	// Test Then
	assert.NoError(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 2", products[1].Name)
}

func TestProductRepository_FindById(t *testing.T) {
	// Test Setup
	repo, mockDB := setupTest()
	migrateError := mockDB.AutoMigrate(&model.Product{})
	if migrateError != nil {
		t.Fatal("failed to migrate database")
	}

	// Test Given
	testProduct := model.Product{
		Name:  "Test Product",
		Price: 99.99,
	}
	mockDB.Create(&testProduct)

	// Test When
	id := fmt.Sprintf("%d", testProduct.ID)
	product, err := repo.FindByID(id)

	// Test Then
	t.Run("Success", func(t *testing.T) {
		assert.NoError(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, testProduct.Name, product.Name)
		assert.Equal(t, testProduct.Price, product.Price)
	})
	t.Run("Not Found", func(t *testing.T) {
		product, err := repo.FindByID("999") // ID que não existe

		assert.Error(t, err)
		assert.Nil(t, product)
	})
}

func TestProductRepository_Create(t *testing.T) {
	// Test Setup
	repo, mockDB := setupTest()
	migrateError := mockDB.AutoMigrate(&model.Product{})
	if migrateError != nil {
		t.Fatal("failed to migrate database")
	}

	// Test Given
	testProduct := model.Product{
		Name:  "Test Product",
		Price: 99.99,
	}
	// Test When
	err := repo.Create(&testProduct)
	// Test Then
	var product model.Product
	mockDB.First(&product, testProduct.ID)

	t.Run("Success", func(t *testing.T) {
		assert.NoError(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, testProduct.Name, product.Name)
		assert.Equal(t, testProduct.Price, product.Price)
	})
	t.Run("Not Found", func(t *testing.T) {
		err := repo.Create(nil)

		assert.Error(t, err)
	})
}

func TestProductRepository_Update(t *testing.T) {
	// Test Setup
	repo, mockDB := setupTest()
	migrateError := mockDB.AutoMigrate(&model.Product{})
	if migrateError != nil {
		t.Fatal("failed to migrate database")
	}
	// Test Given

	initialProduct := model.Product{
		Name:  "Product 1",
		Price: 10,
	}
	mockDB.Create(&initialProduct)

	t.Run("Success", func(t *testing.T) {

		id := fmt.Sprintf("%d", initialProduct.ID)
		updatedProduct := &model.Product{
			Name:  "Updated Product",
			Price: 149.99,
		}

		err := repo.Update(id, updatedProduct)
		assert.NoError(t, err)

		var product model.Product
		mockDB.First(&product, initialProduct.ID)

		assert.Equal(t, "Updated Product", product.Name)
		assert.Equal(t, 149.99, product.Price)
	})
}

func TestProductRepository_Delete(t *testing.T) {
	// Test Setup
	repo, mockDB := setupTest()
	migrateError := mockDB.AutoMigrate(&model.Product{})
	if migrateError != nil {
		t.Fatal("failed to migrate database")
	}
	// Test Given
	initialProduct := model.Product{Name: "Product 1", Price: 10}

	// Test When
	err := repo.Delete(initialProduct.ID)

	// Test Then
	assert.NoError(t, err)
	var product model.Product
	mockDB.First(&product, initialProduct.ID)
	assert.Empty(t, product)

}
