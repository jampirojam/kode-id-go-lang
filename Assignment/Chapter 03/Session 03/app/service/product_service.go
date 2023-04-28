package service

import (
	"ojam-test/c3/s2/app/dto/request"
	"ojam-test/c3/s2/app/dto/response"
)

type ProductService interface {
	Add(req request.ProductRequest, productTypeId int, userRoleId int) (response.ProductResponse, error)
	UpdateById(id int, req request.ProductRequest, productTypeId int, userRoleId int) (response.ProductResponse, error)
	DeleteById(id int, productTypeId int, userRoleId int) (response.ProductResponse, error)
	GetById(id int, productTypeId int, userRoleId int) (response.ProductResponse, error)
	GetAll(productTypeId int, userRoleId int) ([]response.ProductResponse, error)
}