.PHONY: test build pack upload

BINARY_NAME=gl-mail-grps-server
GOARCH=amd64
CGO_ENABLED=0
GOOS=linux

test:
	go test ./...

build:
	go build -ldflags "-X main.version=$(VER)" -o $(BINARY_NAME) .

pack:
	GOOS=linux make build
	docker-compose build gl-mail-grpc-server-golang

upload:
	docker push chapal/gl-mail-grpc-server-golang:$(VER)

serve:
	DB_SOURCE="root:root@tcp(localhost:3306)/gl_mail_api?parseTime=true" KEY_SERVER_CRT=server.crt KEY_SERVER_KEY=server.key SERVER_SOURCE=localhost:50051 go run main.go

docker-serve:
	DB_SOURCE="test:test@tcp(gl-mail-grpc-server-mysql:3306)/test?parseTime=true" KEY_SERVER_CRT=server.crt KEY_SERVER_KEY=server.key SERVER_SOURCE=:50051 go run main.go

prod: build
	DB_SOURCE="root:root@tcp(localhost:3306)/gl_mail_api?parseTime=true" KEY_SERVER_CRT=server.crt KEY_SERVER_KEY=server.key SERVER_SOURCE=127.0.0.1:50051 $(BINARY_NAME)

lint:
	gometalinter ./...

clean:
	go clean
	rm -rf $(BINARY_NAME)

generate-keys:
	openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes -subj '/CN=localhost'

docker-generate-keys:
	openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes -subj '/CN=gl-mail-grpc-server-golang'