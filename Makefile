run-service:
	go build -o ./client/cli ./client/main.go &&\
	docker compose up -d --build

stop-service:
	rm ./client/cli &&\
	docker compose down

grpc-gui:
	grpcui -plaintext localhost:9000

before-push:
	go mod tidy &&\
	gofumpt -l -w . &&\
	go build ./...&&\
	golangci-lint run ./... &&\
	go test -v ./tests/...

protoc-books:
	protoc -I ./api --go_out=./internal/modules --go-grpc_out=./internal/modules ./api/books.proto
