package ClienteTelefones

import (
	"GoMysql"
)

const ConsNomeTabela = "ClienteTelefones"

type ClienteTelefonesST struct {
	dbConexao   *GoMysql.ConexaoST
	Fields      []ClienteTelefoneDadosST
	Field       ClienteTelefoneDadosST
	Response    interface{}
	RecordCount int
}

func NewClienteTelefonesST(dbConexaoIn *GoMysql.ConexaoST) *ClienteTelefonesST {
	s := &ClienteTelefonesST{}
	s.RecordCount = 0
	s.dbConexao = dbConexaoIn
	return s
}
