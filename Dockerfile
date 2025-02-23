# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /build

# Install build dependencies
RUN apk add --no-cache make

# Copy source files
COPY . .

# Build the application
RUN make build

# Final stage
FROM alpine:3.19

WORKDIR /app

# Copy binary from builder
COPY --from=builder /build/bin/raft-server /app/

ENV RAFTIG_API_INTERNAL_PORT=8081
# Expose the API port
EXPOSE 8081

# Run the server
CMD ["/app/raft-server"]
