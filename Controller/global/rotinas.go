package global

import (
	"encoding/json"
	"net/http"
)

func ResponseReturn(w http.ResponseWriter, Retorno interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	jsonstr, err := json.Marshal(Retorno)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.Write(jsonstr)
	return nil
}
