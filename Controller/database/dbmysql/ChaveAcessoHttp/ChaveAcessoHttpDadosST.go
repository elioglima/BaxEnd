package ChaveAcessoHttp

import "time"

/* **********************************************************************
	STRUCT ChaveAcessoHttpDadosST
	Usuários e Fields
** ********************************************************************** */

type ChaveAcessoHttpDadosST struct {
	RegistroID      int64     `db_autprimary:"true"`
	EmpresaID       int64     `db_notnull:"true" db_unsigned:"true"`
	DataCadastro    time.Time `db_notnull:"true"`
	DataAtualizacao time.Time `db_notnull:"true"`
	Descricao       string    `db_notnull:"true" db_unique:"true"`
	KeyAPI          string    `db_notnull:"true" db_tm1:"500"`
	KeyAPP          string    `db_notnull:"true" db_tm1:"500"`
}

func NewChaveAcessoHttpDadosST() *ChaveAcessoHttpDadosST {
	s := &ChaveAcessoHttpDadosST{}
	return s
}
