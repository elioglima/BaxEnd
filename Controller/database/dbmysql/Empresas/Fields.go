package Empresas

import (
	"GoLibs"
	"GoMysql"
)

func (s *EmpresaST) MarshalResult(Results []map[string]interface{}) error {
	s.Field = EmpresaDadosST{}
	s.Fields = []EmpresaDadosST{}

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

func (s *EmpresaST) MarshalResultToField(Results []map[string]interface{}) error {
	s.Field = EmpresaDadosST{}
	s.Field.Id = GoMysql.FirstValueToInt64(Results, "Id")
	s.Field.DataCadastro = GoMysql.FirstValueToTime(Results, "DataCadastro")
	s.Field.DataAtualizacao = GoMysql.FirstValueToTime(Results, "DataAtualizacao")
	s.Field.DataAtivacao = GoMysql.FirstValueToTime(Results, "DataAtivacao")
	s.Field.Ativado = GoMysql.FirstValueToBool(Results, "Ativado")
	s.Field.Nome = GoMysql.FirstValueToStr(Results, "Nome")
	s.Field.Email = GoMysql.FirstValueToStr(Results, "Email")
	s.Field.CategoriaID = GoMysql.FirstValueToInt64(Results, "CategoriaID")
	s.Field.CategoriaDesc = GoMysql.FirstValueToStr(Results, "CategoriaDesc")

	s.Field.TipoPessoaID = GoMysql.FirstValueToInt64(Results, "TipoPessoaID")
	s.Field.TipoPessoaDesc = GoMysql.FirstValueToStr(Results, "TipoPessoaDesc")

	if s.Field.TipoPessoaID == 0 {
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

func (s *EmpresaST) MarshalResultToFields(Results []map[string]interface{}) error {

	s.Fields = []EmpresaDadosST{}
	for _, Result := range Results {
		FieldTemp := EmpresaDadosST{}
		FieldTemp.Id = GoMysql.GetValueToInt64(Result, "Id")
		FieldTemp.DataCadastro = GoMysql.GetValueToTime(Result, "DataCadastro")
		FieldTemp.DataAtualizacao = GoMysql.GetValueToTime(Result, "DataAtualizacao")
		FieldTemp.DataAtivacao = GoMysql.GetValueToTime(Result, "DataAtivacao")
		FieldTemp.Ativado = GoMysql.GetValueToBool(Result, "Ativado")
		FieldTemp.Nome = GoMysql.GetValueToStr(Result, "Nome")
		FieldTemp.Email = GoMysql.GetValueToStr(Result, "Email")

		FieldTemp.CategoriaID = GoMysql.GetValueToInt64(Result, "CategoriaID")
		FieldTemp.CategoriaDesc = GoMysql.GetValueToStr(Result, "CategoriaDesc")

		FieldTemp.TipoPessoaID = GoMysql.GetValueToInt64(Result, "TipoPessoaID")
		FieldTemp.TipoPessoaDesc = GoMysql.GetValueToStr(Result, "TipoPessoaDesc")

		if s.Field.TipoPessoaID == 0 {
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
