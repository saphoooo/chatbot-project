FROM golang:1.14rc1-buster as builder

ENV GO111MODULE=on

COPY * /app/

WORKDIR /app

RUN CGO_ENABLED=0 go build -o chat

FROM alpine:3.9

COPY avatars /avatars/

COPY templates /templates/

COPY --from=builder /app/chat .

EXPOSE 8080

ENTRYPOINT [ "./chat" ]