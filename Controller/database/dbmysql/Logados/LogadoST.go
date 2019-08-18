package Logados

import (
	"BaxEnd/Controller/database/dbmysql/Empresas"
	"GoMysql"

)

const ConsNomeTabela = "logado"

type LogadoST struct {
	dbConexao   *GoMysql.ConexaoST
	Fields      []LogadoDadosST
	Field       LogadoDadosST
	Response    interface{}
	Empresa     *Empresas.EmpresaST
	RecordCount int
}

func NewLogadoST(dbConexaoIn *GoMysql.ConexaoST) *LogadoST {
	s := &LogadoST{}
	s.RecordCount = 0
	s.dbConexao = dbConexaoIn
	s.Empresa = Empresas.NewEmpresaST(s.dbConexao)
	return s
}
