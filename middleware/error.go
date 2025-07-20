package middleware

import (
	"backend_golang/utils"
	"backend_golang/utils/customerror"
	"backend_golang/utils/errormessage"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func ErrorMiddleware(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors[0].Err

		log.WithFields(logrus.Fields{
			"error":  err.Error(),
			"path":   c.Request.URL.Path,
			"method": c.Request.Method,
			"ip":     c.ClientIP(),
		}).Error("Request error")

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			utils.ResponseJSON(c, false, errormessage.ErrorInvalidBody, nil, nil, http.StatusBadRequest)
			return
		}

		var customErr *customerror.CustomError
		if errors.As(err, &customErr) {
			httpStatus := customerror.ToHttpStatus(customErr.Codes())
			utils.ResponseJSON(c, false, customErr.Message, nil, nil, httpStatus)
			return
		}

		utils.ResponseJSON(c, false, errormessage.ErrorInternalError, nil, nil, http.StatusInternalServerError)
	}
}
