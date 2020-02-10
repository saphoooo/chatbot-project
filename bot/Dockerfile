FROM golang:rc-alpine3.11 as builder

ENV GO111MODULE=on

COPY * /app/

WORKDIR /app

RUN CGO_ENABLED=0 go build -o bot

FROM alpine:3.9

COPY --from=builder /app/bot .

EXPOSE 9090

ENTRYPOINT [ "./bot" ]