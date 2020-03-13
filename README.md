# kin
quick start:

```go
func main(){
	engine := kin.New()

	//direct route
	engine.Get("/hello", func(ctx *kin.Context) {
		ctx.ToString(http.StatusOK,"hello\n")
	})

	//use group route
	v1 := engine.Group("/v1")
	v1.POST("/hello", func(ctx *kin.Context) {
		ctx.ToJson(http.StatusOK,kin.M{"err_msg":"success"})
	})

	//dynamic route
	v1.GET("/hello/:id", func(ctx *kin.Context) {
		ctx.ToString(http.StatusOK, "hello id = %s\n", ctx.GetParam("id"))
	})
	
	//middleware
	v2 := engine.Group("/v2")
	v2.Use(onlyForV2())
	v2.GET("/lala", func(ctx *kin.Context) {
		ctx.ToString(http.StatusOK,"111\n")
	})
	
	//recovery
	engine.Get("/panic", func(ctx *kin.Context) {
		str := []string{"abc"}
		ctx.ToString(http.StatusOK, str[10])
	})
}
```
