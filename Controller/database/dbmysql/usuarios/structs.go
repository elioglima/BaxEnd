package usuarios

import (
	"GoMysql"
	"time"
)

/* **********************************************************************
	STRUCT UsuarioDadosST
	Usuários e Fields
** ********************************************************************** */

type UsuarioDadosST struct {
	Id              int64     `db_autprimary:"true"`
	DataCadastro    time.Time `db_notnull:"true"`
	Email           string    `db_notnull:"true"`
	Senha           string    `db_notnull:"true" db_tm1:"500"`
	Nome            string    `db_notnull:"true"`
	Doc1            string    `db_comm:"CNPJ ou CPF"`
	Doc2            string    `db_comm:"IE ou RG"`
	Ativado         int       `db_default:"0" db_comm:"Ativação por email"`
	TipoPessoa_ID   int64     `db_comm:"Tipo de pessoa ID - 0 Fisica, 1 juridica" db_default:"0"`
	TipoPessoa_Desc string    `db_comm:"Tipo de pessoa ID - 0 Fisica, 1 juridica" db_default:"Pessoa Física"`
	Categoria_ID    int64     `db_comm:"Código da Categoria" db_default:"0"`
	Categoria_Desc  string    `db_comm:"Descrição da Categoria" db_default:"Definir"`
}

type UsuarioST struct {
	dbConexao   *GoMysql.ConexaoST
	Fields      []UsuarioDadosST
	Field       UsuarioDadosST
	RecordCount int
}

func NewUsuarioST(dbConexaoIn *GoMysql.ConexaoST) *UsuarioST {
	s := &UsuarioST{}
	s.RecordCount = 0
	s.dbConexao = dbConexaoIn
	return s
}

func NewUsuarioDadosST() *UsuarioDadosST {
	s := &UsuarioDadosST{}
	return s
}
