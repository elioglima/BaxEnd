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
	routes.HandleFunc("/api/usuarios/novo/", use(usuario.Novo, basicAuth))
	routes.HandleFunc("/api/usuario/atualizar/{id}", use(usuario.Atualizar, basicAuth))
	routes.HandleFunc("/api/usuario/hash/{email}/{documento}", use(usuario.ColherHash, basicAuth))
	routes.HandleFunc("/api/usuario/ativar/{id}", use(usuario.AtivarCadastro, basicAuth))

	routes.HandleFunc("/api/usuario/apagar/{id}", use(usuario.Apagar, basicAuth))
	routes.HandleFunc("/api/usuario/pesquisa/codigo/{id}", use(usuario.PesquisaCodigo, basicAuth))
	routes.HandleFunc("/api/usuario/pesquisa/nome/{value}", use(usuario.PesquisaNome, basicAuth))
	routes.HandleFunc("/api/usuario/pesquisa/email/{value}", use(usuario.PesquisaEmail, basicAuth))
}
