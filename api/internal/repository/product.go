package repository

import (
	"api/internal/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepositoryInterface {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) FindAll() ([]model.Product, error) {
	var products []model.Product
	result := r.db.Find(&products)
	return products, result.Error
}

func (r *ProductRepository) FindByID(id string) (*model.Product, error) {
	var product model.Product
	result := r.db.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, result.Error
}

func (r *ProductRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}
func (r *ProductRepository) Update(id string, product *model.Product) error {
	return r.db.Model(&model.Product{}).Where("id = ?", id).Updates(product).Error
}
func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}
