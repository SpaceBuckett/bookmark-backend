postgres:
	docker run --name postgres17 -p 5432:5432 postgres_user=root -e postgres_password=secret -d postgress:17-alpine

dbversion:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/bookmarker?sslmode=disable" version

forcemigrate:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bookmarker?sslmode=disable" force 1

closeallconnections:
	docker exec -it postgres17 psql -U root -c "SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = 'bookmaker' AND pid <> pg_backend_pid();"

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root bookmarker

dropdb:
	docker exec -it postgres17 dropdb --username=root bookmarker

createmigration:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bookmarker?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bookmarker?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

start:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown closeallconnections forcemigrate dbversion sqlc test start createmigration