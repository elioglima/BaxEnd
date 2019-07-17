package usuarios

import (
	"BaxEnd/Controller/database/dbmysql/interno/tipo_pessoa"
	"GoLibs"
	"GoLibs/logs"
	"encoding/json"
	"errors"
	"strings"
)

type StUsersAlteraUnicoIn struct {
	ID    string  `bson:"id" json:"id"`
	Nome  *string `bson:"nome" json:"nome"`
	Email *string `bson:"email" json:"email"`
}

func (s *UsuarioST) AlteraUnico(ArrayByteIn []byte) (string, error) {

	dados := NewUsuarioDadosInST(s.dbConexao)

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

	smsg := "Usuario atualizado com sucesso."
	return smsg, nil
}

func (s *UsuarioST) ValidacaoAlterar(dados *UsuarioDadosInST) (string, error) {

	/*
		VERIFICAR O DOCUMENTO DIGITADO
		* SE CONDIS COM O TIPO DE PESSOA

	*/

	if dados.Id == nil {
		if s.Field.Id == 0 {
			smsg := "Paramêtro ID não informado."
			err := errors.New(smsg)
			return err.Error(), err
		}

		dados.Id = &s.Field.Id
	}

	if s.RecordCount == 0 {
		if err := s.PesquisaCodigo(*dados.Id); err != nil {
			smsg := "Erro ao pesquisar o id não foi localizado"
			err := errors.New(smsg)
			return err.Error(), err
		}
	}

	if dados.TipoPessoa_ID == nil {

		// quendo o tipo de pessoa nao for informado
		// verificar o tipo de pessoa do banco de dados
		// para efetuar as validações do documento passado
		// se for o caso

		if s.Field.TipoPessoa_ID == 0 {
			// cadastro de pessoa fisica

			if dados.Doc1 != nil {
				if err := GoLibs.IsCPF(*dados.Doc1); err != nil {
					// verificação de cpf
					smsg := "O CPF informado não é válido."
					err := errors.New(smsg)
					return err.Error(), err
				}
			}

		} else if s.Field.TipoPessoa_ID == 0 {
			// cadastro de pessoa juridica

			if dados.Doc1 != nil {
				if err := GoLibs.IsCPF(*dados.Doc1); err != nil {
					// verificação de cnpj
					smsg := "O CNPJ informado não é válido."
					err := errors.New(smsg)
					return err.Error(), err
				}
			}
		}

	} else if dados.TipoPessoa_ID != nil {

		// verificar a existencia do registro de tipo de pessoa
		TipoPessoa := tipo_pessoa.New()
		if err := TipoPessoa.PesquisaID(*dados.TipoPessoa_ID); err != nil {
			smsg := "O tipo de pessoa informado não é válido."
			err := errors.New(smsg)
			return err.Error(), err
		}

		// o tipo da pessoa tenha sido informada como parametro
		if *dados.TipoPessoa_ID == 0 {
			// cadastro de pessoa fisica

			if dados.Doc1 != nil {
				if err := GoLibs.IsCPF(*dados.Doc1); err != nil {
					// verificação de cpf
					smsg := "O CPF informado não é válido."
					err := errors.New(smsg)
					return err.Error(), err
				}
			}

		} else if *dados.TipoPessoa_ID == 0 {
			// cadastro de pessoa juridica

			if dados.Doc1 != nil {
				if err := GoLibs.IsCPF(*dados.Doc1); err != nil {
					// verificação de cnpj
					smsg := "O CNPJ informado não é válido."
					err := errors.New(smsg)
					return err.Error(), err
				}
			}
		}
	}

	if dados.Doc2 != nil {
		if len(strings.TrimSpace(*dados.Doc2)) == 0 {
			dados.Doc2 = nil

		} else {
			if smsg, err := GoLibs.SoNumeros(*dados.Doc2); err != nil {
				smsg := "Erro ao retirar letras do Doc2 [" + smsg + "]"
				err := errors.New(smsg)
				return err.Error(), err
			}
		}
	}

	return "", nil
}
