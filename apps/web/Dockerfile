# ---- Protobuf Generation Stage ----
# This stage uses the same Go-based setup as the server's Dockerfile to ensure
# consistency in tooling for protobuf generation.
FROM golang:1.24-alpine AS protobuf-gen

# Set the working directory.
WORKDIR /app

# Install buf CLI for code generation.
# Pinning the version in a real-world scenario would be a best practice.
RUN go install github.com/bufbuild/buf/cmd/buf@latest

# Copy only the necessary files for protobuf generation to leverage Docker caching.
COPY packages/protos ./packages/protos
COPY buf.gen.yaml .
COPY buf.yaml ./

# Generate protobuf files for all defined plugins (Go server and Web client).
# The generated web client files will be copied to the builder stage.
RUN buf generate

# ---- Builder Stage ----
# This stage builds the React application.
FROM node:20-alpine AS builder

# Set the working directory.
WORKDIR /app

# Install pnpm, the package manager used in this monorepo.
RUN npm install -g pnpm

# Copy workspace dependency definition files.
COPY package.json pnpm-lock.yaml pnpm-workspace.yaml ./

# Copy application-specific dependency definitions.
COPY apps/web/package.json ./apps/web/

# Install all dependencies for the entire monorepo.
# This ensures that any workspace-linked dependencies are available.
RUN pnpm install

# Copy the web application's source code.
COPY apps/web ./apps/web

# Copy the generated protobuf files from the protobuf-gen stage.
# This is the critical step that provides the generated client code to the build.
COPY --from=protobuf-gen /app/apps/web/src/gen ./apps/web/src/gen

# Build the web application, filtering to the 'web' package.
# This runs the "build" script defined in apps/web/package.json.
RUN pnpm --filter web run build

# ---- Final Stage ----
# This stage serves the built static assets using Nginx. It's a minimal, secure
# image containing only the web server and the application's production assets.
FROM nginx:stable-alpine AS final

# Copy the custom Nginx configuration.
COPY apps/web/nginx.conf /etc/nginx/conf.d/default.conf

# Copy the production-ready static assets from the builder stage.
COPY --from=builder /app/apps/web/dist /usr/share/nginx/html

# Expose the port Nginx will listen on, matching Cloud Run's requirements.
EXPOSE 8080

# The default Nginx CMD is ["nginx", "-g", "daemon off;"], which is exactly
# what is needed to run the server in the foreground.
