package usuarios

import (
	"GoLibs"
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
	s.Field.Categoria_ID = GoMysql.FirstValueToInt64(Results, "Categoria_ID")
	s.Field.Categoria_Desc = GoMysql.FirstValueToStr(Results, "Categoria_Desc")

	s.Field.TipoPessoa_ID = GoMysql.FirstValueToInt64(Results, "TipoPessoa_ID")
	s.Field.TipoPessoa_Desc = GoMysql.FirstValueToStr(Results, "TipoPessoa_Desc")

	if s.Field.TipoPessoa_ID == 0 {
		Doc1DB := GoMysql.FirstValueToStr(Results, "Doc1")
		Doc1Formatado, err := GoLibs.ImprimeCPF(Doc1DB)
		if err != nil {
			Doc1Formatado = err.Error()
		}
		s.Field.Doc1 = Doc1Formatado
	} else {
		Doc1DB := GoMysql.FirstValueToStr(Results, "Doc1")
		Doc1Formatado, err := GoLibs.ImprimeCNPJ(Doc1DB)
		if err != nil {
			Doc1Formatado = err.Error()
		}
		s.Field.Doc1 = Doc1Formatado
	}

	s.Field.Doc2 = GoMysql.FirstValueToStr(Results, "Doc2")

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

		FieldTemp.Categoria_ID = GoMysql.GetValueToInt64(Result, "Categoria_ID")
		FieldTemp.Categoria_Desc = GoMysql.GetValueToStr(Result, "Categoria_Desc")

		FieldTemp.TipoPessoa_ID = GoMysql.GetValueToInt64(Result, "TipoPessoa_ID")
		FieldTemp.TipoPessoa_Desc = GoMysql.GetValueToStr(Result, "TipoPessoa_Desc")

		if s.Field.TipoPessoa_ID == 0 {
			Doc1DB := GoMysql.GetValueToStr(Result, "Doc1")
			Doc1Formatado, err := GoLibs.ImprimeCPF(Doc1DB)
			if err != nil {
				Doc1Formatado = err.Error()
			}
			s.Field.Doc1 = Doc1Formatado
		} else {
			Doc1DB := GoMysql.GetValueToStr(Result, "Doc1")
			Doc1Formatado, err := GoLibs.ImprimeCNPJ(Doc1DB)
			if err != nil {
				Doc1Formatado = err.Error()
			}
			s.Field.Doc1 = Doc1Formatado
		}

		FieldTemp.Doc2 = GoMysql.GetValueToStr(Result, "Doc2")
		s.Fields = append(s.Fields, FieldTemp)
	}
	return nil
}
