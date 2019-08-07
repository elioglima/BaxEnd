package Controller

import (
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

	// sRotaUsuario := "/api/{EmpresaID:[0-9]+}/usuario"
	// sRotaUsuario := "/api/"

	routes.HandleFunc("/api/usuario/novo/", usuario.Test)
	// routes.HandleFunc("/api/usuario/pesquisar/todos/", use(usuario.PesquisaTodos, basicAuth))
	// routes.HandleFunc("/api/usuario/atualizar/{id:[0-9]+}", use(usuario.Atualizar, basicAuth))
	// routes.HandleFunc("/api/usuario/hash/{email}/{documento}", use(usuario.ColherHash, basicAuth))
	// routes.HandleFunc("/api/usuario/ativar/{id:[0-9]+}", use(usuario.AtivarCadastro, basicAuth))

	// routes.HandleFunc("/api/usuario/apagar/{id}", use(usuario.Apagar, basicAuth))
	// routes.HandleFunc("/api/usuario/pesquisa/codigo/{id}", use(usuario.PesquisaCodigo, basicAuth))
	// routes.HandleFunc("/api/usuario/pesquisa/nome/{value}", use(usuario.PesquisaNome, basicAuth))
	// routes.HandleFunc("/api/usuario/pesquisa/email/{value}", use(usuario.PesquisaEmail, basicAuth))
}
