package router

import "net/http"

type adapter func(http.HandlerFunc) http.HandlerFunc

func adapt(handler http.HandlerFunc, adapters ...adapter) http.HandlerFunc {
	for _, adapter := range adapters {
		handler = adapter(handler)
	}
	return handler
}
