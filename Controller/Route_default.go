package Controller

import (
	"BaxEnd/Controller/routes/api/ChaveAcessoRoute"
	"BaxEnd/Controller/routes/api/empresa"
	"BaxEnd/Controller/routes/api/usuario"
	"BaxEnd/Controller/routes/views"
	"net/http"

	"github.com/gorilla/mux"
)

func setRoutes() {

	routes = NewRouter()
	routes.StrictSlash(true)

	SetAuthenticationMiddleware(routes)
	SetInterceptorInput(routes)
	SetRoutesWalk(routes)

	SetRoutesUsuario(routes)
	// SetRoutesViews(routes)
	routes.NotFoundHandler = http.HandlerFunc(views.NotFound)

}

func SetRoutesViews(routes *mux.Router) {
	routes.HandleFunc("/docs/", use(views.Docs, DocBasicAuth))
	routes.HandleFunc("/docs/{nivel1}", use(views.Docs, DocBasicAuth))
	routes.HandleFunc("/", views.Home)
}

func SetRoutesUsuario(routes *mux.Router) {

	routes.HandleFunc("/api/chave/acesso/pesquisa/todos", ChaveAcessoRoute.PesquisaTodos)

	routes.HandleFunc("/api/usuario/pesquisa/todos", usuario.PesquisaTodos)
	routes.HandleFunc("/api/usuario/pesquisa/nome", usuario.PesquisaNome)
	routes.HandleFunc("/api/usuario/pesquisa/codigo", usuario.PesquisaCodigo)
	routes.HandleFunc("/api/usuario/pesquisa/email", usuario.PesquisaEmail)
	routes.HandleFunc("/api/usuario/atualizar", usuario.Atualizar)
	routes.HandleFunc("/api/usuario/novo", usuario.Novo)
	routes.HandleFunc("/api/usuario/apagar", usuario.Apagar)

	routes.HandleFunc("/api/empresa/pesquisa/todos", empresa.PesquisaTodos)
	routes.HandleFunc("/api/empresa/pesquisa/nome", empresa.PesquisaNome)
	routes.HandleFunc("/api/empresa/pesquisa/codigo", empresa.PesquisaCodigo)
	routes.HandleFunc("/api/empresa/pesquisa/email", empresa.PesquisaEmail)
	routes.HandleFunc("/api/empresa/atualizar", empresa.Atualizar)
	routes.HandleFunc("/api/empresa/novo", empresa.Novo)
	routes.HandleFunc("/api/empresa/apagar", empresa.Apagar)

}
