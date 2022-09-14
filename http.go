package main

import (
	"github.com/gin-gonic/gin"
)

func Serve(host string, port string) {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		author, exist := ctx.GetQuery("author")
		if !exist {
			ctx.JSON(200, QueryRandom())
		} else {
			ctx.JSON(200, QueryByAuthorRandom(author))
		}
	})
	r.Run(host + ":" + port)
}
