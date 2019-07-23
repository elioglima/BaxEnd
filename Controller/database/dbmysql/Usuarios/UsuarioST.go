package Usuarios

import "GoMysql"

type UsuarioST struct {
	dbConexao   *GoMysql.ConexaoST
	Fields      []UsuarioDadosST
	Field       UsuarioDadosST
	Response    interface{}
	RecordCount int
}

func NewUsuarioST(dbConexaoIn *GoMysql.ConexaoST) *UsuarioST {
	s := &UsuarioST{}
	s.RecordCount = 0
	s.dbConexao = dbConexaoIn
	return s
}
