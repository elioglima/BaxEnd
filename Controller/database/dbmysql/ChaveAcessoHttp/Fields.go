package ChaveAcessoHttp

import (
	"GoMysql"
)

func (s *ChaveAcessoHttpST) MarshalResult(Results []map[string]interface{}) error {
	s.Field = ChaveAcessoHttpDadosST{}
	s.Fields = []ChaveAcessoHttpDadosST{}

	if len(Results) == 0 {
		return nil
	}

	if err := s.MarshalResultToField(Results); err != nil {
		return err
	}

	if err := s.MarshalResultToFields(Results); err != nil {
		return err
	}

	return nil
}

func (s *ChaveAcessoHttpST) MarshalResultToField(Results []map[string]interface{}) error {
	s.Field = ChaveAcessoHttpDadosST{}
	s.Field.RegistroID = GoMysql.FirstValueToInt64(Results, "RegistroID")
	s.Field.EmpresaID = GoMysql.FirstValueToInt64(Results, "EmpresaId")
	s.Field.DataCadastro = GoMysql.FirstValueToTime(Results, "DataCadastro")
	s.Field.DataAtualizacao = GoMysql.FirstValueToTime(Results, "DataAtualizacao")
	s.Field.KeyAPI = GoMysql.FirstValueToStr(Results, "KeyAPI")
	s.Field.KeyAPP = GoMysql.FirstValueToStr(Results, "KeyAPP")
	return nil
}

func (s *ChaveAcessoHttpST) MarshalResultToFields(Results []map[string]interface{}) error {
	s.Fields = []ChaveAcessoHttpDadosST{}
	for _, Result := range Results {
		FieldTemp := ChaveAcessoHttpDadosST{}
		FieldTemp.RegistroID = GoMysql.GetValueToInt64(Result, "RegistroID")
		FieldTemp.EmpresaID = GoMysql.GetValueToInt64(Result, "EmpresaID")
		FieldTemp.DataCadastro = GoMysql.GetValueToTime(Result, "DataCadastro")
		FieldTemp.DataAtualizacao = GoMysql.GetValueToTime(Result, "DataAtualizacao")
		FieldTemp.KeyAPI = GoMysql.GetValueToStr(Result, "KeyAPI")
		FieldTemp.KeyAPP = GoMysql.GetValueToStr(Result, "KeyAPP")
		s.Fields = append(s.Fields, FieldTemp)
	}
	return nil
}
