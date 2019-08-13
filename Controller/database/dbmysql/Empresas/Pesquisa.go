package Empresas

/*
	CLASSE DE CADASTRO DE EMPRESAS

*/

import (
	"GoLibs"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

func (s *EmpresaST) PesquisaTodos(ArrayByteIn []byte) error {
	type CDados struct {
	}

	dados := CDados{}
	if err := json.Unmarshal(ArrayByteIn, &dados); err != nil {
		return err
	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		return errors.New("Banco de dados não conectado.")
	}

	sSQL := "select * from " + ConsNomeTabela
	sSQL += " limit 0,1000 "
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

func (s *EmpresaST) PesquisaCodigo(ID int64) error {

	if ID == 0 {
		return nil
	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		return errors.New("Banco de dados não conectado.")
	}

	s.RecordCount = 0

	sSQL := " select * from " + ConsNomeTabela
	sSQL += " where id = " + strconv.FormatInt(ID, 10)
	sSQL += " limit 0,1"
	RecordCount, Results, err := s.dbConexao.Query(sSQL)
	if err != nil {
		return err
	}

	s.RecordCount = RecordCount
	if s.RecordCount == 0 {
		return nil
	}

	if err := s.MarshalResult(Results); err != nil {
		return err
	}

	return nil
}

func (s *EmpresaST) PesquisaNome(nome_in string) error {

	if len(strings.TrimSpace(nome_in)) == 0 {
		return errors.New("Não é possível pesquisar o email, o mesmo não foi informado.")
	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		return errors.New("Banco de dados não conectado.")
	}

	s.RecordCount = 0

	sSQL := " select * from " + ConsNomeTabela
	sSQL += " where nome like " + GoLibs.Asp(nome_in+"%")
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

func (s *EmpresaST) PesquisaWhere(WhereIn string) error {

	if len(strings.TrimSpace(WhereIn)) == 0 {
		return errors.New("Paramêtros de pesquisa não localizado.")

	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		return errors.New("Banco de dados não conectado.")
	}

	s.RecordCount = 0

	sSQL := " select * from " + ConsNomeTabela
	sSQL += WhereIn
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
