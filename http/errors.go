package http

import (
	"encoding/json"
	"net/http"
)

func WriteInternalErr(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// encode
func WriteJson(w http.ResponseWriter, v any) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(v)
	WriteInternalErr(w)
	return err
}
