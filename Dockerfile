FROM golang:1.24-alpine AS build

ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o server ./cmd/server

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/server .

COPY --from=build /app/.env .env

CMD ["./server"]
