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

When you make ```docker-compose up``` it will run watcher on go files, so each
change will trigger re-build application. In a terminal output you may see the
following error:
```
gl-mail-grpc-server-mysql     | 2018-04-04T09:25:23.950720Z 2 [Note] Aborted connection 2 to db: 'test' user: 'test' host: '172.18.0.3' (Got an error reading communication packets)
```
So, don't worry about it, when watcher makes re-build it somehow affects mysql connection.
But it doesn't affect the program itself.

Makefile:
```bash
make serve
```

TODO:
-----

Aritcles:

https://mycodesmells.com/post/authentication-in-grpc
https://bbengfort.github.io/programmer/2017/03/03/secure-grpc.html
