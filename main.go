package main

import (
	"kin/kin"
	"net/http"
)

func main() {
	engine := kin.New()

	v1 := engine.Group("/v1")
	v1.GET("/hello/:id", func(ctx *kin.Context) {
		ctx.ToString(http.StatusOK, "hello id = %s\n", ctx.GetParam("id"))
	})

	engine.Run(":9999")
}
