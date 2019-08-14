package ClienteEmails

import "time"

type ClienteEmailDadosST struct {
	Id              int64     `db_autprimary:"true"`
	EmpresaID       int64     `db_notnull:"true" db_unsigned:"true"`
	ClienteID       int64     `db_notnull:"true" db_unsigned:"true"`
	DataCadastro    time.Time `db_notnull:"true"`
	DataAtualizacao time.Time `db_notnull:"true"`
	Email           string
}

func NewClienteEmailDadosST() *ClienteEmailDadosST {
	s := &ClienteEmailDadosST{}
	return s
}
