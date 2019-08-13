package empresa

import (
	"BaxEnd/Controller/database"
	"GoLibs/logs"
	"errors"
	"io/ioutil"
	"net/http"
)

func Apagar(w http.ResponseWriter, r *http.Request) {

	logs.Branco("empresa/Apagar/")
	Retorno := sRetorno{}
	Retorno.Ini()

	ArrayByteIn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Retorno.Erro = true
		Retorno.Msg = errors.New("Erro ao receber body. \n " + err.Error()).Error()
		Retorno.Dados = nil
		logs.Erro(Retorno.Msg)
		responseReturn(w, Retorno)
		return
	}

	msg, err := database.MySql.Empresa.Apagar(ArrayByteIn)
	if err != nil {
		Retorno.Erro = false
		Retorno.Msg = msg
		responseReturn(w, Retorno)
		return
	}

	Retorno.Msg = "Registro deletado com sucesso."
	responseReturn(w, Retorno)
}
