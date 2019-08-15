package Token

import (
	"GoLibs/logs"
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
	EmpresaID    int64
	DataCriacao  time.Time
	DataValidade time.Time
	KeyAPI       string
	KeyAPP       string
}

func (s *TokenST) Encode() {

	algoritimo := [50]int{31, 37, 47, 9, 18, 25, 40, 6, 0, 44, 11, 12, 39, 28, 24, 45, 16, 8, 38, 15, 41, 29, 35, 26, 13, 33, 3, 7, 21, 49, 5, 1, 10, 2, 46, 27, 20, 23, 43, 36, 48, 32, 34, 17, 19, 14, 4, 22, 42, 30}

	EMPRESA_ID := "00000001"           // 8
	DT_CRIACAO := "20060102150405000"  // 17
	DT_VALIDADE := "20060102150405000" // 17
	KEYAPP := EMPRESA_ID + DT_CRIACAO + DT_VALIDADE

	for index := len(KEYAPP); index < 50; index++ {
		KEYAPP += "0"
	}

	KeyAppNew := ""
	for i := range algoritimo {
		KeyAppNew += KEYAPP[algoritimo[i] : algoritimo[i]+1]
	}

	s.KeyAPP = KeyAppNew

}

func (s *TokenST) Decode(KeyApp string) error {

	algoritimo := [50]int{31, 37, 47, 9, 18, 25, 40, 6, 0, 44, 11, 12, 39, 28, 24, 45, 16, 8, 38, 15, 41, 29, 35, 26, 13, 33, 3, 7, 21, 49, 5, 1, 10, 2, 46, 27, 20, 23, 43, 36, 48, 32, 34, 17, 19, 14, 4, 22, 42, 30}

	KeyAppNewDec := make([]string, 200)
	for a := 0; a < len(algoritimo); a++ {
		KeyAppNewDec[algoritimo[a]] = KeyApp[a : a+1]
	}

	KeyAPIDecode := ""
	for _, valor := range KeyAppNewDec {
		KeyAPIDecode += valor
	}

	s.KeyAPI = KeyAPIDecode

	EMPRESA_ID := KeyAPIDecode[0:8]
	// 2006-01-02T15:04:05.000Z - 8:25
	DT_CRIACAO := KeyAPIDecode[8:12] + "-" + KeyAPIDecode[12:14] + "-" + KeyAPIDecode[14:16] + "T" + KeyAPIDecode[16:18] + ":" + KeyAPIDecode[18:20] + ":" + KeyAPIDecode[20:22] + "." + KeyAPIDecode[22:25] + "Z"
	DT_VALIDADE := KeyAPIDecode[25:42]

	logs.Cyan(EMPRESA_ID, "EMPRESA_ID")
	logs.Cyan(DT_CRIACAO, "DT_CRIACAO")
	logs.Cyan(DT_VALIDADE, "DT_VALIDADE")
	logs.Cyan(KeyAPIDecode)
	return nil
}
