package usecase

import (
	"backend_golang/internal/cache"
	"backend_golang/internal/entity"
	"backend_golang/internal/repository"
	"backend_golang/utils/customerror"
	"backend_golang/utils/errormessage"
	"context"
)

type ProductUsecase interface {
	GetAllProduct(ctx context.Context, f entity.ProductFilter) ([]entity.Product, error)
	AddProduct(ctx context.Context, p entity.Product) error
}

type productUsecaseImpl struct {
	productRepo  repository.ProductRepo
	productCache cache.ProductCache
}

func NewProductUsecase(productRepo repository.ProductRepo, cache cache.ProductCache) *productUsecaseImpl {
	return &productUsecaseImpl{
		productRepo:  productRepo,
		productCache: cache,
	}
}

func (u *productUsecaseImpl) GetAllProduct(ctx context.Context, f entity.ProductFilter) ([]entity.Product, error) {
	version, err := u.productCache.GetVersion(ctx)
	if err != nil {
		version = "1"
	}

	products, err := u.productCache.GetAll(ctx, version, f.SortBy, f.SortOrder)
	if err == nil && products != nil {
		return products, nil
	}

	products, err = u.productRepo.GetAllProduct(ctx, f)
	if err != nil {
		return nil, customerror.NewWithLastCustomError(customerror.ERRPRODUSECASEGETALLPROD, err)
	}

	if err := u.productCache.SetAll(ctx, version, f.SortBy, f.SortOrder, products); err != nil {
		return nil, customerror.New(customerror.ERRPRODUSECASEGETALLPROD, errormessage.ErrorFailToSetRedisCache, err)
	}

	return products, nil
}

func (u *productUsecaseImpl) AddProduct(ctx context.Context, p entity.Product) error {
	if err := u.productRepo.InsertProduct(ctx, p); err != nil {
		return customerror.NewWithLastCustomError(customerror.ERRUSECASEADDPROD, err)
	}

	if err := u.productCache.BumpVersion(ctx); err != nil {
		return customerror.New(customerror.ERRUSECASEADDPROD, errormessage.ErrorFailToBumpRedisVersion, err)
	}

	return nil
}
