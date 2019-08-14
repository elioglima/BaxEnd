package ClienteTipoEndereco

import (
	"GoMysql"
)

func (s *ClienteTipoEnderecoST) MarshalResult(Results []map[string]interface{}) error {
	s.Field = ClienteTipoEnderecoDadosST{}
	s.Fields = []ClienteTipoEnderecoDadosST{}

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

func (s *ClienteTipoEnderecoST) MarshalResultToField(Results []map[string]interface{}) error {
	s.Field = ClienteTipoEnderecoDadosST{}
	s.Field.Id = GoMysql.FirstValueToInt64(Results, "Id")
	s.Field.DataCadastro = GoMysql.FirstValueToTime(Results, "DataCadastro")
	s.Field.DataAtualizacao = GoMysql.FirstValueToTime(Results, "DataAtualizacao")
	s.Field.TipoID = GoMysql.FirstValueToInt64(Results, "TipoID")
	s.Field.TipoDesc = GoMysql.FirstValueToStr(Results, "TipoDesc")
	s.Field.Endereco = GoMysql.FirstValueToStr(Results, "Endereco")
	s.Field.Numero = GoMysql.FirstValueToStr(Results, "Numero")
	s.Field.Complemento = GoMysql.FirstValueToStr(Results, "Complemento")
	s.Field.CEP = GoMysql.FirstValueToStr(Results, "CEP")
	s.Field.Bairro = GoMysql.FirstValueToStr(Results, "Bairro")
	s.Field.Cidade = GoMysql.FirstValueToStr(Results, "Cidade")
	s.Field.Estado = GoMysql.FirstValueToStr(Results, "Estado")
	s.Field.UF = GoMysql.FirstValueToStr(Results, "UF")
	return nil
}

func (s *ClienteTipoEnderecoST) MarshalResultToFields(Results []map[string]interface{}) error {

	s.Fields = []ClienteTipoEnderecoDadosST{}
	for _, Result := range Results {
		FieldTemp := ClienteTipoEnderecoDadosST{}
		FieldTemp.Id = GoMysql.GetValueToInt64(Result, "Id")
		FieldTemp.DataCadastro = GoMysql.GetValueToTime(Result, "DataCadastro")
		FieldTemp.DataAtualizacao = GoMysql.GetValueToTime(Result, "DataAtualizacao")
		FieldTemp.TipoID = GoMysql.GetValueToInt64(Result, "TipoID")
		FieldTemp.TipoDesc = GoMysql.GetValueToStr(Result, "TipoDesc")
		FieldTemp.Endereco = GoMysql.GetValueToStr(Result, "Endereco")
		FieldTemp.Numero = GoMysql.GetValueToStr(Result, "Numero")
		FieldTemp.Complemento = GoMysql.GetValueToStr(Result, "Complemento")
		FieldTemp.CEP = GoMysql.GetValueToStr(Result, "CEP")
		FieldTemp.Bairro = GoMysql.GetValueToStr(Result, "Bairro")
		FieldTemp.Cidade = GoMysql.GetValueToStr(Result, "Cidade")
		FieldTemp.Estado = GoMysql.GetValueToStr(Result, "Estado")
		FieldTemp.UF = GoMysql.GetValueToStr(Result, "UF")
	}
	return nil
}
