package common

import (
	"encoding/json"
	"log"
	"net/http"
)

// Error data type
type Error struct {
	Code int    `json:"errorCode" example:"35"`
	Msg  string `json:"message" example:"Invalid input"`
}

// MarshallAndSend errors over http channel.
func (e *Error) MarshallAndSend(w http.ResponseWriter) {
	log.Println(e.Msg)
	w.WriteHeader(e.Code)
	if e.Code == http.StatusInternalServerError {
		e.Msg = "Internal server error"
	}
	json.NewEncoder(w).Encode(e)
}
