package Controller

import (
	"BaxEnd/Controller/routes/api/empresa"
	"BaxEnd/Controller/routes/api/usuario"
	"BaxEnd/Controller/routes/views"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func setRoutes() {

	routes = NewRouter()
	routes.StrictSlash(true)

	routes.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
		var body map[string]interface{}
		type Cdata struct {
			Nome string
		}

		DataTipo := Cdata{}

		data, err := ioutil.ReadAll(r.Body)
		if err == nil && data != nil {
			err = json.Unmarshal(data, &DataTipo)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
		}

		fmt.Printf("%s", DataTipo.Nome)
		return

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		fmt.Fprintf(w, "body: %s\n", body)
	})

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

	// routes.HandleFunc("/api/usuario/pesquisar/todos/", use(usuario.PesquisaTodos, basicAuth))

}
