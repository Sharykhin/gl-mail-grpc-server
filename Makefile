.PHONY: test build docker-serve serve lint clean generate-keys

BINARY_NAME=gl-mail-grps-server
GOARCH=amd64
CGO_ENABLED=0
GOOS=linux

test:
	go test ./...

build:
	go build -o $(BINARY_NAME) .

local-serve:
	DB_SOURCE="root:root@tcp(localhost:3306)/gl_mail_api?parseTime=true" KEY_SERVER_CRT=server.crt KEY_SERVER_KEY=server.key SERVER_SOURCE=localhost:50051 APP_ENV=dev go run main.go

serve:
	docker-compose exec gl-mail-grpc-server-golang go run *.go

lint:
	gometalinter ./...

clean:
	go clean
	rm -rf $(BINARY_NAME)

generate-keys:
	openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes -subj '/CN=localhost'