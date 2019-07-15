package views

import (
	logger "GoLibs/logs"
	"BaxEnd/Controller/global"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	logger.Erro("NotFound", r.Method, r.URL)
	path := global.DirPrivate()

	t := time.Now()
	Estrutura := stEstrutura{}
	Estrutura.Data = fmt.Sprintf("%02d/%02d/%d %02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
	Estrutura.Versao = global.Versao

	Templ, _ := template.ParseFiles(path + "NotFound.html")
	Templ.Execute(w, Estrutura)
}
