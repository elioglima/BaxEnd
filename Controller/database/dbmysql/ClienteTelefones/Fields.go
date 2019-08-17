package ClienteTelefones

import (
	"GoMysql"
)

func (s *ClienteTelefonesST) MarshalResult(Results []map[string]interface{}) error {
	s.Field = ClienteTelefoneDadosST{}
	s.Fields = []ClienteTelefoneDadosST{}

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

func (s *ClienteTelefonesST) MarshalResultToField(Results []map[string]interface{}) error {
	s.Field = ClienteTelefoneDadosST{}
	s.Field.Id = GoMysql.FirstValueToInt64(Results, "Id")
	s.Field.DataCadastro = GoMysql.FirstValueToTime(Results, "DataCadastro")
	s.Field.DataAtualizacao = GoMysql.FirstValueToTime(Results, "DataAtualizacao")
	s.Field.TipoID = GoMysql.FirstValueToInt64(Results, "TipoID")
	s.Field.TipoDesc = GoMysql.FirstValueToStr(Results, "TipoDesc")
	s.Field.PAIS = GoMysql.FirstValueToInt(Results, "PAIS")
	s.Field.DDD = GoMysql.FirstValueToInt(Results, "DDD")
	s.Field.Numero = GoMysql.FirstValueToStr(Results, "Numero")
	s.Field.Ramal = GoMysql.FirstValueToInt(Results, "Ramal")
	s.Field.Contato = GoMysql.FirstValueToStr(Results, "Contato")
	return nil
}

func (s *ClienteTelefonesST) MarshalResultToFields(Results []map[string]interface{}) error {

	s.Fields = []ClienteTelefoneDadosST{}
	for _, Result := range Results {
		FieldTemp := ClienteTelefoneDadosST{}
		FieldTemp.Id = GoMysql.GetValueToInt64(Result, "Id")
		FieldTemp.DataCadastro = GoMysql.GetValueToTime(Result, "DataCadastro")
		FieldTemp.DataAtualizacao = GoMysql.GetValueToTime(Result, "DataAtualizacao")
		FieldTemp.TipoID = GoMysql.GetValueToInt64(Result, "TipoID")
		FieldTemp.TipoDesc = GoMysql.GetValueToStr(Result, "TipoDesc")
		FieldTemp.PAIS = GoMysql.GetValueToInt(Result, "PAIS")
		FieldTemp.DDD = GoMysql.GetValueToInt(Result, "DDD")
		FieldTemp.Numero = GoMysql.GetValueToStr(Result, "Numero")
		FieldTemp.Ramal = GoMysql.GetValueToInt(Result, "Ramal")
		FieldTemp.Contato = GoMysql.GetValueToStr(Result, "Contato")
		s.Fields = append(s.Fields, FieldTemp)
	}
	return nil
}
