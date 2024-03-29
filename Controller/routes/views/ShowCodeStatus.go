package views

import (
	logger "GoLibs/logs"
	"BaxEnd/Controller/global"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func ShowCodeStatus(w http.ResponseWriter, r *http.Request, EstruturaIn *stEstrutura) {

	logger.Rosa("ShowCodeStatus")
	path := global.DirPrivate()

	t := time.Now()

	Estrutura := &stEstrutura{}

	if EstruturaIn == nil {
		Estrutura.Data = fmt.Sprintf("%02d/%02d/%d %02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
		Estrutura.Versao = global.Versao
	} else {
		Estrutura = EstruturaIn
	}

	Templ, _ := template.ParseFiles(path + "Docs/ShowCodeStatus.html") //setp 1
	Templ.Execute(w, Estrutura)                                        //step 2
}
