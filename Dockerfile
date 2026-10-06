FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/bin/any-business ./cmd/any-business

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /app/bin/any-business /any-business
COPY --from=builder /app/.env.example /.env

EXPOSE 18000

ENTRYPOINT ["/any-business"]
