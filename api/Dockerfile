# Build Stage
FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Final Stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /docker-gs-ping /app/.env ./
CMD ["./docker-gs-ping"]