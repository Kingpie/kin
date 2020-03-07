package kin

import (
	"net/http"
	"strings"
)

//处理函数
type HandlerFunc func(ctx *Context)

//ServeHTTP实例
type Engine struct {
	router *router
	*RouterGroup
	groups []*RouterGroup
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

//设置路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

//增加GET方法
func (engine *Engine) Get(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

//增加POST方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

//interface实现
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//中间件
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	ctx := newContext(w, req)
	ctx.handlers = middlewares
	engine.router.handle(ctx)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
