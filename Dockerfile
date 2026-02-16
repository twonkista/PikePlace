# Stage 1: The Builder (The "Compiler" state)
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
# Compile the "91-grade" binary
RUN go build -o main ./cmd/api

# Stage 2: The Runner (The "Execution" state)
FROM alpine:latest
WORKDIR /root/
# Only copy the binary from the builder
COPY --from=builder /app/main .
# Expose the port your mux is listening on
EXPOSE 8080
CMD ["./main"]