package request

type CommentRequest struct {
	Message string `json:"message" validate:"required"`
	PhotoId int    `json:"photo_id" validate:"required"`
}
