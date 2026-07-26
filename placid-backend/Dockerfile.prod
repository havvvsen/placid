# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/

RUN CGO_ENABLED=0 GOOS=linux go build -o placid-api ./cmd/server/main.go

# Stage 2
FROM alpine:latest  

# Install security certificates (critical for HTTPS APIs)
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/placid-api .

EXPOSE ${API_PORT}

CMD ["./placid-api"]
