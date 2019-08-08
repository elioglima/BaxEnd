package usuario

import (
	"BaxEnd/Controller/database"
	"GoLibs/logs"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

)

func PesquisaTodos(w http.ResponseWriter, r *http.Request) {

	logs.Branco("usuario/pesquisa/todos/")
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

	if err := database.MySql.Conectar(); err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	if err := database.MySql.Usuario.PesquisaTodos(ArrayByteIn); err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	if database.MySql.Usuario.RecordCount == 0 {
		Retorno.Msg = "Nenhum usuário não localizado"
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	Retorno.Dados = database.MySql.Usuario.Fields
	Retorno.Msg = "Usuários localizado com sucesso."
	responseReturn(w, Retorno)
}

func PesquisaCodigo(w http.ResponseWriter, r *http.Request) {

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
		Retorno.Msg = "Nenhum usuário não localizado"
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	Retorno.Dados = database.MySql.Usuario.Field
	Retorno.Msg = "Usuários localizado com sucesso."
	responseReturn(w, Retorno)
}

func PesquisaEmail(w http.ResponseWriter, r *http.Request) {

	Retorno := sRetorno{}
	params := mux.Vars(r)
	value := params["value"]

	if err := database.MySql.Usuario.PesquisaEmail(value); err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	if database.MySql.Usuario.RecordCount == 0 {
		Retorno.Msg = "Nenhum usuário não localizado"
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	Retorno.Dados = database.MySql.Usuario.Field
	Retorno.Msg = "Usuário localizado com sucesso."
	responseReturn(w, Retorno)
}

func PesquisaNome(w http.ResponseWriter, r *http.Request) {

	Retorno := sRetorno{}
	params := mux.Vars(r)
	value := params["value"]

	if len(strings.TrimSpace(value)) == 0 {
		Retorno.Erro = true
		Retorno.Msg = errors.New("Paramêtro não informado").Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	if err := database.MySql.Usuario.PesquisaNome(value); err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	if database.MySql.Usuario.RecordCount == 0 {
		Retorno.Msg = "Nenhum usuário não localizado"
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	Retorno.Dados = database.MySql.Usuario.Fields
	Retorno.Msg = "Usuário localizado com sucesso."
	responseReturn(w, Retorno)
}
