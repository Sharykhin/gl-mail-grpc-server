Mail GRPC Server:
=================

This server is responsible for managing mails. Currently is saves failed mails 
into a storage. Service provides grpc api for creating and getting list of failed
mails

Requirements:
-------------

[docker](https://www.docker.com)

Attention:
----------
By default app environment is set to **dev** and communication
with grpc server goes with insecure way. To switch to secure
change *APP_ENV* to **prod**

Generating keys:
-----------------
In case you need to generate secure keys locally:
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
For docker use the following name: gl-mail-grpc-server-golang

Or just use Makefile command:
```bash
make generate-keys
```

Usage:
------
##### For Development:
1. Build all containers:
```bash
docker-compose build
```

2. Run containers:
```bash
docker-compose up
```

3. Run migrations:
```bash
make migrate
```

4. Run the server:
```bash
make serve
```

Makefile:
```bash
make serve
```

TODO:
-----

Aritcles:

https://mycodesmells.com/post/authentication-in-grpc
https://bbengfort.github.io/programmer/2017/03/03/secure-grpc.html
