# Base stage for shared dependencies
FROM golang:1.24-alpine AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Development stage
FROM base AS development
RUN go install github.com/air-verse/air@latest
WORKDIR /app/scheduler
ENV GO111MODULE=on
CMD ["air"]

# Build stage
FROM base AS builder
WORKDIR /app
COPY pkg/ ./pkg/
COPY scheduler/ ./scheduler/
RUN CGO_ENABLED=0 GOOS=linux go build -o /scheduler ./scheduler

# Production stage
FROM alpine:3.19 AS production
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /scheduler .
RUN adduser -D appuser
USER appuser
ENTRYPOINT ["./scheduler"]
CMD ["wait"]
