package Controller

import (
	"GoLibs/logs"
	"net/http"

	"github.com/gorilla/mux"
)

func SetInterceptorInput(routes *mux.Router) {
	routes.Use(InterceptorInput)
}
func InterceptorInput(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logs.Atencao(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
