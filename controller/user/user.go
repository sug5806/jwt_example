package user

import (
	"github.com/gin-gonic/gin"
	"jwt-example/model"
	"jwt-example/service"
	"jwt-example/utils/errors"
	"jwt-example/utils/response"
)

func SignUp(c *gin.Context) {
	var signUp model.SignUpRequest
	err := c.BindJSON(&signUp)

	if err != nil {
		errors.BadRequestErrorResponse(c, "Form Invalid")
		return
	}

	resp, apiErr := service.UserService.SignUp(signUp)
	if apiErr != nil {
		errors.ServerInternalErrorResponse(c, apiErr.Error)
		return
	}

	response.SuccessResponse(c, resp)

}
