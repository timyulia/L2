package handler

import (
	"calendar"
	"encoding/json"
	"net/http"
)

type responseOK struct {
	Message string           `json:"message"`
	Events  []calendar.Event `json:"events"`
}

type responseError struct {
	Error string `json:"error"`
}

func throwError(w http.ResponseWriter, status int, err error) {
	resp := responseError{
		Error: err.Error(),
	}

	out, _ := json.MarshalIndent(resp, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)
}

func writeResponse(w http.ResponseWriter, status int, msg string, events []calendar.Event) {
	resp := responseOK{
		Message: msg,
		Events:  events,
	}

	out, _ := json.MarshalIndent(resp, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)
}
