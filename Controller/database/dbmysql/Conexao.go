package dbmysql

import (
	"BaxEnd/Controller/database/dbmysql/usuarios"
	"GoLibs/logs"
	"GoMysql"
)

type ConexaoST struct {
	ParamsConexao GoMysql.ParamsConexaoST
	dbConexao     *GoMysql.ConexaoST
	Usuario       *usuarios.UsuarioST
}

func NewConexao() *ConexaoST {
	s := &ConexaoST{}
	s.ParamsConexao.IP = "localhost"
	s.ParamsConexao.PORTA = 3306
	s.ParamsConexao.BANCO = "DBBaxEnd"
	s.ParamsConexao.USUARIO = "root"
	s.ParamsConexao.SENHA = "AB@102030"
	s.dbConexao = GoMysql.NewConexao(s.ParamsConexao)

	s.Usuario = usuarios.NewUsuarioST(s.dbConexao)
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

	Usuario := usuarios.NewUsuarioDadosST()
	if err := s.dbConexao.CreateTable(Usuario); err != nil {
		logs.Erro(err)
		return err
	}

	return nil
}
