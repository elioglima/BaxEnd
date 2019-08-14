package ClienteTelefones

import (
	"BaxEnd/Controller/database/dbmysql/interno/tipo_pessoa"
	"GoLibs"
	"GoLibs/logs"
	"encoding/json"
	"errors"
	"fmt"
)

func (s *ClienteTelefonesST) Novo(ArrayByteIn []byte) (string, error) {

	dados := NewClienteTelefoneDadosInST(s.dbConexao)

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

	Results, err := dados.Inserir()

	if err != nil {
		logs.Erro(err)
		smsg := "Erro ao inserir cadastro"
		err := errors.New(smsg)
		return err.Error(), err

	} else {

		id, err := Results.LastInsertId()
		if err != nil {
			logs.Erro(err)
			smsg := "Erro ao localizar cadastro inserido"
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

	return "Cadastrado efetuado com sucesso.", nil
}

func (s *ClienteTelefonesST) ValidacaoNovo(dados *ClienteTelefoneDadosInST) (string, error) {

	if dados.Nome == nil {
		smsg := "O nome não foi informado. [nome]"
		err := errors.New(smsg)
		return err.Error(), err

	} else if dados.Email == nil {
		smsg := "O e-mail não foi informado. [email]"
		err := errors.New(smsg)
		return err.Error(), err

	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		smsg := "Banco de dados não conectado."
		err := errors.New(smsg)
		return err.Error(), err
	}

	sWhere := "email = " + GoLibs.Asp(*dados.Email)
	if err := s.PesquisaWhere(sWhere); err != nil {
		logs.Erro(err)
		smsg := "Erro ao pesquisar o email se existe:" + *dados.Email
		err := errors.New(smsg)
		return err.Error(), err
	}

	if s.RecordCount > 0 {
		smsg := "Email [" + *dados.Email + "] já consta na cadastrado na base de dados."
		err := errors.New(smsg)
		return err.Error(), err
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

	return "Cadastro validado com sucesso", nil
}
