package Usuarios

import (
	"BaxEnd/Controller/database/dbmysql/Empresas"
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
	Empresa     *Empresas.EmpresaST
	RecordCount int
}

func NewUsuarioST(dbConexaoIn *GoMysql.ConexaoST) *UsuarioST {
	s := &UsuarioST{}
	s.RecordCount = 0
	s.dbConexao = dbConexaoIn
	s.Empresa = Empresas.NewEmpresaST(s.dbConexao)
	return s
}

func (s *UsuarioST) LoadEmpresa(EmpresaID int64) error {
	if err := s.Empresa.PesquisaCodigo(EmpresaID); err != nil {
		logs.Erro(err)
		return err
	}

	if s.Empresa.RecordCount == 0 {
		return errors.New("Empresa não localizada")
	}

	if s.Empresa.Field.Ativado == false {
		return errors.New("O cadastro da Empresa está desativado.")
	}

	return nil
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
