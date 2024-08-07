include .env

.SILENT:



run:
	go run cmd/main.go

migrate-up:
	sqlite3 $(DATABASE_URL) < migrations/up.sql

migrate-down:
	sqlite3 $(DATABASE_URL) < migrations/down.sql

migrate-mock:
	sqlite3 $(DATABASE_URL) < migrations/mock_information.sql
