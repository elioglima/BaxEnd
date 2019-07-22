package usuarios

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
	Id              *int64             // chave não alteravel
	DataCadastro    time.Time          // campo automatico
	Email           *string            // chave não alteravel
	Nome            *string            // nome compledo do usuario
	Doc1            *string            // 0 CPF ou 1 CNPJ
	Doc2            *string            // 0 RG ou 1 IE
	TipoPessoa_ID   *int64             // campo de tabela statica
	TipoPessoa_Desc *string            // campo colhe dados automatico
	Categoria_ID    *int64             // campo de tabela statica
	Categoria_Desc  *string            // campo colhe dados automatico
	dbConexao       *GoMysql.ConexaoST // classe de conexão, instanciada no inicio da aplicação
	SQLResult       sql.Result
}

func NewUsuarioDadosInST(dbConexao *GoMysql.ConexaoST) *UsuarioDadosInST {
	s := new(UsuarioDadosInST)
	s.dbConexao = dbConexao
	return s
}

func (s *UsuarioDadosInST) Inserir() (sql.Result, error) {

	numUp := 0
	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Insert("usuario")
	s.dbConexao.SQL.Add("DataCadastro", time.Now())

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

	if s.TipoPessoa_ID != nil {
		numUp++
		s.dbConexao.SQL.Add("TipoPessoa_ID", *s.TipoPessoa_ID)
		s.dbConexao.SQL.Add("TipoPessoa_Desc", *s.TipoPessoa_Desc)
	}

	if s.Categoria_ID != nil {
		numUp++
		s.dbConexao.SQL.Add("Categoria_ID", *s.Categoria_ID)
		s.dbConexao.SQL.Add("Categoria_Desc", *s.Categoria_Desc)
	}

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

	if s.TipoPessoa_ID != nil {
		s.dbConexao.SQL.Add("TipoPessoa_ID", *s.TipoPessoa_ID)
		s.dbConexao.SQL.Add("TipoPessoa_Desc", *s.TipoPessoa_Desc)
		numUp++
	}

	if s.Categoria_ID != nil {
		numUp++
		s.dbConexao.SQL.Add("Categoria_ID", *s.Categoria_ID)
		s.dbConexao.SQL.Add("Categoria_Desc", *s.Categoria_Desc)
	}

	if numUp == 0 {
		return nil, errors.New("Nenhum campo informado para atualização")
	}

	return s.dbConexao.SQL.Execute()
}
