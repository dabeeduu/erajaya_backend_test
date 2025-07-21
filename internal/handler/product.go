package handler

import (
	"backend_golang/internal/dto"
	"backend_golang/internal/entity"
	"backend_golang/internal/usecase"
	"backend_golang/utils"
	"backend_golang/utils/customerror"
	"backend_golang/utils/errormessage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) *productHandler {
	return &productHandler{
		productUsecase: productUsecase,
	}
}

func (h *productHandler) ListAllProduct(c *gin.Context) {
	var filterDTO dto.GetProductFilter
	if err := c.ShouldBindQuery(&filterDTO); err != nil {
		c.Error(customerror.New(customerror.ERRPRODHANDLERLISTALLPRODBIND, errormessage.ErrorInvalidQueryParam, err))
		return
	}

	var f entity.ProductFilter
	if filterDTO.SortBy != nil {
		f.SortBy = *filterDTO.SortBy
	}
	if filterDTO.SortOrder != nil {
		f.SortOrder = *filterDTO.SortOrder
	}

	products, err := h.productUsecase.GetAllProduct(c.Request.Context(), f)
	if err != nil {
		c.Error(customerror.NewWithLastCustomError(customerror.ERRPRODHANDLERLISTALLPROD, err))
		return
	}

	datas := []dto.GetProductResponse{}
	for _, p := range products {
		data := dto.GetProductResponse{
			ID:          p.ID,
			Name:        p.Name,
			Price:       p.Price,
			Description: p.Description,
			Quantity:    p.Quantity,
		}
		datas = append(datas, data)
	}

	utils.ResponseJSON(c, true, "successful", datas, nil, http.StatusOK)
}

func (h *productHandler) AddProduct(c *gin.Context) {
	var req dto.AddProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(customerror.New(customerror.ERRPRODHANDLERADDPRODBIND, errormessage.ErrorInvalidBody, err))
		return
	}

	product := entity.Product{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Quantity:    req.Quantity,
	}

	if err := h.productUsecase.AddProduct(c.Request.Context(), product); err != nil {
		c.Error(customerror.NewWithLastCustomError(customerror.ERRPRODHANDLERADDPROD, err))
		return
	}

	utils.ResponseJSON(c, true, "product added succesfully", nil, nil, http.StatusCreated)
}
