package Empresas

/*
	19/07/2019 16:34

	obs:
		* atualização de email e senha será efetuado
		  por uma rota especifica por questões de
		  segurança.

*/

import (
	"BaxEnd/Controller/database/dbmysql/interno/tipo_pessoa"
	"GoLibs"
	"GoLibs/logs"
	"encoding/json"
	"errors"
	"strings"
)

func (s *EmpresaST) Atualizar(ArrayByteIn []byte) (string, error) {

	dados := NewEmpresaDadosInST(s.dbConexao)

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

func (s *EmpresaST) ValidacaoAlterar(dados *EmpresaDadosInST) (string, error) {

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

	if s.Field.Ativado == false {

		// verificar possiblidade de enviar email
		// verificar possibilidade de conferir configurações
		smsg := "Registro não esta ativo, verifique o email principal."
		err := errors.New(smsg)
		return err.Error(), err
	}

	if dados.TipoPessoaID == nil {

		// quendo o tipo de pessoa nao for informado
		// verificar o tipo de pessoa do banco de dados
		// para efetuar as validações do documento passado
		// se for o caso

		if s.Field.TipoPessoaID == 0 {
			// cadastro de pessoa fisica

			if dados.Doc1 != nil {
				if err := GoLibs.IsCPF(*dados.Doc1); err != nil {
					// verificação de cpf
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

		} else if s.Field.TipoPessoaID == 1 {
			// cadastro de pessoa juridica

			if dados.Doc1 != nil {
				if err := GoLibs.IsCNPJ(*dados.Doc1); err != nil {
					// verificação de cnpj
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
		}

	} else if dados.TipoPessoaID != nil {
		// verificar a existencia do registro de tipo de pessoa
		TipoPessoa := tipo_pessoa.New()
		if err := TipoPessoa.PesquisaID(*dados.TipoPessoaID); err != nil {
			smsg := "O tipo de pessoa informado não é válido."
			err := errors.New(smsg)
			return err.Error(), err
		}

		dados.TipoPessoaDesc = &TipoPessoa.Field.Descricao

		// o tipo da pessoa tenha sido informada como parametro
		if *dados.TipoPessoaID == 0 {

			// cadastro de pessoa fisica
			if dados.Doc1 != nil {
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

			} else {
				if len(strings.TrimSpace(s.Field.Doc1)) > 0 {
					if err := GoLibs.IsCPF(s.Field.Doc1); err != nil { // verificação de cpf
						smsg := "O CPF cadastrado não é válido. " + s.Field.Doc1
						err := errors.New(smsg)
						return err.Error(), err
					}
				}
			}

		} else if *dados.TipoPessoaID == 1 {

			// cadastro de pessoa juridica
			if dados.Doc1 != nil {
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

			} else {
				if len(strings.TrimSpace(s.Field.Doc1)) > 0 {
					if err := GoLibs.IsCNPJ(s.Field.Doc1); err != nil { // verificação de cpf
						smsg := "O documento " + s.Field.Doc1 + " cadastrado não é válido, "

						if err := GoLibs.IsCPF(s.Field.Doc1); err == nil { // verificação de cpf
							smsg += "deveria ser um CNPJ pois é um CPF."
							smsg += "Informe um CNPJ na transação."
						}

						err := errors.New(smsg)
						return err.Error(), err
					}
				}
			}
		}
	}

	if dados.Doc2 != nil {

		// verificação do documento doc2 caso tenha sido informado
		if len(strings.TrimSpace(*dados.Doc2)) == 0 {
			dados.Doc2 = nil

		} else {

			// retira os numeros caso exista
			doc2, err := GoLibs.SoNumeros(*dados.Doc2)
			if err != nil {
				smsg := "Erro ao retirar letras do Doc2 [" + doc2 + "]"
				err := errors.New(smsg)
				return err.Error(), err
			}

			dados.Doc2 = &doc2
		}

	} else {

		// caso o documento esteja preenchido e contenha caracteres
		// retira e deixa apenas numeros
		if err := GoLibs.StrJustNumber(s.Field.Doc2); err != nil {
			doc2, err := GoLibs.SoNumeros(s.Field.Doc2)
			if err == nil {
				dados.Doc2 = &doc2
			}
		}

	}
	return "", nil
}
