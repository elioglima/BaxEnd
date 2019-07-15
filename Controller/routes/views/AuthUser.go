package views

import (
	"BaxEnd/Controller/global"
	libs "GoLibs"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func AuthUser(w http.ResponseWriter, r *http.Request) (*stEstrutura, error) {

	type stspwdr struct {
		Hash string `json:"hash"`
	}

	t := time.Now()
	path := global.DirPrivate()
	Estrutura := &stEstrutura{}
	Estrutura.Data = fmt.Sprintf("%02d/%02d/%d %02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
	Estrutura.Versao = global.Versao

	data, err := ioutil.ReadFile(path + "Docs/menulateral.html")
	if err == nil {
		Estrutura.MenuLateral = template.HTML(string(data))
	}

	data, err = ioutil.ReadFile(path + "Docs/rodape.html")
	if err == nil {
		Estrutura.Rodape = template.HTML(string(data))
	}

	inhash := r.FormValue("hash")
	usuario := r.FormValue("nmacs")
	senha := r.FormValue("pwacs")

	spwdr := &stspwdr{}

	b, err := getHash()
	if err != nil {
		Estrutura.CodeStatus = 500
		Estrutura.Mensagem = template.HTML(err.Error())
		ShowCodeStatus(w, r, Estrutura)
		return Estrutura, err
	}

	err = json.Unmarshal(b, spwdr)
	if err != nil {
		Estrutura.CodeStatus = 500
		Estrutura.Mensagem = template.HTML(err.Error())
		ShowCodeStatus(w, r, Estrutura)
		return Estrutura, err
	}

	if len(strings.TrimSpace(inhash)) > 0 {

		if inhash != spwdr.Hash {
			UserNotAuth(w, r)
			return Estrutura, errors.New("Autenticação de usuario inválida")
		}

		Estrutura.Hash = inhash
		return Estrutura, nil
	}

	hash, err := libs.HashEncode(usuario + senha)
	if err != nil {
		Estrutura.CodeStatus = 500
		Estrutura.Mensagem = template.HTML(err.Error())
		ShowCodeStatus(w, r, Estrutura)
		return Estrutura, err
	} else if hash != spwdr.Hash {
		UserNotAuth(w, r)
		return Estrutura, errors.New("Autenticação de usuario inválida")
	}

	Estrutura.Hash = hash
	return Estrutura, nil
}
