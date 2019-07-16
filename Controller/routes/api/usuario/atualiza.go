package usuario

import (
	"BaxEnd/Controller/database"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AlteraUnico(w http.ResponseWriter, r *http.Request) {

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

	msg, err := database.MySql.Usuario.AlteraUnico(ArrayByteIn)
	if err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	Retorno.Msg = msg
	Retorno.Dados = database.MySql.Usuario.Field
	responseReturn(w, Retorno)
}
