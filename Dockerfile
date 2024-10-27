# Build stage
FROM golang:1.23-alpine AS builder
RUN apk add --no-cache make gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build

# Runtime stage
FROM alpine:3.18

# Create non-root user
RUN adduser -D -g '' appuser

WORKDIR /app

COPY --from=builder /app/bin/number-search .

COPY --from=builder /app/input.txt .

RUN chown -R appuser:appuser /app

USER appuser

ENV PORT=":8080" \
    LOG_LEVEL="info" \
    VARIATION="10" \
    FILE_PATH="./input.txt"

EXPOSE 8080

CMD ["./number-search"]