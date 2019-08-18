package Token

import (
	"GoLibs"
	"GoLibs/logs"
	"errors"
	"fmt"
	"strings"
	"time"

)

/*

	KeyAPI é o composto

	EmpresaID
	1 - 0 0 0 0 0 0 0 0 1
		1 2 3 4 5 6 7 8 9


	DATA CREIACAO
	2 - 2006-01-02T15:04:05.000Z
		1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17
		2 0 0 6 0 1 0 2 1  5  0  4  0  5  0  0  0

	DATA VALIDADE
	3 - 2006-01-02T15:04:05.000Z
		1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17
		2 0 0 6 0 1 0 2 1  5  0  4  0  5  0  0  0

	KEYAPP
	4 - Lya0Q+hdkLoBDOk8wQdikuWV3g0EMxIKGjGER5Ul2bOvFpHFq5ml7TI5Rhxk7qh00LmtJAiT1FGjRt1MiywVg==

*/

type TokenST struct {
	EmpresaID      int64
	EmpresaTipoDoc int64
	EmpresaDoc     string
	DataCadastro   time.Time
	DataValidade   time.Time
	KeyAPI         string
	KeyAPIHash     string
	KeyAPP         string
}

func NewToken() *TokenST {
	s := &TokenST{}
	s.EmpresaID = -1
	s.EmpresaTipoDoc = -1
	s.EmpresaDoc = ""
	s.DataCadastro = time.Now()
	s.DataValidade = time.Now().AddDate(-365, 0, 0)
	return s
}

func (s *TokenST) Encode() error {

	algoritimo := [60]int{41, 27, 47, 59, 1, 18, 25, 20, 16, 0, 14, 31, 42, 29, 28, 51, 5, 17, 26, 35, 6, 8, 58, 7, 30, 15, 48, 11, 9, 56, 57, 45, 13, 10, 34, 3, 4, 19, 33, 39, 40, 43, 36, 46, 2, 23, 38, 54, 37, 53, 22, 12, 50, 52, 24, 21, 49, 44, 55, 32}

	if s.EmpresaID <= 0 {
		err := errors.New("Erro ao gerar token:Código da empresa não informado.")
		return err
	}

	sEmpresaID, err := GoLibs.FormatLeft(fmt.Sprintf("%v", s.EmpresaID), 8, "0")
	if err != nil {
		return err
	}

	if s.EmpresaTipoDoc < 0 || s.EmpresaTipoDoc > 1 {
		err := errors.New("Erro ao gerar token:Tipo de Documento não informado." + fmt.Sprintf("%v", s.EmpresaTipoDoc))
		return err
	}

	s.EmpresaDoc = GoLibs.JustNumber(s.EmpresaDoc)
	if len(strings.TrimSpace(s.EmpresaDoc)) == 0 {
		err := errors.New("Erro ao gerar token:Documento não informado.")
		return err
	}

	if s.EmpresaTipoDoc == 0 {
		if err := GoLibs.IsCPF(s.EmpresaDoc); err != nil {
			return err
		}
	} else if s.EmpresaTipoDoc == 1 {
		if err := GoLibs.IsCNPJ(s.EmpresaDoc); err != nil {
			return err
		}
	}

	sEmpresaDoc, err := GoLibs.FormatLeft(s.EmpresaDoc, 15, "0")
	if err != nil {
		return err
	}

	s.DataCadastro = time.Now()
	s.DataValidade = time.Now().AddDate(1, 0, 0)

	KEYAPPTemp := sEmpresaID
	KEYAPPTemp += fmt.Sprintf("%v", s.EmpresaTipoDoc)
	KEYAPPTemp += sEmpresaDoc
	KEYAPPTemp += GoLibs.FormatDateTime("JustNumber2", s.DataCadastro)
	KEYAPPTemp += GoLibs.FormatDateTime("JustNumber2", s.DataValidade)
	KEYAPP, err := GoLibs.FormatRigth(KEYAPPTemp, len(algoritimo), "0")
	if err != nil {
		return err
	}

	if len(algoritimo) > len(KEYAPP) {
		return errors.New(fmt.Sprintf("%v :: %v - %v %v", "KeyAPP gerada é inválida", KEYAPP, len(KEYAPP), len(algoritimo)))
	}

	KeyAppNew := ""
	for i := range algoritimo {
		KeyAppNew += KEYAPP[algoritimo[i] : algoritimo[i]+1]
	}

	s.KeyAPP = KeyAppNew

	// adicionar dados da chave
	KeyAPI, err := GoLibs.HashEncode("KeyAPI=" + s.KeyAPP + "=KeyAPP")
	if err != nil {
		smsg := err.Error()
		logs.Erro(smsg)
		return errors.New(smsg)
	}
	s.KeyAPI = strings.ReplaceAll(KeyAPI, "/", "")

	return nil
}

