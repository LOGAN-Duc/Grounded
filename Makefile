start:
	go run cmd/dev/main.go
build:
	go build -o app cmd/dev/main.go
migration:
	./scripts/make-migration.sh
migrate:
	sql-migrate up -config=configs/dbconfig.yml
migrate-rollback:
	sql-migrate down -config=configs/dbconfig.yml
migrate-rollback-all:
	sql-migrate down -config=configs/dbconfig.yml -limit=0