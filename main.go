package main

import (
	"fmt"
	"kin/kin"
	"net/http"
)

func main() {
	engine := kin.New()
	engine.Get("/hehe", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL=%q\n", r.URL.Path)
	})

	engine.Get("/haha", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	engine.Run(":9999")
}
