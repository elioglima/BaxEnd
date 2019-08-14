package ClienteTelefones

import (
	"GoMysql"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

/* **********************************************************************
	STRUCT ClienteTelefoneDadosInST
	Classe de Entrada e alterações de dados da Cliente e filiais

** ********************************************************************** */

type ClienteTelefoneDadosInST struct {
	Id             *int64
	EmpresaID      *int64
	Email          *string
	Nome           *string
	Doc1           *string
	Doc2           *string
	TipoPessoaID   *int64
	TipoPessoaDesc *string
	CategoriaID    *int64
	CategoriaDesc  *string
	dbConexao      *GoMysql.ConexaoST
	SQLResult      sql.Result
}

func NewClienteTelefoneDadosInST(dbConexao *GoMysql.ConexaoST) *ClienteTelefoneDadosInST {
	s := new(ClienteTelefoneDadosInST)
	s.dbConexao = dbConexao
	return s
}

func (s *ClienteTelefoneDadosInST) Inserir() (sql.Result, error) {

	numUp := 0
	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Insert("Cliente")
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

func (s *ClienteTelefoneDadosInST) Update() (sql.Result, error) {

	numUp := 0
	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Update("Cliente")
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

func (s *ClienteTelefoneDadosInST) Apagar() (sql.Result, error) {

	if s.Id == nil {
		return nil, errors.New("Erro interno ao verificar a id, na hora de apagar registro.")
	} else if *s.Id == 1 {
		return nil, errors.New("O primeiro registro não pode ser alterado.")
	} else if *s.Id <= 0 {
		return nil, errors.New("Erro interno ao verificar a id, na hora de apagar registro.")
	}

	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Delete("Cliente")
	sWhere := "Id = " + fmt.Sprintf("%v", *s.Id)
	s.dbConexao.SQL.Where(sWhere)
	return s.dbConexao.SQL.Execute()
}
