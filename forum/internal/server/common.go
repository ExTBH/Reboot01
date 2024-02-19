package server

import (
	"encoding/json"
	"net/http"
)

// Returns false if an error happens
func parseBody(data any, w http.ResponseWriter, r *http.Request) bool {
	return json.NewDecoder(r.Body).Decode(&data) == nil
}

// converts `v`	to json and writes it to the response writer
func writeToJson(v any, w http.ResponseWriter) error {
	buff, err := json.Marshal(v)
	if err != nil {
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(buff)
	return err
}

func IsAcceptedMethod(w http.ResponseWriter, r *http.Request, allowedMethod string) bool {
	if r.Method != allowedMethod {
		errorServer(w, r, http.StatusMethodNotAllowed, true)
		return false
	}
	return true
}
