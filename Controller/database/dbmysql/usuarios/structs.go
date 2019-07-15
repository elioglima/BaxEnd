package usuarios

import "time"

type UsuarioDadosST struct {
	Id             int       `db_autprimary:"true"`
	DataCadastro   time.Time `db_notnull:"true"`
	Email          string    `db_notnull:"true"`
	Senha          string    `db_notnull:"true"`
	Nome           string    `db_notnull:"true"`
	Doc1           string    `db_comm:"CNPJ ou CPF"`
	Doc2           string    `db_comm:"IE ou RG"`
	Categoria_ID   int       `db_comm:"Código da Categoria"`
	Categoria_Desc string    `db_comm:"Descrição da Categoria"`
}

type UsuarioST struct {
	Field []UsuarioDadosST
}

func NewUsuarioST() *UsuarioST {
	s := &UsuarioST{}
	return s
}

func NewUsuarioDadosST() *UsuarioDadosST {
	s := &UsuarioDadosST{}
	return s
}
