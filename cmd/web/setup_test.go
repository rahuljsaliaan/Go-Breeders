package main

import (
	"fmt"
	"log"
	"os"
	"rahuljsaliaan/go-breeders/models"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	testApp.loadEnv()

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", testApp.config.dbUser, testApp.config.dbPassword, testApp.config.dbPort, testApp.config.dbName)

	db, err := initMySQLDB(dsn)

	if err != nil {
		log.Panic(err)
	}

	testApp = application{
		DB:     db,
		Models: *models.New(db),
	}

	os.Exit(m.Run())
}
