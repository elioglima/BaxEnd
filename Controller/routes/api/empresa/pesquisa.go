package empresa

import (
	"BaxEnd/Controller/database"
	"GoLibs/logs"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func PesquisaTodos(w http.ResponseWriter, r *http.Request) {

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

	type CDados struct {
		Nome *string
	}

	dados := CDados{}
	if err := json.Unmarshal(ArrayByteIn, &dados); err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	if dados.Nome == nil {
		Retorno.Erro = true
		Retorno.Msg = errors.New("Nenhum paramêtro infomado.").Error()
		Retorno.Dados = nil
		logs.Erro(Retorno.Msg)
		responseReturn(w, Retorno)
		return
	} else if len(strings.TrimSpace(*dados.Nome)) <= 0 {
		Retorno.Erro = true
		Retorno.Msg = errors.New("Paramêtro inválido.").Error()
		Retorno.Dados = nil
		logs.Erro(Retorno.Msg)
		responseReturn(w, Retorno)
		return
	}

	if err := database.MySql.Empresa.PesquisaArrayByteIn(ArrayByteIn); err != nil {
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

func PesquisaCodigo(w http.ResponseWriter, r *http.Request) {
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
		Id *int64
	}

	dados := CDados{}
	if err := json.Unmarshal(ArrayByteIn, &dados); err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	if dados.Id == nil {
		Retorno.Erro = true
		Retorno.Msg = errors.New("Nenhum paramêtro infomado.").Error()
		Retorno.Dados = nil
		logs.Erro(Retorno.Msg)
		responseReturn(w, Retorno)
		return
	} else if *dados.Id <= 0 {
		Retorno.Erro = true
		Retorno.Msg = errors.New("Paramêtro inválido.").Error()
		Retorno.Dados = nil
		logs.Erro(Retorno.Msg)
		responseReturn(w, Retorno)
		return
	}

	if err := database.MySql.Empresa.PesquisaCodigo(*dados.Id); err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	if database.MySql.Empresa.RecordCount == 0 {
		Retorno.Msg = "Nenhum registro localizado."
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	Retorno.Dados = database.MySql.Empresa.Field
	Retorno.Msg = "Registro localizado com sucesso."
	responseReturn(w, Retorno)
}

func PesquisaEmail(w http.ResponseWriter, r *http.Request) {
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
		Email *string
	}

	dados := CDados{}
	if err := json.Unmarshal(ArrayByteIn, &dados); err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		logs.Erro(Retorno.Msg)
		responseReturn(w, Retorno)
		return
	}

	if dados.Email == nil {
		Retorno.Erro = true
		Retorno.Msg = "Paramêtro não informado."
		Retorno.Dados = nil
		logs.Erro(Retorno.Msg)
		responseReturn(w, Retorno)
		return
	}

	if err := database.MySql.Empresa.PesquisaArrayByteIn(ArrayByteIn); err != nil {
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

	Retorno.Dados = database.MySql.Empresa.Field
	Retorno.Msg = "Registro localizado com sucesso."
	responseReturn(w, Retorno)
}
