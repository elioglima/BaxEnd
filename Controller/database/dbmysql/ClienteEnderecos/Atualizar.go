package ClienteEnderecos

/*
	19/07/2019 16:34

	obs:
		* atualização de email e senha será efetuado
		  por uma rota especifica por questões de
		  segurança.

*/

import (
	"GoLibs/logs"
	"encoding/json"
	"errors"
)

func (s *ClienteEnderecosST) Atualizar(ArrayByteIn []byte) (string, error) {

	dados := NewClienteEnderecoDadosInST(s.dbConexao)

	err := json.Unmarshal(ArrayByteIn, &dados)
	if err != nil {
		smsg := "Json recebido é inválido. \n" + err.Error()
		err := errors.New(smsg)
		return err.Error(), err
	}

	if smsg, err := s.ValidacaoAlterar(dados); err != nil {
		logs.Erro(err)
		err := errors.New(smsg)
		return err.Error(), err
	}

	if _, err := dados.Update(); err != nil {
		return err.Error(), err
	}

	if err := s.PesquisaCodigo(s.Field.Id); err != nil {
		logs.Erro(err)
		smsg := "Erro ao pesquisar o id localizado e inserido:"
		err := errors.New(smsg)
		return err.Error(), err
	}

	s.Response = nil
	smsg := "Usuario atualizado com sucesso."
	return smsg, nil
}

func (s *ClienteEnderecosST) ValidacaoAlterar(dados *ClienteEnderecoDadosInST) (string, error) {

	/*
		VERIFICAR O DOCUMENTO DIGITADO
		* SE CONDIS COM O TIPO DE PESSOA

	*/

	if dados.Id == nil {
		smsg := "Paramêtro ID não informado."
		err := errors.New(smsg)
		return err.Error(), err
	}

	if *dados.Id <= 0 {
		smsg := "Paramêtro ID não informado."
		err := errors.New(smsg)
		return err.Error(), err
	}

	if *dados.Id == 1 {
		smsg := "O primeiro registro não pode ser alterado."
		err := errors.New(smsg)
		return err.Error(), err
	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		smsg := "Banco de dados não conectado."
		err := errors.New(smsg)
		return err.Error(), err
	}

	if err := s.PesquisaCodigo(*dados.Id); err != nil {
		smsg := "Erro ao pesquisar o id não foi localizado"
		err := errors.New(smsg)
		return err.Error(), err
	}

	if s.RecordCount == 0 {
		smsg := "Registro não localizado"
		err := errors.New(smsg)
		return err.Error(), err
	}

	return "", nil
}
