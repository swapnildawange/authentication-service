FROM golang:1.19.3-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 &&  go build -o authApp ./cmd/main.go

RUN chmod +x /app/authApp


FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/authApp /app
COPY --from=builder /app/env  /app

CMD ["/app/authApp"]