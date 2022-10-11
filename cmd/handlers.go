package main

import (
	"net/http"

	"github.com/adedaramola/golang-auth/utils"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

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

	utils.ResponseJson(w, data, http.StatusOK)
}
