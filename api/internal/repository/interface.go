package repository

import "api/internal/model"

type ProductRepositoryInterface interface {
	FindAll() ([]model.Product, error)
	FindByID(id string) (*model.Product, error)
	Create(product *model.Product) error
	Update(id string, product *model.Product) error
	Delete(id uint) error // Note que aqui é uint conforme seu código
}
