package http

import (
	"encoding/json"
	"net/http"
)

func WriteError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

// write json to http respose, if error then write internal server error
func WriteJson(w http.ResponseWriter, v any) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(v)
	if err != nil {
		WriteError(w, http.StatusInternalServerError)
	}
	return err
}
