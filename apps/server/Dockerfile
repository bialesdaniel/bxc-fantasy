# ---- Build Stage ----
# Use a specific version of the Go Alpine image as a builder.
# This aligns with the `go.mod` version and improves build reproducibility
# by pinning to a more specific tag instead of a floating one.
FROM golang:1.24.0-alpine3.19 AS builder

# Install build tools. git is required for 'go mod download'.
RUN apk add --no-cache git
# Install buf CLI for code generation
RUN go install github.com/bufbuild/buf/cmd/buf@latest

# Set the working directory inside the container.
WORKDIR /app

# Copy go.mod and go.sum to leverage Docker's layer caching.
# This step will only be re-run if these files change.
COPY apps/server/go.mod apps/server/go.sum ./
# Download dependencies.
RUN go mod download

# Copy the protobuf files.
COPY packages/protos/ /app/protos/
COPY buf.gen.yaml /app/buf.gen.yaml

# Generate protobuf files.
RUN buf generate protos

# Copy the server application source code.
COPY ./apps/server/internal/ ./apps/server/cmd/ ./apps/server/

# Build the Go application into a static binary.
# CGO_ENABLED=0 is crucial for creating a static binary that can run in a minimal image like Alpine.
RUN CGO_ENABLED=0 go build -o /server ./apps/server/cmd/main.go

# ---- Final Stage ----
# Use a distroless static image for a minimal and secure final container.
# It contains only our static binary and its direct runtime dependencies,
# and runs as a non-root user by default, significantly reducing the attack surface.
FROM gcr.io/distroless/static-debian12

# Copy the static binary from the builder stage.
COPY --from=builder /server /server

# Expose the port the application will run on.
EXPOSE 8080

# Set the command to run the application. The user is non-root by default.
CMD ["/server"]
