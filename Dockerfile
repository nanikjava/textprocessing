FROM golang:1.17 AS build

WORKDIR /app
RUN mkdir -p /app/test-files

COPY . .
RUN go mod download

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s"  -o rockt-parser

CMD ["/app/rockt-parser","-d","/app/test-files"]