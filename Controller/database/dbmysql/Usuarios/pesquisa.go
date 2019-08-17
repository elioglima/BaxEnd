package Usuarios

import (
	"GoLibs"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func (s *UsuarioST) PesquisaTodos(ArrayByteIn []byte) error {
	type CDados struct {
		EmpresaID int64
	}

	dados := CDados{}
	fmt.Printf("%+s\n", ArrayByteIn)
	if err := json.Unmarshal(ArrayByteIn, &dados); err != nil {
		return err
	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		return errors.New("Banco de dados não conectado.")
	}

	if err := s.LoadEmpresa(dados.EmpresaID); err != nil {
		return err
	}

	sSQL := "select * from " + ConsNomeTabela
	sSQL += " where EmpresaID = " + strconv.FormatInt(s.Empresa.Field.Id, 10)
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

func (s *UsuarioST) PesquisaNome(ArrayByteIn []byte) error {

	type CDados struct {
		EmpresaID int64
		Nome      string
	}

	dados := CDados{}
	fmt.Printf("%+s\n", ArrayByteIn)
	if err := json.Unmarshal(ArrayByteIn, &dados); err != nil {
		return err
	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		return errors.New("Banco de dados não conectado.")
	}

	if err := s.LoadEmpresa(dados.EmpresaID); err != nil {
		return err
	}

	s.RecordCount = 0

	sSQL := " select * from usuario "
	sSQL += " where EmpresaID = " + strconv.FormatInt(s.Empresa.Field.Id, 10)
	sSQL += " and nome like " + GoLibs.Asp(dados.Nome+"%")
	sSQL += " limit 0,100"
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

func (s *UsuarioST) PesquisaCodigo(ID int64) error {

	if ID == 0 {
		return nil
	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		return errors.New("Banco de dados não conectado.")
	}

	s.RecordCount = 0

	sSQL := " select * from usuario "
	sSQL += " where EmpresaID = " + strconv.FormatInt(s.Empresa.Field.Id, 10)
	sSQL += " and id = " + fmt.Sprintf("%v", ID)
	sSQL += " limit 0,1"
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

func (s *UsuarioST) PesquisaEmail(email_in string) error {

	if len(strings.TrimSpace(email_in)) == 0 {
		return errors.New("Não é possível pesquisar o email, o mesmo não foi informado.")
	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		return errors.New("Banco de dados não conectado.")
	}

	s.RecordCount = 0

	sSQL := " select * from usuario "
	sSQL += " where EmpresaID = " + strconv.FormatInt(s.Empresa.Field.Id, 10)
	sSQL += " and email like " + GoLibs.Asp(email_in+"%")
	sSQL += " limit 0,1"
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

func (s *UsuarioST) PesquisaWhere(WhereIn string) error {

	if len(strings.TrimSpace(WhereIn)) == 0 {
		return errors.New("Paramêtros de pesquisa não localizado.")

	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		return errors.New("Banco de dados não conectado.")
	}

	s.RecordCount = 0

	sSQL := " select * from " + ConsNomeTabela
	sSQL += " where EmpresaID = " + strconv.FormatInt(s.Empresa.Field.Id, 10)
	sSQL += " and " + strings.Replace(WhereIn, "where", "", -1)
	sSQL += " limit 0,1"
	RecordCount, Results, err := s.dbConexao.Query(sSQL)
	if err != nil {
		return err
	}

	if err := s.MarshalResult(Results); err != nil {
		return err
	}

	s.RecordCount = RecordCount
	if s.RecordCount == 0 {
		return errors.New("Usuário não foi localizado.")
	}

	return nil
}
