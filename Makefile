createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

postgres:
	docker run --name postgres12 --network smyik-network -p 5432:5432 -e POSTGRES_USER=root -e  POSTGRES_PASSWORD=secret -d postgres:12-alpine

migrateup:
	migrate -path  ./db/migration/ --database="postgres://root:soLLdaYy5pock4eFdQVa@smyik-db.cvumq0oimn62.ap-south-1.rds.amazonaws.com:5432/postgres?" --verbose up

migrateup1:
	migrate -path  ./db/migration/ --database="postgres://root:soLLdaYy5pock4eFdQVa@smyik-db.cvumq0oimn62.ap-south-1.rds.amazonaws.com:5432/postgres" --verbose up 1

migratedown:
	migrate -path  ./db/migration/ --database="postgres://root:soLLdaYy5pock4eFdQVa@smyik-db.cvumq0oimn62.ap-south-1.rds.amazonaws.com:5432/postgres" --verbose down

migratedown1:
	migrate -path  ./db/migration/ --database="postgres://root:soLLdaYy5pock4eFdQVa@smyik-db.cvumq0oimn62.ap-south-1.rds.amazonaws.com:5432/postgres" --verbose down 1

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

# you can simply run the application by this command
start_app:
	docker run --name smyik --network smyik-network -p  8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres12:5432/simple_bank?sslmode=disable" smyik:latest

start_prod:
	docker run --name smyik --network smyik-network -p  8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres12:5432/simple_bank?sslmode=disable" 1shubham7/smyik:v1.0.0

db_docs:
	dbdocs build doc/database.dbml --project Smyik

# to generate sql schema from  dbml
sql_schema:
	dbml2sql --postgres -o doc/schema.sql doc/database.dbml

.PHONY: createdb dropdb postgres migrateup migratedown test server mock i_love_you migrateup1 migratedown1 run_docker_image sqlc db_docs  sql_schema
