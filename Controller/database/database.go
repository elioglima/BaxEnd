package database

import (
	"BaxEnd/Controller/database/dbmysql"
	"GoLibs/logs"
)

var (
	MySql *dbmysql.ConexaoST
)

func Iniciar() {
	MySql = dbmysql.NewConexao()
	if err := MySql.RepararBanco(); err != nil {
		logs.Erro(err)
	}

	if err := MySql.Conectar(); err != nil {
		logs.Erro(err)
	}

}
