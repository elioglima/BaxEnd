package usuario

import (
	"BaxEnd/Controller/database"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*
	{
		"hash":"YeqlJsoXK8jfZcXU+hysFQ+gH59tq/lUDLuKBrxW1LOBoOnEQWUXd0QmPK0HdIHXjn1k+U/KZRYIsDKwujUuoA==",
		"senha":"AB@102030",
		"senha_conf":"AB@102030"
	}
*/

func AtivarCadastro(w http.ResponseWriter, r *http.Request) {

	Retorno := sRetorno{}
	Retorno.Ini()
	params := mux.Vars(r)

	ID, err := strconv.Atoi(params["id"])
	if err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	if err := database.MySql.Usuario.PesquisaCodigo(int64(ID)); err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	if database.MySql.Usuario.RecordCount == 0 {
		Retorno.Msg = "Usuário não localizado"
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	ArrayByteIn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Retorno.Erro = true
		Retorno.Msg = errors.New("Erro ao receber body. \n " + err.Error()).Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return

	} else if len(ArrayByteIn) == 0 {
		Retorno.Erro = true
		Retorno.Msg = "Erro ao receber parametros"
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	msg, err := database.MySql.Usuario.AtivarCadastro(ArrayByteIn)
	if err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	Retorno.Msg = msg
	Retorno.Dados = database.MySql.Usuario.Response
	responseReturn(w, Retorno)
}
