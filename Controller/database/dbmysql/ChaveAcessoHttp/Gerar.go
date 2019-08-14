package ChaveAcessoHttp

import (
	"BaxEnd/Controller/database/dbmysql/Empresas"
	"GoLibs"
	"GoLibs/logs"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func (s *ChaveAcessoHttpST) Gerar(ArrayByteIn []byte) error {

	dados := NewChaveAcessoHttpDadosInST(s.dbConexao)
	if err := json.Unmarshal(ArrayByteIn, &dados); err != nil {
		logs.Erro(err)
		return err
	}

	if dados.EmpresaID == nil {
		smsg := "Código da emrepsa não foi informado."
		logs.Erro(smsg)
		return errors.New(smsg)

	} else if *dados.EmpresaID == 0 {
		smsg := "Código da emrepsa informado não é válido."
		logs.Erro(smsg)
		return errors.New(smsg)

	} else if dados.Descricao == nil {
		smsg := "Decrição da chave não foi informado."
		logs.Erro(smsg)
		return errors.New(smsg)

	} else if len(strings.TrimSpace(*dados.Descricao)) == 0 {
		smsg := "Decrição da chave não pode ser em branco."
		logs.Erro(smsg)
		return errors.New(smsg)

	}

	if err := s.dbConexao.CheckConnect(); err != nil {
		smsg := "Banco de dados não conectado."
		logs.Erro(smsg)
		return errors.New(smsg)
	}

	EmpresaDB := Empresas.NewEmpresaST(s.dbConexao)
	if err := EmpresaDB.PesquisaCodigo(*dados.EmpresaID); err != nil {
		smsg := err.Error()
		logs.Erro(smsg)
		return err

	} else if EmpresaDB.RecordCount == 0 {
		smsg := "Código da empresa informado (" + fmt.Sprintf("%v", *dados.EmpresaID) + "), não foi localiado."
		logs.Erro(smsg)
		return errors.New(smsg)
	}

	// adicionar dados da chave
	KeyAPI, err := GoLibs.HashEncode("KeyAPI=" + GoLibs.NowToDecimal() + "=KeyAPP")
	if err != nil {
		smsg := err.Error()
		logs.Erro(smsg)
		return errors.New(smsg)
	}
	KeyAPI = strings.Replace(KeyAPI, "/", "", -1)
	dados.KeyAPI = &KeyAPI

	KeyAPP, err := GoLibs.HashEncode("KeyAPP=" + GoLibs.NowToDecimal() + "=KeyAPI")
	if err != nil {
		smsg := err.Error()
		logs.Erro(smsg)
		return errors.New(smsg)
	}
	KeyAPP = strings.Replace(KeyAPP, "/", "", -1)
	dados.KeyAPP = &KeyAPP

	Results, err := dados.Inserir()

	if err != nil {
		logs.Erro(err)
		smsg := "Erro ao gerar key :: " + err.Error()
		err := errors.New(smsg)
		return err

	} else {

		RegistroID, err := Results.LastInsertId()
		if err != nil {
			logs.Erro(err)
			smsg := "Erro ao localizar cadastro inserido"
			err := errors.New(smsg)
			return err
		}

		type SDadosPesq struct {
			RegistroID int64
		}

		DadosPesq := SDadosPesq{}
		DadosPesq.RegistroID = RegistroID
		ArrayByteIn, err := json.Marshal(DadosPesq)
		if err != nil {
			logs.Erro(err)
			smsg := "Erro ao localizar cadastro inserido"
			err := errors.New(smsg)
			return err
		}

		if err := s.Pesquisa(ArrayByteIn); err != nil {
			logs.Erro(err)
			smsg := "Erro ao pesquisar o id localizado e inserido:" + fmt.Sprintf("%v", RegistroID)
			err := errors.New(smsg)
			return err
		}
	}

	return nil
}
