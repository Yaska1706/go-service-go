package repository

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"go-mbv-go/internal"
	"os"
	"path/filepath"
)

type Storage interface {
	Migration
}

type storage struct {
	db        *sql.DB
	migration *migrate.Migrate
}

func NewStorage(db *sql.DB) Storage {
	basePath := internal.GetBasePath()
	migrationsDir := filepath.Join("file://", basePath, "/migrations")
	migrate, err := migrate.New(migrationsDir, os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("could not setup migration for Storage")
	}

	return &storage{db: db, migration: migrate}
}
