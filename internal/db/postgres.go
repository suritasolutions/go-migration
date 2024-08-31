package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func NewPostgresDB(ctx context.Context) Database {
	return &postgresDB{
		ctx: ctx,
		dns: fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			ctx.Value("database").(string),
		),
	}
}

type postgresDB struct {
	ctx context.Context
	dns string
}

func (d *postgresDB) ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", d.dns)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (d *postgresDB) Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (d *postgresDB) GetCreateMigrationTableSQL() string {
	return `
		CREATE TABLE IF NOT EXISTS migrations (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
}

func (d *postgresDB) GetDatabaseExistsSQL() string {
	return `
		SELECT datname
		FROM pg_catalog.pg_database
		WHERE lower(datname) = lower('` + d.ctx.Value("database").(string) + `');
	`
}

func (d *postgresDB) GetCreateDatabaseSQL() string {
	return `
		CREATE DATABASE ` + d.ctx.Value("database").(string) + ` WITH ENCODING 'UTF8';
	`
}
