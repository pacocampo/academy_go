FROM golang:1.8.5

RUN mkdir -p /app

WORKDIR /app

ADD . /app

RUN go build ./app.go

EXPOSE 8080

CMD ["./app"]
