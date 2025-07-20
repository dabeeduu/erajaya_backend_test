package dto

type GetProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type AddProductRequest struct {
	Name        string `json:"name" binding:"required,notblank"`
	Price       int    `json:"price" binding:"gt=0"`
	Description string `json:"description" binding:"required,notblank"`
	Quantity    int    `json:"quantity" binding:"gte=0"`
}
