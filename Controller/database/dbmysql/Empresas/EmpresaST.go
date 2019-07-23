package Empresas

import (
	"GoLibs/logs"
	"GoMysql"
	"time"
)

const ConsNomeTabela = "empresa"

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

func (s *EmpresaST) Demo() error {
	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Insert(ConsNomeTabela)
	s.dbConexao.SQL.Add("DataCadastro", time.Now())
	s.dbConexao.SQL.Add("DataAtualizacao", time.Now())
	s.dbConexao.SQL.Add("nome", "Maxtime Info")
	s.dbConexao.SQL.Add("doc1", "21639921877")
	s.dbConexao.SQL.Add("doc2", "321666318")
	s.dbConexao.SQL.Add("ativado", "1")
	s.dbConexao.SQL.Add("DataAtivacao", time.Now())
	if _, err := s.dbConexao.SQL.Execute(); err != nil {
		logs.Erro("Erro ao criar empresa de demonstração.")
		return err
	}

	return nil
}
