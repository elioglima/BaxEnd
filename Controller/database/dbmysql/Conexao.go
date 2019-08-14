package dbmysql

/*

	CLASSE ConexaoST

	Responsável por carregar objeto e classe para uso em geral
	e manipulação de dados do banco de dados.

*/

import (
	"BaxEnd/Controller/database/dbmysql/ChaveAcessoHttp"
	"BaxEnd/Controller/database/dbmysql/Empresas"
	"BaxEnd/Controller/database/dbmysql/Usuarios"
	"BaxEnd/Controller/database/dbmysql/interno/tipo_pessoa"
	"GoLibs/logs"
	"GoMysql"
	"os"
)

type ConexaoST struct {
	ParamsConexao   GoMysql.ParamsConexaoST
	dbConexao       *GoMysql.ConexaoST
	Empresa         *Empresas.EmpresaST
	ChaveAcessoHttp *ChaveAcessoHttp.ChaveAcessoHttpST
	Usuario         *Usuarios.UsuarioST
	TipoPessoa      tipo_pessoa.TipoPessoaST
}

func NewConexao() *ConexaoST {
	s := &ConexaoST{}
	s.ParamsConexao.IP = "localhost"
	s.ParamsConexao.PORTA = 3306
	s.ParamsConexao.BANCO = "DBBaxEnd"
	s.ParamsConexao.USUARIO = "root"
	s.ParamsConexao.SENHA = "AB@102030"
	s.dbConexao = GoMysql.NewConexao(s.ParamsConexao)

	s.Empresa = Empresas.NewEmpresaST(s.dbConexao)
	s.ChaveAcessoHttp = ChaveAcessoHttp.NewChaveAcessoHttpST(s.dbConexao)
	s.Usuario = Usuarios.NewUsuarioST(s.dbConexao)

	return s
}

func (s *ConexaoST) Conectar() error {

	if err := s.dbConexao.ConectarSystem(); err != nil {
		return err
	}

	if err := s.dbConexao.CheckDBExist(); err != nil {
		return err
	}

	if err := s.dbConexao.Conectar(); err != nil {
		return err
	}

	return nil
}

func (s *ConexaoST) RepararBanco() error {

	if err := s.dbConexao.CheckConnectSys(); err != nil {

		if err := s.dbConexao.ConectarSystem(); err != nil {
			return err
		}

		if err := s.dbConexao.CheckConnectSys(); err != nil {
			return err
		}

	}

	if err := s.dbConexao.CheckDBExist(); err != nil {
		if err := s.dbConexao.CreateDB(); err != nil {
			return err
		}
	}

	if err := s.dbConexao.CheckConnect(); err != nil {

		if err := s.dbConexao.Conectar(); err != nil {
			return err
		}

		if err := s.dbConexao.CheckConnect(); err != nil {
			return err
		}
	}

	if err := s.CriaEstrutura(); err != nil {
		logs.Erro(err)
		return err
	}

	return nil
}

func (s *ConexaoST) CriaEstrutura() error {

	/*
		Neste item permite

			* checar a estrutura do banco
			* criar as tabelas que falta e
			* criar dados de exemplos e testes

	*/

	ObjetoUsuario := Usuarios.NewUsuarioDadosST()
	if err := s.dbConexao.DropTable(ObjetoUsuario); err != nil {
		logs.Erro(err)
		return err
	}

	ObjetoEmpresa := Empresas.NewEmpresaDadosST()
	if err := s.dbConexao.DropTable(ObjetoEmpresa); err != nil {
		logs.Erro(err)
		return err
	}

	// inicio da limpeza da base de dados
	ObjetoChaveAcessoHttp := ChaveAcessoHttp.NewChaveAcessoHttpDadosST()
	if err := s.dbConexao.DropTable(ObjetoChaveAcessoHttp); err != nil {
		logs.Erro(err)
		return err
	}

	// inicio da criação de tabelas
	if err := s.dbConexao.CreateTable(ObjetoEmpresa); err != nil {
		logs.Erro(err)
		return err
	}

	if err := s.dbConexao.CreateTable(ObjetoChaveAcessoHttp); err != nil {
		logs.Erro(err)
		return err
	}

	if err := s.dbConexao.CreateTable(ObjetoUsuario); err != nil {
		logs.Erro(err)
		return err
	}

	if err := s.dbConexao.ForeignKey("usuario", "empresaid", "empresa", "id", true, true); err != nil {
		logs.Erro(err)
		return err
	}

	// importação de dados iniciais para teste
	Empresa := Empresas.NewEmpresaST(s.dbConexao)
	if err := Empresa.Root(); err != nil {
		logs.Erro(err)
		os.Exit(0)
	}

	Usuario := Usuarios.NewUsuarioST(s.dbConexao)
	if err := Usuario.Root(); err != nil {
		logs.Erro(err)
		os.Exit(0)
	}

	return nil
}
