# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app
COPY assets/ assets/
COPY go.mod ./

COPY main.go main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o placid-file-server ./main.go

# Stage 2
FROM alpine:latest  

# Install security certificates (critical for HTTPS APIs)
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/placid-file-server .

EXPOSE ${FILE_SERVER_PORT}

CMD ["./placid-file-server"]
