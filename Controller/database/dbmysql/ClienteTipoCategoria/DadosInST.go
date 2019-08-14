package ClienteTipoCategoria

import (
	"GoMysql"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

/* **********************************************************************
	STRUCT ClienteTipoCategoriaDadosInST
	Classe de Entrada e alterações de dados da Cliente e filiais

** ********************************************************************** */

type ClienteTipoCategoriaDadosInST struct {
	Id        *int64
	EmpresaID *int64
	ClienteID *int64
	Descricao *string
	dbConexao *GoMysql.ConexaoST
	SQLResult sql.Result
}

func NewClienteTipoCategoriaDadosInST(dbConexao *GoMysql.ConexaoST) *ClienteTipoCategoriaDadosInST {
	s := new(ClienteTipoCategoriaDadosInST)
	s.dbConexao = dbConexao
	return s
}

func (s *ClienteTipoCategoriaDadosInST) Inserir() (sql.Result, error) {

	numUp := 0
	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Insert("Cliente")
	s.dbConexao.SQL.Add("DataCadastro", time.Now())
	s.dbConexao.SQL.Add("DataAtualizacao", time.Now())

	if s.Descricao != nil {
		numUp++
		s.dbConexao.SQL.Add("Descricao", *s.Descricao)
	}

	if numUp == 0 {
		return nil, errors.New("Nenhum campo informado para atualização")
	}

	return s.dbConexao.SQL.Execute()
}

func (s *ClienteTipoCategoriaDadosInST) Update() (sql.Result, error) {

	numUp := 0
	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Update("Cliente")
	s.dbConexao.SQL.Where("id=" + fmt.Sprintf("%v", *s.Id))
	s.dbConexao.SQL.Add("DataCadastro", time.Now())
	s.dbConexao.SQL.Add("DataAtualizacao", time.Now())

	if s.Descricao != nil {
		s.dbConexao.SQL.Add("Descricao", *s.Descricao)
		numUp++
	}

	if numUp == 0 {
		return nil, errors.New("Nenhum campo informado para atualização")
	}

	return s.dbConexao.SQL.Execute()
}

func (s *ClienteTipoCategoriaDadosInST) Apagar() (sql.Result, error) {

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
