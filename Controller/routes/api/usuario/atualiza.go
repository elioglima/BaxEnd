package usuario

import (
	"BaxEnd/Controller/database"
	logger "GoLibs/logs"
	"net/http"
)

func AlteraUnico(w http.ResponseWriter, r *http.Request) {
	logger.Atencao("Rota AlteraUnico", "iniciando processo")

	Retorno := sRetorno{}
	// params := mux.Vars(r)
	// id := params["id"]

	// b, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	Retorno.Erro = err
	// 	logger.Erro("Rota AlteraUnico", err)
	// 	responseReturn(w, Retorno)
	// 	return
	// } else if len(b) == 0 {
	// 	Retorno.Msg = "Erro Parametros não informado"
	// 	Retorno.Erro = errors.New(Retorno.Msg)
	// 	logger.Erro("Rota AlteraUnico", Retorno.Msg)
	// 	responseReturn(w, Retorno)
	// 	return
	// }

	// msg, err := database.MySql.Usuario.AlteraUnico(u)
	// if err != nil {
	// 	Retorno.Erro = err
	// 	Retorno.Msg = msg
	// 	responseReturn(w, Retorno)
	// 	return
	// }

	Retorno.Msg = "Usuário atualizado com sucesso."
	Retorno.Dados = database.MySql.Usuario.Field
	responseReturn(w, Retorno)
}
