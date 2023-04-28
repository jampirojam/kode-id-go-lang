package impl

import (
	"net/http"
	"ojam-test/c3/s4/app/config"
	"ojam-test/c3/s4/app/dto/request"
	"ojam-test/c3/s4/app/dto/response"
	"ojam-test/c3/s4/app/service"
	"ojam-test/c3/s4/app/util"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "ojam-test/c3/s4/docs"
)

type PhotoHandler struct {
	Svc service.PhotoService
}

func PhotoHandlerImpl(photoService service.PhotoService) *PhotoHandler {
	return &PhotoHandler{
		Svc: photoService,
	}
}

// @Summary      add new photo
// @Description  add new photo
// @Tags         photo
// @Accept       json
// @Produce      json
// @Param        request        body 	request.PhotoRequest	true	"photo request is mandatory"
// @Param 		 Authorization 	header 	string 	true 					"Bearer"
// @Param 		 api-secret 	header	string 	true					"api-secret"
// @Success      200  			{array} 		response.PhotoResponse
// @Failure		 400			{array}			response.ErrorResponse	
// @Router       /v1/photo/ 					[post]
func (ph *PhotoHandler) Add(ctx *gin.Context) {
	var req request.PhotoRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if err := util.ValidateJSON(ctx, req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Error: err,
		})
		return
	}
	
	userId := config.UserAuth(ctx)
	res, err := ph.Svc.Add(req, userId)
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

// @Summary      update photo
// @Description  update photo by id
// @Tags         photo
// @Accept       json
// @Produce      json
// @Param        id   			path  	int  	true					"photo_id"
// @Param        request        body 	request.PhotoRequest	true	"photo request is mandatory"
// @Param 		 Authorization 	header 	string 	true 					"Bearer"
// @Param 		 api-secret 	header	string 	true					"api-secret"
// @Success      200  			{array} 		response.PhotoResponse
// @Failure		 400			{array}			response.ErrorResponse	
// @Router       /v1/photo/{id} 		[put]
func (ph *PhotoHandler) UpdateById(ctx *gin.Context) {
	var req request.PhotoRequest

	paramId := ctx.Param("id")
	id, _ := strconv.Atoi(paramId)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	
	if err := util.ValidateJSON(ctx, req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Error: err,
		})
		return
	}

	userId := config.UserAuth(ctx)
	res, err := ph.Svc.UpdateById(id, req, userId)
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

// @Summary      delete photo
// @Description  delete photo by id
// @Tags         photo
// @Accept       json
// @Produce      json
// @Param        id   			path  	int  	true	"photo_id"
// @Param 		 Authorization 	header 	string 	true 	"Bearer"
// @Param 		 api-secret 	header	string 	true	"api-secret"
// @Success      200  			{array}  		response.PhotoResponse
// @Failure		 400			{array}			response.ErrorResponse	
// @Router       /v1/photo/{id} 				[delete]
func (ph *PhotoHandler) DeleteById(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, _ := strconv.Atoi(paramId)

	userId := config.UserAuth(ctx)
	res, err := ph.Svc.DeleteById(id, userId)
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

// @Summary      get photo
// @Description  get photo by id
// @Tags         photo
// @Accept       json
// @Produce      json
// @Param        id   			path  	int  	true	"photo_id"
// @Param 		 Authorization 	header 	string 	true 	"Bearer"
// @Param 		 api-secret 	header	string 	true	"api-secret"
// @Success      200  			{array}  		response.PhotoResponse
// @Failure		 400			{array}			response.ErrorResponse	
// @Router       /v1/photo/{id} 				[get]
func (ph *PhotoHandler) GetById(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, _ := strconv.Atoi(paramId)

	res, err := ph.Svc.GetById(id)
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
// @Description  get all photo
// @Tags         photo
// @Accept       json
// @Produce      json
// @Param 		 Authorization 	header 	string 	true 	"Bearer"
// @Param 		 api-secret 	header	string 	true	"api-secret"
// @Success      200  			{array}  		response.PhotoResponse
// @Failure		 400			{array}			response.ErrorResponse		
// @Router       /v1/photo/ 					[get]
func (ph *PhotoHandler) GetAll(ctx *gin.Context) {
	resList, err := ph.Svc.GetAll()
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