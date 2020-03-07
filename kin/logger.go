package kin

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(ctx *Context) {
		t := time.Now()

		ctx.Next()

		log.Printf("uri:%s|code:%d|time:%v", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}
