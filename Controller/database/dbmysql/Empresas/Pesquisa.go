package Empresas

/*
	CLASSE DE CADASTRO DE EMPRESAS

*/

import (
	"GoLibs"
	"encoding/json"
	"errors"
	"fmt"
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
	sSQL += " where id = " + fmt.Sprintf("%v", ID)
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

func (s *EmpresaST) PesquisaArrayByteIn(ArrayByteIn []byte) error {

	type CDados struct {
		Email *string
		Nome  *string
	}

	dados := CDados{}
	if err := json.Unmarshal(ArrayByteIn, &dados); err != nil {
		return err
	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		return errors.New("Banco de dados não conectado.")
	}

	sWhere := ""
	if dados.Nome != nil {
		if len(strings.TrimSpace(*dados.Nome)) > 0 {
			sWhere += "nome like " + GoLibs.Asp(*dados.Nome+"%")
		}
	}

	if dados.Email != nil {
		if len(strings.TrimSpace(*dados.Email)) > 0 {
			if len(strings.TrimSpace(sWhere)) > 0 {
				sWhere += " and "
			}
			sWhere += " email like " + GoLibs.Asp(*dados.Email+"%")
		}
	}

	if len(strings.TrimSpace(sWhere)) <= 0 {
		return errors.New("Nenhum paramêtro informado.")
	}

	s.RecordCount = 0
	sSQL := " select * from " + ConsNomeTabela
	sSQL += " where " + sWhere
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
		return errors.New("Registro não foi localizado.")
	}

	return nil
}
