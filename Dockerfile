FROM golang:1.22.5 as builder

WORKDIR /app

COPY go.mod go.sum ./ 

RUN go mod download

COPY . .

RUN go build -o http-proxy-task1 .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/http-proxy-task1 .

CMD ["./http-proxy-task1"]