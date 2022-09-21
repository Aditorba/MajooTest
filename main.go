package main

import (
	"github.com/gin-gonic/gin"
	"majooTest/app"
)

func main() {

	// Uncomment if up to production
	gin.SetMode(gin.DebugMode)
	app.Route()

}
