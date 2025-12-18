package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/FelipePn10/panossoerp/internal/infrastructure/config"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database/sqlc"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DB struct {
	SQL *sql.DB
}

func NewDB(cfg *config.Config) (*DB, error) {
	dbSQL, err := sql.Open("pgx", cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := dbSQL.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping DB: %w", err)
	}

	log.Println("Connected to the database successfully")

	return &DB{SQL: dbSQL}, nil
}

func (db *DB) Queries() *sqlc.Queries {
	return sqlc.New(db.SQL)
}

func (db *DB) Close() {
	if db.SQL != nil {
		_ = db.SQL.Close()
	}
}
