package dto

type GetProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
}
