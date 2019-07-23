package usuario

import (
	"BaxEnd/Controller/database"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Atualizar(w http.ResponseWriter, r *http.Request) {

	Retorno := sRetorno{}
	Retorno.Ini()
	params := mux.Vars(r)

	EmpresaID, err := strconv.Atoi(params["EmpresaID"])
	if err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	ID, err := strconv.Atoi(params["id"])
	if err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
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

	if err := database.MySql.Conectar(); err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	if err := database.MySql.Usuario.LoadEmpresa(int64(EmpresaID)); err != nil {
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

	msg, err := database.MySql.Usuario.Atualizar(ArrayByteIn)
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
