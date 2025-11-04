# ---- Protobuf Generation Stage ----
FROM golang:1.24.0-alpine3.19 AS protobuf-gen

# Set the working directory inside the container.
WORKDIR /app

# Install buf CLI for code generation
RUN go install github.com/bufbuild/buf/cmd/buf@latest

# Copy the protobuf files.
COPY packages/protos/ ./protos/
COPY buf.gen.yaml .

# Generate protobuf files.
RUN buf generate protos

# ---- Builder Stage ----
FROM golang:1.24.0-alpine3.19 AS builder

# Set the working directory inside the container.
WORKDIR /app

# Install build tools. git is required for 'go mod download'.
RUN apk add --no-cache git

COPY apps/server/go.mod apps/server/go.sum /app/
# Download dependencies.
RUN go mod download

# Copy generated files from the protobuf-gen stage.
COPY --from=protobuf-gen /app/apps/server/gen /app/gen

# Copy the server application source code.
COPY ./apps/server/ /app/

# Build the Go application into a static binary.
# CGO_ENABLED=0 is crucial for creating a static binary that can run in a minimal image like Alpine.
RUN CGO_ENABLED=0 go build -o /server /app/cmd/main.go

# ---- Final Stage ----
# Use a distroless static image for a minimal and secure final container.
# It contains only our static binary and its direct runtime dependencies,
# and runs as a non-root user by default, significantly reducing the attack surface.
FROM gcr.io/distroless/static-debian12

# Copy the static binary from the builder stage.
COPY --from=builder /server /server

# Expose the port the application will run on.
ARG PORT=8080
EXPOSE ${PORT}

# Set the command to run the application. The user is non-root by default.
CMD ["/server"]
