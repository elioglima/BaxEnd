package usuarios

import (
	"GoLibs"
	"GoLibs/logs"
	"encoding/json"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StUsersNovoIn struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nome  *string            `bson:"nome" json:"nome"`
	Email *string            `bson:"email" json:"email"`
	Senha *string            `bson:"senha" json:"senha"`
}

func (s *UsuarioST) NovoUnico(ArrayByteIn []byte) (string, error) {

	dados := NewUsuarioDadosInST(s.dbConexao)

	err := json.Unmarshal(ArrayByteIn, &dados)
	if err != nil {
		smsg := "Json recebido é inválido. \n" + err.Error()
		err := errors.New(smsg)
		return err.Error(), err
	}

	if dados.Nome == nil {
		smsg := "O nome não foi informado."
		err := errors.New(smsg)
		return err.Error(), err

	} else if dados.Email == nil {
		smsg := "O e-mail não foi informado."
		err := errors.New(smsg)
		return err.Error(), err

	} else if dados.TipoPessoa_ID == nil {
		smsg := "O tipo de pessoa não foi informado."
		err := errors.New(smsg)
		return err.Error(), err
	}

	if *dados.TipoPessoa_ID == 0 {
		if dados.Doc1 == nil {
			smsg := "O CPF não foi informado."
			err := errors.New(smsg)
			return err.Error(), err

		}

		if err := GoLibs.IsCPF(*dados.Doc1); err != nil {
			smsg := "O CPF não é válido."
			err := errors.New(smsg)
			return err.Error(), err
		}

	} else if *dados.TipoPessoa_ID == 1 {

		if dados.Doc1 == nil {
			smsg := "O CNPJ não foi informado."
			err := errors.New(smsg)
			return err.Error(), err
		}

		if err := GoLibs.IsCNPJ(*dados.Doc1); err != nil {
			smsg := "O CNPJ não é válido."
			err := errors.New(smsg)
			return err.Error(), err
		}

	} else {
		smsg := "O tipo de pessoa informado não é válido."
		err := errors.New(smsg)
		return err.Error(), err
	}

	if dados.Senha == nil {
		smsg := "O senha não foi informado."
		err := errors.New(smsg)
		return err.Error(), err

	}

	if err := s.PesquisaEmail(*dados.Email); err != nil {
		logs.Erro(err)
		smsg := "Erro ao pesquisar o email se existe:" + *dados.Email
		err := errors.New(smsg)
		return err.Error(), err
	}

	if s.RecordCount > 0 {
		return "Email " + *dados.Email + " já consta na cadastrado na base de dados.", nil
	}

	Results, err := dados.Inserir()

	if err != nil {
		logs.Erro(err)
		smsg := "Erro ao inserir usuario"
		err := errors.New(smsg)
		return err.Error(), err

	} else {

		id, err := Results.LastInsertId()
		if err != nil {
			logs.Erro(err)
			smsg := "Erro ao localizar usuario inserido"
			err := errors.New(smsg)
			return err.Error(), err
		}

		if err := s.PesquisaCodigo(id); err != nil {
			logs.Erro(err)
			smsg := "Erro ao pesquisar o id localizado e inserido:" + fmt.Sprintf("%v", id)
			err := errors.New(smsg)
			return err.Error(), err
		}
	}

	return "Usuario cadastrado com sucesso.", nil
}

func (s *UsuarioST) InserirVarios(u []UsuarioDadosST) (string, error) {
	return "Usuario cadastrado com sucesso.", nil
}
