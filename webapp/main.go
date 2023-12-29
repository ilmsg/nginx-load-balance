package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	key := os.Getenv("KEY")

	app := gin.Default()
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
			"key":     key,
			"host":    ctx.Request.Host + ctx.Request.URL.String(),
		})
	})
	app.Run(":3000")
}
