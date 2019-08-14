package ChaveAcessoHttp

import "time"

/* **********************************************************************
	STRUCT ChaveAcessoHttpDadosST
	Usuários e Fields
** ********************************************************************** */

type ChaveAcessoHttpDadosST struct {
	Id              int64     `db_autprimary:"true"`
	EmpresaID       int64     `db_notnull:"true" db_unsigned:"true"`
	DataCadastro    time.Time `db_notnull:"true"`
	DataAtualizacao time.Time `db_notnull:"true"`
	KeyAPI          string    `db_notnull:"true"`
	KeyAPP          string    `db_notnull:"true"`
}

func NewChaveAcessoHttpDadosST() *ChaveAcessoHttpDadosST {
	s := &ChaveAcessoHttpDadosST{}
	return s
}
