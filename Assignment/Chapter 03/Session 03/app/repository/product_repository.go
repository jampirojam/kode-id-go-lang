package repository

import (
	"ojam-test/c3/s2/app/entity"
)

type ProductRepository interface {
	Create(product entity.Product) (entity.Product, error)
	UpdateById(id int, product entity.Product) (entity.Product, error)
	DeleteById(id int, product entity.Product) (entity.Product, error)
	FindByIdAndProductTypeId(id, productTypeId int) (entity.Product, error)
	FindById(id int) (entity.Product, error)
	FindAllByProductTypeId(productTypeId int) ([]entity.Product, error)
	FindAll() ([]entity.Product, error)
}
