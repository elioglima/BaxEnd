package Usuarios

import (
	"GoLibs"
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
	Id             *int64             // chave não alteravel
	EmpresaID      *int64             // chave não alteravel - indica a qual empresa o usuario pertence
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

func NewUsuarioDadosInST(dbConexao *GoMysql.ConexaoST) *UsuarioDadosInST {
	s := new(UsuarioDadosInST)
	s.dbConexao = dbConexao
	return s
}

func (s *UsuarioDadosInST) Inserir() (sql.Result, error) {

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

	if s.Email != nil {
		numUp++
		s.dbConexao.SQL.Add("Email", *s.Email)
		hash, err := GoLibs.HashEncode(*s.Email + *s.Nome)
		if err != nil {
			return nil, errors.New("Erro ao gerar hash para senha temporaria, " + err.Error())
		}
		s.dbConexao.SQL.Add("Senha", hash)
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

	// if s.CategoriaID != nil {
	// 	numUp++
	// 	s.dbConexao.SQL.Add("CategoriaID", *s.CategoriaID)
	// 	s.dbConexao.SQL.Add("CategoriaDesc", *s.CategoriaDesc)
	// }

	if numUp == 0 {
		return nil, errors.New("Nenhum campo informado para atualização")
	}

	return s.dbConexao.SQL.Execute()
}

func (s *UsuarioDadosInST) Update() (sql.Result, error) {

	numUp := 0
	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Update("usuario")
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

	// if s.CategoriaID != nil {
	// 	numUp++
	// 	s.dbConexao.SQL.Add("CategoriaID", *s.CategoriaID)
	// 	s.dbConexao.SQL.Add("CategoriaDesc", *s.CategoriaDesc)
	// }

	if numUp == 0 {
		return nil, errors.New("Nenhum campo informado para atualização")
	}

	return s.dbConexao.SQL.Execute()
}
