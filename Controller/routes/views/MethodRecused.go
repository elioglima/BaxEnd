package views

import (
	logger "GoLibs/logs"
	"BaxEnd/Controller/global"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func MethodRecused(w http.ResponseWriter, r *http.Request) {

	logger.Rosa("MethodRecused")
	path := global.DirPrivate()

	t := time.Now()
	Estrutura := stEstrutura{}
	Estrutura.Data = fmt.Sprintf("%02d/%02d/%d %02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
	Estrutura.Versao = global.Versao

	Templ, _ := template.ParseFiles(path + "MethodRecused.html") //setp 1
	Templ.Execute(w, Estrutura)                                  //step 2
}
