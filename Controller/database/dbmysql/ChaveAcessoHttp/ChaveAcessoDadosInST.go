package ChaveAcessoHttp

import (
	"GoMysql"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type ChaveAcessoDadosInST struct {
	Id        *int64
	EmpresaID *int64
	KeyAPI    *string
	KeyAPP    *string
	dbConexao *GoMysql.ConexaoST
	SQLResult sql.Result
}

func NewChaveAcessoDadosInST(dbConexao *GoMysql.ConexaoST) *ChaveAcessoDadosInST {
	s := new(ChaveAcessoDadosInST)
	s.dbConexao = dbConexao
	return s
}

func (s *ChaveAcessoDadosInST) Inserir() (sql.Result, error) {

	if s.EmpresaID == nil {
		return nil, errors.New("Erro interno ao verificar a empresaid.")
	} else if *s.EmpresaID == 0 {
		return nil, errors.New("Erro interno ao verificar a empresaid.")

	}

	numUp := 0
	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Insert("usuario")
	s.dbConexao.SQL.Add("empresaid", *s.EmpresaID)
	s.dbConexao.SQL.Add("DataCadastro", time.Now())
	s.dbConexao.SQL.Add("DataAtualizacao", time.Now())

	if s.KeyAPI != nil {
		numUp++
		s.dbConexao.SQL.Add("KeyAPI", *s.KeyAPI)
	}

	if s.KeyAPP != nil {
		numUp++
		s.dbConexao.SQL.Add("KeyAPP", *s.KeyAPP)
	}

	if numUp == 0 {
		return nil, errors.New("Nenhum campo informado para atualização")
	}

	return s.dbConexao.SQL.Execute()
}

func (s *ChaveAcessoDadosInST) Update() (sql.Result, error) {
	numUp := 0
	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Update("usuario")
	s.dbConexao.SQL.Where("id=" + fmt.Sprintf("%v", *s.Id))
	s.dbConexao.SQL.Add("DataCadastro", time.Now())
	s.dbConexao.SQL.Add("DataAtualizacao", time.Now())

	if s.KeyAPI != nil {
		numUp++
		s.dbConexao.SQL.Add("KeyAPI", *s.KeyAPI)
	}

	if s.KeyAPP != nil {
		numUp++
		s.dbConexao.SQL.Add("KeyAPP", *s.KeyAPP)
	}

	if numUp == 0 {
		return nil, errors.New("Nenhum campo informado para atualização")
	}

	return s.dbConexao.SQL.Execute()
}

func (s *ChaveAcessoDadosInST) Apagar() (sql.Result, error) {

	if s.EmpresaID == nil {
		return nil, errors.New("Erro interno ao verificar a empresaid, na hora de apagar registro.")
	} else if *s.EmpresaID == 0 {
		return nil, errors.New("Erro interno ao verificar a empresaid, na hora de apagar registro.")
	} else if s.Id == nil {
		return nil, errors.New("Erro interno ao verificar a id, na hora de apagar registro.")
	} else if *s.Id == 0 {
		return nil, errors.New("Erro interno ao verificar a id, na hora de apagar registro.")
	}

	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Delete("ChaveAcesso")
	sWhere := "EmpresaID = " + fmt.Sprintf("%v", *s.EmpresaID)
	sWhere += " and Id = " + fmt.Sprintf("%v", *s.Id)
	s.dbConexao.SQL.Where(sWhere)
	return s.dbConexao.SQL.Execute()
}
