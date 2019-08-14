package ClienteEnderecos

import (
	"GoMysql"
	"database/sql"
)

/* **********************************************************************
	STRUCT ClienteEnderecoDadosInST
	Classe de Entrada e alterações de dados da Cliente e filiais

** ********************************************************************** */

type ClienteEnderecoDadosInST struct {
	Id          *int64
	EmpresaID   *int64
	ClienteID   *int64
	TipoID      *int64
	TipoDesc    *string
	Endereco    *string
	Numero      *string
	Complemento *string
	CEP         *string
	Bairro      *string
	Cidade      *string
	Estado      *string
	UF          *string
	dbConexao   *GoMysql.ConexaoST
	SQLResult   sql.Result
}

func NewClienteEnderecoDadosInST(dbConexao *GoMysql.ConexaoST) *ClienteEnderecoDadosInST {
	s := new(ClienteEnderecoDadosInST)
	s.dbConexao = dbConexao
	return s
}
