package ClienteTelefones

import (
	"GoMysql"
	"database/sql"
)

/* **********************************************************************
	STRUCT ClienteTelefoneDadosInST
	Classe de Entrada e alterações de dados da Cliente e filiais

** ********************************************************************** */

type ClienteTelefoneDadosInST struct {
	Id        *int64
	EmpresaID *int64
	ClienteID *int64
	TipoID    *int64
	TipoDesc  *string
	PAIS      *int
	DDD       *int
	Numero    *string
	Ramal     *int
	Contato   *string
	dbConexao *GoMysql.ConexaoST
	SQLResult sql.Result
}

func NewClienteTelefoneDadosInST(dbConexao *GoMysql.ConexaoST) *ClienteTelefoneDadosInST {
	s := new(ClienteTelefoneDadosInST)
	s.dbConexao = dbConexao
	return s
}
