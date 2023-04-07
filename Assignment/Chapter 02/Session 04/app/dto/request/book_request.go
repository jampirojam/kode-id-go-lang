package request

type BookRequest struct {
	Title          string `json:"title"`
	Author         string `json:"author"`
	Publisher      string `json:"publisher"`
	Year           int    `json:"year"`
	BookCategoryId int    `json:"book_category_id"`
}
