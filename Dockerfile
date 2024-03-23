FROM golang:1.22-alpine3.19 as builder

WORKDIR /usr/src/app

COPY . .

RUN go build -o main cmd/responder/main.go

FROM alpine:3.19

WORKDIR /usr/app

COPY --from=builder /usr/src/app/main .

CMD ["/usr/app/main"]