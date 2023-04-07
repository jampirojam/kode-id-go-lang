package impl

import (
	"ojam-test/c2/s4/app/dto/request"
	"ojam-test/c2/s4/app/dto/response"
	"ojam-test/c2/s4/app/entity"
	"ojam-test/c2/s4/app/repository"
	"time"
)

type BookService struct {
	BookRepo repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) *BookService {
	return &BookService{
		BookRepo: bookRepository,
	}
}

func (svc *BookService) GetAll() ([]response.BookResponse, error) {
	var resFinal []response.BookResponse

	resList, err := svc.BookRepo.FindAll()
	if err != nil {
		return resFinal, err
	}

	for _, res := range resList {
		resFinal = append(resFinal, response.BookResponse(res))
	}

	return resFinal, nil
}

func (svc *BookService) UpdateById(id int, req request.BookRequest) (response.BookResponse, error) {
	updatedAt := time.Now()

	updatedBook := entity.Book{
		Title:          req.Title,
		Author:         req.Author,
		Publisher:      req.Publisher,
		Year:           req.Year,
		BookCategoryId: req.BookCategoryId,
		UpdatedAt:      &updatedAt,
	}

	res, err := svc.BookRepo.UpdateById(id, updatedBook)
	if err != nil {
		return response.BookResponse{}, err
	}

	return generateBookResponse(&res)
}

func (svc *BookService) GetById(id int) (response.BookResponse, error) {
	res, err := svc.BookRepo.FindById(id)
	if err != nil {
		return response.BookResponse{}, err
	}

	return generateBookResponse(&res)
}

func (svc *BookService) DeleteById(id int) (response.BookResponse, error) {
	setTime := time.Now()

	deletedBook := entity.Book{
		UpdatedAt: &setTime,
		DeletedAt: &setTime,
		Deleted:   true,
	}

	res, err := svc.BookRepo.DeleteById(id, deletedBook)
	if err != nil {
		return response.BookResponse{}, err
	}

	return generateBookResponse(&res)
}

func (svc *BookService) Create(req request.BookRequest) (response.BookResponse, error) {
	setTime := time.Now()

	createdBook := entity.Book{
		Title:          req.Title,
		Author:         req.Author,
		Publisher:      req.Publisher,
		Year:           req.Year,
		BookCategoryId: req.BookCategoryId,
		CreatedAt:      &setTime,
		UpdatedAt:      &setTime,
	}

	res, err := svc.BookRepo.Create(createdBook)
	if err != nil {
		return response.BookResponse{}, err
	}

	return generateBookResponse(&res)
}

func generateBookResponse(req *entity.Book) (response.BookResponse, error) {
	return response.BookResponse{
		Id:             req.Id,
		Title:          req.Title,
		Author:         req.Author,
		Publisher:      req.Publisher,
		Year:           req.Year,
		BookCategoryId: req.BookCategoryId,
		CreatedAt:      req.CreatedAt,
		UpdatedAt:      req.UpdatedAt,
		DeletedAt:      req.DeletedAt,
		Deleted:        req.Deleted,
	}, nil
}
