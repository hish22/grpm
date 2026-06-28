# Stage 1: Build the binary (using the Go image)
FROM docker.io/library/golang:1.26 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/app-binary ./cmd

# Stage 2: Create the Ubuntu production image
FROM docker.io/library/ubuntu:24.04
# Install basic certificates (needed if your app makes HTTPS requests)
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the binary from the builder stage
WORKDIR /app
COPY --from=builder /app/app-binary /app/app-binary

# Set the entrypoint
ENTRYPOINT ["/bin/bash"]
