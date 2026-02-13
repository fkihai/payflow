package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

type PostgresMigrations struct {
	db *sql.DB
}

func (m *PostgresMigrations) MigrateUp() error {
	ctx := context.Background()
	if err := goose.UpContext(ctx, m.db, "migrations"); err != nil {
		fmt.Printf("migration upgrade error: %v\n", err)
	}
	return nil
}

func (m *PostgresMigrations) MigrateDown() error {
	ctx := context.Background()
	if err := goose.DownContext(ctx, m.db, "migrations"); err != nil {
		fmt.Printf("migration downgrade error: %v\n", err)
	}
	return nil
}

func NewPostgresMigrations(db *sql.DB) *PostgresMigrations {
	return &PostgresMigrations{
		db: db,
	}
}
