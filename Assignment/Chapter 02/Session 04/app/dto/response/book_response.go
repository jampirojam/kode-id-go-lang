package response

import "time"

type BookResponse struct {
	Id             int        `json:"id"`
	Title          string     `json:"title"`
	Author         string     `json:"author"`
	Publisher      string     `json:"publisher"`
	Year           int        `json:"year"`
	BookCategoryId int        `json:"book_category_id"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
	Deleted        bool       `json:"deleted"`
}
