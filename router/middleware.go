package router

import (
	"log"
	"net/http"
	"strings"
)

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func validateContentType() adapter {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if !strings.HasPrefix(r.Header.Get("content-type"), "application/json") {
				w.Write([]byte(`{"error": "unaccepted content type", "code":400}`))
				return
			}
			h.ServeHTTP(w, r)
		}
	}
}
