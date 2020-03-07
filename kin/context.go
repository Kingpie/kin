package kin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type M map[string]interface{}

type Context struct {
	//基础数据
	Writer http.ResponseWriter
	Req    *http.Request

	//请求信息
	Path   string
	Method string

	//响应信息
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (ctx *Context) FormValue(key string) string {
	return ctx.Req.FormValue(key)
}

func (ctx *Context) Query(key string) string {
	return ctx.Req.URL.Query().Get(key)
}

func (ctx *Context) SetStatus(code int) {
	ctx.StatusCode = code
	ctx.Writer.WriteHeader(code)
}

func (ctx *Context) SetHeader(key string, value string) {
	ctx.Writer.Header().Set(key, value)
}

func (ctx *Context) ToString(code int, format string, values ...interface{}) {
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.SetStatus(code)
	ctx.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (ctx *Context) ToJson(code int, obj interface{}) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.SetStatus(code)
	encoder := json.NewEncoder(ctx.Writer)
	err := encoder.Encode(obj)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), 500)
	}
}

func (ctx *Context) ToBinary(code int, data []byte) {
	ctx.SetStatus(code)
	ctx.Writer.Write(data)
}

func (ctx *Context) ToHTML(code int, html string) {
	ctx.SetHeader("Content-Type", "text/html")
	ctx.SetStatus(code)
	ctx.Writer.Write([]byte(html))
}
