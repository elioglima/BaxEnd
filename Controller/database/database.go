package database

import "BaxEnd/Controller/database/dbmysql"

var (
	MySql *dbmysql.ConexaoST
)

func Conectar() {
	MySql = dbmysql.NewConexao()
}
