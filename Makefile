BINARY_NAME=gl-mail-grps-server

serve:
	DB_SOURCE="root:root@tcp(localhost:3306)/gl_mail_api?parseTime=true" KEY_SERVER_CRT=server.crt KEY_SERVER_KEY=server.key SERVER_SOURCE=localhost:50051 go run main.go

docker-serve:
	DB_SOURCE="test:test@tcp(gl-mail-grpc-server-mysql:3306)/test?parseTime=true" KEY_SERVER_CRT=server.crt KEY_SERVER_KEY=server.key SERVER_SOURCE=:50051 go run main.go

prod: build
	DB_SOURCE="root:root@tcp(localhost:3306)/gl_mail_api?parseTime=true" KEY_SERVER_CRT=server.crt KEY_SERVER_KEY=server.key SERVER_SOURCE=127.0.0.1:50051 $(BINARY_NAME)

lint:
	gometalinter ./...

build:
	go build -o $(BINARY_NAME) -v

clean:
	go clean
	rm -rf $(BINARY_NAME)

generate-keys:
	openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes -subj '/CN=127.0.0.1'