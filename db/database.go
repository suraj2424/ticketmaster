package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func ConnectDB() {
	connString := os.Getenv("DATABASE_URL")
	var err error

	Pool, err = pgxpool.New(context.Background(), connString)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully connected to NeonDB!")
}
