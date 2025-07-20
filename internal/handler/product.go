package handler

import (
	"backend_golang/internal/dto"
	"backend_golang/internal/entity"
	"backend_golang/internal/usecase"
	"backend_golang/utils"
	"backend_golang/utils/customerror"
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
	var f entity.ProductFilter
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
