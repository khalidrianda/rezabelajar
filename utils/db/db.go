package db

import (
	"fmt"
	"time"

	"belajar/config.go"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

// untuk open connection ke db
func NewDBConnection(conf *config.AppConfig) *sqlx.DB {

	db, err := sqlx.Connect("postgres", fmt.Sprintf("user= %s host=%s port=%d password=%s dbname=%s sslmode=%s",
		conf.DBUser,
		conf.DBHost,
		conf.DBPort,
		conf.DBPass,
		conf.DBName,
		conf.SSLMode,
	))
	if err != nil {
		log.Fatal("Failed to connect database " + conf.DBName + ", err: " + err.Error())
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(30 * time.Minute)
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to connect database " + conf.DBName + ", err: " + err.Error())
	}
	log.Info("Connection Opened to Database " + conf.DBName)

	return db
}
