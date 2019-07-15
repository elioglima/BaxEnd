package usuario

import (
	"encoding/json"
	"net/http"
)

type sRetorno struct {
	Erro  error
	Msg   string
	Dados interface{}
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
