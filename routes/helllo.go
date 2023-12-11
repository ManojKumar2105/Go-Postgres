package routes

import "github.com/gin-gonic/gin"

func Helloroute(router *gin.Engine) {
	router.GET("/api", func(ctx *gin.Context) {
		ctx.String(200, "Hello World!")
	})
}
