package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"majooTest/helpers"
	"majooTest/util"
	"net/http"
	"strings"
)

func AuthMiddleware(c *gin.Context) {
	// get header
	var response helpers.Response
	accessToken := c.Request.Header.Get(util.Authorization)
	arr := strings.Split(accessToken, " ")
	if len(arr) != 2 {
		response = helpers.ResponseError("Authorization not found", http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	fmt.Println("Authorization : ", arr[1])

	if arr[1] == "" {
		response = helpers.ResponseError("Authorization not found", http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	c.Next()
}
