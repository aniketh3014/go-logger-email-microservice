package main

import (
	"logger/data"
	"net/http"
)

type JSONpayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestpayload JSONpayload

	_ = app.readJSON(w, r, &requestpayload)

	// insert data

	event := data.LogEntry{
		Name: requestpayload.Name,
		Data: requestpayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errJSON(w, err)
		return
	}

	res := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	app.writeJSON(w, http.StatusAccepted, res)
}
