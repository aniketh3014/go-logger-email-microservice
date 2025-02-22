# build the base image
FROM golang:1.24-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o brokerService ./cmd/api

RUN chmod +x /app/brokerService

# build and run the small image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/brokerService /app

CMD ["/app/brokerService"]