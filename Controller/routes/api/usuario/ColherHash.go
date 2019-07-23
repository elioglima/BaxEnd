package usuario

import (
	"BaxEnd/Controller/MsgsTexto"
	"BaxEnd/Controller/database"
	"GoLibs"
	"errors"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func ColherHash(w http.ResponseWriter, r *http.Request) {

	Retorno := sRetorno{}
	Retorno.Ini()
	params := mux.Vars(r)

	email := params["email"]
	documento := params["documento"]

	if len(strings.TrimSpace(email)) == 0 {
		err := errors.New("E-mail não informado.")
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return

	} else if len(strings.TrimSpace(documento)) == 0 {
		err := errors.New("Documento não informado.")
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return

	}

	DocSoNumero, err := GoLibs.SoNumeros(documento)
	if err != nil {
		err := errors.New("Documento informado inválido:" + err.Error())
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	sSQL := " where email = " + GoLibs.Asp(email)
	sSQL += " and doc1 = " + GoLibs.Asp(DocSoNumero)
	err = database.MySql.Usuario.PesquisaWhere(sSQL)
	if err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	if database.MySql.Usuario.Field.Ativado {
		err := errors.New(MsgsTexto.MsgContaAtivada())
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	Hash, err := GoLibs.HashEncode(database.MySql.Usuario.Field.Email + database.MySql.Usuario.Field.Nome)
	if err != nil {
		err := errors.New("Erro ao gerar hash de verificação, " + err.Error())
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	type RetornoST = struct {
		Hash string
	}

	Retorno.Msg = "Hash encontrado com sucesso."
	Retorno.Dados = Hash
	responseReturn(w, Retorno)
}
