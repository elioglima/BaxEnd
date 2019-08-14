package ChaveAcessoHttp

import (
	"GoMysql"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type ChaveAcessoHttpDadosInST struct {
	RegistroID *int64
	EmpresaID  *int64
	Descricao  *string
	KeyAPI     *string
	KeyAPP     *string
	dbConexao  *GoMysql.ConexaoST
	SQLResult  sql.Result
}

func NewChaveAcessoHttpDadosInST(dbConexao *GoMysql.ConexaoST) *ChaveAcessoHttpDadosInST {
	s := new(ChaveAcessoHttpDadosInST)
	s.dbConexao = dbConexao
	return s
}

func (s *ChaveAcessoHttpDadosInST) Inserir() (sql.Result, error) {

	if s.EmpresaID == nil {
		return nil, errors.New("Erro interno ao verificar a empresaid.")
	} else if *s.EmpresaID == 0 {
		return nil, errors.New("Erro interno ao verificar a empresaid.")

	}

	numUp := 0
	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Insert("ChaveAcessoHttp")
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

func (s *ChaveAcessoHttpDadosInST) Update() (sql.Result, error) {
	numUp := 0
	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Update("ChaveAcessoHttp")
	s.dbConexao.SQL.Where("RegistroID=" + fmt.Sprintf("%v", *s.RegistroID))
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

func (s *ChaveAcessoHttpDadosInST) Apagar() (sql.Result, error) {

	if s.EmpresaID == nil {
		return nil, errors.New("Erro interno ao verificar a empresaid, na hora de apagar registro.")
	} else if *s.EmpresaID == 0 {
		return nil, errors.New("Erro interno ao verificar a empresaid, na hora de apagar registro.")
	} else if s.RegistroID == nil {
		return nil, errors.New("Erro interno ao verificar a id, na hora de apagar registro.")
	} else if *s.RegistroID == 0 {
		return nil, errors.New("Erro interno ao verificar a id, na hora de apagar registro.")
	}

	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Delete("ChaveAcesso")
	sWhere := "EmpresaID = " + fmt.Sprintf("%v", *s.EmpresaID)
	sWhere += " and RegistroID = " + fmt.Sprintf("%v", *s.RegistroID)
	s.dbConexao.SQL.Where(sWhere)
	return s.dbConexao.SQL.Execute()
}
