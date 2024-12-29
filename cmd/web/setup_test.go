package main

import (
	"os"
	"rahuljsaliaan/go-breeders/configuration"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	testApp = application{
		App: configuration.New(nil),
	}

	os.Exit(m.Run())
}
