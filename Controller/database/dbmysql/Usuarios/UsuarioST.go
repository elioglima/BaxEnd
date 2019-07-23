package Usuarios

import (
	"GoLibs"
	"GoLibs/logs"
	"GoMysql"
	"errors"
	"time"
)

const ConsNomeTabela = "usuario"

type UsuarioST struct {
	dbConexao   *GoMysql.ConexaoST
	Fields      []UsuarioDadosST
	Field       UsuarioDadosST
	Response    interface{}
	RecordCount int
}

func NewUsuarioST(dbConexaoIn *GoMysql.ConexaoST) *UsuarioST {
	s := &UsuarioST{}
	s.RecordCount = 0
	s.dbConexao = dbConexaoIn
	return s
}

func (s *UsuarioST) Demo() error {
	nome := "Elio Gonçalves de Lima"
	email := "diretoria@maxtime.info"

	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Insert(ConsNomeTabela)
	s.dbConexao.SQL.Add("DataCadastro", time.Now())
	s.dbConexao.SQL.Add("DataAtualizacao", time.Now())
	s.dbConexao.SQL.Add("EmpresaID", 1)
	s.dbConexao.SQL.Add("nome", nome)
	s.dbConexao.SQL.Add("email", email)
	s.dbConexao.SQL.Add("doc1", "21639921877")
	s.dbConexao.SQL.Add("doc2", "321666318")

	Hash, err := GoLibs.HashEncode(email + nome)
	if err != nil {
		err := errors.New("Erro ao gerar hash de verificação, " + err.Error())
		return err
	}

	s.dbConexao.SQL.Add("senha", Hash)
	s.dbConexao.SQL.Add("ativado", "1")
	s.dbConexao.SQL.Add("DataAtivacao", time.Now())
	if _, err := s.dbConexao.SQL.Execute(); err != nil {
		logs.Erro("Erro ao criar empresa de demonstração.")
		return err
	}

	return nil
}
