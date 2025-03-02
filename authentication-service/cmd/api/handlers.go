package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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

	// log to logger service
	err = app.logRequest("authentication", fmt.Sprintf("%s authenticated", usr.Email))
	if err != nil {
		app.errJSON(w, err)
		return
	}

	payLoad := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", usr.Email),
		Data:    usr,
	}

	app.writeJSON(w, http.StatusAccepted, payLoad)
}

func (app *Config) logRequest(name, data string) error {
	var entry struct {
		Name string `json:"name"`
		Data string `josn:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, err := json.MarshalIndent(entry, "", "\t")
	if err != nil {
		log.Println(err)
		return err
	}

	req, err := http.NewRequest("POST", "http://logger-service/log", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
