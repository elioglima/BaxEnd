package ClienteEnderecos

import (
	"GoMysql"
)

const ConsNomeTabela = "ClienteEnderecos"

type ClienteEnderecosST struct {
	dbConexao   *GoMysql.ConexaoST
	Fields      []ClienteEnderecoDadosST
	Field       ClienteEnderecoDadosST
	Response    interface{}
	RecordCount int
}

func NewClienteEnderecosST(dbConexaoIn *GoMysql.ConexaoST) *ClienteEnderecosST {
	s := &ClienteEnderecosST{}
	s.RecordCount = 0
	s.dbConexao = dbConexaoIn
	return s
}
