package impl

import (
	"fmt"
	"ojam-test/c3/s4/app/constant"
	"ojam-test/c3/s4/app/dto/request"
	"ojam-test/c3/s4/app/dto/response"
	"ojam-test/c3/s4/app/entity"
	"ojam-test/c3/s4/app/repository"
	"ojam-test/c3/s4/app/util"
	"time"
)

type PhotoService struct {
	PhotoRepo   repository.PhotoRepository
	CommentRepo repository.CommentRepository
}

func PhotoServiceImpl(
		photoRepository repository.PhotoRepository, 
		commentRepository repository.CommentRepository,
	) *PhotoService {

	return &PhotoService{
		PhotoRepo:   photoRepository,
		CommentRepo: commentRepository,
	}
}

func (svc *PhotoService) Add(req request.PhotoRequest, userId int) (response.PhotoResponse, error) {
	setTime := time.Now()
	id := util.GenerateUUID()
	url := fmt.Sprintf(constant.BASE_URL_PHOTO, userId, id)

	newPhoto := entity.Photo{
		Title:     req.Title,
		Caption:   req.Caption,
		Url:       url,
		UserId:    userId,
		CreatedAt: &setTime,
		UpdatedAt: &setTime,
	}

	res, err := svc.PhotoRepo.Create(newPhoto, userId)
	if err != nil {
		return response.PhotoResponse{}, err
	}

	return generatePhotoResponse(&res)
}

func (svc *PhotoService) UpdateById(id int, req request.PhotoRequest, userId int) (response.PhotoResponse, error) {
	setTime := time.Now()

	updatedPhoto := entity.Photo{
		Title:     req.Title,
		Caption:   req.Caption,
		UpdatedAt: &setTime,
	}

	res, err := svc.PhotoRepo.UpdateById(id, updatedPhoto, userId)
	if err != nil {
		return response.PhotoResponse{}, err
	}

	return generatePhotoResponse(&res)
}

func (svc *PhotoService) DeleteById(id int, userId int) (response.PhotoResponse, error) {
	setTime := time.Now()

	deletedPhoto := entity.Photo{
		UpdatedAt: &setTime,
		DeletedAt: &setTime,
		Deleted:   true,
	}

	res, err := svc.PhotoRepo.DeleteById(id, deletedPhoto, userId)
	if err != nil {
		return response.PhotoResponse{}, err
	}

	return generatePhotoResponse(&res)
}

func (svc *PhotoService) GetById(id int) (response.PhotoResponse, error) {
	var commentsResponse []response.PhotoCommentResponse

	res, err := svc.PhotoRepo.FindById(id)
	if err != nil {
		return response.PhotoResponse{}, err
	}

	comments, err := svc.CommentRepo.FindAllByPhotoId(res.Id)
	if err != nil {
		return response.PhotoResponse{}, err
	}

	for _, c := range comments {

		response := response.PhotoCommentResponse{
			Id:      c.Id,
			Message: c.Message,
			UserId:  c.UserId,
		}

		commentsResponse = append(commentsResponse, response)
	}

	return response.PhotoResponse{
		Id:        res.Id,
		Title:     res.Title,
		Caption:   res.Caption,
		Url:       res.Url,
		UserId:    res.UserId,
		Comments:  commentsResponse,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		DeletedAt: res.DeletedAt,
		Deleted:   res.Deleted,
	}, nil
}

func (svc *PhotoService) GetAll() ([]response.PhotoResponse, error) {
	var resFinal []response.PhotoResponse

	resList, err := svc.PhotoRepo.FindAll()
	if err != nil {
		return resFinal, err
	}

	for _, res := range resList {
		var commentsResponse []response.PhotoCommentResponse
		comments, err := svc.CommentRepo.FindAllByPhotoId(res.Id)
		if err != nil {
			return []response.PhotoResponse{}, err
		}

		for _, c := range comments {

			response := response.PhotoCommentResponse{
				Id:      c.Id,
				Message: c.Message,
				UserId:  c.UserId,
				Url:     c.Url,
			}

			commentsResponse = append(commentsResponse, response)
		}

		response := response.PhotoResponse{
			Id:        res.Id,
			Title:     res.Title,
			Caption:   res.Caption,
			Url:       res.Url,
			UserId:    res.UserId,
			Comments:  commentsResponse,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
			DeletedAt: res.DeletedAt,
			Deleted:   res.Deleted,
		}

		resFinal = append(resFinal, response)
	}

	return resFinal, nil
}

func generatePhotoResponse(photo *entity.Photo) (response.PhotoResponse, error) {
	return response.PhotoResponse{
		Id:        photo.Id,
		Title:     photo.Title,
		Caption:   photo.Caption,
		Url:       photo.Url,
		UserId:    photo.UserId,
		CreatedAt: photo.CreatedAt,
		UpdatedAt: photo.UpdatedAt,
		DeletedAt: photo.DeletedAt,
		Deleted:   photo.Deleted,
	}, nil
}
