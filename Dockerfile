FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /http-proxy-task1 ./main.go

FROM alpine:latest

COPY --from=builder /http-proxy-task1 /http-proxy-task1

EXPOSE 8080

CMD ["/http-proxy-task1"]
