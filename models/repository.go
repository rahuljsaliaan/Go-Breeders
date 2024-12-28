package models

import "database/sql"

// Repository is the database repository. Anything that implements this interface must implement all the methods included here
type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
}

// mysqlRepository is a simple wrapper for the *sql.DB type. This is used to return a MySQL/MariaDB repository
type mysqlRepository struct {
	DB *sql.DB
}

// newMysqlRepository is a convenience factory method to return a mysqlRepository
func newMysqlRepository(conn *sql.DB) Repository {
	return &mysqlRepository{
		DB: conn,
	}
}

// testRepository is a simple wrapper for the *sql.DB type. This is used to return a test repository
type testRepository struct {
	DB *sql.DB
}

// newTestRepository is a convenience factory method to return a testRepository
func newTestRepository() Repository {
	return &testRepository{
		DB: nil,
	}
}
