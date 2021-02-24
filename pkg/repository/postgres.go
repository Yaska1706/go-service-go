package repository

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"go-mbv-go/internal"
	"path/filepath"
)

type Storage interface {
	RunMigrations(connectionString string) error
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{db: db}
}

func (s storage) RunMigrations(connectionString string) error {
	if connectionString == "" {
		return errors.New("repository: the connectionString was empty")
	}

	basePath := internal.GetBasePath()

	migrationsPath := filepath.Join("file://", basePath, "/migrations")

	m, err := migrate.New(migrationsPath, connectionString)

	if err != nil {
		return err
	}

	err = m.Up()

	if err != nil {
		if err.Error() == "no change" {
			return nil
		}
		return err
	}

	return nil
}