func (s *TokenST) Decode(KeyAPP string) error {

	s.KeyAPP = KeyAPP
	algoritimo := [60]int{41, 27, 47, 59, 1, 18, 25, 20, 16, 0, 14, 31, 42, 29, 28, 51, 5, 17, 26, 35, 6, 8, 58, 7, 30, 15, 48, 11, 9, 56, 57, 45, 13, 10, 34, 3, 4, 19, 33, 39, 40, 43, 36, 46, 2, 23, 38, 54, 37, 53, 22, 12, 50, 52, 24, 21, 49, 44, 55, 32}
	if len(KeyAPP) < len(algoritimo) {
		return errors.New("token inválido.")
	}
	KeyAppNewDec := make([]string, 200)
	for a := 0; a < len(algoritimo); a++ {
		KeyAppNewDec[algoritimo[a]] = KeyAPP[a : a+1]
	}

	KeyAPPDecode := ""
	for _, valor := range KeyAppNewDec {
		KeyAPPDecode += valor
	}

	if len(algoritimo) > len(KeyAPPDecode) {
		return errors.New(fmt.Sprintf("%v :: %v - %v %v", "KeyAPP decodificado é inválido", KeyAPPDecode, len(KeyAPPDecode), len(algoritimo)))
	}

	// s.KeyAPP = KeyAPPDecode
	// adicionar dados da chave
	KeyAPI, err := GoLibs.HashEncode("KeyAPI=" + s.KeyAPP + "=KeyAPP")
	if err != nil {
		smsg := err.Error()
		logs.Erro(smsg)
		return errors.New(smsg)
	}
	s.KeyAPI = strings.ReplaceAll(KeyAPI, "/", "")

	DTLayoutDtCriacao := KeyAPPDecode[24:28] + "-" + KeyAPPDecode[28:30] + "-" + KeyAPPDecode[30:32] + "T" + KeyAPPDecode[32:34] + ":" + KeyAPPDecode[34:36] + ":" + KeyAPPDecode[36:38] + "." + KeyAPPDecode[38:41] + "Z"
	DataCadastro, err := GoLibs.StrParseTime(DTLayoutDtCriacao)
	if err != nil {
		return errors.New("token inválido, data criação não localizada." + DTLayoutDtCriacao)
	}

	s.DataCadastro = DataCadastro
	DTLayoutDtValidade := KeyAPPDecode[41:45] + "-" + KeyAPPDecode[45:47] + "-" + KeyAPPDecode[47:49] + "T" + KeyAPPDecode[49:51] + ":" + KeyAPPDecode[51:53] + ":" + KeyAPPDecode[53:55] + "." + KeyAPPDecode[55:58] + "Z"
	DataValidade, err := GoLibs.StrParseTime(DTLayoutDtValidade)
	if err != nil {
		return errors.New("validade do token expirou." + DTLayoutDtValidade)
	}

	s.DataValidade = DataValidade
	return nil
}
