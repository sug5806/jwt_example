package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiResponse struct {
	Result uint64      `json:"result"`
	Data   interface{} `json:"data"`
}

func (a ApiResponse) ApiResponse(c *gin.Context, data interface{}) {
	c.Abort()
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ApiResponse{
		Result: 0,
		Data:   data,
	})
}
