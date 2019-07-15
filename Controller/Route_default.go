package Controller

import (
	"BaxEnd/Controller/routes/api/usuario"
	"BaxEnd/Controller/routes/views"
	"net/http"

	"github.com/gorilla/mux"
)

func setRoutes() {

	routes = NewRouter()

	SetRoutesWalk(routes)
	SetRoutesUsuario(routes)
	routes.NotFoundHandler = http.HandlerFunc(views.NotFound)

}

func SetRoutesUsuario(routes *mux.Router) {
	routes.HandleFunc("/api/usuario/novo/unico/", use(usuario.NovoUnico, basicAuth))
	routes.HandleFunc("/api/usuario/novo/varios/", use(usuario.NovoVarios, basicAuth))
	routes.HandleFunc("/api/usuario/atualiza/{id}", use(usuario.AlteraUnico, basicAuth))
	routes.HandleFunc("/api/usuario/apagar/{id}", use(usuario.Apagar, basicAuth))
	routes.HandleFunc("/api/usuario/pesquisa/todos/", use(usuario.PesquisaTodos, basicAuth))
	routes.HandleFunc("/api/usuario/pesquisa/codigo/{id}", use(usuario.PesquisaCodigo, basicAuth))
	routes.HandleFunc("/api/usuario/pesquisa/nome/{value}", use(usuario.PesquisaNome, basicAuth))
	routes.HandleFunc("/api/usuario/pesquisa/email/{value}", use(usuario.PesquisaEmail, basicAuth))
}
