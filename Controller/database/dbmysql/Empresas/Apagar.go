package Empresas

import (
	"GoLibs/logs"
	"encoding/json"
	"errors"
	"fmt"
)

func (s *EmpresaST) Apagar(ArrayByteIn []byte) (string, error) {

	dados := NewEmpresaDadosInST(s.dbConexao)

	err := json.Unmarshal(ArrayByteIn, &dados)
	if err != nil {
		smsg := "Json recebido é inválido. \n" + err.Error()
		err := errors.New(smsg)
		return err.Error(), err
	}

	smsg, err := s.ValidacaoApagar(dados)
	if err != nil {
		logs.Erro(err)
		err := errors.New(smsg)
		return err.Error(), err
	}

	_, err = dados.Apagar()
	if err != nil {
		logs.Erro(err)
		smsg := "Erro ao apagar registro [" + fmt.Sprintf("%s", *dados.Id) + "]."
		err := errors.New(smsg)
		return err.Error(), err
	}
	return "Usuario deletado com sucesso.", nil
}

func (s *EmpresaST) ValidacaoApagar(dados *EmpresaDadosInST) (string, error) {

	if dados.Id == nil {
		smsg := "Código do registro não informado."
		err := errors.New(smsg)
		return err.Error(), err

	} else if *dados.Id <= 0 {
		smsg := "Código do registro é inválido."
		err := errors.New(smsg)
		return err.Error(), err

	} else if *dados.Id == 1 {
		smsg := "O primeiro registro nao pode ser apagado."
		err := errors.New(smsg)
		return err.Error(), err

	}

	if err := s.dbConexao.Conectar(); err != nil {
		logs.Erro(err)
		return err.Error(), err
	}

	if err := s.PesquisaCodigo(*dados.Id); err != nil {
		logs.Erro(err)
		smsg := "Erro ao localizar o registro:" + fmt.Sprintf("%v", *dados.Id)
		err := errors.New(smsg)
		return err.Error(), err
	}

	if s.RecordCount == 0 {
		smsg := "Registro não existe no banco de dados."
		err := errors.New(smsg)
		return err.Error(), err
	}

	return "Registro validado com sucessos", nil
}
