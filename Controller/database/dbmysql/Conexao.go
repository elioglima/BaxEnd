package dbmysql

/*

	CLASSE ConexaoST

	Responsável por carregar objeto e classe para uso em geral
	e manipulação de dados do banco de dados.

*/

import (
	"BaxEnd/Controller/database/dbmysql/Empresas"
	"BaxEnd/Controller/database/dbmysql/Usuarios"
	"BaxEnd/Controller/database/dbmysql/interno/tipo_pessoa"
	"GoLibs/logs"
	"GoMysql"
)

type ConexaoST struct {
	ParamsConexao GoMysql.ParamsConexaoST
	dbConexao     *GoMysql.ConexaoST
	TipoPessoa    tipo_pessoa.TipoPessoaST
	Empresa       *Empresas.EmpresaST
	Usuario       *Usuarios.UsuarioST
}

func NewConexao() *ConexaoST {
	s := &ConexaoST{}
	s.ParamsConexao.IP = "localhost"
	s.ParamsConexao.PORTA = 3306
	s.ParamsConexao.BANCO = "DBBaxEnd"
	s.ParamsConexao.USUARIO = "root"
	s.ParamsConexao.SENHA = "AB@102030"
	s.dbConexao = GoMysql.NewConexao(s.ParamsConexao)

	s.Usuario = Usuarios.NewUsuarioST(s.dbConexao)
	s.Empresa = Empresas.NewEmpresaST(s.dbConexao)
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

	// criação de tabelas caso não exista
	Empresa := Empresas.NewEmpresaDadosST()
	if err := s.dbConexao.CreateTable(Empresa); err != nil {
		logs.Erro(err)
		return err
	}

	Usuario := Usuarios.NewUsuarioDadosST()
	if err := s.dbConexao.CreateTable(Usuario); err != nil {
		logs.Erro(err)
		return err
	}

	// importação de dados iniciais para teste

	return nil
}
