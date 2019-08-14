package ClienteTipoEndereco

import (
	"GoMysql"
)

const ConsNomeTabela = "ClienteTipoEndereco"

type ClienteTipoEnderecoST struct {
	dbConexao   *GoMysql.ConexaoST
	Fields      []ClienteTipoEnderecoDadosST
	Field       ClienteTipoEnderecoDadosST
	Response    interface{}
	RecordCount int
}

func NewClienteTipoEnderecoST(dbConexaoIn *GoMysql.ConexaoST) *ClienteTipoEnderecoST {
	s := &ClienteTipoEnderecoST{}
	s.RecordCount = 0
	s.dbConexao = dbConexaoIn
	return s
}
