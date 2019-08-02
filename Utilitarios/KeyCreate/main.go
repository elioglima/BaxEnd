package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type AulaEncode struct {
	Nome string
}

func main() {

	Dados := &AulaEncode{}
	Dados.Nome = "Elio"

	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	if err := e.Encode(Dados); err != nil {
		panic(err)
	}
	fmt.Println("Struct []byte ", b)

	arr := b.Bytes()
	var nbyte []byte

	for index := len(arr) - 1; index != -1; index-- {
		nbyte = append(nbyte, arr[index])
	}

	str := base64.StdEncoding.EncodeToString(nbyte)

	fmt.Println(nbyte)

	b.Reset()
	b.Write([]byte(str))
	err := ioutil.WriteFile("output.txt", b.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

	Dados2 := &AulaEncode{}
	d := gob.NewDecoder(&b)
	if err := d.Decode(&Dados2); err != nil {
		panic(err)
	}

	fmt.Println("Decoded Struct ", Dados2.Nome)

}
