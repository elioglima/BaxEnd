package usuario

import (
	"BaxEnd/Controller/database"
	"GoLibs/logs"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func Apagar(w http.ResponseWriter, r *http.Request) {

	logs.Branco("usuario/Apagar/")
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

	type CDados struct {
		EmpresaID *int64
		Id        *int64
	}

	dados := CDados{}
	if err := json.Unmarshal(ArrayByteIn, &dados); err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
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

	if err := database.MySql.Usuario.LoadEmpresa(*dados.EmpresaID); err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	if dados.Id == nil {
		Retorno.Erro = true
		Retorno.Msg = errors.New("Id do usuário não informado.").Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	msg, err := database.MySql.Usuario.Apagar(*dados.Id)
	if err != nil {
		Retorno.Erro = false
		Retorno.Msg = msg
		responseReturn(w, Retorno)
		return
	}
	Retorno.Msg = "Usuário deletado com sucesso."
	responseReturn(w, Retorno)
}
