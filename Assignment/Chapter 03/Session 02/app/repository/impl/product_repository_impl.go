package impl

import (
	"errors"
	"fmt"
	"ojam-test/c3/s2/app/constant"
	"ojam-test/c3/s2/app/entity"
	"ojam-test/c3/s2/app/repository"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func ProductRepositoryImpl(db *gorm.DB) repository.ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (pr *ProductRepository) Create(product entity.Product) (entity.Product, error) {
	newProduct := entity.Product{
		Name:          product.Name,
		ProductTypeId: product.ProductTypeId,
		Price:         product.Price,
		CreatedAt:     product.CreatedAt,
		UpdatedAt:     product.UpdatedAt,
	}

	err := pr.DB.Create(&newProduct).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			return entity.Product{}, fmt.Errorf(constant.CREATE_PRODUCT_ERROR)
		}

		return entity.Product{}, err
	}

	return newProduct, nil
}

func (pr *ProductRepository) UpdateById(id int, product entity.Product) (entity.Product, error) {
	updatedProduct := entity.Product{}

	err := pr.DB.
		Where("id = ? AND deleted = ?", id, false).
		First(&updatedProduct).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Product{}, fmt.Errorf(constant.PRODUCT_NOT_FOUND, id)
		}

		return entity.Product{}, err
	}

	err = pr.DB.Model(&updatedProduct).
		Where("id = ? AND deleted = ?", id, false).
		Updates(entity.Product{
			Name:          product.Name,
			ProductTypeId: product.ProductTypeId,
			Price:         product.Price,
			UpdatedAt:     product.UpdatedAt,
		}).Error

	if err != nil {
		return entity.Product{}, err
	}

	return updatedProduct, nil
}

func (pr *ProductRepository) DeleteById(id int, product entity.Product) (entity.Product, error) {
	deletedProduct := entity.Product{}

	err := pr.DB.
		Where("id = ? AND deleted = ?", id, false).
		First(&deletedProduct).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Product{}, fmt.Errorf(constant.PRODUCT_NOT_FOUND, id)
		}

		return entity.Product{}, err
	}

	err = pr.DB.Model(&deletedProduct).
		Where("id = ? AND deleted = ?", id, false).
		Updates(entity.Product{
			UpdatedAt: product.UpdatedAt,
			DeletedAt: product.DeletedAt,
			Deleted:   product.Deleted,
		}).Error

	if err != nil {
		return entity.Product{}, err
	}

	return deletedProduct, nil
}

func (pr *ProductRepository) FindById(id int) (entity.Product, error) {
	findProduct := entity.Product{}

	err := pr.DB.
		Where("id = ? AND deleted = ?", id, false).
		First(&findProduct).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Product{}, fmt.Errorf(constant.PRODUCT_NOT_FOUND, id)
		}

		return entity.Product{}, err
	}

	return findProduct, nil
}

func (pr *ProductRepository) FindByIdAndProductTypeId(id, productTypeId int) (entity.Product, error) {
	findProduct := entity.Product{}

	err := pr.DB.
		Where("id = ? AND product_type_id = ? AND deleted = ?", id, productTypeId, false).
		First(&findProduct).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Product{}, fmt.Errorf(constant.PRODUCT_NOT_FOUND_OR_NOT_AUTHORIZED, id)
		}

		return entity.Product{}, err
	}

	return findProduct, nil
}

func (pr *ProductRepository) FindAll() ([]entity.Product, error) {
	productList := []entity.Product{}

	err := pr.DB.
		Where("deleted = ?", false).
		Find(&productList).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.Product{}, fmt.Errorf(constant.PRODUCT_IS_EMPTY)
		}

		return []entity.Product{}, err
	}

	return productList, nil
}

func (pr *ProductRepository) FindAllByProductTypeId(productTypeId int) ([]entity.Product, error) {
	productList := []entity.Product{}

	err := pr.DB.
		Where("product_type_id = ? AND deleted = ?", productTypeId, false).
		Find(&productList).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.Product{}, fmt.Errorf(constant.PRODUCT_IS_EMPTY)
		}

		return []entity.Product{}, err
	}

	return productList, nil
}