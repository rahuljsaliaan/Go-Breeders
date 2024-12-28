package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"rahuljsaliaan/go-breeders/models"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	DB          *sql.DB
	Models      models.Models
}

type appConfig struct {
	useCache   bool
	dsn        string
	port       int
	dbUser     string
	dbName     string
	dbPassword string
	dbPort     int
}

func (a *application) loadEnv() {
	// Load environment variables from .env file (optional, for development)
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Load configuration from environment variables
	a.config = appConfig{
		useCache:   getEnvAsBool("USE_CACHE", false),
		port:       getEnvAsInt("PORT", 8080),
		dbUser:     getEnv("MYSQL_USER", "mariadb"),
		dbName:     getEnv("MYSQL_DATABASE", "testdb"),
		dbPassword: getEnv("MYSQL_PASSWORD", "password"),
		dbPort:     getEnvAsInt("DB_PORT", 3306),
	}
}

// Helper functions to retrieve environment variables
func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func main() {
	app := &application{
		templateMap: make(map[string]*template.Template),
	}
	app.loadEnv()

	// DESC: This is used to parse the command-line flags
	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache")
	flag.StringVar(&app.config.dsn, "dsn", fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", app.config.dbUser, app.config.dbPassword, app.config.dbPort, app.config.dbName), "DSN")
	flag.Parse()

	// get database
	db, err := initMySQLDB(app.config.dsn)
	if err != nil {
		log.Panic(err)
	}
	app.DB = db

	app.Models = *models.New(db)

	// Start the web server
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}
	fmt.Println("Starting web application on port", app.config.port)
	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
