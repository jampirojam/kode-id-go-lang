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

type SocialMediaHandler struct {
	Svc service.SocialMediaService
}

func SocialMediaHandlerImpl(photoService service.SocialMediaService) *SocialMediaHandler {
	return &SocialMediaHandler{
		Svc: photoService,
	}
}

// @Summary      add new social media
// @Description  add new social media
// @Tags         social media
// @Accept       json
// @Produce      json
// @Param        request        body 	request.SocialMediaRequest	true	"social media request is mandatory"
// @Param 		 Authorization 	header 	string 	true 					"Bearer"
// @Param 		 api-secret 	header	string 	true					"api-secret"
// @Success      200  			{array} 		response.SocialMediaResponse
// @Failure		 400			{array}			response.ErrorResponse	
// @Router       /v1/social-media/ 					[post]
func (ph *SocialMediaHandler) Add(ctx *gin.Context) {
	var req request.SocialMediaRequest

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

// @Summary      update social media
// @Description  update social media by id
// @Tags         social media
// @Accept       json
// @Produce      json
// @Param        id   			path  	int  	true					"social_media_id"
// @Param        request        body 	request.SocialMediaRequest	true	"social media request is mandatory"
// @Param 		 Authorization 	header 	string 	true 					"Bearer"
// @Param 		 api-secret 	header	string 	true					"api-secret"
// @Success      200  			{array} 		response.SocialMediaResponse
// @Failure		 400			{array}			response.ErrorResponse	
// @Router       /v1/social-media/{id} 		[put]
func (ph *SocialMediaHandler) UpdateById(ctx *gin.Context) {
	var req request.SocialMediaRequest

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

// @Summary      delete social media
// @Description  delete social media by id
// @Tags         social media
// @Accept       json
// @Produce      json
// @Param        id   			path  	int  	true	"social_media_id"
// @Param 		 Authorization 	header 	string 	true 	"Bearer"
// @Param 		 api-secret 	header	string 	true	"api-secret"
// @Success      200  			{array}  		response.SocialMediaResponse
// @Failure		 400			{array}			response.ErrorResponse	
// @Router       /v1/social-media/{id} 				[delete]
func (ph *SocialMediaHandler) DeleteById(ctx *gin.Context) {
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

// @Summary      get social media
// @Description  get social media by id
// @Tags         social media
// @Accept       json
// @Produce      json
// @Param        id   			path  	int  	true	"social_media_id"
// @Param 		 Authorization 	header 	string 	true 	"Bearer"
// @Param 		 api-secret 	header	string 	true	"api-secret"
// @Success      200  			{array}  		response.SocialMediaResponse
// @Failure		 400			{array}			response.ErrorResponse	
// @Router       /v1/social-media/{id} 				[get]
func (ph *SocialMediaHandler) GetById(ctx *gin.Context) {
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
// @Description  get all social media
// @Tags         social media
// @Accept       json
// @Produce      json
// @Param 		 Authorization 	header 	string 	true 	"Bearer"
// @Param 		 api-secret 	header	string 	true	"api-secret"
// @Success      200  			{array}  		response.SocialMediaResponse
// @Failure		 400			{array}			response.ErrorResponse		
// @Router       /v1/social-media/ 					[get]
func (ph *SocialMediaHandler) GetAll(ctx *gin.Context) {
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