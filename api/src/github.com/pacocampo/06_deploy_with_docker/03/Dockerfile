FROM golang:1.8.5-alpine3.6

RUN mkdir -p /app \
    && apk add --no-cache git mercurial \
    && go get github.com/gorilla/mux \
    && apk del git mercurial

WORKDIR /app

ADD . /app

ENV "SECRET_WORD"=" env desde docker"

RUN go build ./app.go

EXPOSE 8080

CMD ["./app"]
