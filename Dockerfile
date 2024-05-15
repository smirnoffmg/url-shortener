FROM --platform=linux/amd64 golang:1.22-alpine as builder

# Install ca-certificates
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o shortener cmd/main.go

# Second stage: runtime environment
FROM scratch

# Copy the binary
COPY --from=builder /app/shortener .

# Copy ca-certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy config
# COPY --from=builder /app/config.yaml .


CMD ["./shortener"]