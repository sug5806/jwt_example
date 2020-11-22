package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"jwt-example/model"
	"jwt-example/service"
	"jwt-example/utils/errors"
)

func MiddleWare(c *gin.Context) {
	fmt.Println("I'm middleWare")

	jwtToken := c.GetHeader("Authorization")
	if len(jwtToken) < 0 {
		errors.UnAuthorizeErrorResponse(c, "UnAuthorization")
		c.Abort()
		return
	}

	claim := &model.Claims{}

	_, err := jwt.ParseWithClaims(jwtToken, claim, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpeced signing method : %v", token.Header["alg"])
		}
		return []byte("jwt secret"), nil
	})

	if err != nil {
		errors.UnAuthorizeErrorResponse(c, err.Error())
		c.Abort()
		return
	}

	apiErr := service.UserService.UserCheck(claim.Id)

	if apiErr != nil {
		errors.UnAuthorizeErrorResponse(c, apiErr.Error)
	}

	c.Next()
}
