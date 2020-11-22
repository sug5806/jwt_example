package my

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiCall(c *gin.Context) {
	fmt.Println("HiHI")

	c.JSON(http.StatusOK, &gin.H{
		"result":  0,
		"message": "success",
	})
}
