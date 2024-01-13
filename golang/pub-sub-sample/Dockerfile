FROM golang:1.20.4-alpine3.18

RUN apk update

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod vendor

ADD . /go/src/app/

RUN go build -o /go/src/app/binary

CMD ["/go/src/app/binary"]