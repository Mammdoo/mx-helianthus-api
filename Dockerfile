FROM golang:1.21.6-alpine as builder

WORKDIR /app

ADD . /app

RUN go build -o helianthus

FROM golang:1.21.6-alpine

WORKDIR /app

COPY --from=builder /app/helianthus /app/helianthus

CMD ["/app/helianthus"]