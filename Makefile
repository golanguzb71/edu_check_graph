include .env

.SILENT:


run:
	go run main.go

migrate-up:
	PGPASSWORD=$(POSTGRES_PASSWORD) psql -U $(POSTGRES_USER) -d $(POSTGRES_DATABASE) -h $(POSTGRES_HOST) -p $(POSTGRES_PORT) -f migrations/up.sql

migrate-down:
	PGPASSWORD=$(POSTGRES_PASSWORD) psql -U $(POSTGRES_USER) -d $(POSTGRES_DATABASE) -h $(POSTGRES_HOST) -p $(POSTGRES_PORT) -f migrations/down.sql
mock:
	PGPASSWORD=$(POSTGRES_PASSWORD) psql -U $(POSTGRES_USER) -d $(POSTGRES_DATABASE) -h $(POSTGRES_HOST) -p $(POSTGRES_PORT) -f migrations/moc_information.sql
