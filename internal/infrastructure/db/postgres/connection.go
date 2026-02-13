package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/fkihai/payflow/internal/infrastructure/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type PostgresConnection struct {
	cfg *config.DatabaseConfig
}

func (conn *PostgresConnection) Connect() (*sql.DB, error) {

	pgUrl := fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s?sslmode=disable",
		conn.cfg.Primary.Driver,
		conn.cfg.Primary.User,
		conn.cfg.Primary.Password,
		conn.cfg.Primary.Host,
		conn.cfg.Primary.Port,
		conn.cfg.Primary.Name,
	)

	db, err := sql.Open("pgx", pgUrl)

	if err != nil {
		return nil, err
	}
	if err := goose.SetDialect("postgres"); err != nil {
		db.Close()
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func NewPostgresConnection(cfg *config.DatabaseConfig) *PostgresConnection {
	return &PostgresConnection{
		cfg: cfg,
	}
}
