build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/api/main.go

run: build
	docker-compose up -d --build server

migrate:
	migrate -path ./db/migrations -database 'postgres://postgres:postgres@0.0.0.0:5432/postgres?sslmode=disable' up