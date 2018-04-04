FROM golang:1.9

ADD . /go/src/github.com/Sharykhin/gl-mail-grpc-server

WORKDIR /go/src/github.com/Sharykhin/gl-mail-grpc-server

RUN go get . && \
    go get github.com/canthefason/go-watcher && \
    go install github.com/canthefason/go-watcher/cmd/watcher && \
    go get -u github.com/pressly/goose/cmd/goose

ENTRYPOINT watcher

EXPOSE 50051
