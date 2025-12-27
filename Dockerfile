# Build

FROM golang:alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /stugo main.go

# Run

FROM alpine:latest

WORKDIR /app

COPY --from=builder /stugo .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/data ./data

EXPOSE 8080

CMD ["./stugo"]