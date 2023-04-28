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

type SocialMediaService struct {
	SocialMediaRepo repository.SocialMediaRepository
}

func SocialMediaServiceImpl(socialMediaRepository repository.SocialMediaRepository) *SocialMediaService {
	return &SocialMediaService{
		SocialMediaRepo: socialMediaRepository,
	}
}

func (svc *SocialMediaService) Add(req request.SocialMediaRequest, userId int) (response.SocialMediaResponse, error) {
	setTime := time.Now()
	id := util.GenerateUUID()
	url := fmt.Sprintf(constant.BASE_URL_SOCIAL_MEDIA, userId, id)

	newSocialMedia := entity.SocialMedia{
		Name:      req.Name,
		Url:       url,
		UserId:    userId,
		CreatedAt: &setTime,
		UpdatedAt: &setTime,
	}

	res, err := svc.SocialMediaRepo.Create(newSocialMedia, userId)
	if err != nil {
		return response.SocialMediaResponse{}, err
	}

	return generateSocialMediaResponse(&res)
}

func (svc *SocialMediaService) UpdateById(id int, req request.SocialMediaRequest, userId int) (response.SocialMediaResponse, error) {
	setTime := time.Now()

	updatedSocialMedia := entity.SocialMedia{
		Name:      req.Name,
		UserId:    userId,
		UpdatedAt: &setTime,
	}

	res, err := svc.SocialMediaRepo.UpdateById(id, updatedSocialMedia, userId)
	if err != nil {
		return response.SocialMediaResponse{}, err
	}

	return generateSocialMediaResponse(&res)
}

func (svc *SocialMediaService) DeleteById(id int, userId int) (response.SocialMediaResponse, error) {
	setTime := time.Now()

	deletedSocialMedia := entity.SocialMedia{
		UpdatedAt: &setTime,
		DeletedAt: &setTime,
		Deleted:   true,
	}

	res, err := svc.SocialMediaRepo.DeleteById(id, deletedSocialMedia, userId)
	if err != nil {
		return response.SocialMediaResponse{}, err
	}

	return generateSocialMediaResponse(&res)
}

func (svc *SocialMediaService) GetById(id int, userId int) (response.SocialMediaResponse, error) {

	res, err := svc.SocialMediaRepo.FindById(id, userId)
	if err != nil {
		return response.SocialMediaResponse{}, err
	}

	return generateSocialMediaResponse(&res)
}

func (svc *SocialMediaService) GetAll(userId int) ([]response.SocialMediaResponse, error) {
	var resFinal []response.SocialMediaResponse

	resList, err := svc.SocialMediaRepo.FindAll(userId)
	if err != nil {
		return resFinal, err
	}

	for _, res := range resList {

		response := response.SocialMediaResponse{
			Id:        res.Id,
			Name:      res.Name,
			Url:       res.Url,
			UserId:    userId,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
			DeletedAt: res.DeletedAt,
			Deleted:   res.Deleted,
		}

		resFinal = append(resFinal, response)
	}

	return resFinal, nil
}

func generateSocialMediaResponse(socialMedia *entity.SocialMedia) (response.SocialMediaResponse, error) {
	return response.SocialMediaResponse{
		Id:        socialMedia.Id,
		Name:      socialMedia.Name,
		Url:       socialMedia.Url,
		UserId:    socialMedia.UserId,
		CreatedAt: socialMedia.CreatedAt,
		UpdatedAt: socialMedia.UpdatedAt,
		DeletedAt: socialMedia.DeletedAt,
		Deleted:   socialMedia.Deleted,
	}, nil
}
