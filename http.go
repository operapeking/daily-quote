package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Serve(host string, port string) {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		from, existFrom := ctx.GetQuery("from")
		if !existFrom {
			ctx.JSON(200, QueryRandom())
			log.Println("ok")
		} else {
			ctx.JSON(200, QueryByFromRandom(from))
		}
	})
	r.Run(host + ":" + port)
}
