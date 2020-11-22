package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApp() {
	UrlMapping()

	if err := router.Run(":9090"); err != nil {
		panic(err)
	}

}
