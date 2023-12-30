package server

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// JSON sends json response
func JSON(w http.ResponseWriter, r *http.Request, status int, v interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(status)

	w.Write(buf.Bytes())
}

// ErrorJSON sends error as json
func JSONError(w http.ResponseWriter, r *http.Request, httpStatusCode int, err error, details string) {
	JSON(w, r, httpStatusCode, map[string]interface{}{"error": err.Error(), "details": details})
}

// NoContent sends no content response
func NoContent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
