postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:12-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank_db

startdb:
	docker postgres start

migrateup:
	migrate -path db/migration -database "postgres://root:root@localhost:5432/simple_bank_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:root@localhost:5432/simple_bank_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

mockdb:
	mockgen -package mockdb -destination db/mock/store.go github.com/eldersoon/simple-bank/db/sqlc Store

# this comand run tests without cache
test-nocache:
	go test -count=1 -v -cover ./...

server:
	go run main.go

air:
	air

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test test-nocache server air mockdb