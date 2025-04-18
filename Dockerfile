FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go-ads-manager/go.mod go-ads-manager/go.sum ./
RUN go mod tidy

COPY go-ads-manager/ ./

RUN go build -o server ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/server .

EXPOSE 8080

CMD ["./server"]
