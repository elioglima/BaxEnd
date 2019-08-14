package Controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetAuthenticationMiddleware(routes *mux.Router) {
	amw := authenticationMiddleware{make(map[string]string)}
	amw.Populate()
	routes.Use(amw.Middleware)
}

// Define our struct
type authenticationMiddleware struct {
	tokenUsers map[string]string
}

// Initialize it somewhere
func (amw *authenticationMiddleware) Populate() {
	amw.tokenUsers["TK1"] = "user0"
	amw.tokenUsers["TK2"] = "userA"
	amw.tokenUsers["TK3"] = "randomUser"
	amw.tokenUsers["TK4"] = "user0"
}

// Middleware function, which will be called for each request
func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("%v", r.Header)

		// token := r.Header.Get("X-Session-Token")

		// if user, found := amw.tokenUsers[token]; found {
		// 	// We found the token in our map
		// 	log.Printf("Authenticated user %s\n", user)
		// 	// Pass down the request to the next middleware (or final handler)
		// 	next.ServeHTTP(w, r)
		// } else {
		// 	// Write an error and stop the handler chain
		// 	log.Printf("Authenticated user %s\n")
		// 	http.Error(w, "Forbidden", http.StatusForbidden)
		// }
		next.ServeHTTP(w, r)
	})
}
