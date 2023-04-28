package response

import "time"

type ProductResponse struct {
	Id            int        `json:"id"`
	Name          string     `json:"name"`
	ProductTypeId int        `json:"product_type_id"`
	Price         int        `json:"price"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	Deleted       bool       `json:"deleted"`
}
