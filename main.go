package main

import (
	"kin/kin"
	"net/http"
)

func onlyForV1() kin.HandlerFunc {
	return func(ctx *kin.Context) {
		ctx.ToString(http.StatusOK, "hehe\n")
	}
}

func main() {
	engine := kin.New()
	engine.Use(kin.Logger())

	v1 := engine.Group("/v1")
	v1.Use(onlyForV1())
	v1.GET("/hello/:id", func(ctx *kin.Context) {
		ctx.ToString(http.StatusOK, "hello id = %s\n", ctx.GetParam("id"))
	})

	v2 := engine.Group("/v2")
	v2.GET("/hello/:id", func(ctx *kin.Context) {
		ctx.ToString(http.StatusOK, "hello id = %s\n", ctx.GetParam("id"))
	})

	engine.Run(":9999")
}
