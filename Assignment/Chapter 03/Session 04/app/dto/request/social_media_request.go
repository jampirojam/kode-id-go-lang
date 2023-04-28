package request

type SocialMediaRequest struct {
	Name   string `json:"name" validate:"required"`
}
