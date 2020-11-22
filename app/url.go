package app

import (
	"jwt-example/controller/hello"
	"jwt-example/controller/middleware"
	"jwt-example/controller/my"
	"jwt-example/controller/user"
)

func UrlMapping() {
	router.GET("/", hello.Hello)
	router.POST("/sign-up", user.SignUp)

	router.GET("/api-call", middleware.MiddleWare, my.ApiCall)
}
