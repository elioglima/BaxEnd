package ClienteTipoCategoria

import (
	"GoMysql"
)

func (s *ClienteTipoCategoriaST) MarshalResult(Results []map[string]interface{}) error {
	s.Field = ClienteTipoCategoriaDadosST{}
	s.Fields = []ClienteTipoCategoriaDadosST{}

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

func (s *ClienteTipoCategoriaST) MarshalResultToField(Results []map[string]interface{}) error {
	s.Field = ClienteTipoCategoriaDadosST{}
	s.Field.Id = GoMysql.FirstValueToInt64(Results, "Id")
	s.Field.DataCadastro = GoMysql.FirstValueToTime(Results, "DataCadastro")
	s.Field.DataAtualizacao = GoMysql.FirstValueToTime(Results, "DataAtualizacao")
	s.Field.Descricao = GoMysql.FirstValueToStr(Results, "Descricao")
	return nil
}

func (s *ClienteTipoCategoriaST) MarshalResultToFields(Results []map[string]interface{}) error {

	s.Fields = []ClienteTipoCategoriaDadosST{}
	for _, Result := range Results {
		FieldTemp := ClienteTipoCategoriaDadosST{}
		FieldTemp.Id = GoMysql.GetValueToInt64(Result, "Id")
		FieldTemp.DataCadastro = GoMysql.GetValueToTime(Result, "DataCadastro")
		FieldTemp.DataAtualizacao = GoMysql.GetValueToTime(Result, "DataAtualizacao")
		FieldTemp.Descricao = GoMysql.GetValueToStr(Result, "Descricao")
		s.Fields = append(s.Fields, FieldTemp)
	}
	return nil
}
