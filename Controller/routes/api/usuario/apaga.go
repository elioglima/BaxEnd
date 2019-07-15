package usuario

import (
	"BaxEnd/Controller/database"
	"net/http"

	"github.com/gorilla/mux"
)

func Apagar(w http.ResponseWriter, r *http.Request) {
	Retorno := sRetorno{}
	params := mux.Vars(r)
	id := params["id"]
	msg, err := database.MySql.Usuario.Apagar(id)
	if err != nil {
		Retorno.Erro = err
		Retorno.Msg = msg
		responseReturn(w, Retorno)
		return
	}
	Retorno.Msg = "Usuário deletado com sucesso."
	responseReturn(w, Retorno)
}
