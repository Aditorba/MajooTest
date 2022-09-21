package controller

import (
	"github.com/gin-gonic/gin"
	"majooTest/dto"
	"majooTest/helpers"
	"majooTest/service"
	"majooTest/util"
	"net/http"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
}

func NewAuthController(authServiceCheck service.AuthService) AuthController {
	return &authController{
		authService: authServiceCheck,
	}
}

func (controller *authController) Login(context *gin.Context) {
	var loginDTO dto.LoginDTO
	var response helpers.Response
	if err := context.ShouldBindJSON(&loginDTO); err != nil {
		var data = util.ValidationBindJsonField(err)
		context.JSON(http.StatusBadRequest, data)
	} else {
		isValidLogin, _ := controller.authService.Login(loginDTO)
		if isValidLogin {
			result, status := controller.authService.GenerateToken(loginDTO.Username)
			if status != nil {
				response = helpers.ResponseError(status.Error(), http.StatusBadRequest)
				context.JSON(http.StatusBadRequest, response)
			} else {
				response = helpers.ResponseSuccess(result)
				context.JSON(http.StatusOK, response)
			}
		} else {
			response = helpers.ResponseError("Wrong combination Username and Password", http.StatusUnauthorized)
			context.JSON(http.StatusUnauthorized, response)
		}
	}
}
