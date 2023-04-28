package impl

import (
	"fmt"
	"ojam-test/c3/s2/app/constant"
	"ojam-test/c3/s2/app/dto/request"
	"ojam-test/c3/s2/app/dto/response"
	"ojam-test/c3/s2/app/entity"
	"ojam-test/c3/s2/app/repository"
	"ojam-test/c3/s2/app/util"
	"time"
)

type ProductService struct {
	ProductRepo repository.ProductRepository
}

func ProductServiceImpl(productRepository repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepo: productRepository,
	}
}

func (svc *ProductService) Add(req request.ProductRequest, productTypeId int, userRoleId int) (response.ProductResponse, error) {
	setTime := time.Now()

	if !util.CheckSuitabilityProduct(req.Name, req.ProductType) {
		return response.ProductResponse{}, fmt.Errorf(constant.PRODUCT_NOT_MATCH)
	}

	typeId, err := util.GetProductTypeId(req.ProductType)
	if err != nil {
		return response.ProductResponse{}, err
	}

	if userRoleId != 1 && typeId != productTypeId {
		return response.ProductResponse{}, fmt.Errorf(constant.PRODUCT_NOT_AUTHORIZED)
	}

	newProduct := entity.Product{
		Name:          req.Name,
		ProductTypeId: typeId,
		Price:         req.Price,
		CreatedAt:     &setTime,
		UpdatedAt:     &setTime,
	}

	res, err := svc.ProductRepo.Create(newProduct)
	if err != nil {
		return response.ProductResponse{}, err
	}

	return generateProductResponse(&res)
}

func (svc *ProductService) UpdateById(id int, req request.ProductRequest, productTypeId int, userRoleId int) (response.ProductResponse, error) {
	setTime := time.Now()

	if !util.CheckSuitabilityProduct(req.Name, req.ProductType) {
		return response.ProductResponse{}, fmt.Errorf(constant.PRODUCT_NOT_MATCH)
	}

	typeId, err := util.GetProductTypeId(req.ProductType)
	if err != nil {
		return response.ProductResponse{}, err
	}

	if userRoleId != 1 && typeId != productTypeId {
		return response.ProductResponse{}, fmt.Errorf(constant.PRODUCT_NOT_AUTHORIZED)
	}

	updatedProduct := entity.Product{
		Name:          req.Name,
		ProductTypeId: typeId,
		Price:         req.Price,
		UpdatedAt:     &setTime,
	}

	res, err := svc.ProductRepo.UpdateById(id, updatedProduct)
	if err != nil {
		return response.ProductResponse{}, err
	}

	return generateProductResponse(&res)
}

func (svc *ProductService) DeleteById(id int, productTypeId int, userRoleId int) (response.ProductResponse, error) {
	setTime := time.Now()

	if !util.CheckSuitabilityProductTypeIdAndUserRoleId(productTypeId, userRoleId) {
		return response.ProductResponse{}, fmt.Errorf(constant.PRODUCT_NOT_AUTHORIZED)
	}

	deletedProduct := entity.Product{
		UpdatedAt: &setTime,
		DeletedAt: &setTime,
		Deleted:   true,
	}

	res, err := svc.ProductRepo.DeleteById(id, deletedProduct)
	if err != nil {
		return response.ProductResponse{}, err
	}

	return generateProductResponse(&res)
}

func (svc *ProductService) GetById(id int, productTypeId int, userRoleId int) (response.ProductResponse, error) {
	var product entity.Product
	
	if !util.CheckSuitabilityProductTypeIdAndUserRoleId(productTypeId, userRoleId) {
		return response.ProductResponse{}, fmt.Errorf(constant.PRODUCT_NOT_AUTHORIZED)
	}

	if util.CheckHeadRole(userRoleId) {
		res, err := svc.ProductRepo.FindById(id)
		if err != nil {
			return response.ProductResponse{}, err
		}

		product = res
	} else {
		res, err := svc.ProductRepo.FindByIdAndProductTypeId(id, productTypeId)
		if err != nil {
			return response.ProductResponse{}, err
		}

		product = res
	}


	return generateProductResponse(&product)
}

func (svc *ProductService) GetAll(productTypeId int, userRoleId int) ([]response.ProductResponse, error) {
	var resFinal []response.ProductResponse
	var productList []entity.Product

	if !util.CheckSuitabilityProductTypeIdAndUserRoleId(productTypeId, userRoleId) {
		return []response.ProductResponse{}, fmt.Errorf(constant.PRODUCT_NOT_AUTHORIZED)
	}

	if util.CheckHeadRole(userRoleId) {
		resList, err := svc.ProductRepo.FindAll()
		if err != nil {
			return []response.ProductResponse{}, err
		}

		productList = resList
	} else {
		resList, err := svc.ProductRepo.FindAllByProductTypeId(productTypeId)
		if err != nil {
			return []response.ProductResponse{}, err
		}

		productList = resList
	}


	for _, product := range productList {
		resFinal = append(resFinal, response.ProductResponse(product))
	}

	return resFinal, nil
}

func generateProductResponse(product *entity.Product) (response.ProductResponse, error) {
	return response.ProductResponse{
		Id:            product.Id,
		Name:          product.Name,
		ProductTypeId: product.ProductTypeId,
		Price:         product.Price,
		CreatedAt:     product.CreatedAt,
		UpdatedAt:     product.UpdatedAt,
		DeletedAt:     product.DeletedAt,
		Deleted:       product.Deleted,
	}, nil
}
