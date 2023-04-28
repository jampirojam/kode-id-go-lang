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

type CommentHandler struct {
	Svc service.CommentService
}

func CommentHandlerImpl(photoService service.CommentService) *CommentHandler {
	return &CommentHandler{
		Svc: photoService,
	}
}

// @Summary      add new comment
// @Description  add new comment
// @Tags         comment
// @Accept       json
// @Produce      json
// @Param        request        body 	request.CommentRequest	true	"comment request is mandatory"
// @Param 		 Authorization 	header 	string 	true 					"Bearer"
// @Param 		 api-secret 	header	string 	true					"api-secret"
// @Success      200  			{array} 		response.CommentResponse
// @Failure		 400			{array}			response.ErrorResponse	
// @Router       /v1/comment/ 					[post]
func (ph *CommentHandler) Add(ctx *gin.Context) {
	var req request.CommentRequest

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

// @Summary      update comment
// @Description  update comment by id
// @Tags         comment
// @Accept       json
// @Produce      json
// @Param        id   			path  	int  	true					"comment_id"
// @Param        request        body 	request.CommentRequest	true	"comment request is mandatory"
// @Param 		 Authorization 	header 	string 	true 					"Bearer"
// @Param 		 api-secret 	header	string 	true					"api-secret"
// @Success      200  			{array} 		response.CommentResponse
// @Failure		 400			{array}			response.ErrorResponse	
// @Router       /v1/comment/{id} 		[put]
func (ph *CommentHandler) UpdateById(ctx *gin.Context) {
	var req request.CommentRequest

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

// @Summary      delete comment
// @Description  delete comment by id
// @Tags         comment
// @Accept       json
// @Produce      json
// @Param        id   			path  	int  	true	"comment_id"
// @Param 		 Authorization 	header 	string 	true 	"Bearer"
// @Param 		 api-secret 	header	string 	true	"api-secret"
// @Success      200  			{array}  		response.CommentResponse
// @Failure		 400			{array}			response.ErrorResponse	
// @Router       /v1/comment/{id} 				[delete]
func (ph *CommentHandler) DeleteById(ctx *gin.Context) {
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

// @Summary      get comment
// @Description  get comment by id
// @Tags         comment
// @Accept       json
// @Produce      json
// @Param        id   			path  	int  	true	"comment_id"
// @Param 		 Authorization 	header 	string 	true 	"Bearer"
// @Param 		 api-secret 	header	string 	true	"api-secret"
// @Success      200  			{array}  		response.CommentResponse
// @Failure		 400			{array}			response.ErrorResponse	
// @Router       /v1/comment/{id} 				[get]
func (ph *CommentHandler) GetById(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, _ := strconv.Atoi(paramId)

	userId := config.UserAuth(ctx)
	res, err := ph.Svc.GetById(id, userId)
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
// @Description  get all comment
// @Tags         comment
// @Accept       json
// @Produce      json
// @Param 		 Authorization 	header 	string 	true 	"Bearer"
// @Param 		 api-secret 	header	string 	true	"api-secret"
// @Success      200  			{array}  		response.CommentResponse
// @Failure		 400			{array}			response.ErrorResponse		
// @Router       /v1/comment/ 					[get]
func (ph *CommentHandler) GetAll(ctx *gin.Context) {
	userId := config.UserAuth(ctx)
	resList, err := ph.Svc.GetAll(userId)
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