package usecase

import (
	"backend_golang/internal/entity"
	"backend_golang/internal/repository"
	"backend_golang/utils/customerror"
	"context"
)

type ProductUsecase interface {
	GetAllProduct(ctx context.Context, f entity.ProductFilter) ([]entity.Product, error)
}

type productUsecaseImpl struct {
	productRepo repository.ProductRepo
}

func NewProductUsecase(productRepo repository.ProductRepo) *productUsecaseImpl {
	return &productUsecaseImpl{
		productRepo: productRepo,
	}
}

func (u *productUsecaseImpl) GetAllProduct(ctx context.Context, f entity.ProductFilter) ([]entity.Product, error) {
	products, err := u.productRepo.GetAllProduct(ctx, f)
	if err != nil {
		return nil, customerror.NewWithLastCustomError(customerror.ERRPRODUSECASEGETALLPROD, err)
	}
	return products, nil
}
