# ==========================
# Stage 1: Build Application
# ==========================
FROM golang:1.26.5 AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build application
RUN CGO_ENABLED=0 GOOS=linux go build -o task-api ./Cmd/Api

# ==========================
# Stage 2: Runtime
# ==========================
FROM alpine:latest

WORKDIR /app

# Copy executable from builder
COPY --from=builder /app/task-api .


# Expose application port
EXPOSE 8080

# Start application
CMD ["./task-api"]