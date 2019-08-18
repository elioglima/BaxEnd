package Empresas

import (
	"GoMysql"
	"database/sql"
	"errors"
	"fmt"
	"time"

)

/* **********************************************************************
	STRUCT EmpresaDadosInST
	Classe de Entrada e alterações de dados da empresa e filiais

** ********************************************************************** */

type EmpresaDadosInST struct {
	Id             *int64             // chave não alteravel
	Email          *string            // chave não alteravel
	Nome           *string            // nome compledo do usuario
	Doc1           *string            // 0 CPF ou 1 CNPJ
	Doc2           *string            // 0 RG ou 1 IE
	TipoPessoaID   *int64             // campo de tabela statica
	TipoPessoaDesc *string            // campo colhe dados automatico
	CategoriaID    *int64             // campo de tabela statica
	CategoriaDesc  *string            // campo colhe dados automatico
	dbConexao      *GoMysql.ConexaoST // classe de conexão, instanciada no inicio da aplicação
	SQLResult      sql.Result
}

func NewEmpresaDadosInST(dbConexao *GoMysql.ConexaoST) *EmpresaDadosInST {
	s := new(EmpresaDadosInST)
	s.dbConexao = dbConexao
	return s
}

func (s *EmpresaDadosInST) Inserir() (sql.Result, error) {

	numUp := 0
	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Insert("empresa")
	s.dbConexao.SQL.Add("DataCadastro", time.Now())
	s.dbConexao.SQL.Add("DataAtualizacao", time.Now())

	if s.Email != nil {
		numUp++
		s.dbConexao.SQL.Add("Email", *s.Email)
	}

	if s.Nome != nil {
		numUp++
		s.dbConexao.SQL.Add("Nome", *s.Nome)
	}

	if s.Doc1 != nil {
		numUp++
		s.dbConexao.SQL.Add("Doc1", *s.Doc1)
	}

	if s.Doc2 != nil {
		numUp++
		s.dbConexao.SQL.Add("Doc2", *s.Doc2)
	}

	if s.TipoPessoaID != nil {
		numUp++
		s.dbConexao.SQL.Add("TipoPessoaID", *s.TipoPessoaID)
		s.dbConexao.SQL.Add("TipoPessoaDesc", *s.TipoPessoaDesc)
	}

	if numUp == 0 {
		return nil, errors.New("Nenhum campo informado para atualização")
	}

	return s.dbConexao.SQL.Execute()
}

func (s *EmpresaDadosInST) Update() (sql.Result, error) {

	numUp := 0
	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Update("empresa")
	s.dbConexao.SQL.Where("id=" + fmt.Sprintf("%v", *s.Id))
	s.dbConexao.SQL.Add("DataCadastro", time.Now())
	s.dbConexao.SQL.Add("DataAtualizacao", time.Now())

	if s.Nome != nil {
		s.dbConexao.SQL.Add("Nome", *s.Nome)
		numUp++
	}

	if s.Doc1 != nil {
		s.dbConexao.SQL.Add("Doc1", *s.Doc1)
		numUp++
	}

	if s.Doc2 != nil {
		s.dbConexao.SQL.Add("Doc2", *s.Doc2)
		numUp++
	}

	if s.TipoPessoaID != nil {
		s.dbConexao.SQL.Add("TipoPessoaID", *s.TipoPessoaID)
		s.dbConexao.SQL.Add("TipoPessoaDesc", *s.TipoPessoaDesc)
		numUp++
	}

	if numUp == 0 {
		return nil, errors.New("Nenhum campo informado para atualização")
	}

	return s.dbConexao.SQL.Execute()
}

func (s *EmpresaDadosInST) Apagar() (sql.Result, error) {

	if s.Id == nil {
		return nil, errors.New("Erro interno ao verificar a id, na hora de apagar registro.")
	} else if *s.Id == 1 {
		return nil, errors.New("O primeiro registro não pode ser alterado.")
	} else if *s.Id <= 0 {
		return nil, errors.New("Erro interno ao verificar a id, na hora de apagar registro.")
	}

	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Delete("empresa")
	sWhere := "Id = " + fmt.Sprintf("%v", *s.Id)
	s.dbConexao.SQL.Where(sWhere)
	return s.dbConexao.SQL.Execute()
}
