Mail GRPC Server:
=================

This server is responsible for managing mails

Requirements:
-------------

[docker]()

Installation:
-------------
Build all containers:
```bash
docker-compose build
```

Usage:
------

Locally:
```bash
SERVER_SOURCE=127.0.0.1:50051 go run main.go
```

Docker:
```bash
docker-compose up
```

Aritcles:

https://mycodesmells.com/post/authentication-in-grpc
https://bbengfort.github.io/programmer/2017/03/03/secure-grpc.html
