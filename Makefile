postgresinit:
	sudo docker run --name postgres_latest -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:latest

psql:
	sudo docker exec -it postgres_latest psql

createdb:
	sudo docker exec -it postgres_latest createdb --username=root --owner=root library

dropdb:
	sudo docker exec -it postgres_latest dropdb library

migrateup:
	migrate -path pkg/db/migrations -database "postgres://root:root@localhost:5432/library?sslmode=disable" -verbose up

migratedown:
	migrate -path pkg/db/migrations -database "postgres://root:root@localhost:5432/library?sslmode=disable" -verbose down

.PHONY: postgresinit psql createdb dropdb migrateup migratedown