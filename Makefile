postgres:
	docker run --name postgres16 -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:16-alpine

create_db:
	docker exec -it postgres16 createdb --username=postgres --owner=postgres simple_bank

drop_db:
	docker exec -it postgres16 dropdb simple_bank

migrations_up:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrations_down:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

.PHONY: create_db drop_db postgres migrations_up migrations_down