FROM golang:1.11.3-alpine3.8

RUN apk --update --no-cache add git
RUN apk --update --no-cache add protobuf
RUN go get -v -u github.com/golang/protobuf/protoc-gen-go && rm -rf /go/*/*
RUN apk --update --no-cache add make ncurses perl
