package database

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

// RunMigrations applies database migrations using SQL files.
func RunMigrations(dsn, migrationsPath string) error {
	p := &pgx.Postgres{}
	d, err := p.Open(dsn)

	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	defer func() {
		if err := d.Close(); err != nil {
			log.Printf("Failed to close database connection: %v", err)
		}
	}()

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "pgx", d)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	return nil
}
