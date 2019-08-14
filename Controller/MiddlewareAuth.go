package Controller

import (
	"GoLibs/logs"
	"net/http"
)

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logs.Atencao(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
