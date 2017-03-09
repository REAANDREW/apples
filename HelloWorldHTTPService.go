package main

import "github.com/gin-gonic/gin"

func CreateHelloWorldHTTPService(endpoint string) *HTTPService {
	return CreateHTTPService(endpoint, func(router *gin.Engine) {
		router.GET("/", func(ctx *gin.Context) {
			ctx.String(200, "Hello, World!")
		})
	})
}
