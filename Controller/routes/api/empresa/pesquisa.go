package empresa

import (
	"BaxEnd/Controller/database"
	"GoLibs/logs"
	"errors"
	"io/ioutil"
	"net/http"
)

func PesquisaTodos(w http.ResponseWriter, r *http.Request) {

	logs.Branco("empresa/pesquisa/todos/")
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

	if err := database.MySql.Empresa.PesquisaTodos(ArrayByteIn); err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	if database.MySql.Empresa.RecordCount == 0 {
		Retorno.Msg = "Nenhuma registro localizado"
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	Retorno.Dados = database.MySql.Empresa.Fields
	Retorno.Msg = "Empresa localizada com sucesso."
	responseReturn(w, Retorno)
}

func PesquisaNome(w http.ResponseWriter, r *http.Request) {

	logs.Branco("empresa/pesquisa/nome/")
	Retorno := sRetorno{}

	ArrayByteIn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Retorno.Erro = true
		Retorno.Msg = errors.New("Erro ao receber body. \n " + err.Error()).Error()
		Retorno.Dados = nil
		logs.Erro(Retorno.Msg)
		responseReturn(w, Retorno)
		return
	}

	if err := database.MySql.Empresa.PesquisaNome(ArrayByteIn); err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	if database.MySql.Empresa.RecordCount == 0 {
		Retorno.Msg = "Nenhum registro localizado"
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	Retorno.Dados = database.MySql.Empresa.Fields
	Retorno.Msg = "Registro localizado com sucesso."
	responseReturn(w, Retorno)
}
