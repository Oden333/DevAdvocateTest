postgres:
	docker run --name postgresql --network rest-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=qwerty -d postgres:latest

createdb:
	docker exec -it postgresql createdb --username=root --owner=root users

dropdb:
	docker exec -it postgresql dropdb users

migrateup:
	migrate -path db/migration -database "postgresql://root:qwerty@localhost:5437/users?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:qwerty@localhost:5437/users?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown