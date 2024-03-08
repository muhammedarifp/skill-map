package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Message": "welcome to skill map",
		})
	})

	app.GET("/login", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Message": "Login is required",
		})
	})
	app.Run()
}
