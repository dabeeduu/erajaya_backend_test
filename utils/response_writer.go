package utils

import (
	"backend_golang/internal/dto"

	"github.com/gin-gonic/gin"
)

func ResponseJSON(c *gin.Context, success bool, msg string, data any, pagination *dto.Pagination, statusCode int) {
	c.JSON(statusCode, dto.Response{
		Success:    success,
		Message:    msg,
		Pagination: nil,
		Data:       data,
	})
}
