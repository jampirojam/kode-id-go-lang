package mocks

import (
	"fmt"
	"ojam-test/c3/s2/app/constant"
	"ojam-test/c3/s2/app/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (pr *ProductRepositoryMock) FindById(id int) (entity.Product, error) {
	args := pr.Mock.Called(id)

	if id != 1 {
		return entity.Product{}, fmt.Errorf(constant.PRODUCT_NOT_FOUND, id)
	} else {
		product := args.Get(0).(entity.Product)
		return product, nil
	}
}

func (pr *ProductRepositoryMock) Create(product entity.Product) (entity.Product, error) {
	if product.Price != 0 {
		return product, nil
	} else {
		return entity.Product{}, fmt.Errorf(constant.CREATE_PRODUCT_ERROR)
	}
}

func (pr *ProductRepositoryMock) UpdateById(id int, product entity.Product) (entity.Product, error) {
	args := pr.Mock.Called(id)

	if id != 1 {
		return entity.Product{}, fmt.Errorf(constant.PRODUCT_NOT_FOUND, id)
	} else {
		ret := args.Get(0).(entity.Product)
		return ret, nil
	}
}

func (pr *ProductRepositoryMock) DeleteById(id int, product entity.Product) (entity.Product, error) {
	args := pr.Mock.Called(id)

	if id != 1 {
		return entity.Product{}, fmt.Errorf(constant.PRODUCT_NOT_FOUND, id)
	} else {
		ret := args.Get(0).(entity.Product)
		return ret, nil
	}
}

func (pr *ProductRepositoryMock) FindByIdAndProductTypeId(id, productTypeId int) (entity.Product, error) {
	args := pr.Mock.Called(id, productTypeId)

	if id != 1 && productTypeId != 2 {
		return entity.Product{}, fmt.Errorf(constant.PRODUCT_NOT_FOUND_OR_NOT_AUTHORIZED, id)
	} else {
		product := args.Get(0).(entity.Product)
		return product, nil
	}
}

func (pr *ProductRepositoryMock) FindAllByProductTypeId(productTypeId int) ([]entity.Product, error) {
	if productTypeId != 1 {
		return []entity.Product{}, fmt.Errorf(constant.PRODUCT_IS_EMPTY)
	} else {
		return []entity.Product{
			{
				Id: 1,
				ProductTypeId: 1,
			},
			{
				Id: 2,
				ProductTypeId: 1,
			},
		}, nil
	}
}

func (pr *ProductRepositoryMock) FindAll() ([]entity.Product, error) {
	return []entity.Product{
		{
			Id: 1,
			ProductTypeId: 1,
		},
		{
			Id: 2,
			ProductTypeId: 1,
		},
	}, nil
}