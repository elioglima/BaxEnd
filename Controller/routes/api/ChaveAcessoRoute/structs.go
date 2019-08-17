package ChaveAcessoRoute

import (
	"encoding/json"
	"net/http"
)

type sRetorno struct {
	Erro  bool
	Msg   string
	Dados interface{}
}

func (s *sRetorno) Ini() {
	s.Erro = false
	s.Msg = "Aguardando.."
	s.Dados = nil
}

func responseReturn(w http.ResponseWriter, Retorno sRetorno) error {
	w.Header().Set("Content-Type", "application/json")

	jsonstr, err := json.Marshal(Retorno)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.Write(jsonstr)
	return nil
}
