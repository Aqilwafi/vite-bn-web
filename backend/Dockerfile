# ---------- Stage 1: Build ----------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod dan go.sum terlebih dahulu untuk cache layer
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build binary dengan flags untuk ukuran kecil
RUN go build -ldflags="-s -w" -o main .

# ---------- Stage 2: Final ----------
FROM alpine:latest

# Tambah CA certificates (penting untuk koneksi HTTPS/DB SSL)
RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copy binary dari builder
COPY --from=builder /app/main .

# Railway inject PORT otomatis, jadi tidak perlu hardcode
EXPOSE 4000

# Jalankan binary
CMD ["./main"]
