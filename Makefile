DB_URL=postgresql://root:password@localhost:5432/bank?sslmode=disable

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root bank

dropdb:
	docker exec -it postgres12 dropdb bank

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

migrateup:
	migrate -path db/migration -database "${DB_URL}" -verbose up

migratedown:
	migrate -path db/migration -database "${DB_URL}" -verbose down

migrateupby1:
	migrate -path db/migration -database "${DB_URL}" -verbose up 1

migratedownby1:
	migrate -path db/migration -database "${DB_URL}" -verbose down 1

sqlc:
	sqlc generate

.PHONY: postgres createdb dropd new_migration migrateup migratedown migrateupby1 migratedownby1 sqlc