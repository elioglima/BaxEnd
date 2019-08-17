package Usuarios

/*
	19/07/2019 16:34

	obs:
		* atualização de email e senha será efetuado
		  por uma rota especifica por questões de
		  segurança.

*/

import (
	"BaxEnd/Controller/MsgsTexto"
	"GoLibs"
	"GoLibs/logs"
	"GoMysql"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type AtivarCadastroDadosRebebidoST struct {
	Id        int64
	Hash      *string `json:"hash"`
	Senha     *string `json:"senha"`
	SenhaConf *string `json:"senha_conf"`
	dbConexao *GoMysql.ConexaoST
}

func (s *UsuarioST) AtivarCadastro(ArrayByteIn []byte) (string, error) {

	if s.Field.Ativado {
		smsg := MsgsTexto.MsgContaAtivada()
		err := errors.New(smsg)
		return err.Error(), nil
	}

	DadosRecebido := &AtivarCadastroDadosRebebidoST{}
	DadosRecebido.Id = s.Field.Id
	DadosRecebido.dbConexao = s.dbConexao

	err := json.Unmarshal(ArrayByteIn, &DadosRecebido)
	if err != nil {
		smsg := "Json recebido é inválido. \n" + err.Error()
		err := errors.New(smsg)
		return err.Error(), err
	}

	if smsg, err := s.ValidacaoDeAtivacao(DadosRecebido); err != nil {
		logs.Erro(err)
		err := errors.New(smsg)
		return err.Error(), err
	}

	if _, err := DadosRecebido.AtivarAgora(); err != nil {
		return err.Error(), err
	}

	s.Response = nil
	smsg := "Usuario ativado com sucesso."
	return smsg, nil
}

func (s *UsuarioST) ValidacaoDeAtivacao(DadosRecebido *AtivarCadastroDadosRebebidoST) (string, error) {

	if DadosRecebido.Hash == nil {
		smsg := "Campo Hash não informado."
		err := errors.New(smsg)
		return err.Error(), err
	}

	if DadosRecebido.Senha == nil {
		smsg := "Campo Senha não informado."
		err := errors.New(smsg)
		return err.Error(), err
	}

	if DadosRecebido.SenhaConf == nil {
		smsg := "Campo Senha Confirmação não informado."
		err := errors.New(smsg)
		return err.Error(), err
	}

	hash, err := GoLibs.HashEncode(s.Field.Email + s.Field.Nome)
	if err != nil {
		return "", errors.New("Erro ao gerar hash de verificação, " + err.Error())
	}

	if *DadosRecebido.Hash != hash {
		logs.Branco("Hash informado ", *DadosRecebido.Hash)
		logs.Cyan("Hash cadastro ", hash)
		smsg := "Hash inválido, solicite um novo."
		err := errors.New(smsg)
		return err.Error(), err
	}

	return "", nil
}

func (s *AtivarCadastroDadosRebebidoST) AtivarAgora() (sql.Result, error) {
	s.dbConexao.SQL.Clear()
	s.dbConexao.SQL.Update("usuario")
	s.dbConexao.SQL.Where("id=" + fmt.Sprintf("%v", s.Id))
	s.dbConexao.SQL.Add("DataAtualizacao", time.Now())
	s.dbConexao.SQL.Add("DataAtivacao", time.Now())
	s.dbConexao.SQL.Add("Ativado", 1)
	return s.dbConexao.SQL.Execute()
}
