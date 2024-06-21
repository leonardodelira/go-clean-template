package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func postgresURL() string {
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	dbname := os.Getenv("PG_DBNAME")
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)

	return url
}

func config() *pgxpool.Config {
	databaseURL := postgresURL()
	dbConfig, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}
	return dbConfig
}

func NewPostgresConnection() *pgxpool.Pool {
	ctx := context.Background()
	connPoll, err := pgxpool.NewWithConfig(ctx, config())
	if err != nil {
		log.Fatal(fmt.Printf("error on connect to postgres: %v", err))
	}

	err = connPoll.Ping(ctx)
	if err != nil {
		log.Fatal(fmt.Printf("could not ping the database: %v", err))
	}

	fmt.Println("Connected to the database!!")

	return connPoll
}
