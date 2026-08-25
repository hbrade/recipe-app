package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB(databaseURL string) error {
	var err error
	DB, err = pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return fmt.Errorf("Unable to connect to database: %w", err)
	}

	// Test connection
	err = DB.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to ping database: %w", err)
	}

	fmt.Println("✓ Database connected successfully")
	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
