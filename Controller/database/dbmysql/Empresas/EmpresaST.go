package Empresas

import (
	"BaxEnd/Controller/RootBuild"
	"GoLibs/logs"
	"GoMysql"
	"os"
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

func (s *EmpresaST) Root() error {

	if err := s.PesquisaCodigo(1); err != nil {
		return err
	}

	layout := "2006-01-02T15:04:05.000Z"
	DataCompra, err := time.Parse(layout, RootBuild.EmpresaDataCompra)
	if err != nil {
		logs.Erro(err)
		os.Exit(0)
	}

	s.dbConexao.SQL.Clear()

	if s.RecordCount == 0 {
		s.dbConexao.SQL.Insert(ConsNomeTabela)
	} else {
		s.dbConexao.SQL.Update(ConsNomeTabela)
		s.dbConexao.SQL.Where("id=1")
	}

	s.dbConexao.SQL.Add("DataCadastro", DataCompra)
	s.dbConexao.SQL.Add("DataAtualizacao", time.Now())
	s.dbConexao.SQL.Add("nome", RootBuild.EmpresaNome)
	s.dbConexao.SQL.Add("email", RootBuild.EmpresaEmail)
	s.dbConexao.SQL.Add("doc1", RootBuild.EmpresaDoc1)
	s.dbConexao.SQL.Add("doc2", RootBuild.EmpresaDoc2)
	s.dbConexao.SQL.Add("ativado", 1)
	s.dbConexao.SQL.Add("DataAtivacao", DataCompra)

	if _, err := s.dbConexao.SQL.Execute(); err != nil {
		logs.Erro("Erro ao criar empresa de demonstração.")
		return err
	}

	return nil
}
