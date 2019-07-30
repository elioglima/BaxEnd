package token

import (
	"GoLibs"
	"encoding/json"
	"strings"
	"time"
)

func (s *TokenST) Encode(ArrayByteJsonIn []byte) bool {

	var (
		TokenRecive TokenReciveST
		Bloco1      string
		Bloco2      string
		Bloco3      string
		Bloco4      string
	)

	err := json.Unmarshal(ArrayByteJsonIn, &TokenRecive)
	if err != nil {
		s.Erro = true
		s.Status = 5001
		s.Response = "Token inválido."
		return false
	}

	if TokenRecive.CHVA == nil {
		s.Erro = true
		s.Status = 5001
		s.Response = "Authenticação não autorizada."
		return false

	} else if TokenRecive.DVS == nil {
		s.Erro = true
		s.Status = 5002
		s.Response = "Authenticação não autorizada."
		return false

	} else if TokenRecive.CDE == nil {
		s.Erro = true
		s.Status = 5003
		s.Response = "Authenticação não autorizada."
		return false

	} else if TokenRecive.TPI == nil {
		s.Erro = true
		s.Status = 5004
		s.Response = "Authenticação não autorizada."
		return false
	}

	if len(strings.TrimSpace(*TokenRecive.TPI)) != 4 {
		s.Erro = true
		s.Status = 5005
		s.Response = "Authenticação não autorizada."
		return false
	}

	if TokenRecive.IDT == nil {
		s.Erro = true
		s.Status = 5006
		s.Response = "Authenticação não autorizada."
		return false
	}

	if *TokenRecive.CHVA == CHVA {
		s.Erro = true
		s.Status = 5007
		s.Response = "Authenticação não autorizada."
		return false
	}

	if len(strings.TrimSpace(*TokenRecive.DVS)) == 0 {
		s.Erro = true
		s.Status = 5008
		s.Response = "Authenticação não autorizada."
		return false
	}

	DVSDecB64, err := GoLibs.Base64Decode(*TokenRecive.DVS)
	if err != nil {
		s.Erro = true
		s.Status = 5009
		s.Response = "Authenticação não autorizada."
		return false

	} else if len(strings.TrimSpace(DVSDecB64)) <= 5 {
		s.Erro = true
		s.Status = 5010
		s.Response = "Authenticação não autorizada."
		return false
	}

	dia := DVSDecB64[0:2]
	mes := DVSDecB64[2:4]
	ano := DVSDecB64[4:8]
	hora := DVSDecB64[8:10]
	min := DVSDecB64[10:12]

	sData := ano + "-" + mes + "-" + dia + "T" + hora + ":" + min + ":00.000Z"
	data, err := GoLibs.StrParseTime(sData)
	if err != nil {
		s.Erro = true
		s.Status = 5011
		s.Response = "Authenticação não autorizada."
		return false

	}

	if data.Before(data.AddDate(0, 0, -30)) {
		s.Erro = true
		s.Status = 5012
		s.Response = "Authenticação não autorizada."
		return false

	} else if data.Before(data.AddDate(0, 0, 30)) {
		s.Erro = true
		s.Status = 5012
		s.Response = "Authenticação não autorizada."
		return false

	}

	// montagem do token
	// P1 - DATA DO TOKEN
	dataHoje := time.Now()
	sdataHojeDay, sdataHojeMonth, sdataHojeYear, sdataHojeHours, sdataHojeMinute, _ := GoLibs.TimeToDecStr(dataHoje)
	A1 := sdataHojeDay
	A2 := sdataHojeMonth
	A3 := sdataHojeYear
	A4 := sdataHojeHours
	A5 := sdataHojeMinute

	// P2 - OPTIONS
	A6, err := GoLibs.FormatLeft(*TokenRecive.CDE, 4, "0")
	if err != nil {
		s.Erro = true
		s.Status = 5013
		s.Response = "Authenticação não autorizada."
		return false
	}

	A7 := *TokenRecive.TPI
	A8 := *TokenRecive.IDT

	// P3 - DATA VALIDADE TOKEN
	var dataValidade = dataHoje.AddDate(0, 0, 30)

	sDataValidadeDay, sDataValidadeMonth, sDataValidadeYear, sDataValidadeHours, sDataValidadeMinute, _ := GoLibs.TimeToDecStr(dataValidade)
	A9 := sDataValidadeDay
	A10 := sDataValidadeMonth
	A11 := sDataValidadeYear
	A12 := sDataValidadeHours
	A13 := sDataValidadeMinute

	// // P4 - VERIFICADOR - 2=getHours 2=getDate 4=getFullYear 2=getMonth 2=getMinutes
	Verificador := A4 + A1 + A3 + A2 + A5
	strVerificador64 := Verificador
	for i := 0; i < 2; i++ {
		var err error
		strVerificador64, err = GoLibs.Base64Encode(strVerificador64)
		if err != nil {
			s.Erro = true
			s.Status = 5014
			s.Response = "Authenticação não autorizada."
			return false
		}
	}

	A14 := strVerificador64

	Bloco1 = A5 + A11 + A2 + A13 + A7
	Bloco2 = A4 + A6 + A3 + A12 + A10 + A9 + A1
	Bloco3 = A8
	Bloco4 = A14

	strBloco164, err := GoLibs.Base64Encode(Bloco1)
	if err != nil {
		s.Erro = true
		s.Status = 5014
		s.Response = "Authenticação não autorizada."
		return false
	}

	B1 := strBloco164

	strBloco264, err := GoLibs.Base64Encode(Bloco2)
	if err != nil {
		s.Erro = true
		s.Status = 5014
		s.Response = "Authenticação não autorizada."
		return false
	}

	B2 := strBloco264

	strBloco364 := Bloco3
	for i := 0; i < 2; i++ {
		var err error
		strBloco364, err = GoLibs.Base64Encode(strBloco364)
		if err != nil {
			s.Erro = true
			s.Status = 5014
			s.Response = "Authenticação não autorizada."
			return false
		}
	}

	B3 := strBloco364

	strBloco464, err := GoLibs.Base64Encode(Bloco4)
	if err != nil {
		s.Erro = true
		s.Status = 5014
		s.Response = "Authenticação não autorizada."
		return false
	}

	B4 := strBloco464

	B31 := ""
	B32 := ""
	B33 := ""
	B34 := ""
	B35 := ""
	B36 := ""

	B31 = B3[0:GoLibs.IfthenI(GoLibs.Len(B3) < 20, GoLibs.Len(B3), 20)]

	if GoLibs.Len(B3) > 20 {
		B36 = B3[20:GoLibs.IfthenI(GoLibs.Len(B3) > 20 && GoLibs.Len(B3) < 40, GoLibs.Len(B3), 40)]
	}

	if GoLibs.Len(B3) > 40 {
		B33 = B3[40:GoLibs.IfthenI(GoLibs.Len(B3) > 40 && GoLibs.Len(B3) < 60, GoLibs.Len(B3), 60)]
	}

	if GoLibs.Len(B3) > 60 {
		B35 = B3[60:GoLibs.IfthenI(GoLibs.Len(B3) > 60 && GoLibs.Len(B3) < 80, GoLibs.Len(B3), 80)]
	}

	if GoLibs.Len(B3) > 80 {
		B35 = B3[80:GoLibs.IfthenI(GoLibs.Len(B3) > 80 && GoLibs.Len(B3) < 100, GoLibs.Len(B3), 100)]
	}

	if GoLibs.Len(B3) > 1000 {
		B35 = B3[100:GoLibs.IfthenI(GoLibs.Len(B3) > 100 && GoLibs.Len(B3) < 120, GoLibs.Len(B3), 120)]
	}

	s.Status = 200
	s.Response = "Authenticação efetuada com sucesso."
	s.A1.B1 = B1
	s.A1.B2 = B2
	s.A1.B3.B1 = B31
	s.A1.B3.B2 = B32
	s.A1.B3.B3 = B33
	s.A1.B3.B4 = B34
	s.A1.B3.B5 = B35
	s.A1.B3.B6 = B36
	s.A1.B4 = B4
	return true
}
