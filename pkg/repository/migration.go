package repository

import (
	"errors"
	"os"
)

type Migration interface {
	ApplyMigrations() error
	MigrateDown(steps int) error
}

// ApplyMigrations should be used for testing, staging and production
func (s *storage) ApplyMigrations() error {
	environment := os.Getenv("ENVIRONMENT")

	// only run migrations in runtime in staging and production
	if environment == "DEVELOPMENT" {
		return nil
	}

	err := s.migration.Up()

	// TODO: not sure I like this, look into later
	if !errors.Is(err, errors.New("no change")) {
		return err
	}

	return nil
}

// MigrateDown should be used for testing, staging and production
// TODO: Not sure if this is actually needed
func (s *storage) MigrateDown(steps int) error {
	environment := os.Getenv("ENVIRONMENT")

	// only run migrations in runtime in staging and production
	if environment == "DEVELOPMENT" {
		return nil
	}

	err := s.migration.Steps(steps)

	if err != nil {
		return err
	}
	return nil
}
