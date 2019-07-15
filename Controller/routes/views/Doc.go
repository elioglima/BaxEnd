package views

import (
	logger "GoLibs/logs"
	"BaxEnd/Controller/global"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Docs(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	nivel1 := params["nivel1"]
	path := global.DirPrivate()

	Estrutura, err := AuthUser(w, r)
	if err != nil {
		return
	}

	if len(strings.TrimSpace(nivel1)) == 0 {
		logger.Branco("Metodo", r.Method, r.URL, "Docs/index")
		t, _ := template.ParseFiles(path + "Docs/index.html")
		t.Execute(w, Estrutura)
		return
	}

	logger.Branco("Metodo", r.Method, r.URL, "Docs/"+nivel1)
	Templ, err := template.ParseFiles(path + "Docs/" + nivel1 + ".html") //setp 1
	if err != nil {
		NotFound(w, r)
		return
	}

	Templ.Execute(w, Estrutura)
}
