package Usuarios

import (
	"GoLibs/logs"
	"encoding/json"
	"errors"
	"fmt"
)

func (s *UsuarioST) Apagar(ArrayByteIn []byte) (string, error) {

	dados := NewUsuarioDadosInST(s.dbConexao)

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

	dados.EmpresaID = &s.Empresa.Field.Id
	_, err = dados.Apagar()
	if err != nil {
		logs.Erro(err)
		smsg := "Erro ao apagar registro [" + fmt.Sprintf("%s", *dados.Id) + "] do usuario."
		err := errors.New(smsg)
		return err.Error(), err
	}
	return "Usuario deletado com sucesso.", nil
}

func (s *UsuarioST) ValidacaoApagar(dados *UsuarioDadosInST) (string, error) {

	if dados.EmpresaID == nil {
		smsg := "Código da empresa não informado."
		err := errors.New(smsg)
		return err.Error(), err

	} else if *dados.EmpresaID <= 0 {
		smsg := "Código da empresa inválido."
		err := errors.New(smsg)
		return err.Error(), err

	} else if dados.Id == nil {
		smsg := "Código do registro de usuário não informado."
		err := errors.New(smsg)
		return err.Error(), err

	} else if *dados.Id <= 0 {
		smsg := "Código do registro de usuário é inválido."
		err := errors.New(smsg)
		return err.Error(), err

	}

	if err := s.dbConexao.Conectar(); err != nil {
		logs.Erro(err)
		return err.Error(), err
	}

	if s.Empresa.RecordCount == 0 {
		if err := s.LoadEmpresa(int64(*dados.EmpresaID)); err != nil {
			logs.Erro(err)
			return err.Error(), err
		}
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

	return "Usuário validado com sucessos", nil
}
