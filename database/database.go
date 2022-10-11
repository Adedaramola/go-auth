package database

import (
	"embed"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
)

// go:embed "database/migrations/*.sql"
var EmbeddedMigrations embed.FS

type DB struct {
	db *sqlx.DB
}

func NewConnection(dsn string, automigrate bool) (*DB, error) {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if automigrate {
		fs, err := iofs.New(EmbeddedMigrations, "database/migrations")
		if err != nil {
			return nil, err
		}

		migrator, err := migrate.NewWithSourceInstance("iofs", fs, dsn)
		if err != nil {
			return nil, err
		}

		err = migrator.Up()
		switch {
		case errors.Is(err, migrate.ErrNoChange):
			fmt.Println("Nothing to migrate")
			break
		case err != nil:
			return nil, err
		}
	}

	defer db.Close()

	return &DB{db}, nil
}
