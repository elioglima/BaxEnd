package ChaveAcessoHttp

import (
	"GoLibs"
	"GoLibs/logs"
	"encoding/json"
	"errors"
	"strconv"
)

func (s *ChaveAcessoHttpST) Pesquisa(ArrayByteIn []byte) error {

	dados := NewChaveAcessoHttpDadosInST(s.dbConexao)
	if err := json.Unmarshal(ArrayByteIn, &dados); err != nil {
		logs.Erro(err)
		return err
	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		return errors.New("Banco de dados não conectado.")
	}

	sSQL := "select * from " + ConsNomeTabela
	CountCampo := 0

	if dados.EmpresaID != nil {
		CountCampo++
		sSQL += " where EmpresaID = " + strconv.FormatInt(*dados.EmpresaID, 10)
	}

	if dados.Descricao != nil {
		if CountCampo == 0 {
			sSQL += " where "
		} else {
			sSQL += " and "
		}
		sSQL += " descricao = " + GoLibs.Asp(*dados.Descricao)
		CountCampo++
	}

	if CountCampo == 0 {
		sSQL += " limit 0,1000 "
	}

	RecordCount, Results, err := s.dbConexao.Query(sSQL)
	if err != nil {
		return err
	}

	if err := s.MarshalResult(Results); err != nil {
		return err
	}

	s.RecordCount = RecordCount
	if s.RecordCount == 0 {
		return nil
	}

	return nil
}
