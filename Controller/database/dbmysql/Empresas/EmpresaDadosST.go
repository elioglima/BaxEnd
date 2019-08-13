package Empresas

import "time"

type EmpresaDadosST struct {
	Id              int64     `db_autprimary:"true"`
	DataCadastro    time.Time `db_notnull:"true"`
	DataAtualizacao time.Time `db_notnull:"true"`
	Nome            string    `db_notnull:"true"`
	Email           string    `db_notnull:"true"`
	TipoPessoaID    int64     `db_comm:"Tipo de pessoa ID - 0 Fisica, 1 juridica" db_default:"0"`
	TipoPessoaDesc  string    `db_comm:"Tipo de pessoa ID - 0 Fisica, 1 juridica" db_default:"Pessoa Física"`
	Doc1            string    `db_comm:"CNPJ ou CPF"`
	Doc2            string    `db_comm:"IE ou RG"`
	Ativado         bool      `db_default:"0" db_comm:"Ativação por email"`
	DataAtivacao    time.Time
	CategoriaID     int64  `db_comm:"Código da Categoria" db_default:"0"`
	CategoriaDesc   string `db_comm:"Descrição da Categoria" db_default:"Definir"`
}

func NewEmpresaDadosST() *EmpresaDadosST {
	s := &EmpresaDadosST{}
	return s
}
