package impl

import (
	"fmt"
	"ojam-test/c3/s2/app/constant"
	"ojam-test/c3/s2/app/dto/request"
	"ojam-test/c3/s2/app/dto/response"
	"ojam-test/c3/s2/app/entity"
	"ojam-test/c3/s2/test/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &mocks.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{ProductRepo: productRepository}

func TestProductService_GetById_NotAuth(t *testing.T) {
	actual, err := productService.GetById(2, 2, 2)

	assert.NotNil(t, err)
	assert.Empty(t, actual)
	assert.Equal(t, fmt.Sprintf(constant.PRODUCT_NOT_AUTHORIZED), err.Error())
}

func TestProductService_GetById_NotFound(t *testing.T) {
	productRepository.Mock.On("FindById", 2).
		Return(entity.Product{}, fmt.Errorf(constant.PRODUCT_NOT_FOUND, 2))

	actual, err := productService.GetById(2, 1, 1)

	assert.NotNil(t, err)
	assert.Empty(t, actual)
	assert.Equal(t, fmt.Sprintf(constant.PRODUCT_NOT_FOUND, 2), err.Error())
}

func TestProductService_GetById_Found(t *testing.T) {
	expected := entity.Product{
		Id: 1,
	}

	productRepository.Mock.On("FindById", 1).Return(expected, nil)

	actual, err := productService.GetById(1, 1, 1)

	assert.Nil(t, err)
	assert.NotEmpty(t, actual)
	assert.Equal(t, expected.Id, actual.Id)
}

func TestProductService_GetByIdAndProductTypeId_NotFound(t *testing.T) {
	productRepository.Mock.On("FindByIdAndProductTypeId", 2, 3).
		Return(entity.Product{}, fmt.Errorf(constant.PRODUCT_NOT_FOUND_OR_NOT_AUTHORIZED, 2))

	actual, err := productService.GetById(2, 3, 4)

	assert.NotNil(t, err)
	assert.Empty(t, actual)
	assert.Equal(t, fmt.Sprintf(constant.PRODUCT_NOT_FOUND_OR_NOT_AUTHORIZED, 2),  err.Error())
}

func TestProductService_GetByIdAndProductTypeId_Found(t *testing.T) {
	expected := entity.Product{
		Id: 1,
		ProductTypeId: 2,
	}

	productRepository.Mock.On("FindByIdAndProductTypeId", 1, 2).Return(expected, nil)

	actual, err := productService.GetById(1, 2, 3)
	
	assert.Nil(t, err)
	assert.NotEmpty(t, actual)
	assert.Equal(t, expected.Id, actual.Id)
	assert.Equal(t, expected.ProductTypeId, actual.ProductTypeId)
}

func TestProductService_DeleteById_NotAuth(t *testing.T) {
	actual, err := productService.DeleteById(2, 2, 2)

	assert.NotNil(t, err)
	assert.Empty(t, actual)
	assert.Equal(t, fmt.Sprintf(constant.PRODUCT_NOT_AUTHORIZED), err.Error())
}

func TestProductService_DeleteById_NotFound(t *testing.T) {
	expected := response.ProductResponse{}

	productRepository.Mock.On("DeleteById", 2).
		Return(entity.Product{}, fmt.Errorf(constant.PRODUCT_NOT_FOUND, 2))

	actual, err := productService.DeleteById(2, 3, 4)

	assert.NotNil(t, err)
	assert.Empty(t, actual)
	assert.Equal(t, fmt.Sprintf(constant.PRODUCT_NOT_FOUND, 2), err.Error())
	assert.Equal(t, expected, actual)
}

func TestProductService_DeleteById_Success(t *testing.T) {
	expected := entity.Product{
		Id: 1,
		Deleted: true,
	}

	productRepository.Mock.On("DeleteById", 1).Return(expected, nil)
	actual, err := productService.DeleteById(1, 1, 1)

	assert.Nil(t, err)
	assert.NotEmpty(t, actual)
	assert.Equal(t, expected.Id, actual.Id)
	assert.Equal(t, expected.Deleted, actual.Deleted)
}

func TestProductService_UpdateById_NotFound(t *testing.T) {
	request := request.ProductRequest{
		Name: "Product Mobil",
		ProductType: "MOBIL",
	}

	expected := response.ProductResponse{}

	productRepository.Mock.On("UpdateById", 2).
		Return(entity.Product{}, fmt.Errorf(constant.PRODUCT_NOT_FOUND, 2))

	actual, err := productService.UpdateById(2, request, 1, 1)

	assert.NotNil(t, err)
	assert.Empty(t, actual)
	assert.Equal(t, fmt.Sprintf(constant.PRODUCT_NOT_FOUND, 2), err.Error())
	assert.Equal(t, expected, actual)
}

func TestProductService_UpdateById_ProductNotMatch(t *testing.T) {
	request := request.ProductRequest{
		Name: "Product Motor",
		ProductType: "MOBIL",
	}

	expected := response.ProductResponse{}

	productRepository.Mock.On("UpdateById", 1).
		Return(entity.Product{}, fmt.Errorf(constant.PRODUCT_NOT_MATCH))
	actual, err := productService.UpdateById(1, request, 1, 1)

	assert.NotNil(t, err)
	assert.Empty(t, actual)
	assert.Equal(t, fmt.Sprintf(constant.PRODUCT_NOT_MATCH), err.Error())
	assert.Equal(t, expected, actual)
}

func TestProductService_UpdateById_NotAuth(t *testing.T) {
	request := request.ProductRequest{
		Name: "Product Mobil",
		ProductType: "MOBIL",
	}
	
	actual, err := productService.UpdateById(2, request, 2, 2)

	assert.NotNil(t, err)
	assert.Empty(t, actual)
	assert.Equal(t, fmt.Sprintf(constant.PRODUCT_NOT_AUTHORIZED), err.Error())
}

func TestProductService_UpdateById_FailRequest(t *testing.T) {
	request := request.ProductRequest{}

	actual, err := productService.UpdateById(1, request, 1, 1)

	assert.NotNil(t, err)
	assert.Empty(t, actual)
	assert.Equal(t, fmt.Sprintf(constant.PRODUCT_TYPE_NAME_NOT_FOUND), err.Error())
}

func TestProductService_UpdateById_Success(t *testing.T) {
	request := request.ProductRequest{
		Name: "Product Mobil",
		ProductType: "MOBIL",
	}

	expected := entity.Product{
		Id: 1,
		Name: "Product Mobil",
	}

	productRepository.Mock.On("UpdateById", 1).Return(expected, nil)
	actual, err := productService.UpdateById(1, request, 1, 1)

	assert.Nil(t, err)
	assert.NotEmpty(t, actual)
	assert.Equal(t, expected.Id, actual.Id)
	assert.Equal(t, expected.Name, actual.Name)
}

func TestProductService_Create_FailRequest(t *testing.T) {
	request := request.ProductRequest{}

	actual, err := productService.Add(request, 1, 1)

	assert.NotNil(t, err)
	assert.Empty(t, actual)
	assert.Equal(t, fmt.Sprintf(constant.PRODUCT_TYPE_NAME_NOT_FOUND), err.Error())
}

func TestProductService_Create_ProductNotMatch(t *testing.T) {
	request := request.ProductRequest{
		Name: "Product Motor",
		ProductType: "MOBIL",
	}

	expected := response.ProductResponse{}

	productRepository.Mock.On("Create", 1).
		Return(entity.Product{}, fmt.Errorf(constant.PRODUCT_NOT_MATCH))
	actual, err := productService.Add(request, 1, 1)

	assert.NotNil(t, err)
	assert.Empty(t, actual)
	assert.Equal(t, fmt.Sprintf(constant.PRODUCT_NOT_MATCH), err.Error())
	assert.Equal(t, expected, actual)
}

func TestProductService_Create_NotAuth(t *testing.T) {
	request := request.ProductRequest{
		Name: "Product Mobil",
		ProductType: "MOBIL",
	}
	
	actual, err := productService.Add(request, 2, 2)

	assert.NotNil(t, err)
	assert.Empty(t, actual)
	assert.Equal(t, fmt.Sprintf(constant.PRODUCT_NOT_AUTHORIZED), err.Error())
}

func TestProductService_Create_Success(t *testing.T) {
	request := request.ProductRequest{
		Name: "Product Mobil",
		ProductType: "MOBIL",
		Price: 50000,
	}

	expected := entity.Product{
		Name: "Product Mobil",
		ProductTypeId: 1,
		Price: 50000,
	}
	
	productRepository.Mock.On("Create", expected).Return(expected, nil)
	actual, err := productService.Add(request, 1, 1)

	assert.Nil(t, err)
	assert.NotEmpty(t, actual)
	assert.Equal(t, expected.Name, actual.Name)
}

func TestProductService_Create_Fail(t *testing.T) {
	request := request.ProductRequest{
		Name: "Product Mobil",
		ProductType: "MOBIL",
	}

	expected := entity.Product{
		Name: "Product Mobil",
		ProductTypeId: 1,
	}
	
	productRepository.Mock.On("Create", expected).Return(expected, nil)
	actual, err := productService.Add(request, 1, 1)

	assert.NotNil(t, err)
	assert.Empty(t, actual)
	assert.Equal(t, fmt.Sprintf(constant.CREATE_PRODUCT_ERROR), err.Error())
}

func TestProductService_GetAll_NotAuth(t *testing.T) {
	actual, err := productService.GetAll(2, 2)

	assert.NotNil(t, err)
	assert.Empty(t, actual)
	assert.Equal(t, fmt.Sprintf(constant.PRODUCT_NOT_AUTHORIZED), err.Error())
}

func TestProductService_GetAll_NotFound(t *testing.T) {
	productRepository.Mock.On("FindAll").
		Return(entity.Product{}, fmt.Errorf(constant.PRODUCT_NOT_FOUND, 2))

	actual, err := productService.GetById(2, 1, 1)

	assert.NotNil(t, err)
	assert.Empty(t, actual)
	assert.Equal(t, fmt.Sprintf(constant.PRODUCT_NOT_FOUND, 2), err.Error())
}

func TestProductService_GetAll_Found(t *testing.T) {
	expected := []entity.Product{
		{
			Id: 1,
		},
	}

	productRepository.Mock.On("FindAll").Return(expected, nil)

	actual, err := productService.GetAll(1, 1)

	assert.Nil(t, err)
	assert.NotEmpty(t, actual)
}

func TestProductService_GetAllAndProductTypeId_NotFound(t *testing.T) {
	productRepository.Mock.On("FindAllByProductTypeId", 2, 3).
		Return([]entity.Product{}, fmt.Errorf(constant.PRODUCT_IS_EMPTY))

	actual, err := productService.GetAll(2, 3)

	assert.NotNil(t, err)
	assert.Empty(t, actual)
	assert.Equal(t, fmt.Sprintf(constant.PRODUCT_IS_EMPTY),  err.Error())
}

func TestProductService_GetAllAndProductTypeId_Found(t *testing.T) {
	expected := []entity.Product{
		{
			Id: 1,
			ProductTypeId: 2,
		},
	}

	productRepository.Mock.On("FindAllByProductTypeId", 1, 2).Return(expected, nil)

	actual, err := productService.GetAll(1, 2)
	
	assert.Nil(t, err)
	assert.NotEmpty(t, actual)
}