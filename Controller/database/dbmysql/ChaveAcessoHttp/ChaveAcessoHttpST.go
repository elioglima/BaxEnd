package ChaveAcessoHttp

import (
	"GoMysql"
)

const ConsNomeTabela = "ChaveAcessoHttp"

type ChaveAcessoHttpST struct {
	dbConexao   *GoMysql.ConexaoST
	Fields      []ChaveAcessoHttpDadosST
	Field       ChaveAcessoHttpDadosST
	Response    interface{}
	RecordCount int
}

func NewChaveAcessoHttpST(dbConexaoIn *GoMysql.ConexaoST) *ChaveAcessoHttpST {
	s := &ChaveAcessoHttpST{}
	s.RecordCount = 0
	s.dbConexao = dbConexaoIn
	return s
}
