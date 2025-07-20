package entity

type ProductFilter struct {
	SortBy    string
	SortOrder string
}

type Product struct {
	ID          int
	Name        string
	Price       int
	Description string
	Quantity    int
}
