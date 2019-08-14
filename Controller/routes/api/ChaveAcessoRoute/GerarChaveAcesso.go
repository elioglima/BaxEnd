package ChaveAcessoRoute

import (
	"BaxEnd/Controller/database"
	"GoLibs/logs"
	"errors"
	"io/ioutil"
	"net/http"
)

func GerarChaveAcesso(w http.ResponseWriter, r *http.Request) {

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

	if err := database.MySql.ChaveAcessoHttp.Gerar(ArrayByteIn); err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	if database.MySql.ChaveAcessoHttp.RecordCount == 0 {
		Retorno.Msg = "Nenhum registro localizado"
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	Retorno.Dados = database.MySql.ChaveAcessoHttp.Fields
	Retorno.Msg = "Registro localizado com sucesso."
	responseReturn(w, Retorno)
}
