package Token

import (
	"GoLibs"
	"errors"
	"fmt"
	"strconv"
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
	EmpresaTipoDoc int
	EmpresaDoc     string
	DataCriacao    time.Time
	DataValidade   time.Time
	KeyAPI         string
	KeyAPIHash     string
	KeyAPP         string
}

func (s *TokenST) Encode() error {

	algoritimo := [60]int{41, 27, 47, 59, 1, 18, 25, 20, 16, 0, 14, 31, 42, 29, 28, 51, 5, 17, 26, 35, 6, 8, 58, 7, 30, 15, 48, 11, 9, 56, 57, 45, 13, 10, 34, 3, 4, 19, 33, 39, 40, 43, 36, 46, 2, 23, 38, 54, 37, 53, 22, 12, 50, 52, 24, 21, 49, 44, 55, 32}

	s.EmpresaID = 1                  // 8
	s.EmpresaTipoDoc = 1             // 1
	s.EmpresaDoc = "000021639921877" // 15
	s.DataCriacao = time.Now()       // "20060102150405000"  // 17
	s.DataValidade = time.Now()      // "20060102150405000" // 17

	sEmpresaID, err := GoLibs.FormatLeft(fmt.Sprintf("%v", s.EmpresaID), 8, "0")
	if err != nil {
		return err
	}

	sEmpresaDoc, err := GoLibs.FormatLeft(s.EmpresaDoc, 15, "0")
	if err != nil {
		return err
	}

	KEYAPPTemp := sEmpresaID
	KEYAPPTemp += strconv.Itoa(s.EmpresaTipoDoc)
	KEYAPPTemp += sEmpresaDoc
	KEYAPPTemp += GoLibs.FormatDateTime("JustNumber", s.DataCriacao)
	KEYAPPTemp += GoLibs.FormatDateTime("JustNumber", s.DataValidade)

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
	return nil
}

func (s *TokenST) Decode(KeyApp string) error {

	algoritimo := [60]int{41, 27, 47, 59, 1, 18, 25, 20, 16, 0, 14, 31, 42, 29, 28, 51, 5, 17, 26, 35, 6, 8, 58, 7, 30, 15, 48, 11, 9, 56, 57, 45, 13, 10, 34, 3, 4, 19, 33, 39, 40, 43, 36, 46, 2, 23, 38, 54, 37, 53, 22, 12, 50, 52, 24, 21, 49, 44, 55, 32}
	if len(KeyApp) < len(algoritimo) {
		return errors.New("KeyAPP inválido.")
	}
	KeyAppNewDec := make([]string, 200)
	for a := 0; a < len(algoritimo); a++ {
		KeyAppNewDec[algoritimo[a]] = KeyApp[a : a+1]
	}

	KeyAPIDecode := ""
	for _, valor := range KeyAppNewDec {
		KeyAPIDecode += valor
	}
	s.KeyAPI = KeyAPIDecode

	KeyAPIHash, err := GoLibs.HashEncode(s.KeyAPI)
	if err != nil {
		return err
	}
	s.KeyAPIHash = KeyAPIHash

	DTLayoutDtCriacao := KeyAPIDecode[24:28] + "-" + KeyAPIDecode[28:30] + "-" + KeyAPIDecode[30:32] + "T" + KeyAPIDecode[32:34] + ":" + KeyAPIDecode[34:36] + ":" + KeyAPIDecode[36:38] + "." + KeyAPIDecode[38:41] + "Z"
	DataCriacao, err := GoLibs.StrParseTime(DTLayoutDtCriacao)
	if err != nil {
		return err
	}

	s.DataCriacao = DataCriacao

	DTLayoutDtValidade := KeyAPIDecode[41:45] + "-" + KeyAPIDecode[45:47] + "-" + KeyAPIDecode[47:49] + "T" + KeyAPIDecode[49:51] + ":" + KeyAPIDecode[51:53] + ":" + KeyAPIDecode[53:55] + "." + KeyAPIDecode[55:57] + "Z"
	DataValidade, err := GoLibs.StrParseTime(DTLayoutDtValidade)
	if err != nil {
		return err
	}

	s.DataValidade = DataValidade
	return nil
}
