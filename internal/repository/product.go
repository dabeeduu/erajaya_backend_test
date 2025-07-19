package repository

import (
	"backend_golang/internal/entity"
	"context"
	"database/sql"
)

type ProductRepo interface {
	GetAllProduct(ctx context.Context, f entity.ProductFilter) ([]entity.Product, error)
}

type productRepoImpl struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *productRepoImpl {
	return &productRepoImpl{
		db: db,
	}
}

func (r *productRepoImpl) GetAllProduct(ctx context.Context, f entity.ProductFilter) ([]entity.Product, error) {
	q := `
	SELECT
		id,
		name,
		price,
		description,
		quantity
	FROM
		products
	;
	`

	db := r.db

	rows, err := db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []entity.Product{}
	for rows.Next() {
		var p entity.Product
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Price,
			&p.Description,
			&p.Quantity,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
