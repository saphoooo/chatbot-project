FROM golang:1.14rc1-buster as builder

ENV GO111MODULE=on

COPY * /app/

WORKDIR /app

RUN CGO_ENABLED=0 go build -o bot

FROM scratch

COPY --from=builder /app/bot .

EXPOSE 9090

ENTRYPOINT [ "./bot" ]