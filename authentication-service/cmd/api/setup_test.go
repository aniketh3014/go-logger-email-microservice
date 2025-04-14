package main

import (
	"authentication/data"
	"os"
	"testing"
)

var TestApp Config

func TestMain(m *testing.M) {
	repo := data.NewPostgresTestRepository(nil)
	TestApp.Repo = repo
	os.Exit(m.Run())
}
