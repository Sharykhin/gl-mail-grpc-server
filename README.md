Mail GRPC Server:
=================

This server is responsible for managing mails

Requirements:
-------------

[docker]()

Generating keys:
-----------------


Generating files locally:
```bash
openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

Put you host into *Common Name*: 127.0.0.1
```bash
openssl req -new -sha256 -key server.key -out server.csr
openssl x509 -req -sha256 -in server.csr -signkey server.key -out server.crt -days 3650
```

Put you host into *Common Name*: 127.0.0.1

Usage:
------
##### Docker:
Build all containers:
```bash
docker-compose build
```

Run containers:
```bash
docker-compose up
```


##### Locally:
```bash
DB_SOURCE="root:root@tcp(localhost:3306)/gl_mail_api?parseTime=true" SERVER_SOURCE=127.0.0.1:50051 go run main.go
```

or use Makefile:
```bash
make serve
```

TODO:
-----

Current mysql is added for some test, real mysql should be used from
gl-mail-api service.

Aritcles:

https://mycodesmells.com/post/authentication-in-grpc
https://bbengfort.github.io/programmer/2017/03/03/secure-grpc.html
