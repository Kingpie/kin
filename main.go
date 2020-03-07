package main

import (
	"kin/kin"
	"net/http"
)

func main() {
	engine := kin.New()
	engine.Get("/html", func(ctx *kin.Context) {
		ctx.ToHTML(http.StatusOK, "<h1>Hello Kin</h1>")
	})

	engine.POST("/json", func(ctx *kin.Context) {
		ctx.ToJson(http.StatusOK, kin.M{
			"user": ctx.FormValue("user"),
			"pass": ctx.FormValue("pass"),
		})
	})

	engine.Run(":9999")
}
