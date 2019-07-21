package usuario

import (
	"BaxEnd/Controller/database"
	"net/http"

	"github.com/gorilla/mux"
)

func HashAtivacao(w http.ResponseWriter, r *http.Request) {

	Retorno := sRetorno{}
	Retorno.Ini()
	params := mux.Vars(r)

	email := params["email"]
	documento := params["documento"]

	Hash, err := database.MySql.Usuario.PesquisaEmailHash(email, documento)
	if err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	type RetornoST = struct {
		Hash string
	}

	Retorno.Msg = "Consulta efetuada com sucesso."
	Retorno.Dados = RetornoST{Hash}
	responseReturn(w, Retorno)
}
