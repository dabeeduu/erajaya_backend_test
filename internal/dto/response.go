package dto

type Response struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Data       any         `json:"data,omitempty"`
}

type Pagination struct {
	PageCount  int `json:"page_count"`
	ItemCount  int `json:"item_count"`
	TotalCount int `json:"total_count"`
}
