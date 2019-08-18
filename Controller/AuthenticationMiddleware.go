package Controller

import (
	"BaxEnd/Controller/Token"
	"BaxEnd/Controller/database"
	"GoLibs/logs"
	"errors"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

)

func SetAuthenticationMiddleware(routes *mux.Router) {
	amw := authenticationMiddleware{make(map[string]string)}
	routes.Use(amw.Middleware)
}

// Define our struct
type authenticationMiddleware struct {
	tokenUsers map[string]string
}

// Middleware function, which will be called for each request
func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("%v", r.Header)

		KeyApp := r.Header.Get("BaxEnd-Token")
		// logs.Cyan(token)

		if len(strings.TrimSpace(KeyApp)) == 0 {
			err := errors.New("401 - Ops, token não foi informado.")
			logs.Erro(err.Error())
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if err := database.MySql.CheckConnect(); err != nil {
			logs.Erro(err.Error())
			http.Error(w, "401 - Ops, banco de dados está conectado.", http.StatusUnauthorized)
			return
		}

		Tk := Token.NewToken()
		if err := Tk.Decode(KeyApp); err != nil {
			logs.Erro(err.Error())
			http.Error(w, "401 - Ops, "+err.Error(), http.StatusUnauthorized)
			return
		}

		// autenticacao no banco
		if err := database.MySql.ChaveAcessoHttp.Auth(Tk.KeyAPI); err != nil {
			logs.Erro(err.Error())
			logs.Erro("Tk.KeyAPI", Tk.KeyAPI)
			http.Error(w, "401 - Ops, "+err.Error(), http.StatusUnauthorized)
			return
		}

		if database.MySql.ChaveAcessoHttp.RecordCount == 0 {
			err := errors.New("401 - Ops, token não localizado no banco de dados.")
			logs.Erro(err.Error())
			logs.Erro("Tk.KeyAPI", Tk.KeyAPI)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
