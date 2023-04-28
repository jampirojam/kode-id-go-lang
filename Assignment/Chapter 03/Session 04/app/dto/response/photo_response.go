package response

import (
	"time"
)

type PhotoResponse struct {
	Id        int                    `json:"id"`
	Title     string                 `json:"title"`
	Caption   string                 `json:"caption"`
	Url       string                 `json:"url"`
	UserId    int                    `json:"user_id"`
	Comments  []PhotoCommentResponse `json:"comments"`
	CreatedAt *time.Time             `json:"created_at"`
	UpdatedAt *time.Time             `json:"updated_at"`
	DeletedAt *time.Time             `json:"deleted_at"`
	Deleted   bool                   `json:"deleted"`
}

type PhotoCommentResponse struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
	UserId  int    `json:"user_id"`
	Url     string `json:"url"`
}
