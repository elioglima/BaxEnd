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

	ArrayByteIn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Retorno.Erro = true
		Retorno.Msg = errors.New("Erro ao receber body. \n " + err.Error()).Error()
		Retorno.Dados = nil
		logs.Erro(Retorno.Msg)
		responseReturn(w, Retorno)
		return
	}

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

	type Cdata struct {
		Nome string
	}
	// fmt.Printf("%s ok", r.Body)

	DataTipo := Cdata{}

	ArrayByteIn, err := ioutil.ReadAll(r.Body)
	if err == nil && ArrayByteIn != nil {
		err = json.Unmarshal(ArrayByteIn, &DataTipo)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}

	fmt.Printf("%s", DataTipo.Nome)

}
