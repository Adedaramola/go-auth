package main

import (
	"encoding/json"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

}

func AttemptToAuthenticate(w http.ResponseWriter, r *http.Request) {

}

func Logout(w http.ResponseWriter, r *http.Request) {

}

func User(w http.ResponseWriter, r *http.Request) {

}

func Ping(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"message": "Server is alive",
	}

	responseJson(w, data, http.StatusOK)
}

func responseJson(w http.ResponseWriter, data any, status int) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)

	return nil
}
