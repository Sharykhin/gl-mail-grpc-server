BINARY_NAME=gl-mail-grps-server

serve:
	DB_SOURCE="root:root@tcp(localhost:3306)/gl_mail_api?parseTime=true" SERVER_SOURCE=127.0.0.1:50051 go run main.go

docker-serve:
	DB_SOURCE=test:test@tcp(gl-mail-grpc-server-mysql:3306)/test?parseTime=true SERVER_SOURCE=:50051 go run main.go

prod: build
	DB_SOURCE="root:root@tcp(localhost:3306)/gl_mail_api?parseTime=true" SERVER_SOURCE=127.0.0.1:50051 $(BINARY_NAME)

lint:
	gometalinter ./...

build:
	go build -o $(BINARY_NAME) -v

clean:
	go clean
	rm -rf $(BINARY_NAME)