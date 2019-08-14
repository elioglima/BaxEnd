package Clientes

import (
	"GoMysql"
)

const ConsNomeTabela = "Cliente"

type ClienteST struct {
	dbConexao   *GoMysql.ConexaoST
	Fields      []ClienteDadosST
	Field       ClienteDadosST
	Response    interface{}
	RecordCount int
}

func NewClienteST(dbConexaoIn *GoMysql.ConexaoST) *ClienteST {
	s := &ClienteST{}
	s.RecordCount = 0
	s.dbConexao = dbConexaoIn
	return s
}
