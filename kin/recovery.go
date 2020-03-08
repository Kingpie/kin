package kin

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

//打印堆栈
func trace(message string) string {
	var callers [32]uintptr
	n := runtime.Callers(3, callers[:])

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range callers[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Error")
			}
		}()

		c.Next()
	}
}
