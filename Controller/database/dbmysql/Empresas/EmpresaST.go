package Empresas

import "GoMysql"

type EmpresaST struct {
	dbConexao   *GoMysql.ConexaoST
	Fields      []EmpresaDadosST
	Field       EmpresaDadosST
	Response    interface{}
	RecordCount int
}

func NewEmpresaST(dbConexaoIn *GoMysql.ConexaoST) *EmpresaST {
	s := &EmpresaST{}
	s.RecordCount = 0
	s.dbConexao = dbConexaoIn
	return s
}
