package ClienteEmails

import (
	"GoMysql"
)

const ConsNomeTabela = "ClienteEmails"

type ClienteEmailsST struct {
	dbConexao   *GoMysql.ConexaoST
	Fields      []ClienteEmailDadosST
	Field       ClienteEmailDadosST
	Response    interface{}
	RecordCount int
}

func NewClienteEmailsST(dbConexaoIn *GoMysql.ConexaoST) *ClienteEmailsST {
	s := &ClienteEmailsST{}
	s.RecordCount = 0
	s.dbConexao = dbConexaoIn
	return s
}
