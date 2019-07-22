package views

import (
	"BaxEnd/Controller/global"
	logger "GoLibs/logs"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"
)

func NotFound(w http.ResponseWriter, r *http.Request) {

	logger.Erro("NotFound", r.Method, r.URL)
	URLSplit := strings.Split(r.URL.String(), "/")

	if strings.ToLower(URLSplit[1]) == "api" {
		w.Header().Set("Content-Type", "application/json")

		type RetornoST struct {
			Erro     bool
			Response string
		}

		Retorno := RetornoST{}
		Retorno.Erro = true
		Retorno.Response = "Recurso não alocado ou não implantado."

		jsonstr, err := json.Marshal(Retorno)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(jsonstr)
		return
	}

	path := global.DirPrivate()

	t := time.Now()
	Estrutura := stEstrutura{}
	Estrutura.Data = fmt.Sprintf("%02d/%02d/%d %02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
	Estrutura.Versao = global.Versao

	Templ, _ := template.ParseFiles(path + "NotFound.html")
	Templ.Execute(w, Estrutura)
}
