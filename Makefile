up:
	docker compose up 

down:
	docker compose down

build:
	docker compose build

restart: down up

migrate-auth:
	migrate -path services/auth/migrations -database "postgres://postgres:qwerty@localhost:5433/auth_db?sslmode=disable" up


migrate-note:
	migrate -path services/note/migrations -database "postgres://postgres:qwerty@localhost:5434/note_db?sslmode=disable" up

migrate-user:
	migrate -path services/user/migrations -database "postgres://postgres:qwerty@localhost:5435/user_db?sslmode=disable" up

migrate: migrate-auth migrate-note migrate-user 