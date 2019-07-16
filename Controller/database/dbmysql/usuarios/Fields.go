package usuarios

import (
	"GoMysql"
)

func (s *UsuarioST) MarshalResult(Results []map[string]interface{}) error {
	s.Field = UsuarioDadosST{}
	s.Fields = []UsuarioDadosST{}

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

func (s *UsuarioST) MarshalResultToField(Results []map[string]interface{}) error {
	s.Field = UsuarioDadosST{}
	s.Field.Id = GoMysql.FirstValueToInt64(Results, "Id")
	s.Field.DataCadastro = GoMysql.FirstValueToTime(Results, "DataCadastro")
	s.Field.Email = GoMysql.FirstValueToStr(Results, "Email")
	s.Field.Senha = GoMysql.FirstValueToStr(Results, "Senha")
	s.Field.Nome = GoMysql.FirstValueToStr(Results, "Nome")
	s.Field.Doc1 = GoMysql.FirstValueToStr(Results, "Doc1")
	s.Field.Doc2 = GoMysql.FirstValueToStr(Results, "Doc2")
	s.Field.TipoPessoa_ID = GoMysql.FirstValueToInt(Results, "TipoPessoa_ID")
	s.Field.TipoPessoa_Desc = GoMysql.FirstValueToStr(Results, "TipoPessoa_Desc")
	s.Field.Categoria_ID = GoMysql.FirstValueToInt(Results, "Categoria_ID")
	s.Field.Categoria_Desc = GoMysql.FirstValueToStr(Results, "Categoria_Desc")
	return nil
}

func (s *UsuarioST) MarshalResultToFields(Results []map[string]interface{}) error {

	s.Fields = []UsuarioDadosST{}
	for _, Result := range Results {
		FieldTemp := UsuarioDadosST{}
		FieldTemp.Id = GoMysql.GetValueToInt64(Result, "Id")
		FieldTemp.DataCadastro = GoMysql.GetValueToTime(Result, "DataCadastro")
		FieldTemp.Email = GoMysql.GetValueToStr(Result, "Email")
		FieldTemp.Senha = GoMysql.GetValueToStr(Result, "Senha")
		FieldTemp.Nome = GoMysql.GetValueToStr(Result, "Nome")
		FieldTemp.Doc1 = GoMysql.GetValueToStr(Result, "Doc1")
		FieldTemp.Doc2 = GoMysql.GetValueToStr(Result, "Doc2")
		FieldTemp.TipoPessoa_ID = GoMysql.GetValueToInt(Result, "TipoPessoa_ID")
		FieldTemp.TipoPessoa_Desc = GoMysql.GetValueToStr(Result, "TipoPessoa_Desc")
		FieldTemp.Categoria_ID = GoMysql.GetValueToInt(Result, "Categoria_ID")
		FieldTemp.Categoria_Desc = GoMysql.GetValueToStr(Result, "Categoria_Desc")
		s.Fields = append(s.Fields, FieldTemp)
	}
	return nil
}
