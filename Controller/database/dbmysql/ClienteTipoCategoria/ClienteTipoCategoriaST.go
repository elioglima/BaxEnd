package ClienteTipoCategoria

import (
	"GoMysql"
)

const ConsNomeTabela = "ClienteTipoCategoria"

type ClienteTipoCategoriaST struct {
	dbConexao   *GoMysql.ConexaoST
	Fields      []ClienteTipoCategoriaDadosST
	Field       ClienteTipoCategoriaDadosST
	Response    interface{}
	RecordCount int
}

func NewClienteTipoCategoriaST(dbConexaoIn *GoMysql.ConexaoST) *ClienteTipoCategoriaST {
	s := &ClienteTipoCategoriaST{}
	s.RecordCount = 0
	s.dbConexao = dbConexaoIn
	return s
}
