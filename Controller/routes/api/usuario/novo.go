package usuario

import (
	"BaxEnd/Controller/database"
	"GoLibs/logs"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Novo(w http.ResponseWriter, r *http.Request) {

	logs.Branco("usuario/novo")

	Retorno := sRetorno{}
	Retorno.Ini()

	type Cdata struct {
		Nome string
	}

	DataTipo := Cdata{}

	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		err = json.Unmarshal(data, &DataTipo)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}

	fmt.Printf("%s", DataTipo.Nome)
	return

	ArrayByteIn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Retorno.Erro = true
		Retorno.Msg = errors.New("Erro ao receber body. \n " + err.Error()).Error()
		Retorno.Dados = nil
		logs.Erro(Retorno.Msg)
		responseReturn(w, Retorno)
		return
	}

	logs.Erro(ArrayByteIn)
	err = json.Unmarshal(ArrayByteIn, &DataTipo)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// } else if len(ArrayByteIn) == 0 {
	// 	Retorno.Erro = true
	// 	Retorno.Msg = "Nenhum paramêtros informado."
	// 	Retorno.Dados = nil
	// 	logs.Erro(Retorno.Msg)
	// 	responseReturn(w, Retorno)
	// 	return
	// }

	Retorno.Erro = true
	Retorno.Msg = "Nenhum paramêtros informado."
	Retorno.Dados = nil
	logs.Erro(Retorno.Msg)
	responseReturn(w, Retorno)
	return

	if err := database.MySql.Conectar(); err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	msg, err := database.MySql.Usuario.Novo(ArrayByteIn)
	if err != nil {
		Retorno.Erro = true
		Retorno.Msg = err.Error()
		Retorno.Dados = nil
		responseReturn(w, Retorno)
		return
	}

	Retorno.Msg = msg
	Retorno.Dados = database.MySql.Usuario.Field
	responseReturn(w, Retorno)
}

func Test(w http.ResponseWriter, r *http.Request) {
	var body map[string]interface{}
	type Cdata struct {
		Nome string
	}

	DataTipo := Cdata{}

	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		err = json.Unmarshal(data, &DataTipo)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}

	fmt.Printf("%s", DataTipo.Nome)
	return

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Fprintf(w, "body: %s\n", body)
}
