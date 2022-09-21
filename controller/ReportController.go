package controller

import (
	"github.com/gin-gonic/gin"
	"majooTest/dto"
	"majooTest/helpers"
	"majooTest/log"
	"majooTest/service"
	"majooTest/util"
	"net/http"
)

type ReportController interface {
	GetMerchantReport(ctx *gin.Context)
	GetOutletReport(ctx *gin.Context)
}

type reportController struct {
	authService   service.AuthService
	reportService service.ReportService
}

func NewReportController(
	authServiceCheck service.AuthService,
	reportServiceCheck service.ReportService) ReportController {
	return &reportController{
		authService:   authServiceCheck,
		reportService: reportServiceCheck,
	}
}

func (controller *reportController) GetMerchantReport(context *gin.Context) {
	acessKey := context.Request.Header.Get(util.Authorization)
	isValid, userDTO, message := controller.authService.ValidateToken(acessKey)
	if isValid {
		var pageDTO dto.PageDTO

		if err := context.ShouldBindJSON(&pageDTO); err != nil {
			var data = util.ValidationBindJsonField(err)
			context.JSON(http.StatusBadRequest, data)
		} else {
			result, err := controller.reportService.GetMerchantReport(pageDTO, int(userDTO.Id))
			var response helpers.Response

			if err != nil {
				response = helpers.ResponseError(err.Error(), http.StatusNotFound)
			} else {
				response = helpers.ResponseSuccess(result)
			}

			log.Info("Get List Merchant Report Data ", response)
			context.JSON(http.StatusOK, response)
		}
	} else {
		responseMessage := "Token not authorized, " + message
		response := helpers.ResponseError(responseMessage, http.StatusUnauthorized)
		context.JSON(http.StatusUnauthorized, response)
	}
}

func (controller *reportController) GetOutletReport(context *gin.Context) {
	acessKey := context.Request.Header.Get(util.Authorization)
	isValid, userDTO, message := controller.authService.ValidateToken(acessKey)
	if isValid {
		var pageDTO dto.PageDTO

		if err := context.ShouldBindJSON(&pageDTO); err != nil {
			var data = util.ValidationBindJsonField(err)
			context.JSON(http.StatusBadRequest, data)
		} else {
			result, err := controller.reportService.GetOutletReport(pageDTO, int(userDTO.Id))
			var response helpers.Response

			if err != nil {
				response = helpers.ResponseError(err.Error(), http.StatusNotFound)
			} else {
				response = helpers.ResponseSuccess(result)
			}

			log.Info("Get List Merchant Report Data ", response)
			context.JSON(http.StatusOK, response)
		}
	} else {
		responseMessage := "Token not authorized, " + message
		response := helpers.ResponseError(responseMessage, http.StatusUnauthorized)
		context.JSON(http.StatusUnauthorized, response)
	}
}
