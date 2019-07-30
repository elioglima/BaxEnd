package token

import (
	"GoLibs"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

func (s *TokenST) Decode(dados []byte) bool {

	var (
		err              error
		Token            token2ST
		Blocos           BlocosST
		sdataToken       string
		dataToken        time.Time
		sdataValidade    string
		dataValidade     time.Time
		sdataVerificador string
		dataVerificador  time.Time
		dataHoje         time.Time
	)

	Token = token2ST{}
	dataHoje = time.Now()

	Blocos.B1 = *Token.B1
	Blocos.B2 = *Token.B2
	Blocos.B4 = *Token.B4
	Blocos.B5 = *Token.B5
	Blocos.B6 = *Token.B6

	// Bloco3 := ""
	if Token.B3 == nil {
		s.Erro = true
		s.Status = 5001
		s.Response = "Token inválido."
		return false

	} else if Token.B3.B1 == nil {
		s.Erro = true
		s.Status = 5001
		s.Response = "Token inválido."
		return false
	}

	Blocos.B3 = *Token.B3.B1

	if Token.B3.B6 != nil {
		Blocos.B3 += *Token.B3.B6
	}

	if Token.B3.B3 != nil {
		Blocos.B3 += *Token.B3.B3
	}

	if Token.B3.B5 != nil {
		Blocos.B3 += *Token.B3.B5
	}

	if Token.B3.B2 != nil {
		Blocos.B3 += *Token.B3.B2
	}

	if Token.B3.B4 != nil {
		Blocos.B3 += *Token.B3.B4
	}

	StrDecodeB64, err := base64.StdEncoding.DecodeString(*Token.B4)
	if err != nil {
		s.Erro = true
		s.Status = 5001
		s.Response = "Token inválido."
		return false
	}

	Blocos.B4 = fmt.Sprintf("%s", StrDecodeB64)

	A1 := Blocos.B2[16:18]
	A2 := Blocos.B1[6:8]
	A3 := Blocos.B2[6:10]
	A4 := Blocos.B2[0:2]
	A5 := Blocos.B1[0:2]

	A6 := Blocos.B2[2:6]   // CDE
	A7 := Blocos.B1[10:14] // TPI

	A8 := Blocos.B3
	for i := 0; i < 4; i++ {
		// A8 := base64.decode(base64.decode(base64.decode(base64.decode(Bloco3)))) // IDT
		A8, err = GoLibs.Base64Decode(A8)
		if err != nil {
			s.Erro = true
			s.Status = 5001
			s.Response = "Token inválido."
			return false
		}
	}

	A9 := Blocos.B2[14:16]
	A10 := Blocos.B2[12:14]
	A11 := Blocos.B1[2:6]
	A12 := Blocos.B2[10:12]
	A13 := Blocos.B1[8:10]

	A14 := Blocos.B4
	for i := 0; i < 4; i++ {
		// var A14 = base64.decode(base64.decode(Bloco4));
		A14, err = GoLibs.Base64Decode(A14)
		if err != nil {
			s.Erro = true
			s.Status = 5001
			s.Response = "Token inválido."
			return false
		}
	}

	// // bloco1 = 2 4 2 2 4
	// // bloco2 = 2 4 4 2 2 2 2

	sdataToken = A3 + "-" + A2 + "-" + A1 + "T" + A4 + ":" + A5 + ":00.000Z"
	layout := "2006-01-02T15:04:05.000Z"
	dataToken, err = time.Parse(layout, sdataToken)
	if err != nil {
		s.Erro = true
		s.Status = 5001
		s.Response = "Token inválido."
		return false
	}

	sdataValidade = A11 + "-" + A10 + "-" + A9 + "T" + A12 + ":" + A13 + ":00.000Z"
	dataValidade, err = time.Parse(layout, sdataValidade)
	if err != nil {
		s.Erro = true
		s.Status = 5001
		s.Response = "Token inválido."
		return false
	}

	// P4 - VERIFICADOR - 2=getHours 2=getDate 4=getFullYear 2=getMonth 2=getMinutes
	sdataVerificador = A14[4:8] + "-" + A14[8:10] + "-" + A14[2:4] + "T" + A14[0:2] + ":" + A14[10:12] + ":00.000Z"
	dataVerificador, err = time.Parse(layout, sdataVerificador)
	if err != nil {
		s.Erro = true
		s.Status = 5001
		s.Response = "Token inválido."
		return false
	}

	if dataToken.Before(dataVerificador) {
		s.Erro = true
		s.Status = 5002
		s.Response = "Não foi possível validar o token."
		return false
	}

	if dataToken.Before(dataHoje.AddDate(0, 0, -30)) {
		s.Erro = true
		s.Status = 5003
		s.Response = "Authenticação não autorizada."
		return false
	}

	if dataToken.Before(dataHoje.AddDate(0, 0, 30)) {
		s.Erro = true
		s.Status = 5003
		s.Response = "Token inválido."
		return false
	}

	if dataValidade.Before(dataHoje.AddDate(0, 0, -30)) {
		s.Erro = true
		s.Status = 5003
		s.Response = "Token expirou."
		return false
	}

	if dataValidade.Before(dataHoje.AddDate(0, 0, 30)) {
		s.Erro = true
		s.Status = 5003
		s.Response = "Token expirou."
		return false
	}

	// ({ CHVA: CHVA(), DVS: base64.encode(fdate.NowToDecimal()), CDE:A6, TPI:A7, IDT:A8})

	// gerando token
	TokenSend := TokenSendST{}
	TokenSend.CHVA = CHVA
	StrB64, err := GoLibs.Base64Encode(GoLibs.NowToDecimal())
	if err != nil {
		s.Erro = true
		s.Status = 5003
		s.Response = "Falha ao gerar token."
		return false
	}
	TokenSend.DVS = StrB64
	TokenSend.CDE = A6
	TokenSend.TPI = A7
	TokenSend.IDT = A8

	jsonToken, err := json.Marshal(TokenSend)
	if err != nil {
		s.Erro = true
		s.Status = 5003
		s.Response = "Falha ao gerar token."
		return false
	}

	s.Erro = true
	s.Status = 200
	s.Response = "Token decodificado com sucesso."
	s.DataToken = dataToken
	s.DataValidade = dataValidade
	s.CDE = A6
	s.TPI = A7
	s.IDT = A8
	s.ToBytesJson = jsonToken
	return false

}
