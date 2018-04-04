FROM golang:1.9

ADD . /go/src/github.com/Sharykhin/gl-mail-grpc-server

WORKDIR /go/src/github.com/Sharykhin/gl-mail-grpc-server

RUN go get . && \
    go get -u github.com/pressly/goose/cmd/goose && \
    go install github.com/Sharykhin/gl-mail-grpc-server

ENTRYPOINT /go/bin/gl-mail-grpc-server

EXPOSE 50051
