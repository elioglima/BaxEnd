package ChaveAcessoHttp

import (
	"BaxEnd/Controller/config/admin"
	"GoLibs/logs"
	"GoMysql"
	"bytes"
	"encoding/json"
	"time"
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

func (s *ChaveAcessoHttpST) Root() error {

	/*
		RegistroID      int64     `db_autprimary:"true"`
		EmpresaID       int64     `db_notnull:"true" db_unsigned:"true"`
		DataCadastro    time.Time `db_notnull:"true"`
		DataAtualizacao time.Time `db_notnull:"true"`
		Descricao       string    `db_notnull:"true" db_unique:"true"`
		KeyAPI          string    `db_notnull:"true" db_tm1:"500"`
		KeyAPP          string    `db_notnull:"true" db_tm1:"500"`
	*/

	s.Field.EmpresaID = 1
	s.Field.DataValidade = time.Now().AddDate(365, 0, 0)
	s.Field.Descricao = "Chave de Acesso Admin"

	ArrayByteIn := new(bytes.Buffer)
	json.NewEncoder(ArrayByteIn).Encode(s.Field)
	if err := s.Gerar(ArrayByteIn.Bytes()); err != nil {
		logs.Erro("Erro ao criar empresa de demonstração.")
		return err
	}

	AdminBaxEndToken := "REACT_APP_KPP=" + s.Field.KeyAPP
	admin.LocalGenerateFileBaxEndToken(AdminBaxEndToken)

	return nil
}
