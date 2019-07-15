package instalador

import (
	libs "GoLibs"
	"GoLibs/logs"
	"database/sql"
	"errors"
)

var (
	ConexaoSystem *sql.DB
	Conexao       *sql.DB
)

func ConectarSystem() error {
	var err error
	p := mysqldb.STMysqlParams{}
	p.IP = "localhost"
	p.PORTA = 3306
	p.BANCO = "information_schema"
	p.USUARIO = "root"
	p.SENHA = "AB@102030"

	ConexaoSystem, err = mysqldb.Conectar(p)
	if err != nil {
		logs.Erro("Erro ao se conectar no servidor.", err)
		return err
	}

	logs.Sucesso("Mysql Conectado com sucesso.", ConexaoSystem, err)
	return nil
}

func Conectar() error {
	var err error
	p := mysqldb.STMysqlParams{}
	p.IP = "localhost"
	p.PORTA = 3306
	p.BANCO = "xpressapi"
	p.USUARIO = "root"
	p.SENHA = "AB@102030"

	Conexao, err = mysqldb.Conectar(p)
	if err != nil {
		logs.Erro("Erro ao se conectar no servidor.", err)
		return err
	}

	logs.Sucesso("Mysql Conectado com sucesso.", Conexao, err)
	return nil
}

func CheckDBExist() error {
	if ConexaoSystem == nil {
		smsg := "Nenhuma conexão encontrada."
		logs.Erro(smsg, Conexao)
		return errors.New(smsg)
	}

	sSQL := " SELECT schema_name FROM information_schema.schemata"
	sSQL += " WHERE schema_name = " + libs.Asp("XpressAPI")
	sSQL += " limit 0,1"
	rows, err := mysqldb.QueryDB(ConexaoSystem, sSQL)
	if err != nil {
		return err
	}

	RecordCount, _, err := mysqldb.GetRows(*rows)
	if err != nil {
		return err
	}

	if RecordCount == 0 {
		return errors.New("Banco de dados nao existe")
	}

	return nil
}

func CreateDataBase() error {
	sSQL := "CREATE SCHEMA `xpressapi` DEFAULT CHARACTER SET utf8;"
	_, err := mysqldb.Execute(ConexaoSystem, sSQL)
	return err
}

func CheckTBUsuarios() error {
	sSQL := " SELECT table_name, table_type FROM information_schema.tables"
	sSQL += " where  table_schema = 'xpressapi'"
	sSQL += " and table_name = 'usuarios'"
	rows, err := mysqldb.QueryDB(ConexaoSystem, sSQL)
	if err != nil {
		return err
	}

	RecordCount, _, err := mysqldb.GetRows(*rows)
	if err != nil {
		return err
	}

	if RecordCount == 0 {
		return errors.New("Tabela de usuarios não existe.")
	}

	return nil
}

func CheckEstrutura() error {
	err := Conectar()
	if err != nil {
		logs.Erro(err)
		return err
	}

	err = CheckTBUsuarios()
	if err != nil {
		logs.Erro(err)
		return err
	}

	return nil
}
