package usuarios

import (
	"GoMysql"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

/* **********************************************************************
	STRUCT UsuarioST
	Classe de Usuário com os metodos fornecidos para as rotas

** ********************************************************************** */

type UsuarioDadosInST struct {
	Id              *int64
	DataCadastro    time.Time
	Email           *string
	Senha           *string
	Nome            *string
	Doc1            *string
	Doc2            *string
	TipoPessoa_ID   *int64
	TipoPessoa_Desc *string
	Categoria_ID    *int64
	Categoria_Desc  *string
	dbConexao       *GoMysql.ConexaoST
	SQLResult       sql.Result
}

func NewUsuarioDadosInST(dbConexao *GoMysql.ConexaoST) *UsuarioDadosInST {
	s := new(UsuarioDadosInST)
	s.dbConexao = dbConexao
	return s
}

func (s *UsuarioDadosInST) Inserir() (sql.Result, error) {
	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Insert("usuario")
	s.dbConexao.SQL.Add("DataCadastro", time.Now())

	if s.Email != nil {
		s.dbConexao.SQL.Add("Email", *s.Email)
	}

	if s.Senha != nil {
		s.dbConexao.SQL.Add("Senha", *s.Senha)
	}

	if s.Nome != nil {
		s.dbConexao.SQL.Add("Nome", *s.Nome)
	}

	if s.Doc1 != nil {
		s.dbConexao.SQL.Add("Doc1", *s.Doc1)
	}

	if s.Doc2 != nil {
		s.dbConexao.SQL.Add("Doc2", *s.Doc2)
	}

	if s.TipoPessoa_ID != nil {
		s.dbConexao.SQL.Add("TipoPessoa_ID", *s.TipoPessoa_ID)
	}
	if s.Categoria_ID != nil {
		s.dbConexao.SQL.Add("Categoria_ID", *s.Categoria_ID)
	}
	if s.Categoria_Desc != nil {
		s.dbConexao.SQL.Add("Categoria_Desc", *s.Categoria_Desc)
	}

	return s.dbConexao.SQL.Execute()
}

func (s *UsuarioDadosInST) Update() (sql.Result, error) {

	numUp := 0

	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Update("usuario")
	s.dbConexao.SQL.Where("id=" + fmt.Sprintf("%v", *s.Id))
	s.dbConexao.SQL.Add("DataCadastro", time.Now())

	if s.Nome != nil {
		s.dbConexao.SQL.Add("Nome", *s.Nome)
		numUp++
	}

	if s.Doc2 != nil {
		s.dbConexao.SQL.Add("Doc2", *s.Doc2)
		numUp++
	}

	if numUp == 0 {
		return nil, errors.New("Nenhum campo informado para atualização")
	}

	return s.dbConexao.SQL.Execute()
}
