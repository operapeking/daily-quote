package main

import (
	"github.com/gin-gonic/gin"
)

func Serve(host string, port string) {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, QueryByRandom())
	})
	r.Run(host + ":" + port)
}
