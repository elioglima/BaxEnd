package ClienteEnderecos

import "time"

type ClienteEnderecoDadosST struct {
	Id              int64     `db_autprimary:"true"`
	EmpresaID       int64     `db_notnull:"true" db_unsigned:"true"`
	ClienteID       int64     `db_notnull:"true" db_unsigned:"true"`
	DataCadastro    time.Time `db_notnull:"true"`
	DataAtualizacao time.Time `db_notnull:"true"`
	TipoID          int64     `db_default:"0"`
	TipoDesc        string    `db_default:"Celular"`
	PAIS            int
	DDD             int
	Numero          string
	Ramal           int
	Contato         string
}

func NewClienteEnderecoDadosST() *ClienteEnderecoDadosST {
	s := &ClienteEnderecoDadosST{}
	return s
}
