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
	SetRoutesViews(routes)
	routes.NotFoundHandler = http.HandlerFunc(views.NotFound)

}

func SetRoutesViews(routes *mux.Router) {
	routes.HandleFunc("/docs/", use(views.Docs, DocBasicAuth))
	routes.HandleFunc("/docs/{nivel1}", use(views.Docs, DocBasicAuth))
	routes.HandleFunc("/", views.Home)

}

func SetRoutesUsuario(routes *mux.Router) {
	routes.HandleFunc("/api/usuarios/pesquisar/todos/", use(usuario.PesquisaTodos, basicAuth))
	routes.HandleFunc("/api/usuarios/novo/unico/", use(usuario.NovoUnico, basicAuth))
	routes.HandleFunc("/api/usuario/novo/varios/", use(usuario.NovoVarios, basicAuth))
	routes.HandleFunc("/api/usuario/atualiza/{id}", use(usuario.AlteraUnico, basicAuth))
	routes.HandleFunc("/api/usuario/apagar/{id}", use(usuario.Apagar, basicAuth))
	routes.HandleFunc("/api/usuario/pesquisa/codigo/{id}", use(usuario.PesquisaCodigo, basicAuth))
	routes.HandleFunc("/api/usuario/pesquisa/nome/{value}", use(usuario.PesquisaNome, basicAuth))
	routes.HandleFunc("/api/usuario/pesquisa/email/{value}", use(usuario.PesquisaEmail, basicAuth))
}
