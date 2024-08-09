package database

import (
	"context"
	"github.com/joho/godotenv"
	"testing"
)

func TestConnection(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		return
	}
	err = RDB.Ping(context.TODO())
	if err != nil {
		t.Fatalf("Failed to connect to Redis: %v", err)
	}
}
