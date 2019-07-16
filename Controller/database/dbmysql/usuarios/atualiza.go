package usuarios

import (
	"GoLibs/logs"
	"encoding/json"
	"errors"
)

type StUsersAlteraUnicoIn struct {
	ID    string  `bson:"id" json:"id"`
	Nome  *string `bson:"nome" json:"nome"`
	Email *string `bson:"email" json:"email"`
}

func (s *UsuarioST) AlteraUnico(ArrayByteIn []byte) (string, error) {

	dados := NewUsuarioDadosInST(s.dbConexao)
	dados.Id = &s.Field.Id

	err := json.Unmarshal(ArrayByteIn, &dados)
	if err != nil {
		smsg := "Json recebido é inválido. \n" + err.Error()
		err := errors.New(smsg)
		return err.Error(), err
	}

	/*

		VERIFICAR O DOCUMENTO DIGITADO
		* SO NUMERO
		* SE CONDIS COM O TIPO DE PESSOA
		* CRIAR MASCARA PARA DOC1 E DOC2

	*/

	if _, err := dados.Update(); err != nil {
		return err.Error(), err
	}

	if err := s.PesquisaCodigo(s.Field.Id); err != nil {
		logs.Erro(err)
		smsg := "Erro ao pesquisar o id localizado e inserido:"
		err := errors.New(smsg)
		return err.Error(), err
	}

	smsg := "Usuario atualizado com sucesso."
	return smsg, nil
}
