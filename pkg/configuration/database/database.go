package database

import (
	"database/sql"
	"fmt"

	"github.com/barchart/common-go/pkg/validation"
)

var validate = validation.GetValidator()

// Database is a type of database configuration
type Database struct {
	Provider string `validate:"required" json:"provider"`
	Host     string `validate:"required" json:"host"`
	Port     int    `validate:"required" json:"port"`
	Database string `validate:"required" json:"database"`
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

// GetConnectionString returns connection string in format: user=? password=? host=? port=? dbname=?
func (database Database) GetConnectionString() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s", database.Username, database.Password, database.Host, database.Port, database.Database)
}

// GetConnectionURL returns connection string in format: provider://user:password@host:port/database
func (database Database) GetConnectionURL() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s", database.Provider, database.Username, database.Password, database.Host, database.Port, database.Database)
}

// OpenDB return *sql.DB instance
func (database Database) OpenDB() (*sql.DB, error) {
	err := validate.Struct(database)

	if err != nil {
		return nil, err
	}

	db, err := sql.Open(database.Provider, database.GetConnectionString())

	if err != nil {
		return nil, err
	}

	return db, nil
}
