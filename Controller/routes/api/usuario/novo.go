package usuario

import (
	"BaxEnd/Controller/database"
	logger "GoLibs/logs"
	"net/http"
)

func NovoUnico(w http.ResponseWriter, r *http.Request) {
	logger.Atencao("NovoUnico", "Iniciando processo")

	Retorno := sRetorno{}

	logger.Atencao("NovoUnico", "Leitura do Body")
	// b, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	logger.Erro("NovoUnico", err)
	// 	Retorno.Erro = err
	// 	responseReturn(w, Retorno)
	// 	return
	// }

	// msg, err := database.MySql.Usuario.NovoUnico(u)
	// if err != nil {
	// 	logger.Erro("NovoUnico", err)
	// 	Retorno.Erro = err
	// 	Retorno.Msg = msg
	// 	responseReturn(w, Retorno)
	// 	return
	// }

	Retorno.Msg = "Usuário cadastrado com sucesso."
	Retorno.Dados = database.MySql.Usuario.Field
	logger.Sucesso("NovoUnico", Retorno.Msg)
	responseReturn(w, Retorno)
}

func NovoVarios(w http.ResponseWriter, r *http.Request) {
	Retorno := sRetorno{}
	Retorno.Msg = "Função não implantada."
	responseReturn(w, Retorno)
}
