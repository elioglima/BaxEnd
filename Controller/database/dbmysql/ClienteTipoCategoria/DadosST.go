package ClienteTipoCategoria

import "time"

type ClienteTipoCategoriaDadosST struct {
	Id              int64     `db_autprimary:"true"`
	EmpresaID       int64     `db_notnull:"true" db_unsigned:"true"`
	ClienteID       int64     `db_notnull:"true" db_unsigned:"true"`
	DataCadastro    time.Time `db_notnull:"true"`
	DataAtualizacao time.Time `db_notnull:"true"`
	Descricao       string    `db_notnull:"true"`
}

func NewClienteTipoCategoriaDadosST() *ClienteTipoCategoriaDadosST {
	s := &ClienteTipoCategoriaDadosST{}
	return s
}
