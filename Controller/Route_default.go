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

	sRotaUsuario := "/api/{EmpresaID:[0-9]+}/usuario"

	routes.HandleFunc(sRotaUsuario+"/pesquisar/todos/", use(usuario.PesquisaTodos, basicAuth))
	routes.HandleFunc(sRotaUsuario+"/novo/", use(usuario.Novo, basicAuth))
	routes.HandleFunc(sRotaUsuario+"/atualizar/{id:[0-9]+}", use(usuario.Atualizar, basicAuth))
	routes.HandleFunc(sRotaUsuario+"/hash/{email}/{documento}", use(usuario.ColherHash, basicAuth))
	routes.HandleFunc(sRotaUsuario+"/ativar/{id:[0-9]+}", use(usuario.AtivarCadastro, basicAuth))

	routes.HandleFunc("/api/usuario/apagar/{id}", use(usuario.Apagar, basicAuth))
	routes.HandleFunc("/api/usuario/pesquisa/codigo/{id}", use(usuario.PesquisaCodigo, basicAuth))
	routes.HandleFunc("/api/usuario/pesquisa/nome/{value}", use(usuario.PesquisaNome, basicAuth))
	routes.HandleFunc("/api/usuario/pesquisa/email/{value}", use(usuario.PesquisaEmail, basicAuth))
}
