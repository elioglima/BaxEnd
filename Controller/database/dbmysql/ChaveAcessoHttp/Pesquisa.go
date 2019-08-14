package ChaveAcessoHttp

import (
	"encoding/json"
	"errors"
	"strconv"
)

func (s *ChaveAcessoHttpST) PesquisaTodos(ArrayByteIn []byte) error {
	type CDados struct {
		EmpresaID int64
	}

	dados := CDados{}
	if err := json.Unmarshal(ArrayByteIn, &dados); err != nil {
		return err
	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		return errors.New("Banco de dados não conectado.")
	}

	sSQL := "select * from " + ConsNomeTabela
	sSQL += " where EmpresaID = " + strconv.FormatInt(dados.EmpresaID, 10)
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
