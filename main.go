package lemonlog

import (
	"net/http"
	"github.com/zenazn/goji/web"
	"time"
	"fmt"
)

func Logger(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, req)
		finish := time.Now()
		path:=req.URL.String()
		if (len(path) >3) {
			if (path[:4] == "/api") {
				fmt.Printf("[%-3s]%-10s|%-6s\n", req.Method, finish.Sub(start), req.URL.String())
			}
		}
	}
	return http.HandlerFunc(fn)
}
