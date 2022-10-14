package datastore

import (
	"errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	*sqlx.DB
}

func NewConnection(dbSource string, automigrate bool) (*DB, error) {
	db, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		return nil, err
	}

	log.Println("Database connection established")

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if automigrate {
		err := runMigrations("file://datastore/migrations", dbSource)
		if err != nil {
			return nil, err
		}
	}

	defer db.Close()

	return &DB{db}, nil
}

func runMigrations(migrationsPath, dbSource string) error {
	migration, err := migrate.New(migrationsPath, dbSource)
	if err != nil {
		return err
	}

	err = migration.Up()

	switch {
	case errors.Is(err, migrate.ErrNoChange):
		log.Println("Nothing to migrate")
		break
	case err != nil:
		return err
	}

	return nil
}
