package response

import "time"

type CommentResponse struct {
	Id        int        `json:"id"`
	Message   string     `json:"message"`
	UserId    int        `json:"user_id"`
	PhotoId   int        `json:"photo_id"`
	Url       string     `json:"url"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Deleted   bool       `json:"deleted"`
}
