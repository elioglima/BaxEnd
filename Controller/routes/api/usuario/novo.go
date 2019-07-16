package usuario

import (
	"BaxEnd/Controller/database"
	"errors"
	"io/ioutil"
	"net/http"
)

func NovoUnico(w http.ResponseWriter, r *http.Request) {
	Retorno := sRetorno{}
	Retorno.Ini()

	if err := database.MySql.Conectar(); err != nil {
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

	msg, err := database.MySql.Usuario.NovoUnico(ArrayByteIn)
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

func NovoVarios(w http.ResponseWriter, r *http.Request) {
	Retorno := sRetorno{}
	Retorno.Msg = "Função não implantada."
	responseReturn(w, Retorno)
}
