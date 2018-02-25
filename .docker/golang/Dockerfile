FROM golang:1.9

ADD . /go/src/github.com/Sharykhin/gl-mail-grpc-server

WORKDIR /go/src/github.com/Sharykhin/gl-mail-grpc-server

RUN go get .

RUN go install github.com/Sharykhin/gl-mail-grpc-server

ENTRYPOINT /go/bin/gl-mail-grpc-server

#CMD tail -f /dev/null
