package impl

import (
	"net/http"
	"ojam-test/c3/s2/app/config"
	"ojam-test/c3/s2/app/dto/request"
	"ojam-test/c3/s2/app/dto/response"
	"ojam-test/c3/s2/app/service"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "ojam-test/c3/s2/docs"
)

type ProductHandler struct {
	Svc service.ProductService
}

func ProductHandlerImpl(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		Svc: productService,
	}
}

// @Summary      add new product
// @Description  add new product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        request        body 	request.ProductRequest	true	"product request is mandatory"
// @Param 		 Authorization 	header 	string 	true 					"Bearer"
// @Param 		 api-secret 	header	string 	true					"api-secret"
// @Success      200  			{array} 		response.ProductResponse
// @Failure		 400			{array}			response.ErrorResponse	
// @Router       /v1/product/ 					[post]
func (ph *ProductHandler) Add(ctx *gin.Context) {
	var req request.ProductRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	productTypeId, userRoleId := config.UserAuth(ctx)
	res, err := ph.Svc.Add(req, productTypeId, userRoleId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"context" : res,
	})
}

// @Summary      update product
// @Description  update product by id
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        id   			path  	int  	true					"product_id"
// @Param        request        body 	request.ProductRequest	true	"product request is mandatory"
// @Param 		 Authorization 	header 	string 	true 					"Bearer"
// @Param 		 api-secret 	header	string 	true					"api-secret"
// @Success      200  			{array} 		response.ProductResponse
// @Failure		 400			{array}			response.ErrorResponse	
// @Router       /v1/product/{id} 		[put]
func (ph *ProductHandler) UpdateById(ctx *gin.Context) {
	var req request.ProductRequest

	paramId := ctx.Param("id")

	id, _ := strconv.Atoi(paramId)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	
	productTypeId, userRoleId := config.UserAuth(ctx)
	res, err := ph.Svc.UpdateById(id, req, productTypeId, userRoleId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"context" : res,
	})
}

// @Summary      delete product
// @Description  delete product by id
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        id   			path  	int  	true	"product_id"
// @Param 		 Authorization 	header 	string 	true 	"Bearer"
// @Param 		 api-secret 	header	string 	true	"api-secret"
// @Success      200  			{array}  		response.ProductResponse
// @Failure		 400			{array}			response.ErrorResponse	
// @Router       /v1/product/{id} 				[delete]
func (ph *ProductHandler) DeleteById(ctx *gin.Context) {
	paramId := ctx.Param("id")

	id, _ := strconv.Atoi(paramId)

	productTypeId, userRoleId := config.UserAuth(ctx)
	res, err := ph.Svc.DeleteById(id, productTypeId, userRoleId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"context" : res,
	})
}

// @Summary      get product
// @Description  get product by id
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        id   			path  	int  	true	"product_id"
// @Param 		 Authorization 	header 	string 	true 	"Bearer"
// @Param 		 api-secret 	header	string 	true	"api-secret"
// @Success      200  			{array}  		response.ProductResponse
// @Failure		 400			{array}			response.ErrorResponse	
// @Router       /v1/product/{id} 				[get]
func (ph *ProductHandler) GetById(ctx *gin.Context) {
	paramId := ctx.Param("id")

	id, _ := strconv.Atoi(paramId)

	productTypeId, userRoleId := config.UserAuth(ctx)
	res, err := ph.Svc.GetById(id, productTypeId, userRoleId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"context" : res,
	})
}

// @Summary      get all
// @Description  get all product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 Authorization 	header 	string 	true 	"Bearer"
// @Param 		 api-secret 	header	string 	true	"api-secret"
// @Success      200  			{array}  		response.ProductResponse
// @Failure		 400			{array}			response.ErrorResponse		
// @Router       /v1/product/ 					[get]
func (ph *ProductHandler) GetAll(ctx *gin.Context) {
	productTypeId, userRoleId := config.UserAuth(ctx)
	resList, err := ph.Svc.GetAll(productTypeId, userRoleId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"context" : resList,
	})
}