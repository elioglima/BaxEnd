package main

import (
	"BaxEnd/Controller/Token"
	"GoLibs/logs"
)

func main() {

	/*
		EXEMPLO DE REGX SO NUMERO
	*/

	logs.DebugSucesso = true
	logs.DebugErro = true

	// Algoritimo := [10]int{9, 6, 2, 3, 5, 4, 1, 7, 8, 0}

	// ValorTexto := "ABCDEFGHIJ"
	// var ValorNovo string
	// for i := 0; i < len(Algoritimo); i++ {
	// 	// logs.Sucesso(i)
	// 	ValorNovo += ValorTexto[Algoritimo[i] : Algoritimo[i]+1]
	// }

	// logs.Sucesso("ValorNovo", ValorNovo)

	// AlgoritimoDec := make([]string, 10)
	// for a := 0; a < len(Algoritimo); a++ {
	// 	AlgoritimoDec[Algoritimo[a]] = ValorNovo[a : a+1]
	// }

	// logs.Sucesso(AlgoritimoDec)

	// for a := 1; a < 9; a++ {
	// 	for b := 0; b < 9; b++ {
	// 		if Algoritimo[b] == a {
	// 			logs.Cyan(ValorTexto, ValorNovo, b, b+1, ValorNovo[b-1:b])
	// 		}
	// 	}

	// }

	// logs.Sucesso(ValorNovo)

	// var Algoritimo []int
	// iCountAlgoritimo := 0
	// iLimite := 50

	// for iCountAlgoritimo < iLimite {
	// 	bExiste := false
	// 	iPush := rand.Intn(iLimite)
	// 	for _, valor := range Algoritimo {
	// 		if !bExiste && valor == iPush {
	// 			bExiste = true
	// 		}
	// 	}

	// 	if !bExiste {
	// 		Algoritimo = append(Algoritimo, iPush)
	// 		iCountAlgoritimo++
	// 	}
	// }

	// MontaPrint := " var Algoritimo [200] int{"
	// for _, valor := range Algoritimo {
	// 	MontaPrint += strconv.Itoa(valor) + ","
	// }
	// MontaPrint += "}"
	// logs.Cyan("MontaPrint", MontaPrint)

	// logs.Cyan("len", iCountAlgoritimo)
	// logs.Cyan("len conf", len(Algoritimo))
	// logs.Cyan(Algoritimo)

	// 00000200006006001252000011015000000000040255400001
	token := Token.TokenST{}
	token.Decode("00000200006006001252000011015000000000040255400001")
	// token.Encode()

}
