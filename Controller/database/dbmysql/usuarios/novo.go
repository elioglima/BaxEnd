package Usuarios

import (
	"BaxEnd/Controller/database/dbmysql/interno/tipo_pessoa"
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

func (s *UsuarioST) Novo(ArrayByteIn []byte) (string, error) {

	dados := NewUsuarioDadosInST(s.dbConexao)

	err := json.Unmarshal(ArrayByteIn, &dados)
	if err != nil {
		smsg := "Json recebido é inválido. \n" + err.Error()
		err := errors.New(smsg)
		return err.Error(), err
	}

	smsg, err := s.ValidacaoNovo(dados)
	if err != nil {
		logs.Erro(err)
		err := errors.New(smsg)
		return err.Error(), err
	}

	if s.RecordCount > 0 {
		err := errors.New(smsg)
		return err.Error(), nil
	}

	dados.EmpresaID = &s.Empresa.Field.Id
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

func (s *UsuarioST) ValidacaoNovo(dados *UsuarioDadosInST) (string, error) {

	if dados.Nome == nil {
		smsg := "O nome não foi informado. [nome]"
		err := errors.New(smsg)
		return err.Error(), err

	} else if dados.Email == nil {
		smsg := "O e-mail não foi informado. [email]"
		err := errors.New(smsg)
		return err.Error(), err

	}

	if err := s.PesquisaEmail(*dados.Email); err != nil {
		logs.Erro(err)
		smsg := "Erro ao pesquisar o email se existe:" + *dados.Email
		err := errors.New(smsg)
		return err.Error(), nil
	}

	if s.RecordCount > 0 {
		smsg := "Email [" + *dados.Email + "] já consta na cadastrado na base de dados."
		err := errors.New(smsg)
		return err.Error(), nil
	}

	if dados.TipoPessoaID == nil {
		smsg := "O tipo de pessoa não foi informado. [tipopessoaid]"
		err := errors.New(smsg)
		return err.Error(), err
	}

	// verificar a existencia do registro de tipo de pessoa
	TipoPessoa := tipo_pessoa.New()
	if err := TipoPessoa.PesquisaID(*dados.TipoPessoaID); err != nil {
		smsg := "O tipo de pessoa informado não é válido."
		err := errors.New(smsg)
		return err.Error(), err
	}

	dados.TipoPessoaDesc = &TipoPessoa.Field.Descricao

	if *dados.TipoPessoaID == 0 { // o tipo da pessoa tenha sido informada como parametro

		if dados.Doc1 == nil {
			smsg := "CPF não informado. [doc1]"
			err := errors.New(smsg)
			return err.Error(), err

		} else if dados.Doc1 != nil { // cadastro de pessoa fisica

			if err := GoLibs.IsCPF(*dados.Doc1); err != nil { // verificação de cpf
				smsg := "O CPF informado não é válido."
				err := errors.New(smsg)
				return err.Error(), err
			}

			Doc1SoNumeros, err := GoLibs.SoNumeros(*dados.Doc1)
			if err != nil {
				return err.Error(), err
			}

			dados.Doc1 = &Doc1SoNumeros
		}

	} else if *dados.TipoPessoaID == 1 {

		// cadastro de pessoa juridica
		if dados.Doc1 == nil {
			smsg := "CNPJ não informado. [doc1]"
			err := errors.New(smsg)
			return err.Error(), err

		} else if dados.Doc1 != nil {
			if err := GoLibs.IsCNPJ(*dados.Doc1); err != nil { // verificação de cnpj
				smsg := "O CNPJ informado não é válido."
				err := errors.New(smsg)
				return err.Error(), err
			}

			Doc1SoNumeros, err := GoLibs.SoNumeros(*dados.Doc1)
			if err != nil {
				return err.Error(), err
			}

			dados.Doc1 = &Doc1SoNumeros
		}

	} else {
		smsg := "O tipo de pessoa informado não é válido."
		err := errors.New(smsg)
		return err.Error(), err
	}

	return "Usuário validado com sucesso", nil
}
