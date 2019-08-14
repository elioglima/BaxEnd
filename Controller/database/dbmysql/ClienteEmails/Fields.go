package ClienteEmails

import (
	"GoMysql"
)

func (s *ClienteEmailsST) MarshalResult(Results []map[string]interface{}) error {
	s.Field = ClienteEmailDadosST{}
	s.Fields = []ClienteEmailDadosST{}

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

func (s *ClienteEmailsST) MarshalResultToField(Results []map[string]interface{}) error {
	s.Field = ClienteEmailDadosST{}
	s.Field.Id = GoMysql.FirstValueToInt64(Results, "Id")
	s.Field.DataCadastro = GoMysql.FirstValueToTime(Results, "DataCadastro")
	s.Field.DataAtualizacao = GoMysql.FirstValueToTime(Results, "DataAtualizacao")
	s.Field.Email = GoMysql.FirstValueToStr(Results, "Email")
	return nil
}

func (s *ClienteEmailsST) MarshalResultToFields(Results []map[string]interface{}) error {
	s.Fields = []ClienteEmailDadosST{}
	for _, Result := range Results {
		FieldTemp := ClienteEmailDadosST{}
		FieldTemp.Id = GoMysql.GetValueToInt64(Result, "Id")
		FieldTemp.DataCadastro = GoMysql.GetValueToTime(Result, "DataCadastro")
		FieldTemp.DataAtualizacao = GoMysql.GetValueToTime(Result, "DataAtualizacao")
		FieldTemp.Email = GoMysql.GetValueToStr(Result, "Email")
		s.Fields = append(s.Fields, FieldTemp)
	}
	return nil
}
