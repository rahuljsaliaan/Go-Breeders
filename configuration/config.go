package configuration

import (
	"database/sql"
	"rahuljsaliaan/go-breeders/models"
	"sync"
)

type Application struct {
	Models *models.Models
}

var instance *Application
var db *sql.DB
var once sync.Once

func New(pool *sql.DB) *Application {
	db = pool

	return GetInstance()
}

func GetInstance() *Application {
	once.Do(func() {
		instance = &Application{
			Models: models.New(db),
		}
	})

	return instance
}
