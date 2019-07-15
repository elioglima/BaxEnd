package usuario

import (
	"BaxEnd/Controller/database"
	"net/http"

	"github.com/gorilla/mux"
)

func PesquisaTodos(w http.ResponseWriter, r *http.Request) {
	Retorno := sRetorno{}
	database.MySql.Usuario.PesquisaTodos()
	Retorno.Dados = database.MySql.Usuario.Field
	Retorno.Msg = "Usuários localizado com sucesso."
	responseReturn(w, Retorno)
}

func PesquisaCodigo(w http.ResponseWriter, r *http.Request) {
	Retorno := sRetorno{}
	params := mux.Vars(r)
	ID := params["id"]
	database.MySql.Usuario.PesquisaCodigo(ID)
	Retorno.Dados = database.MySql.Usuario.Field
	Retorno.Msg = "Usuário localizado com sucesso."
	responseReturn(w, Retorno)
}

func PesquisaNome(w http.ResponseWriter, r *http.Request) {
	Retorno := sRetorno{}
	params := mux.Vars(r)
	value := params["value"]
	database.MySql.Usuario.PesquisaNome(value)
	Retorno.Dados = database.MySql.Usuario.Field
	Retorno.Msg = "Usuário localizado com sucesso."
	responseReturn(w, Retorno)
}

func PesquisaEmail(w http.ResponseWriter, r *http.Request) {
	Retorno := sRetorno{}
	params := mux.Vars(r)
	value := params["value"]
	database.MySql.Usuario.PesquisaEmail(value)
	Retorno.Dados = database.MySql.Usuario.Field
	Retorno.Msg = "Usuário localizado com sucesso."
	responseReturn(w, Retorno)
}
