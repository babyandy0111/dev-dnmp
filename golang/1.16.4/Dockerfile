FROM golang:1.16.4

WORKDIR /go/src/goweb

COPY /app .

EXPOSE 8080

RUN go build -o main main.go

ENTRYPOINT ./main
