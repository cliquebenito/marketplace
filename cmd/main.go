package main

import (
	"context"
	"log"
	"project/pkg/postgresql"
)

const ()

func main() {
	db, err := postgresql.NewClient(context.Background(), postgresql.DsnConfig{
		Host:        "localhost",
		Port:        "5432",
		User:        "postgres",
		Password:    "postgres",
		Database:    "postgres",
		MaxAttempts: 5,
	})
	if err != nil {
		log.Fatalf("failed to initialize db")
	}

}
