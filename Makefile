createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e  POSTGRES_PASSWORD=secret -d postgres:12-alpine

migrateup:
	migrate -path  ./db/migration/ --database="postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose up

migrateup1:
	migrate -path  ./db/migration/ --database="postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose up 1

migratedown:
	migrate -path  ./db/migration/ --database="postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose down

migratedown1:
	migrate -path  ./db/migration/ --database="postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./... # this means that all the tests in project will run

server:
	go run main.go

mock:
	mockgen --destination db/mock/store.go --package mockdb  github.com/1shubham7/bank/db/sqlc Store

i_love_you:
	go fmt ./...

.PHONY: createdb dropdb postgres migrateup migratedown test server mock i_love_you migrateup1 migratedown1