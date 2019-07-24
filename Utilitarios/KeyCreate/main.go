package main

import (
	"BaxEnd/Controller/database/dbmysql/Empresas"
	"bytes"
	"encoding/gob"
	"fmt"
)

type Student struct {
	Name string
	Age  int32
}

func main() {

	fmt.Println("Gob Example")
	Dados := &Empresas.EmpresaDadosST{}
	Dados.Nome = "Elio"

	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	if err := e.Encode(Dados); err != nil {
		panic(err)
	}

	sEnc := b64.StdEncoding.EncodeToString(b.Bytes())
	fmt.Println(sEnc)

	Dados2 := &Empresas.EmpresaDadosST{}
	d := gob.NewDecoder(&b)
	if err := d.Decode(&Dados2); err != nil {
		panic(err)
	}

	fmt.Println("Decoded Struct ", Dados2.Nome)

}
