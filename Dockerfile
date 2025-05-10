# build stage
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o bot ./cmd/bot

# final image (на базе того же golang)
FROM golang:1.23

WORKDIR /app
COPY --from=builder /app/bot .

CMD ["./bot"]