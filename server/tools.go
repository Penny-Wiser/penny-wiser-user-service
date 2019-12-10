package server

import (
	"encoding/json"
	"net/http"
)

func (s *server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(&v)
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, status int, payload interface{}) error {
	w.WriteHeader(status)
	if payload == nil {
		// Do sth
	}

	err := json.NewEncoder(w).Encode(payload)
	return err
}
