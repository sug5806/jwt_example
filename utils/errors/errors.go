package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func BadRequestErrorResponse(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, ApiError{
		Code:    -1,
		Message: msg,
	})
	c.Abort()
}

func UnAuthorizeErrorResponse(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, ApiError{
		Code:    -1,
		Message: msg,
	})
	c.Abort()
}

func ServerInternalErrorResponse(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, ApiError{
		Code:    -1,
		Message: msg,
	})
	c.Abort()
}
