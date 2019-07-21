package usuarios

import (
	"GoLibs"
	"errors"
	"fmt"
	"strings"
)

func (s *UsuarioST) PesquisaCodigo(ID int64) error {

	if ID == 0 {
		return nil
	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		return errors.New("Banco de dados não conectado.")
	}

	s.RecordCount = 0

	sSQL := " select * from usuario "
	sSQL += " where id = " + fmt.Sprintf("%v", ID)
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

func (s *UsuarioST) PesquisaNome(nome_in string) error {

	if len(strings.TrimSpace(nome_in)) == 0 {
		return errors.New("Não é possível pesquisar o email, o mesmo não foi informado.")
	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		return errors.New("Banco de dados não conectado.")
	}

	s.RecordCount = 0

	sSQL := " select * from usuario "
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

func (s *UsuarioST) PesquisaEmail(email_in string) error {

	if len(strings.TrimSpace(email_in)) == 0 {
		return errors.New("Não é possível pesquisar o email, o mesmo não foi informado.")
	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		return errors.New("Banco de dados não conectado.")
	}

	s.RecordCount = 0

	sSQL := " select * from usuario "
	sSQL += " where email = " + GoLibs.Asp(email_in)
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

func (s *UsuarioST) PesquisaTodos() error {
	s.RecordCount = 0

	if err := s.dbConexao.CheckConnect(); err != nil {
		return errors.New("Banco de dados não conectado.")
	}

	sSQL := "select * from usuario "
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

func (s *UsuarioST) PesquisaEmailHash(email_in, documento_in string) (string, error) {

	if len(strings.TrimSpace(email_in)) == 0 {
		return "", errors.New("Email não informado.")

	} else if len(strings.TrimSpace(documento_in)) == 0 {
		return "", errors.New("Documento não informado.")
	}

	DocSoNumero, err := GoLibs.SoNumeros(documento_in)
	if err != nil {
		return "", errors.New("Documento informado inválido:" + err.Error())
	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		return "", errors.New("Banco de dados não conectado.")
	}

	s.RecordCount = 0

	sSQL := " select ativado, senha from usuario "
	sSQL += " where email = " + GoLibs.Asp(email_in)
	sSQL += " and doc1 = " + GoLibs.Asp(DocSoNumero)
	sSQL += " limit 0,1"
	RecordCount, Results, err := s.dbConexao.Query(sSQL)
	if err != nil {
		return "", err
	}

	if err := s.MarshalResult(Results); err != nil {
		return "", err
	}

	s.RecordCount = RecordCount
	if s.RecordCount == 0 {
		return "", nil
	}

	if s.Field.Ativado == 1 {
		return "", errors.New("Conta de usuário já ativada.")
	}

	if len(strings.TrimSpace(s.Field.Senha)) == 0 {
		return "", errors.New("Hash não definido, erro no servidor.")
	}

	return s.Field.Senha, nil
}
