package configs

import (
	"fmt"
	"os"
)

// DBConfig is the config for the database connection
type DBConfig struct {
	Proto    string
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

// Setup the DatabaseConfig
func (db *DBConfig) Setup() {
	db.Proto = os.Getenv("DB_PROTO")
	db.Host = os.Getenv("DB_HOST")
	db.Port = os.Getenv("DB_PORT")
	db.User = os.Getenv("DB_USER")
	db.Pass = os.Getenv("DB_PASS")
	db.Database = os.Getenv("DB_DATABASE")
}

// DSN creates a dsn string to connect to the DB
func (db *DBConfig) DSN() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s",
		db.Proto,
		db.User,
		db.Pass,
		db.Host,
		db.Port,
		db.Database,
	)
}
