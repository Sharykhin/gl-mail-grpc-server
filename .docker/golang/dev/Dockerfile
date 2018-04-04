FROM golang:1.9

ADD . /go/src/github.com/Sharykhin/gl-mail-grpc-server

WORKDIR /go/src/github.com/Sharykhin/gl-mail-grpc-server

RUN go get . && go get github.com/canthefason/go-watcher && go install github.com/canthefason/go-watcher/cmd/watcher

#ENTRYPOINT watcher

CMD tail -f /dev/null

EXPOSE 50051