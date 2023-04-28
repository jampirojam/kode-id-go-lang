package request

type PhotoRequest struct {
	Title   string `json:"title" validate:"required"`
	Caption string `json:"caption"`
}
