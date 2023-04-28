package impl

import (
	"ojam-test/c3/s4/app/dto/request"
	"ojam-test/c3/s4/app/dto/response"
	"ojam-test/c3/s4/app/entity"
	"ojam-test/c3/s4/app/repository"
	"ojam-test/c3/s4/app/util"
	"time"
)

type CommentService struct {
	CommentRepo repository.CommentRepository
	PhotoRepo   repository.PhotoRepository
}

func CommentServiceImpl(
	photoRepository repository.PhotoRepository, 
	commentRepository repository.CommentRepository,
) *CommentService {

return &CommentService{
	PhotoRepo:   photoRepository,
	CommentRepo: commentRepository,
}
}

func (svc *CommentService) Add(req request.CommentRequest, userId int) (response.CommentResponse, error) {
	setTime := time.Now()
	id := util.GenerateUUID()

	photoRes, err := svc.PhotoRepo.FindById(req.PhotoId)
	if err != nil {
		return response.CommentResponse{}, err
	}

	url := photoRes.Url + id

	newComment := entity.Comment{
		Message:   req.Message,
		PhotoId:   req.PhotoId,
		UserId:    userId,
		Url:       url,
		CreatedAt: &setTime,
		UpdatedAt: &setTime,
	}

	res, err := svc.CommentRepo.Create(newComment, userId)
	if err != nil {
		return response.CommentResponse{}, err
	}

	return generateCommentResponse(&res)
}

func (svc *CommentService) UpdateById(id int, req request.CommentRequest, userId int) (response.CommentResponse, error) {
	setTime := time.Now()

	updatedComment := entity.Comment{
		Message:   req.Message,
		PhotoId:   req.PhotoId,
		UserId:    userId,
		UpdatedAt: &setTime,
	}

	res, err := svc.CommentRepo.UpdateById(id, updatedComment, userId)
	if err != nil {
		return response.CommentResponse{}, err
	}

	return generateCommentResponse(&res)
}

func (svc *CommentService) DeleteById(id int, userId int) (response.CommentResponse, error) {
	setTime := time.Now()

	deletedComment := entity.Comment{
		UpdatedAt: &setTime,
		DeletedAt: &setTime,
		Deleted:   true,
	}

	res, err := svc.CommentRepo.DeleteById(id, deletedComment, userId)
	if err != nil {
		return response.CommentResponse{}, err
	}

	return generateCommentResponse(&res)
}

func (svc *CommentService) GetById(id int, userId int) (response.CommentResponse, error) {

	res, err := svc.CommentRepo.FindById(id, userId)
	if err != nil {
		return response.CommentResponse{}, err
	}

	return generateCommentResponse(&res)
}

func (svc *CommentService) GetAll(userId int) ([]response.CommentResponse, error) {
	var resFinal []response.CommentResponse

	resList, err := svc.CommentRepo.FindAll(userId)
	if err != nil {
		return resFinal, err
	}

	for _, res := range resList {

		response := response.CommentResponse{
			Id:        res.Id,
			Message:   res.Message,
			PhotoId:   res.PhotoId,
			UserId:    res.UserId,
			Url:       res.Url,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
			DeletedAt: res.DeletedAt,
			Deleted:   res.Deleted,
		}

		resFinal = append(resFinal, response)
	}

	return resFinal, nil
}

func generateCommentResponse(comment *entity.Comment) (response.CommentResponse, error) {
	return response.CommentResponse{
		Id:        comment.Id,
		Message:   comment.Message,
		PhotoId:   comment.PhotoId,
		UserId:    comment.UserId,
		Url:       comment.Url,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		DeletedAt: comment.DeletedAt,
		Deleted:   comment.Deleted,
	}, nil
}
