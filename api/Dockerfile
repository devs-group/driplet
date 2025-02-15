# Base stage for shared dependencies
FROM golang:1.24-alpine AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Development stage
FROM base AS development
RUN go install github.com/air-verse/air@latest
WORKDIR /app/api
ENV GO111MODULE=on
CMD ["air"]

# Build stage
FROM base AS builder
WORKDIR /app
COPY pkg/ ./pkg/
COPY api/ ./api/
RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./api

# Production stage
FROM alpine:3.19 AS production
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /api .
ENV PORT=9000
RUN adduser -D appuser
USER appuser
EXPOSE ${PORT}
ENTRYPOINT ["./api"]
CMD ["run"]
