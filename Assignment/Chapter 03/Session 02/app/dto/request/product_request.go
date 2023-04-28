package request

type ProductRequest struct {
	Name        string `json:"name"`
	ProductType string `json:"product_type"`
	Price       int    `json:"price"`
}
