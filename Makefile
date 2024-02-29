postgresinit:
	sudo docker run --name library -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:latest

psql:
	sudo docker exec -it library psql

createdb:
	sudo docker exec -it library createdb --username=root --owner=root library

dropdb:
	sudo docker exec -it library dropdb library

migrateup:
	migrate -path internal/db/migrations -database "postgres://root:root@localhost:5432/library?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migrations -database "postgres://root:root@localhost:5432/library?sslmode=disable" -verbose down

.PHONY: postgresinit psql createdb dropdb migrateup migratedown