package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitzero"`
	Log    LogPayload  `json:"log,omitzero"`
	Mail   MailPayload `json:"mail,omitzero"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {

	payload := jsonResponse{
		Error:   false,
		Message: "Hit the Broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var reqPayload RequestPayload

	err := app.readJSON(w, r, &reqPayload)
	if err != nil {
		app.errJSON(w, err)
		return
	}

	switch reqPayload.Action {
	case "auth":
		app.authenticate(w, reqPayload.Auth)
	case "log":
		app.logData(w, reqPayload.Log)
	case "mail":
		app.sendMail(w, reqPayload.Mail)
	default:
		app.errJSON(w, errors.New("unknown action"))
	}
}

func (app *Config) logData(w http.ResponseWriter, l LogPayload) {
	jsonData, err := json.MarshalIndent(l, "", "\t")
	if err != nil {
		app.errJSON(w, err)
		return
	}

	req, err := http.NewRequest("POST", "http://logger-service/log", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errJSON(w, err)
		return
	}

	req.Header.Set("Context-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		app.errJSON(w, err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusAccepted {
		app.errJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "logged"

	app.writeJSON(w, http.StatusAccepted, payload)

}

func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
	// create some json
	jsonData, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		app.errJSON(w, err)
		return
	}

	// call the service
	req, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errJSON(w, err)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		app.errJSON(w, err)
		return
	}
	defer res.Body.Close()
	//get back currect status code
	if res.StatusCode == http.StatusBadRequest {
		app.errJSON(w, errors.New("invalied creds"))
		return
	} else if res.StatusCode != http.StatusAccepted {
		app.errJSON(w, errors.New("error calling auth service"))
		return
	}

	// read the response
	var jsonFromServce jsonResponse

	err = json.NewDecoder(res.Body).Decode(&jsonFromServce)
	if err != nil {
		app.errJSON(w, err)
		return
	}

	if jsonFromServce.Error {
		app.errJSON(w, errors.New(jsonFromServce.Message), http.StatusUnauthorized)
		return
	}

	var payLoad jsonResponse
	payLoad.Error = false
	payLoad.Message = "authenticated"
	payLoad.Data = jsonFromServce.Data

	app.writeJSON(w, http.StatusAccepted, payLoad)
}

func (app *Config) sendMail(w http.ResponseWriter, m MailPayload) {
	jsonData, _ := json.Marshal(m)

	// call the mail service
	mailServiceURL := "http://mailer-service/send"

	// post to mail service

	req, err := http.NewRequest("POST", mailServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errJSON(w, err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		app.errJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure to get back right status code
	if response.StatusCode != http.StatusAccepted {
		app.errJSON(w, errors.New("error calling main service"))
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "message sent to " + m.To

	app.writeJSON(w, http.StatusAccepted, payload)

}
