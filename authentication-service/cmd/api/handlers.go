package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {

	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user from the database

	usr, err := app.Model.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.errJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := usr.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payLoad := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", usr.Email),
		Data:    usr,
	}

	app.writeJSON(w, http.StatusAccepted, payLoad)
}
