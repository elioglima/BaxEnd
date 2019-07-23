package Usuarios

import (
	"BaxEnd/Controller/RootBuild"
	"BaxEnd/Controller/database/dbmysql/Empresas"
	"GoLibs"
	"GoLibs/logs"
	"GoMysql"
	"errors"
	"os"
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

func (s *UsuarioST) Root() error {

	s.LoadEmpresa(1)
	if err := s.PesquisaCodigo(1); err != nil {
		return err
	}

	layout := "2006-01-02T15:04:05.000Z"
	DataCompra, err := time.Parse(layout, RootBuild.EmpresaDataCompra)
	if err != nil {
		logs.Erro(err)
		os.Exit(0)
	}

	s.dbConexao.SQL.Clear()
	if s.RecordCount == 0 {
		s.dbConexao.SQL.Insert(ConsNomeTabela)
	} else {
		s.dbConexao.SQL.Update(ConsNomeTabela)
		s.dbConexao.SQL.Where("empresaid = 1 and id=1")
	}

	s.dbConexao.SQL.Add("DataCadastro", DataCompra)
	s.dbConexao.SQL.Add("DataAtualizacao", time.Now())
	s.dbConexao.SQL.Add("EmpresaID", 1)
	s.dbConexao.SQL.Add("nome", RootBuild.UsuarioNome)
	s.dbConexao.SQL.Add("email", RootBuild.UsuarioEmail)
	s.dbConexao.SQL.Add("doc1", RootBuild.UsuarioDoc1)
	s.dbConexao.SQL.Add("doc2", RootBuild.UsuarioDoc2)

	Hash, err := GoLibs.HashEncode(RootBuild.UsuarioEmail + RootBuild.UsuarioSenha)
	if err != nil {
		err := errors.New("Erro ao gerar hash de verificação, " + err.Error())
		return err
	}

	s.dbConexao.SQL.Add("senha", Hash)
	s.dbConexao.SQL.Add("ativado", 1)
	s.dbConexao.SQL.Add("DataAtivacao", DataCompra)
	if _, err := s.dbConexao.SQL.Execute(); err != nil {
		logs.Erro("Erro ao criar empresa de demonstração.")
		return err
	}

	return nil
}
