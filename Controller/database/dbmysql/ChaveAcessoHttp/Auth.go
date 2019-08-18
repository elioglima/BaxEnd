package ChaveAcessoHttp

import (
	"GoLibs"
	"GoLibs/logs"
	"errors"
	"strings"

)

func (s *ChaveAcessoHttpST) Auth(KeyAPI string) error {

	if len(strings.TrimSpace(KeyAPI)) == 0 {
		err := errors.New("autenticação no banco de dados inválida.")
		logs.Erro(err.Error())
		return err
	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		err := errors.New("banco de dados não conectado.")
		logs.Erro(err.Error())
		return err
	}

	sSQL := "select * from " + ConsNomeTabela
	sSQL += " where KeyAPI = " + GoLibs.Asp(KeyAPI)
	sSQL += " limit 0,1 "

	RecordCount, Results, err := s.dbConexao.Query(sSQL)
	if err != nil {
		logs.Erro(err.Error())
		return err
	}

	if err := s.MarshalResult(Results); err != nil {
		logs.Erro(err.Error())
		return err
	}

	s.RecordCount = RecordCount
	if s.RecordCount == 0 {
		return nil
	}

	return nil
}
