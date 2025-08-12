package kin

import (
	"html/template"
	"net/http"
	"strings"
)

// 处理函数
type HandlerFunc func(ctx *Context)

// ServeHTTP实例
type Engine struct {
	router *router
	*RouterGroup
	groups        []*RouterGroup
	htmlTemplates *template.Template // for html render
	funcMap       template.FuncMap   // for html render
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// 默认使用log和recovery
func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

// SetFuncMap sets the function map for the engine's templates.
// It allows registering custom functions that can be used within templates.
//
// Parameters:
//   - funcMap: A template.FuncMap containing the functions to be registered
func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

// LoadHTMLGlob loads HTML templates from files matching the given pattern.
// It parses all files that match the pattern and registers them with the engine's function map.
//
// Parameters:
//
//	pattern: A glob pattern string that specifies which template files to load (e.g., "templates/*.html")
//
// The function does not return any value but stores the parsed templates in engine.htmlTemplates.
// If parsing fails, the program will panic due to template.Must().
func (engine *Engine) LoadHTMLGlob(pattern string) {
	// Parse all template files matching the pattern and register them with the engine's function map
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}

// 设置路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// 增加GET方法
func (engine *Engine) Get(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// 增加POST方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// interface实现
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
	ctx.engine = engine
	engine.router.handle(ctx)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
