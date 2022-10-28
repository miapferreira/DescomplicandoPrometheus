FROM golang:1.19.0-alpine3.16 AS buildando

WORKDIR /app
COPY . /app

RUN go build exporter.go



FROM alpine:3.16

COPY --from=buildando /app/exporter /app/exporter
EXPOSE 7788
WORKDIR /app
CMD ["./exporter"]