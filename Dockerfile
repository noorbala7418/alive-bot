FROM golang:1.20.2-alpine3.17 as builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . ./

RUN go build -v -o bot

FROM golang:1.20.2-alpine3.17

ENV TELEGRAM_BOT_DEBUG_MODE=false

COPY --from=builder /app/bot /app/bot

CMD ["/app/bot"]
